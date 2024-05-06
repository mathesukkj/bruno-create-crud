package main

import (
	"os"

	"github.com/mathesukkj/gocat/internal/cli"
)

func main() {
	cli.Cli(os.Args[1:])
}
