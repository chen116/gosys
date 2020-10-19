package main

import (
	"fmt"
	"gosys/pkg/testser"
	"io"
	"log"
	"net/http"
	"rsc.io/quote"
)

func main() {

	res := ser.Serve(3)
	fmt.Println(res)
	fmt.Println(quote.Hello())
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
