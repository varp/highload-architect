package tools

import "os"

type Env struct {
	prefix string "SOCIAL_NET_01"
}

func NewEnv(prefix string) *Env {
	e := &Env{}
	if len(prefix) > 0 {
		e.ChangePrefix(prefix)
	}
	return e
}

func (e *Env) ChangePrefix(prefix string) {
	e.prefix = prefix
}

func (e *Env) GetWithDefault(name, defaultValue string) string {
	v, exists := os.LookupEnv(name)
	if !exists {
		return defaultValue
	}

	return v
}
