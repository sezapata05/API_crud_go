package main

import (
	"api_golang/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8080")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello Word!")
	// })

	// srv := http.Server{
	// 	Addr: ":8080",
	// }

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Sever Started!")

	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("server stopped!")
}
