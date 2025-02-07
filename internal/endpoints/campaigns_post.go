package endpoints

import (
	"batch-email/internal/contract"
	internalErrors "batch-email/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaingService.Create(request)

	if err != nil {
		if errors.Is(err, internalErrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	// status ok created
	render.Status(r, 201)
	// retornando o id na response do client
	render.JSON(w, r, map[string]string{"id": id})

}
