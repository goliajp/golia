package liblog

import (
	"log"
	"os"
)

func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Fatal(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
