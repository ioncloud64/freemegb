# TODO: Create Makefile to automate deployment

.PHONY: all

all:
	windres --input=freemegb.rc --output=freemegb-res.syso
	go build -v -o bin/freemegb.exe -ldflags "-H windowsgui" main
	cp -rf ui bin