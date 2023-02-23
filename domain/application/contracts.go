/* This file is auto-generated and should not be modified */

package application

import (
	"context"
	"github.com/cafrie/domainone/domain/entity"
	. "github.com/orchestd/servicereply"
)

type DomainOneApp interface {
	PathOne(c context.Context, fOne string) (entity.Foo, ServiceReply)
}
