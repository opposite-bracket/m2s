package utils

import "log"

func PanicOnError(format string, err error) {
	if err != nil {
		log.Panicf(format, err)
	}
}

func LogOnError(format string, err error) {
	if err != nil {
		log.Printf(format, err)
	}
}
