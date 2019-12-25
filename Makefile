# TODO: Create Makefile to automate deployment

arch := $(shell go env GOARCH)
os := $(shell go env GOOS)

.PHONY: all
# linux_i386 linux_amd64 additional gcc targets
all:	clean linux_arm windows_i386 windows_amd64

windows_amd64:
	windres --input=freemegb_windows.rc --output=freemegb_windows.syso
ifeq ($(arch),arm)
	PKG_CONFIG=x86_64-linux-gnu-pkg-config CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go b$
endif
	cp -rf ui bin
windows_i386:
	windres --input=freemegb_windows.rc --output=freemegb_windows.syso
ifeq ($(arch),arm)
	PKG_CONFIG=i686-linux-gnu-pkg-config CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -v -o bin/windows_386/freemegb.exe -ldflags "-H windowsgui"
endif
	cp -rf ui bin
linux_arm:
	GOOS=linux GOARCH=arm go build -v -o bin/linux_arm/freemegb
	cp -rf ui bin/linux_arm
clean:
	rm -rf bin
	rm -f *.syso
