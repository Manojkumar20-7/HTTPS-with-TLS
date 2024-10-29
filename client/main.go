package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

const(
	url="https://localhost:8000"
)

func main(){
	cert,err:=os.ReadFile("certificates/root-ca-certificates/root-ca-cert.pem")
	if err!=nil{
		log.Fatalln("Error in reading cert")
	}
	caCertPool:=x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	tlsConfig:=&tls.Config{
		RootCAs: caCertPool,
	}
	tr:=&http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client:=&http.Client{
		Transport: tr,
	}
	res,err:=client.Get(url)
	if err!=nil{
		log.Fatalln("Failed to get response")
	}
	defer res.Body.Close()

	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		log.Fatalln("Failed to read response body",err)
	}
	log.Printf("Response body: %s\n",body)
}