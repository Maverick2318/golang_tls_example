package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "hello world!")
}

// our main function
func main() {
    // Create a CA certificate pool and add cert.pem to it
    caCert, err := ioutil.ReadFile("ca.pem")
    if err != nil {
        log.Fatal(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    // Create the TLS Config with the CA pool and enable Client certificate validation
    tlsConfig := &tls.Config{
        ClientCAs: caCertPool,
        //ClientAuth: tls.RequireAndVerifyClientCert,
        InsecureSkipVerify: true, //Equivalent to verify server hostname = false
    }

    tlsConfig.BuildNameToCertificate()

    router := mux.NewRouter()
    router.HandleFunc("/hello", hello).Methods("GET")

    server := &http.Server{
                Addr:      ":9999",
                Handler:   router,
                TLSConfig: tlsConfig,
    }
    log.Fatal(server.ListenAndServeTLS("server.pem", "server-key.pem"))
}
