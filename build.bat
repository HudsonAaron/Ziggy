rem 打包成linux可执行文件
set goos=linux
set goarch=386
go build -o bin/ziggy

rem 打包成windows可执行文件
rem set goos=windows
rem set goarch=amd64
rem go build -o bin/ziggy.exe

rem 打包成mac可执行文件
rem set goos=darwin
rem set goarch=amd64
rem go build -o bin/ziggy