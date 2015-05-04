// A simple web server

package main

import (
	"fmt"
	"log"
	"net/http"
	"gowiki/wiki"
)

func printHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There, I love %s!", r.URL.Path[len("/print/"):])
}

func viewHandler(w http.ResponseWriter, r *http.Request){

	title := r.URL.Path[len("/view/"):]
	p, _ := wiki.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/print/", printHandler)
	http.HandleFunc("/view/",viewHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}

}
