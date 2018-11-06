package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yaml []byte, mux http.Handler) (http.HandlerFunc, error) {
	yamlPaths, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	return MapHandler(parseMapHandler(yamlPaths), mux), nil
}

func parseMapHandler(yamlPaths []yamlPath) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, yamlPath := range yamlPaths {
		pathsToUrls[yamlPath.Path] = yamlPath.URL
	}
	return pathsToUrls
}

func parseYAML(data []byte) ([]yamlPath, error) {
	yamlPaths := make([]yamlPath, 0)
	err := yaml.Unmarshal(data, &yamlPaths)
	if err != nil {
		return yamlPaths, err
	}
	return yamlPaths, nil
}
