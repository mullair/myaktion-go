package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mullair/myaktion-go/src/myaktion/model"
	"github.com/mullair/myaktion-go/src/myaktion/service"
	log "github.com/sirupsen/logrus"
)

func AddDonation(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	donation, err := requestToDonation(r)
	if err != nil {
		log.Errorf("Failure updating campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = service.AddDonation(id, donation)
	if err != nil {
		log.Errorf("Failure updating campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	sendJson(w, *donation)
}

func requestToDonation(r *http.Request) (*model.Donation, error) {
	var donation model.Donation
	err := json.NewDecoder(r.Body).Decode(&donation)
	if err != nil {
		log.Errorf("Can't serialize request body to campaign struct: %v", err)
		return nil, err
	}
	return &donation, nil
}
