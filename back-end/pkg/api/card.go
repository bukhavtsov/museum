package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bukhavtsov/museum/back-end/pkg/data"

	"github.com/gorilla/mux"
)

type ArtifactData interface {
	ReadAll() ([]*data.ArtifactMaster, error)
}

type artifactAPI struct {
	data ArtifactData
}

func ServerArtifactResource(r *mux.Router, data ArtifactData) {
	api := &artifactAPI{data: data}
	r.HandleFunc("/artifacts", api.getArtifacts).Methods("GET")
}

func (api artifactAPI) getArtifacts(writer http.ResponseWriter, request *http.Request) {
	cards, err := api.data.ReadAll()
	if err != nil {
		log.Println("artifacts haven't been read")
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
