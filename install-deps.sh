#! /bin/bash

go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0
go install github.com/99designs/gqlgen@v0.17.76

# For Linux x86_64 (amd64)
rm -f ./skaffold
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/
rm -f ./skaffold
