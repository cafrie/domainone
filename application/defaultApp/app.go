package defaultApp

import (
	"github.com/cafrie/domainone/domain/application"
)

type domainOneApp struct {
}

func NewDomainOneApp() application.DomainOneApp {
	return &domainOneApp{}
}
