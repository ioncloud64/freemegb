# TODO: Create Makefile to automate deployment
GTK			?= gtk_3_20
arch		:= $(shell go env GOARCH)
os			:= $(shell go env GOOS)
version	:= $(shell cat VERSION)

.PHONY: all
# linux_i386 linux_amd64 additional gcc targets
all: clean host #cross-tools linux_i386 linux_amd64 linux_arm windows_i386 windows_amd64

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
linux_i386:
	GOOS=linux GOARCH=386 go build -v -o bin/linux_i386/freemegb
	cp -rf ui bin/linux_i386
linux_amd64:
	GOOS=linux GOARCH=amd64 go build -v -o bin/linux_amd64/freemegb
	cp -rf ui bin/linux_amd64
travisci:
	go get -t -tags gtk_$GTK_VERSION ./...
	cp -rf ui bin/freemegb-$(version)_$(arch)
host:
	go build -v -o bin/freemegb-$(version)_$(arch)/freemegb
	cp -rf ui bin/freemegb-$(version)_$(arch)
host_deb:
	mkdir -p build/freemegb-$(version)_$(arch)
	cp -rf installer_files/$(os)_$(arch)/* build/freemegb-$(version)_$(arch)
	mkdir -p build/freemegb-$(version)_$(arch)/usr/bin/
	cp bin/freemegb-$(version)_$(arch)/freemegb build/freemegb-$(version)_$(arch)/usr/bin/
	mkdir -p build/freemegb-$(version)_$(arch)/usr/share/freemegb
	cp -rf ui build/freemegb-$(version)_$(arch)/usr/share/freemegb
	sudo chown -R root:root build/freemegb-$(version)_$(arch)/
	mkdir dist
	dpkg -b build/freemegb-$(version)_$(arch)
	mv build/freemegb-$(version)_$(arch).deb dist
clean:
	rm -rf bin
	sudo rm -rf build
	sudo rm -rf dist
	rm -f *.syso
