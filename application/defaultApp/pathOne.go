package defaultApp

import (
	"context"
	"github.com/cafrie/domainone/domain/entity"
	. "github.com/orchestd/servicereply"
)

func (app domainOneApp) PathOne(c context.Context, fOne string) (entity.Foo, ServiceReply) {
	r := entity.Foo{}
	return r, nil
}
