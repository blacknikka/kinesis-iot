#!/bin/bash

aws iot create-keys-and-certificate \
    --no-set-as-active \
    --certificate-pem-outfile ./iot-cert.cert.pem \
    --public-key-outfile ./iot-motor.public.key \
    --private-key-outfile ./iot-motor.private.key

# copy files.
cp ./iot-cert.cert.pem ../../../../../producer/producer/src/cert/
cp ./iot-motor.private.key ../../../../../producer/producer/src/cert/

# download the root CA file.
wget https://www.amazontrust.com/repository/AmazonRootCA1.pem -O ./AmazonRootCA1.pem
cp ./AmazonRootCA1.pem ../../../../../producer/producer/src/cert/

