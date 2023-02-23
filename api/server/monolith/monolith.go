/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/cafrie/domainone/api"
	"github.com/cafrie/domainone/domain/application"
	"github.com/orchestd/dependencybundler/interfaces/validations"
	. "github.com/orchestd/servicereply"
)

type DomainOneInterface struct {
	domainOneApp application.DomainOneApp
	validation   validations.Validations
}

func NewDomainOneInterface(domainOneApp application.DomainOneApp, validation validations.Validations) DomainOneInterface {
	return DomainOneInterface{domainOneApp: domainOneApp, validation: validation}
}

func (i DomainOneInterface) PathOne(c context.Context, req api.PathOneRequest) (api.PathOneResponse, ServiceReply) {
	if vErr := i.validation.Validate(req); vErr != nil && !vErr.IsSuccess() {
		return api.PathOneResponse{}, vErr
	}
	res, err := i.domainOneApp.PathOne(c, req.FOne)
	return api.PathOneResponse(res), err
}
