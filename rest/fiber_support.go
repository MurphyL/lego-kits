package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/MurphyL/lego-kits/core"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"gorm.io/gorm"
)

type ActionEnum string

const (
	ActionCreate ActionEnum = "create"
	ActionUpdate ActionEnum = "update"
	ActionRead   ActionEnum = "read"
	ActionDelete ActionEnum = "delete"
)

type Model interface {
}

type ActionResult struct {
	Action  ActionEnum
	Message string
	Payload any
}

type PageQuery struct {
	PageNum  int
	PageSize int
}

type SkipQuery struct {
	Key   string
	Value string
	Count int
}

type QueryResult[T Model, Q any] struct {
	Params  Q
	Total   int64
	Records []T
}

func NewFiberApp(prefork bool) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		Immutable:             true,
		DisableStartupMessage: true,
		Prefork:               prefork,
		AppName:               "lego-kits",
		ServerHeader:          "lego-core v1.0.1",
		ErrorHandler:          DefaultErrorHandler,
	})
	// 添加压缩插件
	fiberApp.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 或使用 compress.LevelZstd
	}))
	return fiberApp
}

func NewActionResult(action ActionEnum, payload any, errs ...error) *ActionResult {
	return &ActionResult{Action: action, Payload: payload, Message: errors.Join(errs...).Error()}
}

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	// Return status code with error message
	return c.Status(code).SendString(err.Error())
}

func DumpHandler(c *fiber.Ctx) error {
	buf := make([]byte, 1<<20) // 1MB buffer
	n := runtime.Stack(buf, true)
	return c.Send(buf[:n])
}

func HandleModelRequests[R Model](db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var err error
		switch ctx.Method() {
		case http.MethodOptions:
			return ctx.Next()
		case http.MethodTrace:
			return ctx.SendStatus(http.StatusForbidden)
		case http.MethodGet: // 不推荐（无语义）有 Body
			if strings.HasSuffix(ctx.Path(), "/paging") { // 分页查询
				pageInfo := PageQuery{
					PageNum:  ctx.QueryInt("pageNum", 1),
					PageSize: ctx.QueryInt("pageSize", 10),
				}
				tx, cnt := core.NewPageQuery(db, &pageInfo, new(R))
				records := make([]R, 0)
				if err = tx.Find(&records).Error; err == nil {
					return ctx.JSON(QueryResult[R, PageQuery]{Params: pageInfo, Total: cnt, Records: records})
				}
			} else if strings.HasSuffix(ctx.Path(), "/skip") { // 跳表查询
				skipVars := SkipQuery{
					Key:   ctx.Query("key", "id"),
					Value: ctx.Query("value"),
					Count: ctx.QueryInt("count", 10),
				}
				tx, cnt := core.NewSkipQuery(db, &skipVars, new(R))
				records := make([]R, 0)
				if err = tx.Find(&records).Error; err == nil {
					return ctx.JSON(QueryResult[R, SkipQuery]{Params: skipVars, Total: cnt, Records: records})
				}
			} else if strings.HasSuffix(ctx.Path(), "/list") {
				records := make([]R, 0)
				query := ctx.Queries()
				if len(query) > 0 {
					err = db.Find(&records, query).Error
				} else {
					err = db.Find(&records).Error
				}
				if err == nil {
					return ctx.JSON(QueryResult[R, any]{Params: query, Total: int64(len(records)), Records: records})
				}
			} else {
				if query := ctx.Queries(); len(query) > 0 {
					record := new(R)
					err = db.First(record, query).Error
					if err == nil {
						return ctx.JSON(record)
					}
				} else {
					return errors.New("query string is empty")
				}
			}
		case http.MethodPost: // 创建资源/提交数据，有 Body
			payload := new(R)
			if err = json.Unmarshal(ctx.BodyRaw(), payload); err == nil {
				ret := db.Create(payload)
				if err = ret.Error; err == nil && ret.RowsAffected > 0 {
					return ctx.JSON(payload)
				}
			}
		case http.MethodDelete: // 删除资源
			if query := ctx.Queries(); len(query) > 0 {
				ret := db.Model(new(R)).Delete(query)
				if err = ret.Error; err == nil && ret.RowsAffected > 0 {
					return ctx.SendStatus(http.StatusOK)
				} else {
					data, _ := json.Marshal(query)
					return fmt.Errorf("record not exists: %v", string(data))
				}
			} else {
				return errors.New("query string is empty")
			}
		case http.MethodPut: // 整体替换资源，有 Body
			payload := new(R)
			if err = json.Unmarshal(ctx.BodyRaw(), payload); err == nil {
				ret := db.Save(payload)
				if err = ret.Error; err == nil && ret.RowsAffected > 0 {
					return ctx.SendStatus(http.StatusOK)
				} else {
					data, _ := json.Marshal(payload)
					return fmt.Errorf("record not exists: %v", string(data))
				}
			}
		case http.MethodPatch: // 部分更新资源，有 Body
			payload := new(R)
			if err = ctx.QueryParser(payload); err == nil {
				id := ctx.QueryInt("id", 0)
				ret := db.Model(&payload).Where("id = ?", id).Updates(payload)
				if err = ret.Error; err == nil && ret.RowsAffected > 0 {
					return ctx.SendStatus(http.StatusOK)
				} else {
					return fmt.Errorf("record(id = %v) not exists", id)
				}
			}
		}
		if err == nil {
			return errors.New("no impl")
		}
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
}
