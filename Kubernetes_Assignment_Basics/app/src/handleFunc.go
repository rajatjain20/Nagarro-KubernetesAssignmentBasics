package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {

	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "Unknown"
	}

	fmt.Fprintf(w,
		`<!DOCTYPE html>
    <html>
    <head>
      <title>Kubernetes Assignment Basics</title>
      <style>
        body {
          background: linear-gradient(to right, #00c6ff, #0072ff);
          color: white;
          font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          height: 100vh;
          margin: 0;
          text-align: center;
        }
        h1 {
          font-size: 3em;
          text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
          margin-top: 30px;
        }
        footerX {
          font-size: 0.9em;
          color: #e0e0e0;
          position: absolute;
          bottom: 40px;
        }
        footer {
          font-size: 0.9em;
          color: #e0e0e0;
          position: absolute;
          bottom: 20px;
        }
      </style>
    </head>
    <body>
      <h1>This is my first Kubernetes basic assignment</h1>
      <footerX>POD: %s</footerX>
      <footer>created by Rajat Jain</footer>
    </body>
    </html>`, podName)
}

func getHealth(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(started)

	if duration.Seconds() > 30 && duration.Seconds() < 50 {
		w.WriteHeader(500) // this code (500) is considered as error in kubernetes probes (liveness and readiness)
		w.Write([]byte(fmt.Sprintf("Failed: %v secs elapsed since server started.", duration.Seconds())))

		log.Printf("getHealth() -> Failed: %v secs elapsed since server started.\n", duration.Seconds())
	} else {
		w.WriteHeader(200) // any code between 200 to 400 is considered as success in kubernetes probes (liveness and readiness)
		w.Write([]byte("Success"))

		log.Printf("getHealth() -> Success: %v secs elapsed since server started.\n", duration.Seconds())
	}
}
