package model

import "io/ioutil"

type Page struct {
	Title string

	Body []byte
} //Page struct

func (p *Page) Save() error {

	filename := "storage/" + p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)

} //save

func LoadPage(title string) (*Page, error) {

	filename := "storage/" + title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil

} // loadPage

