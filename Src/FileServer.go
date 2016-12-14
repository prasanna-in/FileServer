package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

func main() {
	location := os.Args[1:][0]
	fmt.Println("Listening on Port: 9099")
	fmt.Println("Display Folder   : "+location)
	http.Handle("/",http.FileServer(http.Dir(location)))
	log.Fatal(http.ListenAndServe(":9099",nil))
}
