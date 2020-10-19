package main

import (
	"fmt"
	"gosys/pkg/testser"
)

func main() {

	helloHandler := ser.Serve(3)
	fmt.Println(helloHandler)
}
