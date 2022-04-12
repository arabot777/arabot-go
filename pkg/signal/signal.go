package signal

import (
	"github.com/arabot777/arabot-go/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Wait() {
	if !logger.IsInitialized() {
		panic("signal: fail to get a valid logger")
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-ch
		logger.Infof("signal: service got signal: %s!", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logger.Infof("signal: service exit now!")
			os.Exit(0)
		case syscall.SIGHUP:
			logger.Infof("signal: got signal hup!")
		default:
			logger.Infof("signal: got unknown signal %s!", sig.String())
			os.Exit(0)
		}
	}
}
