package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/segmentio/ksuid"
)

const DEFAULT_GAME_STATE string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -";
type GameInstance struct{
	GameUID ksuid.KSUID
	GameState string
	CurrentTurn bool // false if white, true if black
	Player1UID ksuid.KSUID //white player UID
	Player2UID ksuid.KSUID //black player UID
	Player1Name string //white player username, limited to 32 alphanum chars
	Player2Name string //black player username, limited to 32 alphanum chars
	ExpirationTime time.Time
	Result string
}
var gameMap map[ksuid.KSUID]GameInstance = make(map[ksuid.KSUID]GameInstance);

func CreateGame(conn *websocket.Conn, r *http.Request){
	chessUID := getChessUID(r);
	if chessUID == ksuid.Nil{
		log.Printf("User does not have chessUID, cannot create game. Returning...");
		return;
	}

	gameUID := getGameUID(r);
	newGameInstance := CreateNewGameInstance(chessUID);
	gameMap[newGameInstance.GameUID] = newGameInstance;

	if val, has := gameMap[gameUID]; has{
		if (val.ExpirationTime.Before(time.Now()) ||
					val.GameState != "Undetermined"){
						delete(gameMap, gameUID);
					} else {
						//TODOS: Create response to indicate game already exists and user will need to resign or finish game.
						//TODOS: Or I could have creating a new game be an implied resignation.
					}
	}
		
	

	resPayload := CreateGameResponsePayload{
		GameState: newGameInstance.GameState,
		GameUIDCookie: generateGameUIDCookie(&newGameInstance),
	}

	jsonifiedResPayload, err := json.Marshal(resPayload);
	if err!=nil{
		log.Printf("Error marshalling CreateGameResponsePayload, %s", err.Error());
	}

	resAction := ActionType{
		Type: "CreateGameResponse",
		Payload: jsonifiedResPayload,
	}
	
	jsonifiedResAction, err := json.Marshal(resAction);
	if err!=nil{
		log.Printf("Error marshalling Action for CreateGameResponse, %s", err.Error());
	}
	
	conn.WriteMessage(1, jsonifiedResAction)
}

func CreateNewGameInstance(chessUID ksuid.KSUID)(GameInstance){
	newGameUID, err := ksuid.NewRandom();
		if err !=nil{
			log.Panicf("Hello, we have a big issue, our computer can't generate random numbers. Run.");
		}
	return GameInstance{
		GameUID: newGameUID,
		GameState: DEFAULT_GAME_STATE,
		CurrentTurn: false,
		Player1UID: chessUID, //TODOS: Make it so side is random or choosable
		Player2UID: ksuid.Nil,
		Player1Name: "",
		Player2Name: "",
		ExpirationTime: time.Now().Add(86400*1000*1000), //one day after Now
		Result: "Undetermined",
	}
}


func RequestDraw(w http.ResponseWriter, r *http.Request){
	// gameUID := getGameUID(r);
	// chessUID := getChessUID(r);

	//todos: actually write the function


}

func Resign(w http.ResponseWriter, r *http.Request){
	// todos: actually write the function
}

func ChangeName(w http.ResponseWriter, r *http.Request){
	//todos: actually write the function
}

func SubmitMove(w http.ResponseWriter, r *http.Request){
	//todos: actually write the function
}

func RequestLegalMoves(w http.ResponseWriter, r *http.Request){
	//todos: actually write the function
	//idk i'm thinking legal moves should be calculated on both front-end and back-end
	//to reduce number of messages
}

func RequestGameState(w http.ResponseWriter, r *http.Request){
	//todos: actually write the function

}

func RequestChessUID(w http.Response, r *http.Request){
	return; //i don't think i need to do anything, middleware should take care of it.
}