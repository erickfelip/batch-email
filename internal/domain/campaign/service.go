package campaign

import (
	"batch-email/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	// validando se houve erro na criação da campanha
	if err != nil {
		return "", err
	}
	// validando se houve erro no mock para salvar no "BD"
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", err
	}

	// não houve erro, campanha é criada e o ID da campanha é retornado
	return campaign.ID, nil
}
