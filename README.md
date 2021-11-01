# ts-common
A repository to contain common utility code for timelyship.com

go test  -v -race github.com/mahmudaZaman/commonutil/test -coverpkg=./... -coverprofile=coverage.out
go tool cover -html=coverage.out

mockery --all
golint -set_exit_status $(go list ./... | grep -v /vendor/)
golangci-lint run
