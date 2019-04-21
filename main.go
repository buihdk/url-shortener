package main

import (
	"fmt"
	"log"
	"os"
)

const (
	ConfigPath = "config_data.yml"
)

var (
	Help = func() {
		fmt.Println(`
		 - append to the list: ./url-shortener configure -a dogs -u http://www.dogs.com 
		 - remove from the list: ./url-shortener -d dogs
		 - list redirections: ./url-shortener -l
		 - run http server on a given port: ./url-shortener run -p 8080
		 - help: ./url-shortener -h
		`)
	}
)

func main() {
	if len(os.Args) < 2 {
		Help()
		return
	}

	cfg := LoadConfigStore(ConfigPath)
	defer cfg.Save(ConfigPath)

	switch os.Args[1] {
	case "configure":
		log.Print("configuring")

		if os.Args[2] != "-a" || os.Args[4] != "-u" {
			Help()
			return
		}
		path, url := os.Args[3], os.Args[5]
		info := RedirectionInfo{path, url}

		cfg.Add(&info)
	case "-d":
		log.Print("deleting")

		path := os.Args[2]
		cfg.Remove(path)
	case "-l":
		log.Print("listing")

		for path, url := range *cfg {
			fmt.Printf("%s\t->\t%s\n", path, url)
		}
	case "run":
		log.Print("running server")

		if os.Args[2] != "-p" {
			Help()
			return
		}
		port := os.Args[3]

		StartServer(*cfg, port)
	default:
		Help()
	}
}
