package model

import (
	"time"

	"github.com/volatiletech/null/v8"
)

var (
	DefaultRedisExpiration time.Duration = 5 * time.Minute
	DefaultPageLimit                     = 10
	MustRevalidate                       = "must-revalidate"
	SuperAdminScope        string        = "sup"
	StoreScope             string        = "sto"
	CustomerScope          string        = "cus"
)

type BaseInformation struct {
	CreatedBy int64      `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedBy int64      `json:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedBy null.Int64 `json:"deleted_by" swaggertype:"primitive,string"`
	DeletedAt null.Time  `json:"deleted_at" swaggertype:"primitive,string" example:"2019-08-13T00:00:00Z"`
}

type Pagination struct {
	CurrentPage     int64  `json:"current_page"`
	CurrentElements int64  `json:"current_element"`
	TotalPages      int64  `json:"total_pages"`
	TotalElements   int64  `json:"total_elements"`
	SortBy          string `json:"sort_by"`
}
