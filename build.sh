#!/bin/bash
go build -o inputGPT -ldflags="-s -w" main.go darwin.go language.go userCore.go config.go utils.go renderMessages.go systray.go core.go keytap_darwin.go 
