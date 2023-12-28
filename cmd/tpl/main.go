package main

import (
	"math/rand"
	"time"

	_ "go.uber.org/automaxprocs"

	tpl "github.com/enamespace/tpl/internal"
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	tpl.NewApp("tpl").Run()
}
