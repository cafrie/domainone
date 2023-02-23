/* This file is auto-generated and should not be modified */

package http

import (
	"context"
	"github.com/cafrie/domainone/api"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	. "github.com/orchestd/servicereply"
)

const serviceName = "domainone"

type domainOneHTTPClient struct {
	client transport.HttpClient
}

func NewDomainOneApiHTTPClient(client transport.HttpClient) api.DomainOneApi {
	return domainOneHTTPClient{client: client}
}

func (h domainOneHTTPClient) PathOne(c context.Context, req api.PathOneRequest) (api.PathOneResponse, ServiceReply) {
	var res api.PathOneResponse
	if reply := h.client.Call(c, req, serviceName, api.PathOneMethod, &res, nil); !reply.IsSuccess() {
		return res, reply
	}
	return res, nil
}
