package serve

import "net/http"

func Start(port string) {
  http.ListenAndServe(port, nil)
}
