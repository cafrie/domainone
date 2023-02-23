/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/cafrie/domainone/api"
	"github.com/cafrie/domainone/api/server/monolith"
	. "github.com/orchestd/servicereply"
)

type DomainOneMonolithClient struct {
	MonolithInterface monolith.DomainOneInterface
}

func NewDomainOneMonolithClient(monolithInterface monolith.DomainOneInterface) api.DomainOneApi {
	return DomainOneMonolithClient{MonolithInterface: monolithInterface}
}

func (p DomainOneMonolithClient) PathOne(c context.Context, req api.PathOneRequest) (api.PathOneResponse, ServiceReply) {
	return p.MonolithInterface.PathOne(c, req)
}
