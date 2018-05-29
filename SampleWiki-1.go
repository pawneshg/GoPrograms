package main

import(
	"fmt"
	"io/ioutil"
)
//The Page struct describes how page data will be stored in memory.
type Page struct {
	Title string
	Body []byte
}
// But what about persistent storage? 
//We can address that by creating a save method on Page:
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
//Above method: This method's signature reads: "This is a method named save that takes as its receiver p, a pointer to Page .
// It takes no parameters, and returns a value of type error.

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main(){
	p1 := &Page{Title: "TestPage", Body: []byte("This is sample page")}
	p1.save()
	p3 := &Page{Title: "TestPage1", Body: []byte("This is sample page1")}
	p3.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}





