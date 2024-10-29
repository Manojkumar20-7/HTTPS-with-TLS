package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

const(
	port =":8000"
	responseBody="Hello, TLS!"
)

func main(){
	crt,err:=tls.LoadX509KeyPair("certificates/server-certificates/server-cert.pem","certificates/server-certificates/server-key.pem")
	if err!=nil{
		log.Fatalln("Error in LoadX509KeyPair",err)
	}
	trustedCA,err:=os.ReadFile("certificates/root-ca-certificates/root-ca-cert.pem")
	if err!=nil{
		log.Fatalln("Error in reading rootCA")
	}
	rootCA:=x509.NewCertPool()
	rootCA.AppendCertsFromPEM(trustedCA)
	config:=&tls.Config{
		Certificates: []tls.Certificate{crt},
		RootCAs: rootCA,
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS13,
	}
	router:=http.NewServeMux()
	router.HandleFunc("/",handleRequest)
	server:=&http.Server{
		Addr: port,
		Handler: router,
		TLSConfig: config,
	}
	log.Println("Server is listening on localhost",port)
	err=server.ListenAndServeTLS("","")
	if err!=nil{
		log.Fatalln("Failed to start server", err)
	}
}

func handleRequest(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
	log.Println("Hello from client...")
	w.Write([]byte(responseBody))
}