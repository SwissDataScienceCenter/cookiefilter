package cookiefilter_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SwissDataScienceCenter/cookiefilter"
)

func TestRemoveAllCookies(t *testing.T) {
	cfg := cookiefilter.CreateConfig()
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := cookiefilter.New(ctx, next, cfg, "test")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	cookies := []*http.Cookie{
		{Name: "test1", Value: "value1", Path: "/"},
		{Name: "test2", Value: "value2", Path: "/"},
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	handler.ServeHTTP(recorder, req)

	cookies = req.Cookies()
	if len(cookies) != 0 {
		t.Errorf("there should be no cookies in the request, found %d", len(cookies))
	}
}


func TestRemoveSomeCookies(t *testing.T) {
	cookieToRemove := &http.Cookie{Name: "test1", Value: "value1", Path: "/"}
	cookieToKeep := &http.Cookie{Name: "test2", Value: "value2", Path: "/"}
	cookies := []*http.Cookie{cookieToRemove, cookieToKeep}
	cfg := &cookiefilter.Config{
		KeepCookies: []string{cookieToKeep.Name},
	}
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := cookiefilter.New(ctx, next, cfg, "test")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	handler.ServeHTTP(recorder, req)


	cookies = req.Cookies()
	if len(cookies) != 1 {
		t.Errorf("there should be 1 cookies in the request, found %d", len(cookies))
	}
	if cookies[0].Value != cookieToKeep.Value || cookies[0].Name != cookieToKeep.Name  {
		t.Error("the expected cookie that should be kept does not match", len(cookies))
	}
}