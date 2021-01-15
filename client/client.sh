#!/bin/bash
go run client/gClient.go -port=server:8080 <<EOD
2 + 2
EOD