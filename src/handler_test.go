package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPeriodicTaskHandlerYearsSuccess(t *testing.T) {
	expectedResult := TimeSlice{"20210714T200000Z", "20220714T200000Z"}
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=1y&tz=Europe/Athens&t1=20200714T204603Z&t2=20220715T204603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	var d TimeSlice
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("return code should be 200 OK, got %v", res.StatusCode)
	}
	if !reflect.DeepEqual(d, expectedResult) {
		t.Errorf("time lists not equal, got\nresp: %v,\nexpected: %v", d, expectedResult)
	}
}

func TestPeriodicTaskHandlerDaysSuccess(t *testing.T) {
	expectedResult := TimeSlice{"20201022T200000Z", "20210130T200000Z", "20210510T200000Z"}
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=100d&tz=Europe/Athens&t1=20200714T204603Z&t2=20210715T204603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	var d TimeSlice
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("return code should be 200 OK, got %v", res.StatusCode)
	}
	if !reflect.DeepEqual(d, expectedResult) {
		t.Errorf("time lists not equal, got\nresp: %v,\nexpected: %v", d, expectedResult)
	}
}

func TestPeriodicTaskHandlerMonthsSuccess(t *testing.T) {
	expectedResult := TimeSlice{"20210414T200000Z", "20220114T200000Z"}
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=9mo&tz=Europe/Athens&t1=20200714T204603Z&t2=20220715T204603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	var d TimeSlice
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("return code should be 200 OK, got %v", res.StatusCode)
	}
	if !reflect.DeepEqual(d, expectedResult) {
		t.Errorf("time lists not equal, got\nresp: %v,\nexpected: %v", d, expectedResult)
	}
}
func TestPeriodicTaskHandlerHoursSuccess(t *testing.T) {
	expectedResult := TimeSlice{"20200715T040000Z", "20200715T120000Z", "20200715T200000Z"}
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=8h&tz=Europe/Athens&t1=20200714T204603Z&t2=20200715T204603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	var d TimeSlice
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("return code should be 200 OK, got %v", res.StatusCode)
	}
	if !reflect.DeepEqual(d, expectedResult) {
		t.Errorf("time lists not equal, got\nresp: %v,\nexpected: %v", d, expectedResult)
	}
}
func TestPeriodicTaskHandlerBadStartTimeFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=100d&tz=Europe/Athens&t1=20BADFORMAT200714T204603Z&t2=20210715T204603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("return code should be 400 Bad Request, got %v", res.StatusCode)
	}
}

func TestPeriodicTaskHandlerBadEndTimeFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=100d&tz=Europe/Athens&t1=20200714T204603Z&t2=20210715T20BADFORMAT4603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("return code should be 400 Bad Request, got %v", res.StatusCode)
	}
}

func TestPeriodicTaskHandlerBadPeriodFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=100badperiod&tz=Europe/Athens&t1=20BADFORMAT200714T204603Z&t2=20210715T204603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("return code should be 400 Bad Request, got %v", res.StatusCode)
	}
}
func TestPeriodicTaskHandlerBadTimeZoneFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ptlist?period=100badperiod&tz=WHATEVER&t1=20BADFORMAT200714T204603Z&t2=20210715T204603Z", nil)
	w := httptest.NewRecorder()
	PeriodicTaskHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("return code should be 400 Bad Request, got %v", res.StatusCode)
	}
}
