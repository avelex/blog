#!/bin/bash

openssl genrsa -des3 -passout pass:x -out server.pass.key 2048
openssl rsa -passin pass:x -in server.pass.key -out server.key
rm server.pass.key
openssl req -new -key server.key -out server.csr \
    -subj "/C=RU/ST=Moscow/L=Moscow/O=Avelex/OU=avelex.dairy/CN=avelex.dairy"
openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt