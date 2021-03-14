package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Tiny-Paws/zen/internal/pkg/hosts"
	"github.com/pelletier/go-toml"
)

type Config struct {
	Websites []string
}

var confFilePath string
var toggle bool

func init() {
	flag.StringVar(&confFilePath, "conf", "", "Path to the configuration file")
    flag.BoolVar(&toggle, "toggle", false, "Toggle zen status")
	flag.Parse()
}

func main() {
	confFile, err := os.Open(confFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer confFile.Close()
	content, err := ioutil.ReadAll(confFile)
	if err != nil {
		panic(err)
	}

	config := Config{}
	toml.Unmarshal(content, &config)
    if toggle {
        if hosts.IsZenInstalled() {
            hosts.RemoveZen()
        } else {
            hosts.InstallZen(config.Websites)
        }
    }
}

