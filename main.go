package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", viewHandler)
	fmt.Println("starting at :8080 =>")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	ifaces, err := net.Interfaces()
	if err != nil {
		data := fmt.Sprintf("Error %s\n", err)
		fmt.Fprintf(w, data)
		return
	}

	var ip net.IP
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
		}
	}

	fmt.Fprintf(w, "[Version 6] [%s] username %s, password %s, at %s\n", ip.String(), os.Getenv("USER"), os.Getenv("PASS"), time.Now().String())
}
