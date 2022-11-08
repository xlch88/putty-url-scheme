@echo off
set CGO_ENABLED=0
rsrc -manifest putty-url-scheme.exe.manifest -o putty-url-scheme.exe.syso
go build -ldflags="-H windowsgui -w -s" .
