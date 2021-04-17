package service

import (
	"github.com/mullair/myaktion-go/src/myaktion/db"
	"github.com/mullair/myaktion-go/src/myaktion/model"
	log "github.com/sirupsen/logrus"
)

/* my inclass solution
func AddDonation(id uint, donation *model.Donation) (*model.Campaign, error) {
	existingCampaign, _ := GetCampaignById(id)
	result := db.DB.Model(&existingCampaign).Select("donations").Create(donation)
	if result.Error != nil {
		return nil, result.Error
	}
	return existingCampaign, nil
}
*/
func AddDonation(campaignId uint, donation *model.Donation) error {
	donation.CampaignID = campaignId
	result := db.DB.Create(donation)
	if result.Error != nil {
		return result.Error
	}
	entry := log.WithField("ID", campaignId)
	entry.Info("Successfully added new donation to campaign in database.")
	entry.Tracef("Stored: %v", donation)
	return nil
}

/* old implementation
func AddDonation(id uint, donation *model.Donation) (*model.Campaign, error) {
		if existingCampaign, ok := campaignStore[id]; ok {
		existingCampaign.Donations = append(existingCampaign.Donations, *donation)
		entry := log.WithField("ID", id)
		entry.Info("Successfully updated campaign.")
		entry.Tracef("Updated: %v", existingCampaign)
		return existingCampaign, nil
	}
	return nil, fmt.Errorf("campaign for id not found: %d", id)
}
*/
