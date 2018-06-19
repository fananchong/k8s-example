set GOPATH=%~dp0\..\..\..\..\..\
set GOBIN=%~dp0\bin

go install service_a2.go k8s_api.go
go install service_b2.go
