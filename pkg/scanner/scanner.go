package scanner

import (
	"github.com/tobiaszuercher/vreni/config"
)

type Scanner struct {
	*config.Config
}

func New(cfg *config.Config) *Scanner {
	return &Scanner{
		cfg,
	}
}
