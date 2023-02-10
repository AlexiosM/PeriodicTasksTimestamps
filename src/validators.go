package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

// GetParams reads and validates query input parameters
func GetParams(r *http.Request) (*Params, error) {
	period := r.URL.Query().Get("period")
	tz := r.URL.Query().Get("tz")
	t1 := r.URL.Query().Get("t1")
	t2 := r.URL.Query().Get("t2")

	periodRe, err := regexp.MatchString(`\d*(h|d|mo|y)`, period)
	if err != nil || !periodRe {
		msg := "period bad format"
		log.Error().Err(err).Msg(msg)
		return nil, fmt.Errorf(msg)
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		msg := "tz bad format"
		log.Error().Err(err).Msg(msg)
		return nil, fmt.Errorf(msg)
	}

	startDate, err1 := TimestampToDate(t1)
	if err1 != nil {
		msg := "t1 bad format"
		log.Error().Err(err).Msg(msg)
		return nil, fmt.Errorf(msg)
	}

	endDate, err2 := TimestampToDate(t2)
	if err2 != nil {
		msg := "t2 bad format"
		log.Error().Err(err).Msg(msg)
		return nil, fmt.Errorf(msg)
	}
	if startDate.After(endDate) {
		msg := "end time is before start time"
		log.Error().Err(err).Msg(msg)
		return nil, fmt.Errorf(msg)
	}

	return &Params{
		Period:    period,
		TimeZone:  loc,
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}

// ValidatePeriod gets the period number and chosen period kind (h,d,mo,y)
func ValidatePeriod(params *Params) (int, string, error) {

	periodNumExp := regexp.MustCompile("[0-9]+")
	numOfPeriod := periodNumExp.FindString(params.Period)
	periodExp := regexp.MustCompile("[a-z]+")
	kindOfPeriod := periodExp.FindString(params.Period)

	perNum, err := strconv.Atoi(numOfPeriod)
	if err != nil {
		return -1, "", err
	}
	return perNum, kindOfPeriod, nil
}
