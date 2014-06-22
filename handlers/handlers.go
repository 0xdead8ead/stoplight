/* Firewall Request Management Webapp using Gorilla & HTML Templates */
package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/f47h3r/stoplight/lib"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

type status_page struct {
	Id           string
	FirewallJSON string
}

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	var pass string
	//name := map[string]string{"name": params["name"]}
	var indexTemplate = template.Must(template.New("index").ParseFiles("templates/base.html", "templates/index.html"))
	indexTemplate.ExecuteTemplate(w, "base", pass)
}

//About Handler
func About(w http.ResponseWriter, r *http.Request) {
	var pass string
	var aboutTemplate = template.Must(template.New("about").ParseFiles("templates/base.html", "templates/about.html"))
	aboutTemplate.ExecuteTemplate(w, "base", pass)
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
		//TODO - Append status variable to fwreq
		fwreq["status"] = "networkPending"
		fwRequestId := firewall.SaveFirewall(fwreq)
		statusUrl := fmt.Sprintf("http://%s/status/%s", r.Host, fwRequestId)
		firewall.GenStatusQRCode(statusUrl)

		//Send Status URL to redirect to
		fmt.Fprintf(w, "%s", statusUrl)
	}
}

//Status Request Handler
func Status(w http.ResponseWriter, r *http.Request) {
	var pass string
	var statusTemplate = template.Must(template.New("status").ParseFiles("templates/base.html", "templates/status.html"))

	statusTemplate.ExecuteTemplate(w, "base", pass)
}

func StatusById(w http.ResponseWriter, r *http.Request) {
	//var pass string
	vars := mux.Vars(r)
	fwRequestId := vars["fwRequestId"]

	if bson.IsObjectIdHex(fwRequestId) {
		firewallReq := firewall.GetFirewallStatusByID(fwRequestId)
		//Serialize to JSON
		jsonFirewallReq, err := json.MarshalIndent(firewallReq, "", "    ")
		if err != nil {
			log.Println(err)
		}
		log.Println(string(jsonFirewallReq))

		statusObject := status_page{Id: firewallReq.Id, FirewallJSON: string(jsonFirewallReq)}

		var firewallStatusTemplate = template.Must(template.New("statusbyid").ParseFiles("templates/base.html", "templates/requeststatus.html"))
		err = firewallStatusTemplate.ExecuteTemplate(w, "base", statusObject)
		if err != nil {
			log.Println(err)
		}
	} else {
		http.Redirect(w, r, "/status", 302)
	}

}

//Approval Handler
func Approve(w http.ResponseWriter, r *http.Request) {
	//var pass string

	//TODO - Call Firewall Search for Matching Queues
	netPendingQueue := firewall.GetFirewallRequestByQueue("networkPending")
	//log.Println(requests)
	//for request := range requests {
	//	requestMap := request.(bson.M)
	//	requestMap["_id"] = requestMap["_id"].(bson.ObjectId).Hex()
	//}
	//fwreqJson := firewall.FirewallStructToJson(netPendingQueue.Firewall_Requests[0])
	fwreqJson := netPendingQueue.Firewall_Queue.ToJson()
	//log.Printf("Firewall Req JSON:\n\n%s", fwreqJson)
	log.Printf("Firewall ToJson:\n\n%s", fwreqJson)

	var approveTemplate = template.Must(template.New("approve").ParseFiles("templates/base.html", "templates/approve.html"))

	//approveTemplate = approveTemplate.Funcs(functionMap)
	approveTemplate.ExecuteTemplate(w, "base", netPendingQueue)
}

//Blog Handler
func Blog(w http.ResponseWriter, r *http.Request) {
	var pass string
	var blogTemplate = template.Must(template.New("blog").ParseFiles("templates/base.html", "templates/blog.html"))
	blogTemplate.ExecuteTemplate(w, "base", pass)
}

//Setup Handler
func Setup(w http.ResponseWriter, r *http.Request) {
	var pass string
	var setupTemplate = template.Must(template.New("blog").ParseFiles("templates/base.html", "templates/setup.html"))
	//firewall.InitializeFirewallDB()
	setupTemplate.ExecuteTemplate(w, "base", pass)

}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	var pass string
	var notFoundTemplate = template.Must(template.New("notfound").ParseFiles("templates/404.html"))
	notFoundTemplate.ExecuteTemplate(w, "content", pass)
}
