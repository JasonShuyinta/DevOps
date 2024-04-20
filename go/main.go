package main 

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
  })
  

  http.ListenAndServe(":8081", nil)
}
