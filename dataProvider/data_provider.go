package dataProvider

import (
	"gopkg.in/mgo.v2/bson"
)

const LocationCollection = "serviceLocation"

type DataProvider interface {
	GetLocation(float64, float64) ServiceLocation
	Close()
}

type ServiceLocation struct {
	iD        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	class     string        `bson:"_class,omitempty" json:"_class"`
	Address1  string        `bson:"address1" json:"address1"`
	City      string        `bson:"city" json:"city"`
	Location  Geo           `bson:"location" json:"-"`
	State     string        `bson:"state" json:"state"`
	Zip       string        `bson:"zip" json:"zip"`
	Type      string        `bson:"type" json:"type"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
}

type Geo struct {
	X float64 `bson:"x" json:"x"`
	Y float64 `bson:"y" json:"y"`
}
