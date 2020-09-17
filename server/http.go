package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/caffeines/choto/log"
)

// RunServer will run http server
func RunServer() {
	r := GetRouter()

	srv := &http.Server{
		Addr:         "0.0.0.0:4521",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Log().Errorln(err)
		}
	}()
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c
	wait := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Log().Infoln("shutting down")
	os.Exit(0)

}
