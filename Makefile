#
# Itachi (c) 2022 by Mikhail Kondrashin (mkondrashin@gmail.com)
#
# Makefile
#

itachi_linux_amd64: main.go gmw/dropper.exe gmw/encryptor.exe gmw/spyware.exe gmw/downloader.exe gmw/novirus.exe  gmw/antiav.exe
	GOOS=linux GOARCH=amd64 go build -o itachi_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o itachi_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o itachi_darwin_arm64
	GOOS=windows GOARCH=amd64 go build -o itachi_windows_amd64.exe

gmw/dropper.exe: gmw/dropper/main.go
	GOOS=windows GOARCH=amd64 go build  -o ./gmw/dropper.exe ./gmw/dropper/*.go

gmw/encryptor.exe: gmw/encryptor/main.go
	GOOS=windows GOARCH=amd64 go build -o ./gmw/encryptor.exe  ./gmw/encryptor/*.go

gmw/spyware.exe: gmw/spyware/main.go
	GOOS=windows GOARCH=amd64 go build -o ./gmw/spyware.exe ./gmw/spyware/*.go

gmw/downloader.exe: gmw/downloader/main.go
	GOOS=windows GOARCH=amd64 go build -o ./gmw/downloader.exe ./gmw/downloader/*.go

gmw/novirus.exe: gmw/novirus/main.go
	GOOS=windows GOARCH=amd64 go build -o ./gmw/novirus.exe ./gmw/novirus/*.go

gmw/antiav.exe: gmw/antiav/main.go
	GOOS=windows GOARCH=amd64 go build -o ./gmw/antiav.exe ./gmw/antiav/*.go
