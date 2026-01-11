//  config -> db -> router -> run server

package main

import (
	"log"
	"notes-api/db"
	"notes-api/internal/config"
) 


 func main (){
	cfg , err:= config.Load()

	if err != nil {
		log.Fatalf("config err")
	}

	 client  , database , err := db.Connect(cfg)
	 
	  if err != nil {
			log.Fatalf("Db error")
		}


		defer func  ()  {
			if err:= db.Disconnect(client); err != nil{
				log.Println("mongo disconnect error"  , err)
			}
		}()


 }