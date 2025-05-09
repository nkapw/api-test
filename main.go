// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	// Port dari environment variable atau default 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Handler untuk root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Menerima request dari: %s", r.RemoteAddr)
		
		// Mendapatkan hostname
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		
		// Waktu server
		currentTime := time.Now().Format(time.RFC1123)
		
		// HTML response
		html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hello World PaaS App</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					max-width: 800px;
					margin: 0 auto;
					padding: 20px;
					line-height: 1.6;
					color: #333;
				}
				.container {
					background-color: #f9f9f9;
					border-radius: 5px;
					padding: 20px;
					box-shadow: 0 2px 4px rgba(0,0,0,0.1);
				}
				.header {
					background-color: #4CAF50;
					color: white;
					padding: 10px 20px;
					border-radius: 5px 5px 0 0;
					margin-bottom: 20px;
				}
				.info-row {
					display: flex;
					justify-content: space-between;
					border-bottom: 1px solid #eee;
					padding: 8px 0;
				}
				.label {
					font-weight: bold;
					color: #555;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Hello World!</h1>
				</div>
				<div class="content">
					<p>Selamat datang di aplikasi contoh untuk PaaS Container dengan Load Balancing Round-Robin.</p>
					
					<h2>Informasi Server:</h2>
					<div class="info-row">
						<span class="label">Hostname:</span>
						<span>%s</span>
					</div>
					<div class="info-row">
						<span class="label">Waktu Server:</span>
						<span>%s</span>
					</div>
					<div class="info-row">
						<span class="label">Go Version:</span>
						<span>%s</span>
					</div>
					<div class="info-row">
						<span class="label">OS/Arch:</span>
						<span>%s/%s</span>
					</div>
					<div class="info-row">
						<span class="label">Remote Address:</span>
						<span>%s</span>
					</div>
					<div class="info-row">
						<span class="label">Request Path:</span>
						<span>%s</span>
					</div>
					<div class="info-row">
						<span class="label">User Agent:</span>
						<span>%s</span>
					</div>
					
					<h2>Headers Request:</h2>
					<pre style="background-color: #f0f0f0; padding: 10px; border-radius: 3px; overflow-x: auto;">
%s
					</pre>
				</div>
			</div>
		</body>
		</html>
		`, hostname, currentTime, runtime.Version(), runtime.GOOS, runtime.GOARCH, 
		r.RemoteAddr, r.URL.Path, r.UserAgent(), formatHeaders(r.Header))
		
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	// Handler untuk health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"healthy"}`)
	})

	// Route untuk menunjukkan info
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		info := fmt.Sprintf(`
		{
			"app": "hello-world",
			"hostname": "%s",
			"platform": "%s",
			"go_version": "%s"
		}
		`, hostname, runtime.GOOS, runtime.Version())
		
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, info)
	})

	// Jalankan web server
	serverAddr := ":" + port
	log.Printf("Server berjalan di http://localhost%s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

// Format headers HTTP untuk ditampilkan
func formatHeaders(headers http.Header) string {
	result := ""
	for name, values := range headers {
		for _, value := range values {
			result += fmt.Sprintf("%s: %s\n", name, value)
		}
	}
	return result
}
