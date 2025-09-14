package main

import (
	"os"

	"github.com/kaushtubhkanishk/streamlens/SQLEngine"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	//log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	//s := "SELECT * FROM USERS"
	SQLEngine.Tokenize(os.Stdin, os.Stdout)
}
