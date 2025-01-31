package db

import (
	"batch-email/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

// fake mock to "save" on db
func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	// create mock campaign
	return nil

	// force error
	// return errors.New("error")
}
