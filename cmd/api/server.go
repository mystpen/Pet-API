package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mystpen/Pet-API/config"
)

func Start(cfg *config.Config) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      nil, // TODO: add handlers
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	
	log.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
