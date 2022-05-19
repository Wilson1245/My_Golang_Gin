package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

var MgoConnect *mgo.Collection

func MD() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	MgoConnect = session.DB("GolangApi").C("TestGolangApi")
}
