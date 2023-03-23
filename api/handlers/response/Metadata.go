package response

import "math"

type Metadata struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
}

func NewMetadata(totalCount int, page int, perPage int) *Metadata {
	actualPerPage := int(math.Min(float64(perPage), float64(totalCount)))

	return &Metadata{
		TotalCount: totalCount,
		Page:       page,
		PerPage:    actualPerPage,
		TotalPages: totalCount / actualPerPage,
	}
}
