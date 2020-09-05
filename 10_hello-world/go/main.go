/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START all]
package main

import (
	"fmt"
	"log"
	"net/http"
	"net"
	"strings"
	"os"
    "time"
)

func main() {
	// use PORT environment variable, or default to 8000
	port := "8000"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", echo)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// echo responds to the request with a plain-text "Serving request" message 
// followed by some meta-data baout the environment where it is running
func echo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	addrs, err := net.LookupHost(host)
	ipaddresses := ""
	time, _ := time.Now()
	
	if err == nil {
		ipaddresses = strings.Join(addrs, " ")
	}

	fmt.Fprintf(w, "Echo Test\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
	fmt.Fprintf(w, "Host ip-address(es): %s\n", ipaddresses)
	fmt.Fprintf(w, "Current time: %s\n", time)
}
// [END all]
