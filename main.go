package main

import (
	"din"
	"fmt"
	"net/http"
)

func main() {
	d := din.New()
	d.Get("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%q", "this is foo")
	})
	d.Run(":8888")

}
