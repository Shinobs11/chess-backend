package main

import (
	// "fmt"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/websocket"
	"encoding/json"
);

type GameInstance struct{
	GameID string
	GameState string
	Player1UID string //white player UID
	Player2UID string //black player UID
	Player1Name string //white player username, limited to 32 alphanum chars
	Player2Name string //black player username, limited to 32 alphanum chars
}
//https://www.chessprogramming.org/Extended_Position_Description
//My game state is going to be recorded using EPD for compatibility with chess variants
//idk if i'll actually implement chess variants, but it's always good to think ahead
const DEFAULT_GAME_STATE string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -";

var gameMap map[string]GameInstance = make(map[string]GameInstance);

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	//TODOS: Come back and make the CheckOrigin actually check the origin.
	CheckOrigin: func(r *http.Request) bool{
		return (strings.Split(r.RemoteAddr, ":")[0] == "127.0.0.1");
	},
}

func webSocketHandler(w http.ResponseWriter, r *http.Request){
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

	type TestPayloadType struct {
		A string  `json:"a"`
		B []int		`json:"b"`
		C float64	`json:"c"`
	}
	//this is probably one of my least favourite parts about go so far
	type ActionType struct {
		Type string `json:"type"`
		Payload json.RawMessage `json:"payload"`
	}
	type MessageType struct {
		GameID string `json:"gameID"` //LIKE WHAT IS THIS EVEN
		Action ActionType `json:"action"` //i think it's a bit silly to have a language opinionated on what case you use
	}




	if(msg!=nil){
		log.Println("Message received!");
		if(msgType == 1){

			log.Println(string(msg));
			var parsedMsg MessageType;
			
			json.Unmarshal(msg, &parsedMsg);
			parsedAction := parsedMsg.Action;
			log.Println(parsedAction);
			log.Println("gameID: " + parsedMsg.GameID);
			log.Println("type: " + parsedAction.Type);
			var parsedPayload TestPayloadType;
			log.Println(parsedAction.Payload);
			json.Unmarshal(parsedAction.Payload, &parsedPayload);
			log.Println(parsedPayload);
			

		}else if(msgType == 2){
			log.Println("Bytecoded message");
		} else {
			log.Panicln("Message is not in text format.");
		}
		conn.WriteMessage(1, []byte("Hello from Go, message received."));
	}
	
}




// i'd like to implement this in a scalable and modular fashion
//but i just picked up go about 36 hours ago, so i'm not quite sure how to do generics and stuff
func handleAction(actionMsg []byte){
	
	type Action struct{
		Type string
		Payload []byte
	}
	var action Action;
	err := json.Unmarshal(actionMsg, &action);
	if err != nil{
		log.Println("Invalid message schema, returning...");
		return;
	}

	switch(action.Type){
	case "requestMoves":{

	}
	}




	

}




func main(){

	mux := http.NewServeMux();

	mux.HandleFunc("/ws", webSocketHandler);

	println("Now listening on localhost:3001...");
	log.Fatal(http.ListenAndServe("localhost:3001", mux));
}


















