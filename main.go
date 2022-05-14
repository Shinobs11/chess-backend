package main

import (
	// "fmt"
	"log"
	"net/http"
	// "strings"
	// "github.com/gorilla/websocket"
	// "encoding/json"
	// "github.com/segmentio/ksuid"
	
);

func main(){

	mux := CreateMux();
	println("Now listening on localhost:3001...");
	log.Fatal(http.ListenAndServe("localhost:3001", mux));
}


//https://www.chessprogramming.org/Extended_Position_Description
//My game state is going to be recorded using EPD for compatibility with chess variants
//idk if i'll actually implement chess variants, but it's always good to think ahead









// func authenticateChessUID(game *GameInstance, chessUID ksuid.KSUID)bool{
// 	if(chessUID == game.Player1UID ||chessUID == game.Player2UID){
// 				return true;
// 			} else {
// 				return false;
// 			}
// }







// i'd like to implement this in a scalable and modular fashion
//but i just picked up go about 36 hours ago, so i'm not quite sure how to do generics and stuff
// func handleAction(actionMsg []byte){
	
// 	type Action struct{
// 		Type string
// 		Payload []byte
// 	}
// 	var action Action;
// 	err := json.Unmarshal(actionMsg, &action);
// 	if err != nil{
// 		log.Println("Invalid message schema, returning...");
// 		return;
// 	}

// 	switch(action.Type){
// 	case "requestMoves":{

// 	}
// 	}
// }





















