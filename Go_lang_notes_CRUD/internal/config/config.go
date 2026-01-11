package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	MongoURI   string
	MongoDB    string
	ServerPORT string
}

func config() {

},

 func extractEnv (key string) (string  , error){
	 value := os.Getenv(key)

	  if value == ""{
			return  "",fmt.Errorf("Missing req env")
		}

		return  value , nil
 }


func Load() (config, error) {


	//  godotenv.load() reads .env and sets them into process env 
	//  os.getenv -> reads those values 

	if err :=  godotenv.Load() ; err != nil {
		return  config{},fmt.Errorf("Failed to load .env file")
	}

	// connect  and get mongodb uri 
	mongoURI , err := extractEnv("MONGO_URI")
	if err != nil {
		return  config{},err
	}

	// connect and get mongodb name 
	mongoDB , err := extractEnv("MONGO_DB_NAME")
	if err!= nil {
		return  config{}, err
	}
	
	// connect and get port 
	PORT , err := extractEnv("PORT")
	if err!= nil {
		return  config{}, err
	}

	return  config{MongoURI: mongoURI , MongoDB: mongoDB, ServerPORT: PORT}, nil
	

}