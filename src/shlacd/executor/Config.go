package executor

import (
	"log"
	"shlacd/hrontabd"
)

// executor config
type Config struct {
	Type    string `json:"type"`

	Options struct {
		Silent  bool `json:"silent"`
		Async   bool `json:"async"`

	} `json:"options"`
}


func Resolve(conf Config) (exe hrontabd.Executor){

	switch conf.Type {
	case "local":
		exe = NewExecutorLocal(conf.Options.Silent, conf.Options.Async)

	default:
		log.Panicln("Unknown client type")
	}

	return exe
}