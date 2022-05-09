package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestCreateGameHandler(t *testing.T) {
	testGameId := "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
	req := httptest.NewRequest(http.MethodPost, "/create/game/"+testGameId, nil);
	w := httptest.NewRecorder();
	createGameHandler(w, req);
	res := w.Result();
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body);
	if err != nil {
		t.Errorf("expected error to be nil, but got %v" ,err);
	}
	expectedGame := GameInstance{
		GameID: testGameId,
		GameState: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -",
	}
	if gameMap[testGameId] != expectedGame {
		t.Errorf("expected gameMap[%s] to be %v, but got %v instead",
		testGameId, expectedGame, gameMap[testGameId])
	}
}