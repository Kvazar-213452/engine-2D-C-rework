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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("front/static"))))

	http.HandleFunc("/save_file", main_.SaveFileHandler)
	http.HandleFunc("/", main_.Handler)
	http.HandleFunc("/stop", main_.StopProject)
	http.HandleFunc("/open_file", main_.OpenFile)
	http.HandleFunc("/start", main_.StartHandler)

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
