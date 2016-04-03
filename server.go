package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"os"
)

func main() {
	port := flag.String("port", "2222", "port on which server will be start")
	folder := flag.String("folder", "", "folder to serve")

	flag.Parse()

	if *folder == "" {
		log.Fatal("Please specify folder to serve using -folder flag")
	}

	if _, err := os.Stat(*folder); os.IsNotExist(err) {
		log.Fatal("Specified directiry to serve is not exists")
	}

	r := makeRouter(*folder)
	r.Run(":" + *port)
}
