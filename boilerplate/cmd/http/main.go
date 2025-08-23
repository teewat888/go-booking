package main

import (
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/teewat888/go-booking/boilerplate/internal/config"
	"github.com/teewat888/go-booking/boilerplate/internal/dependencies"
	httpServer "github.com/teewat888/go-booking/boilerplate/internal/http"
	"github.com/teewat888/go-booking/msgoutils"
)

func main() {

	cfg := config.FromEnv()

	deps := dependencies.InitDependencies(&cfg)

	logrus.WithField("service_id", cfg.ServiceId).Info("Starting the service")

	srv := httpServer.New(deps)

	srv.Configure()

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt)

	jwtsrv := msgoutils.NewJWTService("sss")
	jwtsrv.ValidateToken("sss")

	go func() {
		if err := srv.Start(); err != nil {
			logrus.WithError(err).Fatal("Cannot start the server")
		}
	}()

	<-osSignal

	if err := srv.Close(); err != nil {
		logrus.WithError(err).Fatal("Cannot close the server correctly")
	}

}
