package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/segmentio/ksuid"
)

func CreateMux()(*WrappedHandler){
	mux := http.NewServeMux();
	mux.HandleFunc("/ws", WebSocketHandler);
	WrappedMux := NewWrappedHandler(mux);
	return WrappedMux;
}




func WebSocketHandler(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err);
		return;
	}


	msgType, msg, err := conn.ReadMessage();
	if(err!=nil){
		log.Println(err);
		return;
	}

	if (msgType == 1){

	} else{
		log.Printf("Unexpected message type, returning...");
	}


}




func MessageHandler(message []byte, w http.ResponseWriter, r *http.Request){

	//first things first, find out the requested game.
	var parsedMessage ActionType;
	err := json.Unmarshal(message, &parsedMessage);
	if err!=nil{
		log.Printf("Error decoding JSON message: %s", err.Error());
	}

	switch(parsedMessage.Type){
		case "CreateGame":
			
			break;
		case "RequestDraw":
			break;
		case "Resign":
			break;
		case "ChangeName":
			break;
		case "SubmitMove":
			break;
		case "RequestLegalMoves":
			break;
		default:
			return;	
	}
	
	//this syntax is disgusting but it's how it's supposed to be done ig
	//(note for me): basically it's `if (initializer, condition) {code...}`

	if game, ok := gameMap[parsedMessage.GameID]; ok {
		
	} else{
		//game was not found, 
	}
	


}

func getChessUID(r *http.Request)(ksuid.KSUID){
		
	chessUIDCookie, err := r.Cookie("chessUID");
	if err!=nil{
		log.Printf("Idk some issue getting the cookie %s", err.Error());
		return ksuid.Nil;
	}
	chessUID, err := ksuid.Parse(chessUIDCookie.Value);
	if err!= nil{
		log.Printf("Error parsing chessUID, likely invalid chessUID %s", err.Error());
		return ksuid.Nil;
	}
	return chessUID;
}

func getGameUID(r *http.Request)(ksuid.KSUID){
	gameUIDCookie, err := r.Cookie("gameUID");
	if err!=nil{
		log.Printf("GameUID cookie not found");
		return ksuid.Nil;
	}
	gameUID, err := ksuid.Parse(gameUIDCookie.Value);
	if err!=nil{
		log.Printf("Error parsing gameUID from cookie, %s", err.Error());
	}
	return gameUID;
}

func setGameUID(w http.ResponseWriter, game *GameInstance){
	gameUIDCookie := http.Cookie{
		Name: "gameUID",
		Value: game.GameUID.String(),
		Domain: "localhost",
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(w, &gameUIDCookie);
}