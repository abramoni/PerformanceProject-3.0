package main

import (
	"fmt"
	"presentation/mypackages"
)


func main(){
	fmt.Println("Collecting the metrics please wait...")
	jsondata := mypackages.GenerateJson()
	fmt.Println()
	mypackages.StoreInElasticDb(jsondata,"longevitycluster")
	// var time_stamp int64
	// time_stamp = 1655717048
	//mypackages.RetriveFronElasticDb(time_stamp)
}