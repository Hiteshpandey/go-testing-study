package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		headerName   string
		headerValue  string
		addr         string
		emptyAddress bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.3.21", "", false},
		{"", "", "hello:world", false},
	}

	var app application
	// create a dummy handler to check context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})

	for _, v := range tests {
		// create handlr to test
		handlerToTest := app.addIpToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)
		if v.emptyAddress {
			req.RemoteAddr = ""
		}

		if len(v.headerName) > 0 {
			req.Header.Add(v.headerName, v.headerValue)
		}

		if len(v.addr) > 0 {
			req.RemoteAddr = v.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	var app application
	var ctx = context.Background()
	ctx = context.WithValue(ctx, contextUserKey, "testname")
	tests := []string{
		"testname",
		"abd",
		"cdr",
	}

	for _, v := range tests {
		ctx = context.WithValue(ctx, contextUserKey, v)
		contextValue := app.ipFromContext(ctx)
		if strings.EqualFold(contextValue, v) {
			t.Errorf("Expected context key %s value to be %v got %v", contextUserKey, v, contextValue)
		}
	}

}
