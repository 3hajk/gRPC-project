VERSION := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date -R)
VCS_URL := $(shell basename `git rev-parse --show-toplevel`)
VCS_REF := $(shell git log -1 --pretty=%h)
NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)


print:
    @echo VERSION=${VERSION}
    @echo BUILD_DATE=${BUILD_DATE}
    @echo VCS_URL=${VCS_URL}
    @echo VCS_REF=${VCS_REF}
    @echo NAME=${NAME}
    @echo VENDOR=${VENDOR}

up:
	@docker-compose up -d

down:
	@docker-compose down