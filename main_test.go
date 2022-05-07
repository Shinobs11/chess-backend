package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestCreateGameHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/create/game/123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", nil);
	w := httptest.NewRecorder();
	createGameHandler(w, req);
	res := w.Result();
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body);
	if err != nil {
		t.Errorf("expected error to be nil, but got %v" ,err);
	}
	
}