//  config -> db -> router -> run server

package main

import (
	"fmt"
	"log"
	"notes-api/db"
	"notes-api/internal/config"
	"notes-api/internal/server"
) 


 func main (){
	cfg , err:= config.Load()

	if err != nil {
		log.Fatalf("config err")
	}

	 client  , _ , err := db.Connect(cfg)
	 
	  if err != nil {
			log.Fatalf("Db error")
		}


		defer func  ()  {
			if err:= db.Disconnect(client); err != nil{
				log.Println("mongo disconnect error"  , err)
			}
		}()

		route := server.NewRouter()
		
		address:= fmt.Sprintf(":%s",cfg.ServerPORT)

		 if err := route.Run( address); err != nil{
			log.Fatalf("server failed")
		 }

 }