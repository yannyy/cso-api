set -x
app_name=api-gateway
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
docker build -f ./Dockerfile . -t ${app_name}:`git describe --always`
rm ${app_name}