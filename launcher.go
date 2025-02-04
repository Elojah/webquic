package webquic

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for couchbase service with config file namespaces.
type Namespaces struct {
	WebQuic services.Namespace
}

// Launcher represents a couchbase launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	s *Service
	m sync.Mutex
}

// NewLauncher returns a new couchbase Launcher.
func (s *Service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		s:       s,
		ns:      ns,
	}
}

// Up starts the couchbase service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	cfg := Config{}
	if err := cfg.Dial(configs[l.ns.WebQuic]); err != nil {
		// Add namespace key when returning error with logrus
		return err
	}
	return l.s.Dial(cfg)
}

// Down stops the couchbase service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return nil
}
