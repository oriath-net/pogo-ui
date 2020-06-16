PROJECT = pogo-ui

export GOOS=windows
export GOARCH=amd64

$(PROJECT).exe: $(PROJECT).syso
	go build -ldflags="-H windowsgui"

$(PROJECT).syso: $(PROJECT).manifest
	rsrc -manifest $(PROJECT).manifest -o $(PROJECT).syso

.PHONY: $(PROJECT).exe
