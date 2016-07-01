package dataProvider

import (
	"gopkg.in/mgo.v2"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type MongoDbProvider struct {
	Session *mgo.Session
}

func NewMongoDbProvider(url string) (*MongoDbProvider, error) {
	mgo.SetDebug(true)

	sess, err := mgo.DialWithTimeout(url, 2*time.Second)
	if err != nil {
		log.Fatalf("Could not connect to database:", err)
		return nil, err
	}
	return &MongoDbProvider{
		Session: sess,
	}, nil
}

func (m *MongoDbProvider) GetLocation(long float64, lat float64) (ServiceLocation, error) {
	var result ServiceLocation
	var err error
	c := m.Session.DB("").C(LocationCollection)
	q := c.Find(bson.M{
		"location": bson.M{
			"$near": []float64{long, lat},
		},
	})
	err = q.One(&result)
	return result, err
}

func (m *MongoDbProvider) Close() {
	m.Session.Close()
}
