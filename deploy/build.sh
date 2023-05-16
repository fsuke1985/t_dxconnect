#/bin/sh

function clean () {
    rm -rf dist/bootstrap
}

GOOS=linux GOARCH=amd64 go build -o dist/bootstrap main.go
zip dist/lambda-handler.zip dist/bootstrap \
    && clean