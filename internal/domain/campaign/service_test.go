package campaign

import (
	"batch-email/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock simulando o BD
type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	//chamando service + repository
	service := Service{}

	//passando DTO
	newCampaign := contract.NewCampaign{
		Name:    "Nome da campanha",
		Content: "Body",
		Emails:  []string{"teste@teste.com", "teste2@email.com"},
	}

	// simulando a criação da campanha
	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_Campaign_Save_On_Db(t *testing.T) {
	// assert := assert.New(t)
	//passando DTO
	newCampaign := contract.NewCampaign{
		Name:    "Nome da campanha",
		Content: "Body",
		Emails:  []string{"teste@teste.com", "teste2@email.com"},
	}
	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	//chamando service + repository
	service := Service{Repository: repositoryMock}

	// simulando a criação da campanha
	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}
