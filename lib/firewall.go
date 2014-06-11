package firewall

import (
	//"fmt"
	"labix.org/v2/mgo"
)

func SaveFirewall(fwreq map[string]interface{}) {
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
