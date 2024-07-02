package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type PasswordGenerator interface {
	Generate(length int) string
}

type RandomPasswordGenerator struct {
	chars      string
	randSource *rand.Rand
}

func NewRandomPasswordGenerator(chars string) *RandomPasswordGenerator {
	return &RandomPasswordGenerator{
		chars:      chars,
		randSource: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

func (rpg *RandomPasswordGenerator) Generate(length int) string {
	lenChars := len(rpg.chars)
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = rpg.chars[rpg.randSource.Intn(lenChars)]
	}
	return string(password)
}

type Config struct {
	DefaultLength int
	CharSet       string
}

type Application struct {
	config            Config
	passwordGenerator PasswordGenerator
}

func NewApplication(config Config, passwordGenerator PasswordGenerator) *Application {
	return &Application{
		config:            config,
		passwordGenerator: passwordGenerator,
	}
}

func (app *Application) Run() {
	length := app.config.DefaultLength
	if len(os.Args) > 1 {
		userLength, err := strconv.Atoi(os.Args[1])
		if err == nil && userLength > 0 {
			length = userLength
		}
	}

	password := app.passwordGenerator.Generate(length)
	fmt.Println(password)
}

func main() {
	config := Config{
		DefaultLength: 12,
		CharSet:       "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	passwordGenerator := NewRandomPasswordGenerator(config.CharSet)
	app := NewApplication(config, passwordGenerator)
	app.Run()
}
