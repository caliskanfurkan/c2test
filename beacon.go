// A simple C2 beaconing emulation tool

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main () {
	for i:=1; i<=10; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println("Just beaconing!")
		http.Get("http://80.240.139.233/c2.php")
	}
}
