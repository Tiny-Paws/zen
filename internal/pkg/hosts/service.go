package hosts

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

const (
    beginZen = "#### zen ####"
    endZen = "#### end ####"
)

func RemoveZen() {
    var buffer strings.Builder
    file, err := os.Open(hostsPath)
    if err != nil {
        fmt.Println("Could not remove zen")
        os.Exit(1)
    }
    defer file.Close()

    zenZone := false
    sc := bufio.NewScanner(file)
    for sc.Scan() {
        line := sc.Text()
        if line == beginZen {
            zenZone = true
        } else if line == endZen {
            zenZone = false
            continue // Otherwise the endZen footer stays in the hosts file
        }

        if !zenZone {
            buffer.WriteString(line + "\n")
        }
    }
    err = os.WriteFile(hostsPath, []byte(buffer.String()), 0644)
    if err != nil {
        fmt.Println("Could not write %s\n%v", hostsPath, err)
    }
}

func InstallZen(websites []string) {
    file, err := os.OpenFile(hostsPath, os.O_APPEND | os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Could not install zen", err)
        os.Exit(1)
    }
    defer file.Close()
    content := websitesToHosts(websites...)
    file.WriteString(content)
}

func IsZenInstalled() bool {
    file, err := os.Open(hostsPath)
    if err != nil {
        fmt.Println("Could not access %s\n%v", hostsPath, err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for sc.Scan() {
        line := sc.Text()
        if line == beginZen {
            return true
        }
    }
    return false
}

func websitesToHosts(websites ...string) string {
    var b strings.Builder
    fmt.Fprintf(&b, beginZen + "\n")
    for _, website := range websites {
        fmt.Fprintf(&b, "127.0.0.1 %s\n", website)
    }
    fmt.Fprintf(&b, endZen + "\n")
    return b.String()
}

