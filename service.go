package webquic

import (
	"crypto/tls"
	"net/http"

	"github.com/devsisters/goquic"
)

// Service represents the webquic service.
type Service struct {
	Server *goquic.QuicSpdyServer
}

// Dial sends the new config to Service.
func (s *Service) Dial(c Config) error {
	var err error
	if s.Server, err = goquic.NewServer(
		c.Address,
		c.CertPem,
		c.CertKey,
		c.Dispatcher,
		http.DefaultServeMux,
		http.DefaultServeMux,
		&tls.Config{MinVersion: tls.VersionSSL30},
	); err != nil {
		return err
	}
	go func() { _ = s.Server.ListenAndServe() }()
	return nil
}

// Healthcheck returns if database responds.
func (s *Service) Healthcheck() error {
	return nil
}
