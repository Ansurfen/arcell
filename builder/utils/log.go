package utils

import (
	"log"
)

const (
	DEBUG = iota
	DEV
	RELEASE
)

type MODE = int

var mode MODE

func init() {
	mode = DEBUG
}

func SetMode(m MODE) {
	mode = m
}

func Log(msg string) {
	switch mode {
	case DEBUG:
		log.Println(msg)
	case DEV:
		log.Println(msg)
	case RELEASE:
	default:
		log.Fatal("Unknown mode")
	}
}
