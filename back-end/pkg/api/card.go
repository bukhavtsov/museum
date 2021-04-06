package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bukhavtsov/museum/back-end/pkg/data"

	"github.com/gorilla/mux"
)

type ArtifactData interface {
	ReadAll() ([]*data.ArtifactMaster, error)
	Add(artifact *data.ArtifactMaster) (int, error)
	Update(artifactId int, newArtifact *data.ArtifactMaster) error
	Delete(artifactId int) error
}

type artifactAPI struct {
	data ArtifactData
}

func ServerArtifactResource(r *mux.Router, data ArtifactData) {
	api := &artifactAPI{data: data}
	r.HandleFunc("/artifacts", api.getArtifacts).Methods("GET")
	r.HandleFunc("/artifacts", api.createArtifact).Methods("POST")
	r.HandleFunc("/artifacts/{id}", api.updateArtifact).Methods("PUT")
	r.HandleFunc("/artifacts/{id}", api.deleteArtifact).Methods("DELETE")
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

func (api artifactAPI) createArtifact(writer http.ResponseWriter, request *http.Request) {
	artifact := new(data.ArtifactMaster)
	err := json.NewDecoder(request.Body).Decode(&artifact)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if artifact == nil {
		log.Printf("failed empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	artifactId, err := api.data.Add(artifact)
	if err != nil {
		log.Println("artifact hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/artifact/%d", artifactId))
	writer.WriteHeader(http.StatusCreated)
}

func (api artifactAPI) updateArtifact(w http.ResponseWriter, req *http.Request) {
	var artifactMaster *data.ArtifactMaster
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(req.Body).Decode(&artifactMaster)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.data.Update(int(id), artifactMaster)
	if err != nil {
		log.Printf("artifact hasn't been updated, err is: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (api artifactAPI) deleteArtifact(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = api.data.Delete(int(id))
	if err != nil {
		log.Println("artifact hasn't been removed")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
