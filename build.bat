@echo off
set CGO_ENABLED=0
go build -ldflags="-H windowsgui -w -s" .
