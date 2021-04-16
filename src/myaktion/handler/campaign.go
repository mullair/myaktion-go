package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/mullair/myaktion-go/src/myaktion/model"
	"github.com/mullair/myaktion-go/src/myaktion/service"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign *model.Campaign
	campaign, err := getCampaign(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(campaign); err != nil {
		log.Errorf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaign)
}

func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Errorf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, campaigns)
}

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		log.Errorf("Error parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := service.GetCampaignById(id)
	if campaign == nil {
		log.Infof("Id not found: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	sendJson(w, campaign)
}

//my function
/*func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		log.Errorf("ID does not exist: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := getCampaign(r)
	if err != nil {
		log.Errorf("Campaign can not be called: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.UpdateCampaign(id, campaign)
	if err != nil {
		log.Errorf("Error updating campaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, campaign)
}*/

//in class code
func UpdateCampaign(w http.ResponseWriter, r *http.Request) {

	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := getCampaign(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err = service.UpdateCampaign(id, campaign)
	if err != nil {
		log.Errorf("Failure updating campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if campaign == nil {
		http.Error(w, "404 campaign not found", http.StatusNotFound)
		return
	}
	sendJson(w, campaign)
}

func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := service.DeleteCampaign(id)
	if err != nil {
		log.Errorf("Error in deletion %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if campaign == nil {
		http.Error(w, "404 campaign not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "ok"})

}

func getCampaign(r *http.Request) (*model.Campaign, error) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Errorf("Can't serialize request body to campaign struct: %v", err)
		return nil, err
	}
	return &campaign, nil
}
