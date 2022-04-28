package main

import (
	"log"
	apiserver "github.com/SilvioJPEG/redflags-restapi/internal/app"
)

type Hand struct {
	ID    string `json:"uid"`
	Perk1 string `json:"perk1"`
	Perk2 string `json:"perk2"`
	Flag3 string `json:"perk3"`
}

func main() {
	config := apiserver.NewConfig()
	if err := apiserver.Initialize(config); err != nil {
		log.Fatal(err)
	}
}
