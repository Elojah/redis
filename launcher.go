package redis

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for redis service with config file namespaces.
type Namespaces struct {
	Redis services.Namespace
}

// Launcher represents a redis launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	s *Service
	m sync.Mutex
}

// NewLauncher returns a new redis Launcher.
func (s *Service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		s:       s,
		ns:      ns,
	}
}

// Up starts the redis service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Redis]); err != nil {
		// Add namespace key when returning error with logrus
		return err
	}

	return l.s.Dial(sconfig)
}

// Down stops the redis service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.s.Close()
}
