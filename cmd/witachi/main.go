package main

import (
	"crypto/tls"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mpkondrashin/itachi/pkg/generate"
)

//go:embed cert.pem
var certPEM []byte

//go:embed key.pem
var keyPEM []byte

var (
	httpPort  = flag.Int("http", 8080, "HTTP port")
	httpsPort = flag.Int("https", 8443, "HTTPS port")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	for _, sample := range Samples {
		mux.HandleFunc("/"+sample.Filename, func(w http.ResponseWriter, r *http.Request) {
			var err error
			if sample.Filename == "eicar.com" {
				err = generate.ExtractEicar(w)
			} else {
				err = generate.ExtractFile(w, sample.Filename)
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}

	go func() {
		log.Printf("Starting HTTP server on :%d", *httpPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), mux); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatalf("Failed to load certificate: %v", err)
	}

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS10, // Allow older TLS versions
		MaxVersion:               tls.VersionTLS13, // Support up to TLS 1.3
		InsecureSkipVerify:       true,
		ClientAuth:               tls.NoClientCert,
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		},
		Certificates: []tls.Certificate{cert},
	}

	httpsServer := &http.Server{
		Addr:      fmt.Sprintf(":%d", *httpsPort),
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	log.Printf("Starting HTTPS server on :%d", *httpsPort)
	if err := httpsServer.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("HTTPS server error: %v", err)
	}
}
