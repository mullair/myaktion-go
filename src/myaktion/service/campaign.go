package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/mullair/myaktion-go/src/myaktion/db"
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
	result := db.DB.Create(campaign)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Tracef("Stored: %v", campaign)
	return nil
}

func GetCampaigns() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	result := db.DB.Preload("Donations").Find(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaigns)
	return campaigns, nil
}

func GetCampaignById(id uint) (*model.Campaign, error) {
	var campaign model.Campaign

	result := db.DB.Preload("Donations").Find(&campaign, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //keine campaign gefunden
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaign)

	return &campaign, nil
}

func UpdateCampaign(id uint, campaign *model.Campaign) (*model.Campaign, error) {
	//var campaign model.Campaign
	existingCampaign, _ := GetCampaignById(id)
	result := db.DB.Model(&existingCampaign).Updates(campaign)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated campaign.")
	entry.Tracef("Updated: %v", &campaign)
	return existingCampaign, nil

}

func DeleteCampaign(id uint) (*model.Campaign, error) {
	campaign, _ := GetCampaignById(id)
	result := db.DB.Delete(&campaign)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted campaign.")
	entry.Tracef("Deleted: %v", &campaign)
	return campaign, nil
}
