package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bukhavtsov/museum/back-end/pkg/data"

	"github.com/gorilla/mux"
)

type CardData interface {
	ReadAll() ([]*data.Artifact_master_phas, error)
}

type cardAPI struct {
	data CardData
}

func ServerCardResource(r *mux.Router, data CardData) {
	api := &cardAPI{data: data}
	r.HandleFunc("/cards", api.getCards).Methods("GET")
}

func (api cardAPI) getCards(writer http.ResponseWriter, request *http.Request) {
	cards, err := api.data.ReadAll()
	if err != nil {
		log.Println("cards haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	err = json.NewEncoder(writer).Encode(cards)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
