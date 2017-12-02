/*
Module for converting request URLs to request handlers.
*/
package urls

import (
  "errors"
  "net/http"
  "regexp"
)

type HandlerMatch struct {
  URL_name string
  Args map[string]string
  Handler http.Handler
}

type RegexURLPattern struct {
  Handler http.Handler
  Args map[string]string
  Name string
  Regex *regexp.Regexp
}

func (r *RegexURLPattern) resolve (path string) (HandlerMatch, error) {
  match := r.Regex.FindStringSubmatch(path)
  names := r.Regex.SubexpNames()
  h := HandlerMatch{}

  if match != nil {
    for i := range match {
      if i == 0 {
        continue
      }
      r.Args[names[i]] = match[i]
    }
    h.URL_name = r.Name
    h.Args =
    return h, nil
  }
}
