package cli

import (
	"log"
)

func Cli(args []string) {
	if len(args) == 0 {
		log.Panic("No name given!")
	}
}
