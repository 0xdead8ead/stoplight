/* Firewall Request Management Webapp using Gorilla & HTML Templates */
package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := map[string]string{"name": params["name"]}
	var indexTemplate = template.Must(template.New("index").ParseFiles("templates/base.html", "templates/index.html"))
	indexTemplate.ExecuteTemplate(w, "base", name)
}

//About Handler
func About(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var aboutTemplate = template.Must(template.New("about").ParseFiles("templates/base.html", "templates/about.html"))
	aboutTemplate.ExecuteTemplate(w, "base", params)
}

//Firewall Request Handler
func Req(w http.ResponseWriter, r *http.Request) {
	var pass string
	if r.Method == "GET" {
		var reqTemplate = template.Must(template.New("req").ParseFiles("templates/base.html", "templates/req.html"))
		reqTemplate.ExecuteTemplate(w, "base", pass)
	} else if r.Method == "POST" {
		responseBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		fwreq := map[string]interface{}{}
		err = json.Unmarshal(responseBody, &fwreq)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprint(w, "\n--- Request Submitted! ---\n\n%s", fwreq)
	}
}

//Status Request Handler
func Status(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var statusTemplate = template.Must(template.New("status").ParseFiles("templates/base.html", "templates/status.html"))
	statusTemplate.ExecuteTemplate(w, "base", params)
}

//Approval Handler
func Approve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var approveTemplate = template.Must(template.New("approve").ParseFiles("templates/base.html", "templates/approve.html"))
	approveTemplate.ExecuteTemplate(w, "base", params)
}

//Blog Handler
func Blog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var blogTemplate = template.Must(template.New("blog").ParseFiles("templates/base.html", "templates/blog.html"))
	blogTemplate.ExecuteTemplate(w, "base", params)
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	var pass string
	var notFoundTemplate = template.Must(template.New("notfound").ParseFiles("templates/404.html"))
	notFoundTemplate.ExecuteTemplate(w, "content", pass)
}
