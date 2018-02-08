package main

import (
	"http_server/crud"
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/view/", crud.ViewHandler)

	http.HandleFunc("/edit/", crud.EditHandler)

	http.HandleFunc("/save/", crud.SaveHandler)

	fmt.Printf("Server listen to 9000 port ... \n")

	http.ListenAndServe(":9000", nil)

} //main
