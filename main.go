package main

import (
	"fmt"
	"head/main_"
	"head/main_/func_"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := func_.FindFreePort()

	func_.WriteStartArticle(port)

	http.HandleFunc("/", main_.Handler)

	go func() {
		fmt.Printf("Server starting at http://localhost:%s\n", strconv.Itoa(port))
		if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	}()

	func_.RunExe()

	os.Exit(0)
}
