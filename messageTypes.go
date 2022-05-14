package main

import (
	"encoding/json"
	"net/http"
)

//i know this will "work"
//but i have no idea if it's remotely close to "best practices"
type ActionType struct {
	Type string `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
//receiving payloads
type CreateGamePayload struct { // done

}

type RequestDrawPayload struct { //

}

type ResignPayload struct {

}

type ChangeNamePayload struct {
	RequestedName string `json:"requestedName"`
}

type SubmitMovePayload struct {
	Move string `json:"move"`
}

type RequestLegalMovesPayload struct {

}

//responding payloads
type CreateGameResponsePayload struct {
	GameState string `json:"gameState"`
	GameUIDCookie http.Cookie `json:"gameUIDCookie"`
}


/*
list of message types ideas:

User:
Create a new game
Resign from current game
Enter in name
Rename


Front-end/ui/game:
Request set of valid moves
Make a move

*/