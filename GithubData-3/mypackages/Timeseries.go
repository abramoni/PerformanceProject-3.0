package mypackages

import (
	"encoding/json"
	"log"
	"math"
	"net/url"
)

type metric struct {
	Name    string `json:"metricName"`
	Datavec []rate `json:"dataPointVec"`
}

type rate struct {
	Time  int64     `json:"timestampMsecs"`
	Value ratevalue `json:"data"`
}

type ratevalue struct {
	Data int64 `json:"int64Value"`
}

func Filecreaterate() (average_file_create_rate int64) {

	metricName := "kCreateFileOps"
	schemaName := "kBridgeViewPerfStats"
	metricUnitType := "5"
	rollupFunction := "rate"
	entityId := "12522945"
	timePeriod := "day"
	startTime_Msecs, endtime_Msecs := TimeStampGneerator(timePeriod, "M")
	rollupIntervalSecs := "180"
	average_file_create_rate = AverageValueOfTimeSeries(endtime_Msecs, entityId, metricName, metricUnitType, timePeriod, rollupFunction, rollupIntervalSecs, schemaName, startTime_Msecs)
	return
}

func UtilizationChangeRate() (average_utilization_change_rate_converted float64) {

	metricName := "kSystemUsageBytes"
	schemaName := "kBridgeClusterStats"
	metricUnitType := "0"
	rollupFunction := "rate"
	entityId := "2790138600742128"
	timePeriod := "day"
	startTime_Msecs, endtime_Msecs := TimeStampGneerator(timePeriod, "M")
	rollupIntervalSecs := "180"
	average_utilization_change_rate := AverageValueOfTimeSeries(endtime_Msecs, entityId, metricName, metricUnitType, timePeriod, rollupFunction, rollupIntervalSecs, schemaName, startTime_Msecs)
	average_utilization_change_rate_converted = ConvertToUnits(average_utilization_change_rate)
	return
}

func GarbageCollection() (average_garbage_collection_converted float64) {

	metricName := "EstimatedGarbageBytes"
	schemaName := "ApolloV2ClusterStats"
	metricUnitType := "0"
	rollupFunction := "average"
	var entityId string
	entityId = "st-longevity+(ID+2790138600742128)"
	timePeriod := "day"
	startTime_Msecs, endtime_Msecs := TimeStampGneerator(timePeriod, "M")
	rollupIntervalSecs := "180"
	average_garbage_collection := AverageValueOfTimeSeries(endtime_Msecs, entityId, metricName, metricUnitType, timePeriod, rollupFunction, rollupIntervalSecs, schemaName, startTime_Msecs)
	average_garbage_collection_converted = ConvertToUnits(average_garbage_collection)
	return
}

func AverageValueOfTimeSeries(endtime_Msecs, entityId, metricName, metricUnitType, timePeriod, rollupFunction, rollupIntervalSecs, schemaName, startTime_Msecs string) (average int64) {

	newURL := GenerateNewURLforTimeSeries(endtime_Msecs, entityId, metricName, metricUnitType, timePeriod, rollupFunction, rollupIntervalSecs, schemaName, startTime_Msecs)
	path, err := url.PathUnescape(newURL)
	if err != nil {
		log.Fatal(err)
	}
	response := PostRequestForAccessToken()
	data := GetRequestForJsonData(response, path)

	var responsedata metric

	json.Unmarshal(data, &responsedata)
	instances_of_data := len(responsedata.Datavec)
	
	for i := 0; i < len(responsedata.Datavec); i++ {
		average += responsedata.Datavec[i].Value.Data
	}
	if instances_of_data > 0 {
		average /= int64(instances_of_data)
	} else {
		average = 0
	}
	return
}

func ConvertToUnits(average int64)(converted_avg float64){
	mega_byte := 1024 * 1024 * 1024
	converted_avg = float64(average) / float64(mega_byte)
	converted_avg = math.Round(converted_avg*100)/100
	return
}