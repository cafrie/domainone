/* This file is auto-generated and should not be modified */

package api

import (
	"context"
	. "github.com/orchestd/servicereply"
)

type DomainOneApi interface {
	PathOne(c context.Context, req PathOneRequest) (PathOneResponse, ServiceReply)
}
