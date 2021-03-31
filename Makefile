API_VERSION?=0.3.2

all: build

setup: bin/swagger

bin/swagger:
	mkdir -p bin
	curl -sL "https://github.com/go-swagger/go-swagger/releases/download/0.16.0/swagger_linux_amd64" --output $@
	chmod +x $@


.PHONY: swagger.yaml
swagger.yaml:
	curl -sL "https://raw.githubusercontent.com/skyway/skyway-webrtc-gateway/${API_VERSION}/api/api.yaml" | \
	sed -e 's|1\.peers|peers|g' \
	    -e 's|2\.data|data|g' \
	    -e 's|3\.media|media|g' \
		-e 's/\r//g' \
		> swagger.yaml

build: setup swagger.yaml
	./bin/swagger generate client -f swagger.yaml --name gateway

clean:
	rm -rf bin
