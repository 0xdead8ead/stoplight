package firewall

import (
	//"fmt"
	"labix.org/v2/mgo"
	"log"
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

//Saves Firewall Request to MongoDB
func SaveFirewall(fwreq map[string]interface{}) {
	log.Println("%s", fwreq)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	requestCollection := session.DB("firewallapp").C("fwreq")
	err = requestCollection.Insert(fwreq)
	if err != nil {
		panic(err)
	}
}
