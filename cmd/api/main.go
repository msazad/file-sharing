package main

import (
	"file-sharing/pkg/config"
	"file-sharing/pkg/di"
	"log"
)


func main(){
	config,configErr:=config.LoadConfig()

	if configErr !=nil{
		log.Fatal("cannot load config",configErr)
	}

	server,err:=di.InitializeAPI(config)
	if err!=nil{
		log.Fatal("couldnt start server")
	}else{
		server.Start()
	}
}