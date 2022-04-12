package main

import (
  "fmt"
  "log"

  "github.com/psymaza/storage/internal/storage"
)

func main() {
  st := storage.NewStorage()

  file, err := st.Upload("test.txt", []byte("Hello world!"))

  if err != nil {
    log.Fatal(err)
  }

  data, err := st.GetById(file.ID)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(data)
}
