package http

import (
	"crypto/tls"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/httpdown"
	"net/http"
	"time"
)

type HttpLocationService struct {
	HttpServer *http.Server
	HttpDown   *httpdown.HTTP
}

func NewHttpRouterService(router http.Handler, port string) *HttpLocationService {
	portFinal := ":" + fmt.Sprintf("%v", port)
	return &HttpLocationService{
		HttpServer: &http.Server{
			Addr:    portFinal,
			Handler: router,
		},
		HttpDown: &httpdown.HTTP{
			StopTimeout: 10 * time.Second,
			KillTimeout: 1 * time.Second,
		},
	}
}

/*
Set TLS config
If validate are not valid or could not be read TLS Config will be empty
*/
func (rs *HttpLocationService) SetTLS(pemFile string, keyFile string) {
	cer, err := tls.LoadX509KeyPair(pemFile, keyFile)
	if err != nil {
		log.Println(err)
	} else {
		rs.HttpServer.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cer}}

	}

}

func (rs *HttpLocationService) ListenAndServe() error {
	return httpdown.ListenAndServe(rs.HttpServer, rs.HttpDown)

}
