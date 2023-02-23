package repository

import (
	"github.com/cafrie/domainone/domain/repository"
)

type domainOneRepository struct {
}

func NewDomainOneRepository() repository.DomainOneRepository {
	return domainOneRepository{}
}
