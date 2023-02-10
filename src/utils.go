package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeSlice []string

type ErrorResponse struct{
		Status string
		Desc string
}

type Params struct{
	Period string
	TimeZone string
	StartDate time.Time
	EndDate time.Time
}

const (
	hourly string = "h"
	daily string = "d"
	monthly string = "mo"
	yearly string = "y"
)

func TimestampToDate(t string) (time.Time,error){
	var (
		y,d int
		M time.Month
		h,m,s int		
	)
	_,err := fmt.Sscanf(t,"%4v%2v%2vT%2v%2v%2vZ",&y,&M,&d,&h,&m,&s)
	return time.Date(y, M, d, h, m, s,0 , time.UTC),err
}

func WriteJSON(w http.ResponseWriter, code int, response any) {
	w.WriteHeader(code)
	resp, err := json.Marshal(response)
	if err!= nil{
		panic(err)
	}
	w.Write(resp)
}