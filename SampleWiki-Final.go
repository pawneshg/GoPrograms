package main

import(
	//"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"html/template"
	"regexp"
	"errors"
)
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

func getTitle(w http.ResponseWriter, r *http.Request) (string, error){
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil{
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil{
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}



func viewHandler(w http.ResponseWriter, r *http.Request, title string){
	//title := r.URL.Path[len("/view/"):]
	//title, err := getTitle(w, r)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
//Notice that we've used almost exactly the same templating code in both handlers.
// Let's remove this duplication by moving the templating code to its own function:
//for each handler i am using template.ParseFiles and execute
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	//t, err := template.ParseFiles(tmpl + ".html")
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	// err = t.Execute(w, p)
	// if err != nil{
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}



func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	//title, err := getTitle(w, r)
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
	// t, err := template.ParseFiles("edit.html")
	// if err != nil{
	// 	fmt.Fprintf(w, "<h6>%error</h6>",err)
	// }
	// t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string){
	// title := r.URL.Path[len("/save/"):]
	//title, err := getTitle(w, r)
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main(){
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}





