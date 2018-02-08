package main

import (
	"net/http"
	"http_server/crud"
)

func main() {

	/** FOR TEST */
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", crud.ViewHandler)

	http.HandleFunc("/edit/", crud.EditHandler)

	//http.HandleFunc("/save/", crud.SaveHandler)

	http.ListenAndServe(":9000", nil)

} //main
