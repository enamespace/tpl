package main

import (
	"log"

	"github.com/enamespace/tpl/internal/db"
)

func main() {

	if err := db.NewDBOptions().Run(); err != nil {
		log.Fatal(err)
	}

}
