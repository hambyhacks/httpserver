package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Config
	var port int

	// Banner
	fmt.Print("====================================================================================\n")
	fmt.Print("\t\t\t\tGo-Http.Server\n\n")
	fmt.Print("Description: Similar to python3 -m http.server [PORT]. Only implemented file server.\n")
	fmt.Print("Usage: httpserver -p [PORT] (defaults to 9000)\n\n")
	fmt.Print("Written by: hambyhacks\n")
	fmt.Print("Twitter: @hambyhaxx\n")
	fmt.Print("Github: github.com/hambyhacks\n")
	fmt.Print("====================================================================================\n\n")
	// Set flags
	flag.IntVar(&port, "p", 9000, "Port number to be used.")
	flag.Parse()

	// Set HTTP server config
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      fsHandler(),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	// Start HTTP server as background process
	go func() {
		log.Printf("Starting HTTP server at port %s\n", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Shutdown server gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	log.Println("Signal: ", sig)

	// Shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
