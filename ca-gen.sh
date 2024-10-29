cd certificates

read -p "Enter the CA name: " rootca
read -p "Enter the CA config file name: " caconfig

openssl genrsa -out ./root-ca-certificates/$rootca-key.pem 4096

openssl req -x509 -new -nodes -key ./root-ca-certificates/$rootca-key.pem -sha256 -days 365 -config ./root-ca-certificates/$caconfig.cnf -out ./root-ca-certificates/$rootca-cert.pem
