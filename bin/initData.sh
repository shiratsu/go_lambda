#!/usr/bin/env bash

aws dynamodb \
  --region ap-northeast-1 \
  --endpoint-url http://dynamodb:8000 \
    create-table \
  --table-name SampleTable \
  --attribute-definitions \
    AttributeName=userId,AttributeType=N \
    AttributeName=userName,AttributeType=S \
  --key-schema \
    AttributeName=userId,KeyType=HASH AttributeName=userName,KeyType=RANGE \
  --billing-mode PAY_PER_REQUEST