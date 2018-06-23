set http_proxy=127.0.0.1:1080
set https_proxy=127.0.0.1:1080

git config http.proxy http://127.0.0.1:1080
git config https.proxy https://127.0.0.1:1080

set TEMP_DIR=D:\temp
if not exist %TEMP_DIR% ( mkdir %TEMP_DIR% )
if not exist %TEMP_DIR%\src (
    gen.bat
    move /y src %TEMP_DIR%\src
)
set CURDIR=%~dp0
set BASEDIR=%~dp0
set BASEDIR=%CURDIR%\..\..\..\..\..\..\
set GOPATH=%BASEDIR%;%TEMP_DIR%
copy /y Godeps.json.template Godeps.json
cd %CURDIR%\..
if exist vendor ( rd /q /s vendor )
godep.exe update
godep.exe save ./...
cd %CURDIR%

git config --unset http.proxy
git config --unset https.proxy