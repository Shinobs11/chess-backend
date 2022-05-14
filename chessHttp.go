package main

import (
	"strings"
	"encoding/json"
	"log"
	"net/http"
	// "net/http/cookiejar"
	"github.com/gorilla/websocket"
	"github.com/segmentio/ksuid"
	"net/url"
)

var localURL url.URL = url.URL{
	Host: "localhost",
}

func CreateMux()(*WrappedHandler){
	mux := http.NewServeMux();
	mux.HandleFunc("/ws", WebSocketHandler);
	WrappedMux := NewWrappedHandler(mux);
	return WrappedMux;
}



var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	//TODOS: Come back and make the CheckOrigin actually check the origin.
	CheckOrigin: func(r *http.Request) bool{
		return (strings.Split(r.RemoteAddr, ":")[0] == "127.0.0.1");
	},
}




// var cookieJar, _ = cookiejar.New(&cookiejar.Options{
// //todos: figure out if you need options
// })

// var dialer = websocket.Dialer{
// 	Jar: cookieJar, 
// }

func WebSocketHandler(w http.ResponseWriter, r *http.Request){
	// header:= http.Header{

	// }
	
	conn, err := upgrader.Upgrade(w, r, w.Header());
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
		MessageHandler(msg, conn, r);
	} else{
		log.Printf("Unexpected message type, returning...");
	}
	



}




func MessageHandler(message []byte, conn *websocket.Conn, r *http.Request){
	log.Printf("Message received");
	//first things first, find out the requested game.
	var parsedMessage ActionType;
	err := json.Unmarshal(message, &parsedMessage);
	if err!=nil{
		log.Printf("Error decoding JSON message: %s", err.Error());
	}

	switch(parsedMessage.Type){
		case "WebSocketOpen":
			log.Printf("WebSocketOpen message received.");
			
		case "CreateGame":
			//No more json unmarshalling required as all data will be from cookies
			log.Printf("CreateGame message received");
			CreateGame(conn, r);
			
			
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
		case "RequestChessUID":
			break;
		default:
			return;	
	}
	
	//this syntax is disgusting but it's how it's supposed to be done ig
	//(note for me): basically it's `if (initializer, condition) {code...}`




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

func generateGameUIDCookie(game *GameInstance)(http.Cookie){
	gameUIDCookie := http.Cookie{
		Name: "gameUID",
		Value: game.GameUID.String(),
		Domain: "localhost",
		SameSite: http.SameSiteDefaultMode,
	}
	return gameUIDCookie;
}