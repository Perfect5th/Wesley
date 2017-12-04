package serve

import (
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

func ShowIndex(w http.ResponseWriter, req *http.Request) {
  base_dir, _ := os.Getwd()
  files, _ := ioutil.ReadDir(base_dir)

  log.Printf("%s %s", req.Method, req.URL.Path)
  for i := range files {
    nameStr := fmt.Sprintf("%s\n", files[i].Name())
    io.WriteString(w, nameStr)
  }
}

func Start(port string) {
  http.HandleFunc("/", ShowIndex)
  log.Fatal(http.ListenAndServe(port, nil))
}
