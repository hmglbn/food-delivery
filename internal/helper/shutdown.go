package helper

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ListenSigInt() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	Log("press ctrl+c to quit")
	<-done
}

func CloseApp(ctxCancelFunc context.CancelFunc) {
	Log("application is closing now")
	ctxCancelFunc()
	Log("wait to finish everything")
	time.Sleep(time.Microsecond * 10)
	Log("exit")
}
