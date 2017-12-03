package webquic

import (
	"github.com/lucas-clemente/quic-go/h2quic"
)

// Service represents the couchbase service.
type Service struct {
	Adress string
}

// Dial sends the new config to Service.
func (s *Service) Dial(c Config) error {
	s.Adress = c.URL + ":" + c.Port
	return h2quic.ListenAndServeQUIC(s.Adress, c.CertPem, c.CertKey, nil)
}

// Healthcheck returns if database responds.
func (s *Service) Healthcheck() error {
	return nil
}
