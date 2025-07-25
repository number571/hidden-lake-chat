GC=go build
APPNAME=hl-chat
BINPATH=./bin
SOURCE=.

.PHONY: default build remove-std clean
default: build remove-std

build: 
	$(GC) -o $(BINPATH) $(SOURCE); \
	for arch in amd64 arm64; \
	do \
		for platform in linux windows darwin; \
		do \
			echo "build $(APPNAME)_$${arch}_$${platform}"; \
			if [[ $$platform == "windows" ]] \
			then \
				CGO_ENABLED=0 GOOS=$${platform} GOARCH=$${arch} go build -o $(BINPATH)/$(APPNAME)_$${arch}_$${platform}.exe $(SOURCE); \
			else \
				CGO_ENABLED=0 GOOS=$${platform} GOARCH=$${arch} go build -o $(BINPATH)/$(APPNAME)_$${arch}_$${platform} $(SOURCE); \
			fi; \
		done; \
	done;

remove-std:
	make -C $(BINPATH) remove-std

clean:
	make -C $(BINPATH) clean
