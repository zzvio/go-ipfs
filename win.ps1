cd cmd/ipfs
$Env:CGO_CFLAGS="-I${PWD}"
go build -o goipfs.dll -buildmode=c-shared
cp goipfs.dll C:\Users\icte-pa-03\.java\packages\lib
