package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"yemeksepeti-golang-rest/module/key_value"
)

//TestKeyValueServerSet test setting key with httptest
func TestKeyValueServerSet(t *testing.T) {
	//Test data to add list
	data := strings.NewReader(`
	{
		"key": "test",
		"value": "test bar"
	}`)
	//Creating request
	req := httptest.NewRequest(http.MethodPost, "localhost:8080/set", data)
	//Creating response recorder for testing
	rr := httptest.NewRecorder()
	//Initialize our KeyValue set handler
	handler := http.HandlerFunc(key_value.SetKeyHandler)

	//Running request
	handler.ServeHTTP(rr, req)

	//Check if status OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Check body of response recorder, convert it to KeyValue Entity and check keys are matched
	if key_value.DecodeJson(rr.Body.String()).Key != "test" {
		t.Errorf("handler returned unexpected body: got %v ",
			rr.Body.String())
	}
}

//TestKeyValueServerGet test getting key with httptest
func TestKeyValueServerGet(t *testing.T) {
	//Our getting handler gets key from url path,making httptest request doesnt pass our path to handler.
	//So we need to create new server and pass same handler to same pattern and test it there
	mux := http.NewServeMux()
	mux.HandleFunc("/get/", key_value.GetKeyHandler)

	//Creating new server with mux
	srv := httptest.NewServer(mux)
	defer srv.Close()

	//Creating request with the url of server
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/get/test", srv.URL), nil)
	if err != nil {
		t.Errorf("Error on http.NewRequest: %v", err)
	}

	//Simple run the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Error on http.DefaultClient.Do(req): %v", err)
	}
	defer res.Body.Close()

	//response body is io.Reader, using ioutil to read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error on ioutil.ReadAll(res.Body): %v", err)
	}

	//Convert body to KeyValue Entity and check keys are matched
	if key_value.DecodeJson(string(body)).Key != "test" {
		t.Errorf("handler returned unexpected body: got %v ",
			string(body))
	}

	//Check if status OK
	if http.StatusOK != res.StatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}

func TestKeyValueServerFlush(t *testing.T) {
	//Creating request
	req := httptest.NewRequest(http.MethodDelete, "localhost:8080/flush", nil)
	//Creating response recorder for testing
	rr := httptest.NewRecorder()
	//Initialize our KeyValue set handler
	handler := http.HandlerFunc(key_value.FlushHandler)

	//Running request
	handler.ServeHTTP(rr, req)

	//Check if status OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Check body of response is matched expected message
	if rr.Body.String() != "Succesfully flushed all data" {
		t.Errorf("handler returned unexpected body: got %v ",
			rr.Body.String())
	}
}
