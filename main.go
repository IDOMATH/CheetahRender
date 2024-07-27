package main

import (
	"fmt"
	"log"
	"net/http"

	renderer "github.com/idomath/CheetahRender/Renderer"
)

func main() {
	fmt.Println("Hello world!")
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	rr := renderer.NewRenderer("./views", ".html", "./views", ".html", false)

	router.HandleFunc("GET /", handleHome)

	log.Fatal(server.ListenAndServe())
}

func handleHome(rr Renderer) {
	rr.Renderer.Redner("home.html")
}
