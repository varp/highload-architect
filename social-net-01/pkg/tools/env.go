package tools

import (
	"os"
	"strings"
)

const DefaultPrefix = "SOCIAL_NET_01_"

type Env struct {
	prefix string
}

func NewEnv(prefix ...string) *Env {
	e := &Env{}
	e.ChangePrefix(DefaultPrefix)

	if len(prefix) > 0 {
		e.ChangePrefix(strings.Join(prefix, "_"))
	}

	return e
}

func (e *Env) ChangePrefix(prefix string) {
	e.prefix = prefix
}

func (e *Env) GetWithDefault(name, defaultValue string) string {
	v, exists := os.LookupEnv(e.prefix + name)
	if !exists {
		return defaultValue
	}

	return v
}
