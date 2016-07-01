package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/shinji62/go-fleet-service-locator/dataProvider"
	httpServ "github.com/shinji62/go-fleet-service-locator/http"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"strconv"
)

var (
	port      = kingpin.Flag("port", "Port to listen").Envar("PORT").Short('p').Required().String()
	serverPem = kingpin.Flag("cert-pem", "Certificate Pem file").Envar("CERT_PEM").Short('s').String()
	serverKey = kingpin.Flag("cert-key", "Certificate Key file").Envar("CERT_KEY").Short('k').String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	data, err := dataProvider.NewMongoDbProvider("HEREMONGOCONNECTION")
	//location, err := data.GetLocation(139.77, 35.69)
	if err != nil {
		log.Fatal("Err:", err)
	}
	defer data.Close()
	router := httprouter.New()
	router.GET("/locationService/:long/:lat", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		long, err := strconv.ParseFloat(ps.ByName("long"), 64)
		lat, err := strconv.ParseFloat(ps.ByName("lat"), 64)
		location, err := data.GetLocation(long, lat)
		if err != nil {
			log.Fatal("Err:", err)
		}

		w.Header().Set("Content-Type", "application/hal+json;charset=UTF-8")
		// Encode
		location.Type = "service"
		location.Latitude = location.Location.Y
		location.Longitude = location.Location.X
		if err := ffjson.NewEncoder(w).Encode(location); err != nil {
			panic(err)
		}
	})
	httpServer := httpServ.NewHttpRouterService(router, *port)
	httpServer.SetTLS(*serverPem, *serverKey)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
