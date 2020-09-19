#!/bin/bash
aws s3 rm s3://$BUCKET_NAME/ --recursive
aws s3 cp /front s3://$BUCKET_NAME/ --recursive
