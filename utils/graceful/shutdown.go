package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GinShutdown 停止gin服务
func GinShutdown(instance *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭 Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := instance.Shutdown(ctx); err != nil {
		log.Fatal("Server 关闭：", err)
	}
	// 超时3秒 ctx.Done()
	select {
	case <-ctx.Done():
		log.Println("超时3秒.")
	}
	log.Println("Server 退出")
}
