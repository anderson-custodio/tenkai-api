package main

import (
	"log"

	"github.com/softplan/tenkai-api/global"
)

func checkFatalError(err error) {
	if err != nil {
		global.Logger.Error(global.AppFields{global.Function: "upload", "error": err}, "erro fatal")
		log.Fatal(err)
	}
}
