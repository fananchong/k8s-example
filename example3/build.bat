set GOPATH=%~dp0\..\..\..\..\..\
set GOBIN=%~dp0\bin

go install -tags debug service_a.go k8s_api_debug.go
go install -tags debug service_b.go k8s_api_debug.go
