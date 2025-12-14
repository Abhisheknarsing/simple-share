package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port := flag.String("p", "8080", "Port to serve on")
	dir := flag.String("d", ".", "Directory to share")
	help := flag.Bool("h", false, "Show help")

	flag.Parse()

	if *help {
		fmt.Println("Usage: simple-share [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}

	// Verify directory exists
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", *dir)
	}

	// Get Local IP
	ip := getLocalIP()

	fmt.Println("🚀 Simple File Share Started")
	fmt.Printf("📂 Sharing Directory: %s\n", *dir)
	fmt.Printf("Local:   http://localhost:%s\n", *port)
	if ip != "" {
		fmt.Printf("Network: http://%s:%s\n", ip, *port)
	}
	fmt.Println("Press Ctrl+C to stop...")

	// Custom handler to add logging
	fs := http.FileServer(http.Dir(*dir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		fs.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback then return it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
