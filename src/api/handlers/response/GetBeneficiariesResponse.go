package response

import (
	"transfeera.backend.developer.test/src/api/domain"
)

type GetBeneficiariesResponse struct {
	Status   string               `json:"status"`
	Code     int                  `json:"code"`
	Metadata *Metadata            `json:"metadata,omitempty"`
	Data     []domain.Beneficiary `json:"data"`
}
