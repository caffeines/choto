package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/caffeines/choto/config"
	"github.com/caffeines/choto/log"
)

// RunServer will run http server
func RunServer() {
	r := GetRouter()
	port, host := config.App().Port, config.App().Base
	addr := fmt.Sprintf("%s:%d", host, port)
	srv := &http.Server{
		Addr:         addr,
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
	fmt.Printf("server running on port %d...\n", port)

	signal.Notify(c, os.Interrupt)
	<-c
	wait := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Log().Infoln("gracefully shutting down")
	os.Exit(0)

}
