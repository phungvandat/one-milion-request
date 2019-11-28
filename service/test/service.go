package test

import (
	"github.com/phungvandat/onemilion/domain"
)

// Service interface
type Service interface {
	Test(req domain.Payload) error
}
