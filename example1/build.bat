set GOPATH=%~dp0
set GOBIN=%~dp0\bin

go install service_a.go
go install service_b.go
go install client.go