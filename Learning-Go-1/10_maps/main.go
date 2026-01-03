package main

import (
	"fmt"
	"maps"
)

// -> maps
 func main(){

	// ! -> creating maps 

	// m:=make(map[string ]string)

	// ! -> Setting an element
	// m["name"]="Rashedul"
	// m["age"]="21"

	// ! -> Get an element


	// fmt.Println(m["name"]  , m["age"])

	//? IMP: if key dose't exists in the map then it returns zero value 


	m:=make(map[string]int)
	m["age"]=45
	m["price"]=50

	// fmt.Println(m["ages"])

	 fmt.Println(len(m))


	//!  -> delete element in array 

	// delete(m,"price")
	fmt.Println(m)



	//!  -> clear map 
	 clear(m)
	 fmt.Println(m)


	//  ! create manual map 

	 products:=map[string]int{"price" :45 ,"model":5 }

	 fmt.Println(products)

	 _,ok :=products["prices"]
	  if ok {
			fmt.Println("all oke")
		}else{
			fmt.Println("not key")

			
		}



		//  check map equal 

		 result1 := map[string]int{"price" :45 , "quntity":8 }
		 result2 := map[string]int{"price" :45 , "quntity":88 }


		 fmt.Println(maps.Equal(result1,result2))



 }