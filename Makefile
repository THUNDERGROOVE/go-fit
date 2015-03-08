all:
	go build
release:
	rm -rf go-fit
	mkdir go-fit
	cp go-fit.exe go-fit
	cp -r template/ go-fit
	cp -r static/ go-fit
	7z a -tzip -mx9 release.zip go-fit