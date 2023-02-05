package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

//templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.tmpl.Execute(w, nil)
}

func main() {
	var addr = flag.String("addr", ":8000", "The addr of the application.")
	flag.Parse() // parse the flags
	
	r := newRoom()
	
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	
	//run the room
	go r.run()
	
	//start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
