package handler

import (
	"net/http"
	"encoding/json"
	yaml "gopkg.in/yaml.v2"
)


//*********************** Map Handler ***********************
// MapHandler function: returns an http Handler function given
// a map between the path and the URL and a fallback handler.
// It redirects the search from the path to the appropiate
// URL.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path		
		if dest, ok := pathsToUrls[path]; ok{
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}


//*********************** YAML Handler ***********************
// YAMLHandler function: returns a custom http handler 
// function given a map between the path and the URL (yaml
// format) and a fallback handler.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

// parseYAML function: takes data in YAML format and converts 
// it to an array of pathUrls objects
func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}



//*********************** JSON Handler ***********************
// JSONHandler function: returns a custom http handler 
// function given a map between the path and the URL (json
// format) and a fallback handler.
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJson(jsonBytes)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

// parseYAML function: takes data in JSON format and converts 
// it to an array of pathUrls objects
func parseJson(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

//*********************** Utils ***********************
// buildMap function: converts an array of pathUrls
// objects to a map between paths and URLs
func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

// pathUrl struct: custom type to help manage the
// (path,url) data
type pathUrl struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}