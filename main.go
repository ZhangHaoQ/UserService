package main

import (
	"UserService/Model"
	"UserService/config"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Panic(err)
	}
	err = Model.Init(cfg)
	if err != nil {
		log.Panic(err)
	}
}
