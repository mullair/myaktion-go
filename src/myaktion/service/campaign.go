package service

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/mullair/myaktion-go/src/myaktion/model"
)

var (
	campaignStore map[uint]*model.Campaign
	actCampaignId uint = 1
)

func init() {
	campaignStore = make(map[uint]*model.Campaign)
}
func CreateCampaign(campaign *model.Campaign) error {
	campaign.ID = actCampaignId
	campaignStore[actCampaignId] = campaign
	actCampaignId += 1
	log.Infof("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Tracef("Stored: %v", campaign)
	return nil
}

func GetCampaigns() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	for _, campaign := range campaignStore {
		campaigns = append(campaigns, *campaign)
	}
	log.Tracef("Retrieved: %v", campaigns)
	return campaigns, nil
}

/*
func UpdateCampaign(id uint, campaign *model.Campaign) error {
	campaignStore[id] = campaign
	return nil
}
*/

func UpdateCampaign(id uint, campaign *model.Campaign) (*model.Campaign, error) {
	existingCampaign, err := GetCampaignById(id)
	if err != nil {
		return existingCampaign, err
	}
	existingCampaign.Name = campaign.Name
	existingCampaign.Organizer = campaign.Organizer
	existingCampaign.TargetAmount = campaign.TargetAmount
	existingCampaign.DonationMinimum = campaign.DonationMinimum
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated campaign.")
	entry.Tracef("Updated: %v", existingCampaign)
	return existingCampaign, nil

}

func GetCampaignById(id uint) (*model.Campaign, error) {
	campaign := campaignStore[id]
	log.Tracef("campaign found and read", campaign)
	return campaign, nil
}

func DeleteCampaign(id uint) (*model.Campaign, error) {
	campaign := campaignStore[id]
	if campaign == nil {
		log.Tracef("404 campaign not found")
		return nil, fmt.Errorf("no campaign with ID %d", id)
	}
	delete(campaignStore, id)
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted campaign.")
	log.Infof("Successfully deleted campaign with ID %d", id)
	return campaign, nil
}
