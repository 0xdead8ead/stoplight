package firewall

import (
	//"fmt"
	"encoding/hex"
	"encoding/json"
	"labix.org/v2/mgo"
	"log"
	"strings"
	//"reflect"
	"code.google.com/p/rsc/qr"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	//"regexp"
)

type Firewall_Requests []Firewall_Request

type Firewall_Rules []Firewall_Rule

//TODO - Convert Map to Struct & Vis Versa
type Firewall_Request struct {
	Id                string
	Requestor         string
	Service_Date      string
	Service_Term_Date string
	Leader            string
	Business_Approver string
	Service_Number    string
	Status            string
	Firewall_Rules    Firewall_Rules
}

type Firewall_Rule struct {
	Source_Zone      string
	Source_Ips       string
	Dest_Zone        string
	Data_Type        string
	Network_Protocol string
	Dest_Ips         string
	Dest_Ports       string
	App_Name         string
	Server_Location  string
}

type Approval_Queue struct {
	Queue_Name     string
	Firewall_Queue Firewall_Requests
}

//Maps Request Data into a Struct
func FirewallMapToStruct(fwreq interface{}) *Firewall_Request {
	fwreqData := fwreq.(bson.M)
	fwreqObject := new(Firewall_Request)

	// Populate Firewall_Request Struct
	fwreqObject.Id = fwreqData["_id"].(bson.ObjectId).Hex()
	fwreqObject.Requestor = fwreqData["requestor"].(string)
	fwreqObject.Service_Date = fwreqData["service_date"].(string)
	fwreqObject.Service_Term_Date = fwreqData["service_term_date"].(string)
	fwreqObject.Leader = fwreqData["leader"].(string)
	fwreqObject.Business_Approver = fwreqData["business_approver"].(string)
	fwreqObject.Status = fwreqData["status"].(string)

	//TODO - I'd like to see this field populated via API
	fwreqObject.Service_Number = fwreqData["service_number"].(string)

	// Iterate through Rules And load into appropriate struct fields
	fwrules := fwreqData["rules"].([]interface{})
	for rule := range fwrules {
		fwreqRuleData := fwrules[rule].(bson.M)
		fwreqRuleObject := new(Firewall_Rule)
		fwreqRuleObject.Source_Zone = fwreqRuleData["source_zone"].(string)
		fwreqRuleObject.Source_Ips = fwreqRuleData["source_ips"].(string)
		fwreqRuleObject.Dest_Zone = fwreqRuleData["dest_zone"].(string)
		fwreqRuleObject.Data_Type = fwreqRuleData["data_type"].(string)
		fwreqRuleObject.Network_Protocol = fwreqRuleData["network_protocol"].(string)
		fwreqRuleObject.Dest_Ips = fwreqRuleData["dest_ips"].(string)
		fwreqRuleObject.Dest_Ports = fwreqRuleData["dest_ports"].(string)
		fwreqRuleObject.App_Name = fwreqRuleData["app_name"].(string)
		fwreqRuleObject.Server_Location = fwreqRuleData["server_loc"].(string)
		//log.Println(fwreqRuleData)

		//Append Firewall_Rule to Firewall_Rules array in Firewall_Request struct
		fwreqObject.Firewall_Rules = append(fwreqObject.Firewall_Rules, *fwreqRuleObject)
	}
	//log.Printf("Firewall Object Struct:\n\n%s", fwreqObject)
	return fwreqObject
}

func ReqToJson(fwreq Firewall_Request) string {
	return fwreq.ToJson()
}

func (fwreq Firewall_Request) ToJson() string {
	jsonFwReq, err := json.MarshalIndent(fwreq, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonFwReq)
}

func (fwreq Firewall_Requests) ToJson() string {
	jsonFwReq, err := json.MarshalIndent(fwreq, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonFwReq)
}

func (fwRule Firewall_Rule) ToJson() string {
	jsonRule, err := json.MarshalIndent(fwRule, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonRule)
}

func (fwRule Firewall_Rules) ToJson() string {
	jsonRule, err := json.MarshalIndent(fwRule, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonRule)
}

/*
func (fwRule []Firewall_Rule) ToJson() string {
	jsonRule, err := json.MarshalIndent(fwRule, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonRule)
}
*/

//Generates QR-Code based on Status URI and places in the images/qrcodes directory
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

//Saves Firewall Request to MongoDB
func UpdateFirewallStatus(id string, statusUpdate string) {
	//var pass string
	log.Printf("Updating status to:\t%s for firewall:\t%s\n\n", statusUpdate, id)
	// TODO - Append Status Variable to fwreq with Current Queue
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	requestCollection := session.DB("firewallapp").C("fwreq")
	hexId, hexErr := hex.DecodeString(id)
	if hexErr != nil {
		log.Printf("\n\nERROR:\t%s\n\n", hexErr)
	}
	updateId := bson.M{"_id": bson.ObjectId(hexId)}
	update := bson.M{"$set": bson.M{"status": statusUpdate}}
	updateErr := requestCollection.Update(updateId, update)
	if updateErr != nil {
		log.Printf("Can't update document %v\n", updateErr)
	}
}

// Get Firewall Status by ID
func GetFirewallStatusByID(id string) *Firewall_Request {
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
	firewallRequest := FirewallMapToStruct(result)
	//log.Printf("\n\n%s\n\n", result)
	//log.Printf("FIREWALL STRUCT:\t%s\n\n", firewallRequest)
	return firewallRequest

}

//Retreives Firewall Request Based on Queue Name
func GetFirewallRequestByQueue(queueName string) *Approval_Queue {
	log.Printf("Retreiving Firewall Request by QUEUE:\t%s", queueName)
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	requestCollection := session.DB("firewallapp").C("fwreq")
	var result []interface{}
	err = requestCollection.Find(bson.M{"status": queueName}).All(&result)
	if err != nil {
		log.Println(err)
	}
	log.Printf("\n\nFUCKING SHITSTICKS%s\n\n", result)

	var queueList Firewall_Requests
	for i := range result {
		firewallRequest := FirewallMapToStruct(result[i])
		queueList = append(queueList, *firewallRequest)
	}
	firewallQueue := new(Approval_Queue)
	firewallQueue.Queue_Name = queueName
	firewallQueue.Firewall_Queue = queueList
	//log.Println(queueList)

	log.Printf("Firewall Queue:\t%s", firewallQueue)

	return firewallQueue
}
