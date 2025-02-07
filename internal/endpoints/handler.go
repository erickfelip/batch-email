package endpoints

import "batch-email/internal/domain/campaign"

type Handler struct {
	CampaingService campaign.Service
}
