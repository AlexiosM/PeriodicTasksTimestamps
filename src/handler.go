package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
) 

func PeriodicTaskHandler(w http.ResponseWriter, r *http.Request) {
	params,err := GetParams(r)
	if err!= nil{
		log.Error().Err(err).Msg("input validation error")
		WriteJSON(w,http.StatusBadRequest,ErrorResponse{
			Status: "error",
			Desc: "Unsupported period",
		})
		return
	}
	periodNum,periodKind,listErr := ValidatePeriod(params)
	if listErr != nil {
		log.Error().Err(listErr).Msg("period given in wrong format")
		WriteJSON(w,http.StatusBadRequest,ErrorResponse{
			Status: "error",
			Desc: "Unsupported period",
		})
		return
	}
	listResponse := ComputeDatesList(periodNum,periodKind,params.StartDate,params.EndDate)
	WriteJSON(w,http.StatusOK,listResponse)
}
