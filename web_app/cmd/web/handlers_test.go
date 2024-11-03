package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/abc", http.StatusNotFound},
	}

	var app application
	routes := app.routes()
	templatesPath = "./../../templates/"

	//create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, v := range theTests {
		resp, err := ts.Client().Get(ts.URL + v.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("Expected status %d got %d", v.expectedStatusCode, resp.StatusCode)
		}
	}
}
