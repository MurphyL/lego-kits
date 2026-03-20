package core

import (
	"context"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/cgi/internal/sugar"
)

var sugarLogger = sugar.NewSugarLogger()

func NewLegoApp(appConfig AppConfig) *LegoApp {
	ac := fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       appConfig.AppTitle(),
	}
	// 应用服务
	l := &LegoApp{App: fiber.New(ac)}
	// 注册关闭前钩子
	l.App.Hooks().OnPreShutdown(func() error {
		sugarLogger.Infoln("Server is shutting down...")
		return nil
	})
	return l
}

type LegoApp struct {
	*fiber.App
}

// AppConfig 是应用配置接口
type AppConfig interface {
	AppTitle() string
}

// AppHandler 是应用处理程序接口，用于注册路由
type AppHandler interface {
	RegisterRoutes(router fiber.Router)
}

func (l *LegoApp) Mount(url string, useRouterGroup func(router fiber.Router)) {
	useRouterGroup(l.App.Group(path.Join("/api", url)))
}

func (l *LegoApp) Serve(addr string) {
	// 启动服务器协程
	go func() {
		if err := l.App.Listen(addr); err != nil {
			sugarLogger.Info("Server Shutdown:", err.Error())
		}
	}()
	sugarLogger.Info("Server started:", addr)
	// 监听中断信号并触发优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	// 创建带超时的上下文，限制最长等待30秒
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭
	if err := l.App.ShutdownWithContext(ctx); err != nil {
		sugarLogger.Info("Server failed:", err)
	}
	sugarLogger.Info("Server exited")
}
