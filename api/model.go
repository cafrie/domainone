/* This file is auto-generated and should not be modified */

package api

import (
	"github.com/cafrie/domainone/domain/entity"
)

type Foo entity.Foo

type PathOneRequest struct {
	FOne string `json:"fOne" validate:"required"`
}

type PathOneResponse Foo

const (
	PathOneMethod = "pathOne"
)
