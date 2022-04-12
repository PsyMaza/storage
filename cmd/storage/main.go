package main

import (
	"fmt"

	"github.com/psymaza/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("It's Worked", st)
}
