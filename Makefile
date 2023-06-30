#
# Itachi (c) 2022 by Michael Kondrashin (mkondrashin@gmail.com)
# Copyright under MIT Lincese. Please see LICENSE file for details
#
#
# Makefile
#

EXE=gmw/dropper.exe gmw/encryptor.exe gmw/spyware.exe gmw/downloader.exe gmw/antiav.exe gmw/autorun.exe gmw/novirus.exe gmw/antivm.exe 

itachi_linux_amd64: main.go $(EXE)
	GOOS=linux GOARCH=amd64 go build -o itachi_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o itachi_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o itachi_darwin_arm64
	GOOS=windows GOARCH=amd64 go build -o itachi_windows_amd64.exe

gmw/dropper.exe: gmw/dropper/main.go
	GOOS=windows GOARCH=amd64 go build  -ldflags "-s -w" -o ./gmw/dropper.exe ./gmw/dropper/*.go

gmw/encryptor.exe: gmw/encryptor/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./gmw/encryptor.exe  ./gmw/encryptor/*.go

gmw/spyware.exe: gmw/spyware/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./gmw/spyware.exe ./gmw/spyware/*.go

gmw/downloader.exe: gmw/downloader/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./gmw/downloader.exe ./gmw/downloader/*.go

gmw/autorun.exe: gmw/autorun/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./gmw/autorun.exe ./gmw/autorun/*.go

gmw/antiav/AvList.txt.gz: gmw/antiav/AvList.txt
	gzip -c gmw/antiav/AvList.txt > gmw/antiav/AvList.txt.gz

gmw/antiav.exe: gmw/antiav/main.go gmw/antiav/AvList.txt.gz
	#curl https://raw.githubusercontent.com/AV1080p/AvList/master/AvList.txt --output gmw/antiav/AvList.txt
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./gmw/antiav.exe ./gmw/antiav/*.go

gmw/antivm.exe: gmw/antivm/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./gmw/antivm.exe ./gmw/antivm/*.go

gmw/novirus.exe: gmw/novirus/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./gmw/novirus.exe ./gmw/novirus/*.go

.PHONY: clean

clean:
	rm itachi_linux_amd64 itachi_darwin_amd64 itachi_darwin_arm64 itachi_windows_amd64.exe $(EXE)
