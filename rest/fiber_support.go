package main

import (
	"errors"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

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
