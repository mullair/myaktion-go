package service

import (
	"fmt"

	"github.com/mullair/myaktion-go/src/myaktion/model"
	log "github.com/sirupsen/logrus"
)

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
