package main

import (
	"log"
	"net/http"
	"github.com/segmentio/ksuid"
)






type WrappedHandler struct {
	handler http.Handler
}

func (wh *WrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	chessUID := getChessUID(r);
	
	if chessUID == ksuid.Nil {
		newID, err := ksuid.NewRandom();
		if err!= nil{
			log.Printf("For whatever reason I can't generate a new random ID, here's the error %s", err.Error());
		}
		
		cookie:= http.Cookie{
			Name: "chessUID",
			Value: string(newID.String()),
			Domain: "localhost",
			SameSite: http.SameSiteDefaultMode,
			//consider checking out rest of options
		}

		http.SetCookie(w, &cookie)
	}
	wh.handler.ServeHTTP(w, r);	
}




func NewWrappedHandler(handlerToWrap http.Handler)*WrappedHandler{
		return &WrappedHandler{handlerToWrap};
}

