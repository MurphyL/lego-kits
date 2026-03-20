package cgi

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"murphyl.com/lego/cgi/internal"
)

/*
 * cgi 模块用于创建和管理基于 gofiber/v3 的服务
 * 主要功能包括：创建应用、挂载路由、启动服务等
 */
var sugarLogger = internal.NewSugarLogger()

/*
*
  - 创建LegoApp实例
  - @param appConfig app.AppConfig 应用配置
  - @return LegoAppInterface 应用接口
*/
func NewLegoApp(appConfig internal.AppConfig) LegoApp {
	return internal.NewLegoApp(appConfig)
}

type LegoApp interface {
	Mount(url string, useRouterGroup func(router fiber.Router))
	Serve(addr string)
}

/**
 * 从数据库中检索所有记录
 * @param c fiber.Ctx 请求上下文
 * @param db *gorm.DB 数据库连接
 * @return error 错误信息
 */
func RetrieveAll[Q any, T any](c fiber.Ctx, db *gorm.DB) error {
	var query = new(Q)
	if err := c.Bind().All(query); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	var records []T
	if err := db.Where(query).Find(&records).Error; err != nil {
		sugarLogger.Error("查询列表出错：", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(records)
}

/**
 * 从数据库中检索一条记录
 * @param c fiber.Ctx 请求上下文
 * @param db *gorm.DB 数据库连接
 * @return error 错误信息
 */
func RetrieveOne[Q any, T any](c fiber.Ctx, db *gorm.DB) error {
	var query = new(Q)
	if err := c.Bind().All(query); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	var record T
	if err := db.Where(query).Take(&record).Error; err != nil {
		sugarLogger.Error("查询记录出错：", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(record)
}

/**
 * 创建一条记录
 * @param c fiber.Ctx 请求上下文
 * @param db *gorm.DB 数据库连接
 * @return error 错误信息
 */
func CreateOne[T any](c fiber.Ctx, db *gorm.DB) error {
	var payload = new(T)
	if err := c.Bind().Body(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	if err := db.Create(&payload).Error; err != nil {
		sugarLogger.Error("创建记录出错：", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(payload)
}

/**
 * 从数据库中删除一条记录
 * @param c fiber.Ctx 请求上下文
 * @param db *gorm.DB 数据库连接
 * @return error 错误信息
 */
func DeleteOne[Q any, T any](c fiber.Ctx, db *gorm.DB) error {
	var query = new(Q)
	if err := c.Bind().All(query); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	if err := db.Where(query).Delete(new(T)).Error; err != nil {
		sugarLogger.Error("删除记录出错：", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
