cd certificates

read -p "Enter the server name: " server
read -p "Enter the server config file name: " serverconfig
read -p "Enter the CA certificate name: " rootca

openssl req -new -nodes -newkey rsa:4096 -keyout ./server-certificates/$server-key.pem -out ./server-certificates/$server.csr -config ./server-certificates/$serverconfig.cnf

openssl x509 -req -in ./server-certificates/$server.csr -copy_extensions=copy -CA ./root-ca-certificates/$rootca-cert.pem -CAkey ./root-ca-certificates/$rootca-key.pem -CAcreateserial -out ./server-certificates/$server-cert.pem -days 365 -sha256 




# openssl req -new -nodes -newkey rsa:4096 -keyout $server-key.pem -out $server.csr -config ./server-certificates/$serverconfig.cnf

# openssl x509 -req -in $server.csr -copy_extensions=copy -CA ./root-ca-certificates/$rootca-cert.pem -CAkey ./root-ca-certificates/$rootca-key.pem -CAcreateserial -out $server-cert.pem -days 365 -sha256 