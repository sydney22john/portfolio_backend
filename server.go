package main

import (
	"fmt"
	"net/http"
)

func runServer(addr string) {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /contactme", func(w http.ResponseWriter, r *http.Request) {
		type ContactForm struct {
			SenderName  string `json:"sender-name"`
			SenderEmail string `json:"sender-email"`
			Message     string `json:"message"`
		}

		body := &ContactForm{}
		if err := decodeRequestParams(r, body); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", *body)
	})

	server := http.Server{
		Addr:    addr,
		Handler: middlewareCors(mux),
	}

	server.ListenAndServe()
}
