package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var (
	port       = pflag.String("p", ":8080", "port for application")
	configPath = pflag.String("c", "config.yaml", "")
	routeFlags = pflag.StringArray("r", nil, "routes")
)

var routesConfig map[string]Config

var baseDir string

type Config struct {
	File string
}

func init() {
	pflag.Parse()
	routesConfig = make(map[string]Config)
	if routeFlags != nil && len(*routeFlags) > 0 {
		for _, r := range *routeFlags {
			rparts := strings.Split(r, ":")
			if len(rparts) != 2 {
				log.Fatal("bad route syntax")
			}
			routesConfig[rparts[0]] = Config{
				File: rparts[1],
			}
		}
	} else if configPath != nil {
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
	// TODO handle args in some better way
	if len(os.Args) > 1 {
		baseDir = os.Args[1]
	}
	mux := http.NewServeMux()
	for route := range routesConfig {
		mux.HandleFunc(route, singleFileHandler)
	}
	log.Printf("start application on %s port", *port)
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
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
