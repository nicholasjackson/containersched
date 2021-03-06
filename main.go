package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"log/syslog"
	"github.com/nicholasjackson/containersched/global"
	"github.com/nicholasjackson/containersched/handlers"

	"github.com/facebookgo/inject"
	"github.com/alexcesaro/statsd"
)

func main() {
	config := os.Args[1]
	rootfolder := os.Args[2]

	global.LoadConfig(config, rootfolder)

	setupInjection()
	setupHandlers()
}

func setupHandlers() {
	http.Handle("/", handlers.GetRouter())

	fmt.Println("Listening for connections on port", 8001)
	http.ListenAndServe(fmt.Sprintf(":%v", 8001), nil)
}

func setupInjection() {
	var g inject.Graph

	var err error

	statsdClient, err := statsd.New(statsd.Address(global.Config.StatsDServerIP)) // reference to a statsd client
	if err != nil {
		panic(fmt.Sprintln("Unable to create StatsD Client: ", err))
	}

	syslogWriter, err := syslog.Dial("udp", global.Config.SysLogIP, syslog.LOG_SYSLOG, "sorcery")
	if err != nil {
		panic(fmt.Sprintln("Unable to connect to syslog: ", err))
	}

	logWriter := log.New(syslogWriter, "containersched: ", log.Lshortfile)

	err = g.Provide(
		&inject.Object{Value: handlers.RouterDependencies},
		&inject.Object{Value: handlers.HealthDependencies},
		&inject.Object{Value: handlers.EchoDependencies},
		&inject.Object{Value: statsdClient, Name: "statsd"},
		&inject.Object{Value: logWriter},
	)

	if err != nil {
		fmt.Println(err)
	}

	// Here the Populate call is creating instances of NameAPI &
	// PlanetAPI, and setting the HTTPTransport on both to the
	// http.DefaultTransport provided above:
	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}
}
