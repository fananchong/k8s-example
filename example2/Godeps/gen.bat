if exist src ( rd /q /s src )
set GOPATH=%~dp0
cd %GOPATH%\..
godep.exe restore
cd %GOPATH%
