package main

import (
    "os"
    "fmt"
    "flag"
    "io/ioutil"

    "github.com/pelletier/go-toml"
    "github.com/Tiny-Paws/zen/internal/pkg/hosts"
)

type Config struct {
    Websites []string
}

var confFilePath string

func init() {
    flag.StringVar(&confFilePath, "conf", "", "Path to the configuration file")
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
    fmt.Println(config)
    //hosts.InstallZen(config.Websites)
    hosts.RemoveZen()
}

