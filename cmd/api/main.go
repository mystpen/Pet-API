package main

import (
	"log"

	"github.com/mystpen/Pet-API/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err, nil)
	}

	// Connect to DB
	db, err := openDB(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Start api
	err = Start(cfg)
	if err != nil{
		log.Fatal(err)
	}
}
