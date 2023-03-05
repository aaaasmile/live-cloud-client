package main

import (
	"flag"
	"fmt"
	"live-cloud-client/conf"
	"log"
	"os"
)

func main() {
	//var list = flag.String("list", "/", "List files on server on the specific path")
	var configfile = flag.String("config", "config.toml", "Configuration file path")
	var relpath = flag.Bool("relpath", false, "Use relative path. Used it in dev mode or when the exe is called in the same folder as the key")
	var ver = flag.Bool("version", false, "Print the current version")
	flag.Parse()

	if *ver {
		fmt.Printf("%s, version: %s", conf.Appname, conf.Buildnr)
		os.Exit(0)
	}

	_, err := conf.ReadConfig(*configfile, *relpath)
	if err != nil {
		log.Fatal("Config file error: ", err)
	}

	log.Println("That's all folks!")
}
