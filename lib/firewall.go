package firewall

import (
	//"fmt"
	//"encoding/json"
	"labix.org/v2/mgo"
	"log"
	"strings"
	//"reflect"
	"code.google.com/p/rsc/qr"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	//"regexp"
)

type firewall_request struct {
	requestor         string
	service_date      string
	service_term_date string
	leader            string
	business_approver string
	service_number    string
	firewall_rules    []firewall_rule
}

type firewall_rule struct {
	source_zone      string
	source_ips       string
	dest_zone        string
	data_type        string
	network_protocol string
	dest_ips         string
	dest_ports       string
	app_name         string
	server_loc       string
}

type approval_queue struct {
	queue_name        string
	firewall_requests []string
	//firewall_requests []firewall_request
}

func GenStatusQRCode(fwidURI string) {
	//re := regexp.MustCompile(`(?<=\/status\/)([0-9a-f]+)`)
	//log.Println(re.FindStringSubmatch(fwidURI))

	fwUriSplit := strings.Split(fwidURI, "/")
	fwid := fwUriSplit[2]
	log.Println(fwUriSplit)
	log.Println(fwid)
	c, err := qr.Encode(fwidURI, qr.L)
	if err != nil {
		log.Println(err)
	} else {
		pngdat := c.PNG()
		s := []string{"static/images/qrcodes/", fwid, ".png"}
		ioutil.WriteFile(strings.Join(s, ""), pngdat, 0666)
	}
}

//Saves Firewall Request to MongoDB
func SaveFirewall(fwreq map[string]interface{}) string {
	log.Println("%s", fwreq)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	requestCollection := session.DB("firewallapp").C("fwreq")
	//err = requestCollection.Insert(fwreq)
	info, inserterr := requestCollection.Upsert(fwreq, fwreq)
	if inserterr != nil {
		panic(inserterr)
	}
	fwreqIDObject := info.UpsertedId.(bson.ObjectId)
	log.Printf("%s", fwreqIDObject.Hex())
	log.Println(fwreqIDObject)

	GetFirewallStatusByID(fwreqIDObject.Hex())

	return string(fwreqIDObject.Hex())
}

/*
//Initializes Firewall Database Approval Queues
func InitializeFirewallDB() {
	log.Println("Initializing Firewall Approval Queues")

	//TODO - Creates Approval Queues
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Create Approval Queues Collection
	approvalQueueCollection := session.DB("firewallapp").C("approvalqueues")

	//Create Approval Queues
	securityApprovalQueue := approval_queue{queue_name: "SecurityApproval", firewall_requests: []string{"hello", "world"}}
	networkingApprovalQueue := approval_queue{queue_name: "NetworkingApproval"}
	implementationQueue := approval_queue{queue_name: "implementationQueue"}
}
*/

// Get Firewall Status by ID
func GetFirewallStatusByID(id string) interface{} {
	log.Printf("Retreiving Firewall Request by ID:\t%s", id)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	requestCollection := session.DB("firewallapp").C("fwreq")
	var result interface{}
	err = requestCollection.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		panic(err)
	}
	log.Printf("\n\n%s\n\n", result)
	return result

}
