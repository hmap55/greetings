# saludos en go
Este paquete proporciona una forma de como obtener saludos personalizados
## instalacion
Ejecuta el siguiente comando para intalar el paquete:
```bash

go get -u github.com/hmap55/greetings

```
# uso
aqui tienes como usar el paquete en tu codigo:

```go
package main

import (
	"fmt"
	"log"

	"github.com/hmap55/greetings"
)

func main() {

	log.SetPrefix("greetings: ")

	names := []string{"Alex", "Juan", "Maria"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
```

