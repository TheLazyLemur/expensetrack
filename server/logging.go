package main

import (
	"log"
)

type Logger struct {
	pkg string
}

func (l Logger) LogInfo(v interface{}){
	log.Println("\x1b[32mINFO", v, "\x1b[0m")
}

func (l Logger) LogWarning(v interface{}){
	log.Println("\x1b[33mWARNING", v, "\x1b[0m")
}

func (l Logger) LogError(v interface{}){
	log.Println("\x1b[31mERROR", v, "\x1b[0m")
}
