package main

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func String2Int64(str string) int64 {
	iValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return int64(0) 
	}

	return iValue 
}

func main() {

	tm := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)
	tz, _ := time.LoadLocation("Pacific/Samoa")
	converted_tm := tm.In(tz)
	log.Info(converted_tm)

	log.Info(converted_tm.Day())   // 28
}
