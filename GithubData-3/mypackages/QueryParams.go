package mypackages

import (
	"strconv"
	"time"
)

func TimeStampGneerator(timeperiod string, unit string) (startTime string, endTime string) {

	tNow := time.Now()
	tNow_minute := tNow.Minute()
	tNow_hour := tNow.Hour()
	tNow_second := tNow.Second()
	unix_time_in_decimal := 	tNow_second + 60*tNow_minute + 3600*tNow_hour
	unix_time_now := tNow.Unix()
	unix_time_decimal_int64 := int64(unix_time_in_decimal)
	unix_end_time := unix_time_now - unix_time_decimal_int64
	unix_start_time := unix_end_time - 86400
	
	if timeperiod == "past4hrs"{
		unix_end_time= tNow.Unix()
		unix_start_time = unix_end_time - 4*3600
	}

	if timeperiod == "day"{
		
	}
	if timeperiod == "week" {
		unix_start_time = unix_start_time - 6*86400
	}
	if timeperiod == "month"{
		unix_start_time = unix_start_time - 29*86400
	}
	if unit == "M"{
		unix_start_time = unix_start_time * 1000
		unix_end_time = unix_end_time * 1000
	}
	if unit == "U"{
		unix_start_time = unix_start_time * 1000000
		unix_end_time = unix_end_time * 1000000
	}
	s := strconv.FormatInt(unix_start_time, 10)
	e := strconv.FormatInt(unix_end_time, 10)
	return s, e
}

func PresentTimeStamp ()(timestamp_now int64){
	tNow := time.Now()
	timestamp_now = tNow.Unix()
	return
}
