package models

import (
	"github.com/gocql/gocql"
	"time"
)

type CommonFields struct {
	Id        gocql.UUID `json:"id" gocql:"id"`
	CreatedAt time.Time  `json:"created_at" gocql:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" gocql:"updated_at"`
	CreatedBy gocql.UUID `json:"created_by" gocql:"created_by"`
	UpdatedBy gocql.UUID `json:"updated_by" gocql:"updated_by"`
}
