
all: build

setup:
	go get -u -d github.com/skyway/skyway-webrtc-gateway || true

bin/swagger:
	mkdir -p bin
	curl -sL "https://github.com/go-swagger/go-swagger/releases/download/0.16.0/swagger_linux_amd64" --output $@
	chmod +x $@


swagger.yaml:
	sed -e 's|1\.peers|peers|g' \
	    -e 's|2\.data|data|g' \
	    -e 's|3\.media|media|g' \
		$(shell go list -e -f "{{.Dir}}" github.com/skyway/skyway-webrtc-gateway/api)/api.yaml > swagger.yaml

build: bin/swagger swagger.yaml
	./bin/swagger generate client -f swagger.yaml

clean:
	rm -rf bin
