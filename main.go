package main

import (
	"log"
	"net/http"
	"strings"
)
const DEFAULT_GAME_STATE string = "hello";
type GameInstance struct{
	GameID string
	GameState string
}


var gameMap map[string]GameInstance = make(map[string]GameInstance);






func createGameHandler(res http.ResponseWriter, req *http.Request){
	if len(req.URL.Path)>256{
		//insert error here when i figure out how to do that
		return;
	}
	reqid := strings.Split(req.URL.Path, "/")[2];
	if reqid == ""{
		//again insert error
		return;
	}

	//create new game by adding it to map
	gameMap[reqid] = GameInstance{GameID: reqid, GameState: DEFAULT_GAME_STATE}




}

func main(){
	mux := http.NewServeMux();
	mux.HandleFunc("/create/game/", createGameHandler);


	println("Now listening on localhost:3001...");
	log.Fatal(http.ListenAndServe("localhost:3001", mux));
}


















