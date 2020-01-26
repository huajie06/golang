package urlshort

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler func
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if v, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, v, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler func
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urlPaths []urlPath
	err := yaml.Unmarshal(yml, &urlPaths)
	if err != nil {
		return nil, err
	}

	// var urlMap = make(map[string]string)
	var urlMap = map[string]string{} // 2nd way
	for _, v := range urlPaths {
		urlMap[v.Path] = v.URL
	}

	return MapHandler(urlMap, fallback), nil
}

type urlPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func ParseJSONurl(jsonPath string) map[string]string {
	jsonByte, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}
	var dat = make([]jsonURLPath, 1)
	if err := json.Unmarshal(jsonByte, &dat); err != nil {
		panic(err)
	}
	var jsonmap = make(map[string]string)
	for _, v := range dat {
		jsonmap[v.Path] = v.URL
	}
	return jsonmap
}

type jsonURLPath struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
