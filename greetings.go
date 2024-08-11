package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// devuleve un saludo de la persona especifica
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(random(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func random() string {
	formats := []string{
		"hola, %v bienvenido",
		"que bueno verte, %v",
		"¡Saludo, %v encantado de conecerte",
	}

	return formats[rand.Intn(len(formats))]
}
