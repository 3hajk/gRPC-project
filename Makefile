VERSION := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date -R)
VCS_URL := $(shell basename `git rev-parse --show-toplevel`)
VCS_REF := $(shell git log -1 --pretty=%h)
NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)

proto-gen:
	find . -name "*.proto" -exec \
	protoc -I. \
	-I${GOPATH}/src \
	--go_out=plugins=grpc:. \
	--proto_path=. "{}" \;

print:
	@echo VERSION=${VERSION}
	@echo BUILD_DATE=${BUILD_DATE}
	@echo VCS_URL=${VCS_URL}
	@echo VCS_REF=${VCS_REF}
	@echo NAME=${NAME}
	@echo VENDOR=${VENDOR}

build-server:
	cd server && docker build --no-cache -t grpc-project/grpc-server --build-arg VERSION="${VERSION}" \
    --build-arg BUILD_DATE="${BUILD_DATE}" \
    --build-arg VCS_URL="${VCS_URL}" \
    --build-arg VCS_REF="${VCS_REF}" \
    --build-arg NAME="${NAME}" \
    --build-arg VENDOR="${VENDOR}" .

up:
	@docker-compose up -d

down:
	@docker-compose down

scale:
	@docker-compose up -d --scale server=3

scale-up:
	@docker-compose up -d --scale server=3 --scale server1=3

scale-down:
	@docker-compose up -d --scale server=1