export CGO_ENABLED=0
export RSA_PUBLIC_PATH=/go/src/github.com/HETIC-MT-P2021/CQRSES_GROUP3/public.pem
export RSA_PRIVATE_PATH=/go/src/github.com/HETIC-MT-P2021/CQRSES_GROUP3/private.pem
export RSA_PASSWORD=password

go test -v ./... > go-cov.log