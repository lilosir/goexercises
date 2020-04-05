package handler

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathMap[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `ymal:"url"`
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathurls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(pathurls)
	return MapHandler(pathMap, fallback), nil
}

func parseYaml(yamlBytes []byte) ([]pathURL, error) {
	var pathurls []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathurls)
	if err != nil {
		fmt.Println("unmarshal ymal file failed")
		return nil, err
	}
	return pathurls, nil
}

func buildMap(pathurls []pathURL) map[string]string {
	pathMap := make(map[string]string)
	for _, pu := range pathurls {
		pathMap[pu.Path] = pu.URL
	}
	return pathMap
}

func IndexFuncHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "user 1")
}
