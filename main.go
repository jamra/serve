package main

import (
	"flag"
	"fmt" 
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"serverepo-tunneling/tunneling"
)

func main() {
	// Define flags
	port := flag.String("port", "8000", "Port to serve on")
	mode := flag.String("mode", "file", "Mode: file, proxy, or tunnel")
	proxyURL := flag.String("proxy", "", "Proxy server URL (for tunnel mode)")
	localPort := flag.String("local-port", "", "Local port to forward (for tunnel mode)")
	subdomain := flag.String("subdomain", "", "Custom subdomain (for tunnel mode)")
	flag.Parse()

	switch *mode {
	case "file":
		startFileServer(*port)
	case "proxy":
		startProxyServer(*port)
	case "tunnel":
		if *proxyURL == "" || *localPort == "" {
			log.Fatal("Tunnel mode requires -proxy and -local-port flags")
		}
		localPortInt, err := strconv.Atoi(*localPort)
		if err != nil {
			log.Fatal("Invalid local port:", err)
		}
		startTunnelClient(*proxyURL, localPortInt, *subdomain)
	default:
		log.Fatal("Invalid mode. Use: file, proxy, or tunnel")
	}
}

// startFileServer starts the original file server
func startFileServer(port string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", LogReq(http.FileServer(http.Dir(dir))))
	fmt.Printf("🗂️  File server running on http://localhost:%s
", port)
	fmt.Printf("📁 Serving files from: %s
", dir)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// startProxyServer starts the proxy server for tunneling
func startProxyServer(port string) {
	proxy := tunneling.NewProxyServer()

	// Handle tunnel WebSocket connections
	http.HandleFunc("/tunnel", proxy.HandleTunnel)

	// Handle HTTP requests for tunneled subdomains
	http.HandleFunc("/", proxy.HandleHTTP)

	fmt.Printf("🚀 Proxy server running on http://localhost:%s
", port)
	fmt.Printf("📡 WebSocket endpoint: ws://localhost:%s/tunnel
", port)
	fmt.Printf("🌐 Subdomain format: http://{subdomain}.localhost:%s
", port)
	fmt.Println("
💡 To create a tunnel, run:")
	fmt.Printf("   go run main.go -mode=tunnel -proxy=http://localhost:%s -local-port={your-app-port}

", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// startTunnelClient starts the tunnel client
func startTunnelClient(proxyURL string, localPort int, subdomain string) {
	fmt.Printf("🔗 Starting tunnel client...
")
	fmt.Printf("📡 Proxy server: %s
", proxyURL)
	fmt.Printf("🏠 Local port: %d
", localPort)
	if subdomain != "" {
		fmt.Printf("🏷️  Custom subdomain: %s
", subdomain)
	}
	fmt.Println()

	client := tunneling.NewTunnelClient(proxyURL, localPort, subdomain)
	client.StartWithRetry()
}

// LogReq logs HTTP requests (original function)
func LogReq(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		handler.ServeHTTP(w, r)
	}
}
