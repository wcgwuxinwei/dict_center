package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
	"github.com/wcgwuxinwei/dict_center/conf"
	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

func init() {
	if conf.MgoAddr == "" {
		log.Fatalf("Mongo address is empty")
	}
	if conf.MgoPort == "" {
		log.Warnf("Mongo port does not setted, using default port (27017)")
		conf.MgoPort = "27017"
	}
	var err error
	if mgoSession, err = mgo.Dial(conf.MgoAddr + ":" + conf.MgoPort); err != nil {
		log.Fatalf("Mongo Session init failed: %s", err)
	}
	defer mgoSession.Close()
}

func main() {
	log.Warnf("Dict Center Started")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// endpoint
	e.GET("/search/:word", nil)
	e.GET("/sync/all", nil)
}
