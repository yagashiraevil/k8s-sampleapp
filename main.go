package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	myhandlers "github.com/yagashiraevil/k8s-sampleapp/handlers"
	"github.com/yagashiraevil/k8s-sampleapp/util"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := hclog.Default()
	var config, err = util.LoadConfig(".")
	if err != nil {
		l.Error("failed to load config", "error", err)
	}

	// create a new server mux and register handlers
	sm := mux.NewRouter()

	// Handle CORs
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	// Routes
	sm.HandleFunc("/", myhandlers.NewHome)

	s := http.Server{
		Addr:    config.BindAddr,
		Handler: ch(sm),
		ErrorLog: l.StandardLogger(&hclog.StandardLoggerOptions{
			InferLevels: true,
		}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		l.Info("starting server", "addr", config.BindAddr)
		if err := s.ListenAndServe(); err != nil {
			l.Error("failed to start server", "error", err)
		}
	}()
	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	l.Info("received signal, exiting", "signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err = s.Shutdown(ctx)
	if err != nil {
		l.Error("failed to shutdown server", "error", err)
	}
	defer cancel()

}
