package main

import (
	"net/http"
	"time"

	"urlshorter/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	defaultHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := handler.YAMLHandler([]byte(yaml), defaultHandler)
	if err != nil {
		panic(err)
	}

	mux.HandleFunc("/", handler.IndexFuncHandler)
	mux.HandleFunc("/user", handler.UserHandler)

	s := &http.Server{
		Addr:    ":8800",
		Handler: yamlHandler,
		// Handler:      defaultHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
