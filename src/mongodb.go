package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/globalsign/mgo"
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Username string
	Password string
}

const (
	integrationName    = "com.newrelic.mongodb"
	integrationVersion = "0.1.0"
)

var (
	args argumentList
)

func main() {

	mongoIntegration, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error("Failed to create integration")
		os.Exit(1)
	}

	fmt.Println(mongoIntegration)

	roots := x509.NewCertPool()

	ca, err := ioutil.ReadFile("/Users/dantewelch/bluemedora/blue_medora.crt")
	if err != nil {
		log.Error("Failed to open crt file")
	}

	roots.AppendCertsFromPEM(ca)

	tlsConfig := &tls.Config{}
	tlsConfig.RootCAs = roots

	dialInfo := mgo.DialInfo{
		Addrs:    []string{"mdb-rh7-rs1-r2.bluemedora.localnet"},
		Username: "admin",
		Password: "password",
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		if err != nil {
			log.Error("%s", err)
		}
		return conn, err
	}

	session, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		panic(err)
	}

	var ss serverStatus
	err = session.DB("admin").Run(map[interface{}]interface{}{"serverStatus": 1}, &ss)
	if err != nil {
		log.Error("%s", err)
	}
	fmt.Printf("%+v", ss)

}
totalCreated
total_created
