package main

import (
	"github.com/cafrie/domainone/api/server/monolith"
	"github.com/cafrie/domainone/application/defaultApp"
)

func deps() []interface{} {
	return append(internalDeps(), externalDeps()...)
}

func internalDeps() []interface{} {
	return []interface{}{
		defaultApp.NewDomainOneApp,
		monolith.NewDomainOneInterface,
	}
}

func externalDeps() []interface{} {
	return []interface{}{}
}
