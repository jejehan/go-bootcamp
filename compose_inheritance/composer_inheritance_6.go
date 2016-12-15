package main

import (
	"log"
	"os"
)

type Job struct {
	Command string
	Logger  *log.Logger
}

func main() {
	job := &Job{"demo", log.New(os.Stderr, "job ", log.Ldate)}
	job.Logger.Print("test")
}
