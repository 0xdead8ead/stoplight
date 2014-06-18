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

//TODO - Convert Map to Struct & Vis Versa
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
	// Turns Out Go doesnt support Positive Lookbehind as it is
	//https://groups.google.com/d/msg/golang-nuts/7qgSDWPIh_E/OHTAm4wRZL0J
	//re := regexp.MustCompile(`(?<=\/status\/)([0-9a-f]+)`)
	//log.Println(re.FindStringSubmatch(fwidURI))
	fwUriSplit := strings.Split(fwidURI, "/")
	fwid := fwUriSplit[4]
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

	// TODO - Append Status Variable to fwreq with Current Queue

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

	//GetFirewallStatusByID(fwreqIDObject.Hex())
	//TODO - Add New fwreqid to queue references

	return string(fwreqIDObject.Hex())
}

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
