#!/bin/bash

aws iot create-keys-and-certificate \
    --no-set-as-active \
    --certificate-pem-outfile ./iot-cert.cert.pem \
    --public-key-outfile ./iot-motor.public.key \
    --private-key-outfile ./iot-motor.private.key
