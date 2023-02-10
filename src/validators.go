package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"github.com/rs/zerolog/log"
)

func GetParams(r *http.Request) (*Params,error) {
	period := r.URL.Query().Get("period")
	tz := r.URL.Query().Get("tz")
	t1 := r.URL.Query().Get("t1")
	t2 := r.URL.Query().Get("t2")

	//check params
	periodRe,err := regexp.MatchString(`\d*(h|d|mo|y)`,period)
	
	if err!= nil || !periodRe {
		msg := "period bad format"
		log.Error().Err(err).Msg(msg)
		return nil,fmt.Errorf(msg)
	}

	tzRe,err := regexp.MatchString(`[A-Z][a-z]+\/[A-Z][a-z]+`,tz)
	if err!= nil || !tzRe {
		msg := "tz bad format"
		log.Error().Err(err).Msg(msg)
		return nil,fmt.Errorf(msg)
	}

	startDate, err1 := TimestampToDate(t1)
	if err1!= nil {
		msg := "t1 bad format"
		log.Error().Err(err).Msg(msg)
		return nil,fmt.Errorf(msg)
	}

	endDate, err2 := TimestampToDate(t2)
	if err2!= nil {
		msg := "t2 bad format"
		log.Error().Err(err).Msg(msg)
		return nil,fmt.Errorf(msg)
	}
	if startDate.After(endDate) {
		msg := "end time is before start time"
		log.Error().Err(err).Msg(msg)
		return nil,fmt.Errorf(msg)
	}

	return &Params{
		Period:    period,
		TimeZone:  tz,
		StartDate: startDate,
		EndDate:   endDate,
	},nil
}

func ValidatePeriod(params *Params) (int, string, error) {

	periodNumExp := regexp.MustCompile("[0-9]+")
	numOfPeriod := periodNumExp.FindString(params.Period)
	periodExp := regexp.MustCompile("[a-z]+")
	kindOfPeriod :=  periodExp.FindString(params.Period)

	perNum, err := strconv.Atoi(numOfPeriod)
	if err != nil{
		return -1,"",err
	}
	return perNum,kindOfPeriod,nil
}