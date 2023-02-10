package main

import "time"

type IDateAdditions interface{
	dateAddition(int,time.Time)time.Time
	dateDifferenceLessThanPeriod(start,end time.Time,numPeriod int) bool
}
type AddYears struct{}
func (y AddYears) dateAddition(numOfPeriod int,timeIter time.Time) time.Time {
	return timeIter.AddDate(numOfPeriod,0,0)
}
func (y AddYears) dateDifferenceLessThanPeriod(start, end time.Time,numOfPeriod int) bool{
	return end.Year() - start.Year() < numOfPeriod
}
type AddMonths struct{}
func (m AddMonths) dateAddition(numOfPeriod int,timeIter time.Time) time.Time {
	return timeIter.AddDate(0,numOfPeriod,0)
}
func (m AddMonths) dateDifferenceLessThanPeriod(start, end time.Time,numOfPeriod int) bool{
	if start.Year()==end.Year(){
		return int(end.Month()) - int(start.Month()) < numOfPeriod
	}
	return false
}
type AddDays struct{}
func (d AddDays) dateAddition(numOfPeriod int,timeIter time.Time) time.Time {
	return timeIter.AddDate(0,0,numOfPeriod)
}
func (d AddDays) dateDifferenceLessThanPeriod(start, end time.Time,numOfPeriod int) bool{
	if start.Year()==end.Year() && start.Month() == end.Month() {
		return end.Day() - start.Day() < numOfPeriod
	}
	return false
}
type AddHours struct{}
func (h AddHours) dateAddition(numOfPeriod int,timeIter time.Time) time.Time {
	return  timeIter.Add(time.Hour*time.Duration(numOfPeriod))
}
func (h AddHours) dateDifferenceLessThanPeriod(start, end time.Time,numOfPeriod int) bool{
	if start.Year()==end.Year() && start.Month() == end.Month() && start.Day() == end.Day() {
		return end.Hour() - start.Hour() < numOfPeriod
	}
	return false
}

func truncateDate(t time.Time)time.Time {
	return time.Date(t.Year(),t.Month(),t.Day(),t.Hour(),0,0,0,t.Location())
}
// ComputeDatesList computes the list of the dates
func ComputeDatesList(numOfPeriod int, kindOfPeriod string, start,end time.Time) TimeSlice {
	
	switch kindOfPeriod {
	case hourly:
		hours := AddHours{}
		return ComputeDates(hours,numOfPeriod,start,end)
	case daily:
		days := AddDays{}
		return ComputeDates(days,numOfPeriod,start,end)
	case monthly:
		months := AddMonths{}
		return ComputeDates(months,numOfPeriod,start,end)
	case yearly:
		years := AddYears{}
		return ComputeDates(years,numOfPeriod,start,end)
	default:
	}
	return nil
}

// ComputeDates computes the dates in relation to the given request parameter (y,mo,d)
func ComputeDates(f IDateAdditions,numOfPeriod int, start, end time.Time) TimeSlice {
	dateSlice := make(TimeSlice,0)
	timeIter := start
	if f.dateDifferenceLessThanPeriod(start,end,numOfPeriod) {
		return TimeSlice{}
	}
	for {
		timeIter = f.dateAddition(numOfPeriod,timeIter)
		if(timeIter.After(end)){
			break
		}
		timeIter := truncateDate(timeIter)
		dateSlice = append(dateSlice, timeIter.Format("20060102T150405Z"))
	}
	return dateSlice
}