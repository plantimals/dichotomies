#!/bin/bash
go build main.go

./main dichotomies.tsv | jq . > dichotomies.json
