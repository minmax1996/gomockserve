package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

var (
	port       = flag.String("port", ":8080", "port for application")
	configPath = flag.String("config", "config.yaml", "")
)

var routesConfig map[string]Config

var baseDir string

type Config struct {
	File string
}

func init() {
	flag.Parse()

	//TODO override by singleRoutes paths from flags
	if configPath != nil {
		yamlFile, err := os.ReadFile(*configPath)
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(yamlFile, &routesConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		baseDir = os.Args[1]
	}
	mux := http.NewServeMux()
	for route := range routesConfig {
		mux.HandleFunc(route, singleFileHandler)
	}
	err := http.ListenAndServe(*port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func singleFileHandler(w http.ResponseWriter, r *http.Request) {
	routeConfig, ok := routesConfig[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := os.ReadFile(path.Join(baseDir, routeConfig.File))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
