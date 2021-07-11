package main 

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/globalsign/mgo"
)

type timeZoneConvertion struct {
	TimeZone string `bson:"timeZone" json:"timeZone"`
	TimeDifference string `bson:"TimeDifference" json:"TimeDifference"`
	Name string `bson:"name" json:"name"`
}

type tzs []timeZoneConvertion

func main() {

	// setting up client and connection string
	client, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Couldn't connect to Mongo")
	}
	
	// setting up the database
	db := client.DB("timezones")
	err = db.DropDatabase()
	if err != nil {
		log.Fatal("Couldn't drop MongoDB database. Err: " + err.Error())
	}

	// setting up the collection
	coll := db.C("timezones")
	// reading file 
	data, err := ioutil.ReadFile("timezones.json")
	if err != nil {
		log.Fatal("Couldn't open file")
	}

	// timezones
	var timeZones tzs 
	err = json.Unmarshal(data, &timeZones)
	if err != nil {
		log.Fatal("Couldn't unmarshall JSON")
	}

	// inserting records
	for _, v := range timeZones {
		coll.Insert(&v)
	}
}