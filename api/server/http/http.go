/* This file is auto-generated and should not be modified */

package http

import (
	"github.com/cafrie/domainone/api"
	"github.com/cafrie/domainone/api/server/monolith"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/transport"
)

func InitHandlers(router transport.IRouter, m monolith.DomainOneInterface, c configuration.Config) {
	router.POST(api.PathOneMethod, transport.HandleFunc(m.PathOne))
}
