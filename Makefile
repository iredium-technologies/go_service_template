.PHONY: all deploy

REGISTRY  = iredium
DDIR      = deploy
ODIR      = $(DDIR)/_output
VERSION   = $(shell git show -q --format=%h)
DIRS      = $(shell cd deploy && ls -d */ | grep -v "_")
PROJECT   = project
SERVICES ?= $(DIRS:/=)

DATE       = $(shell date +'%Y%m%d-%H%M%S')
BRANCH     = $(shell git rev-parse --abbrev-ref HEAD)

all: dep compile build deploy

test:
	govendor test -v +local,^program

dep:
	govendor sync -v +outside

compile:
	$(foreach var, $(SERVICES), GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(var)/bin/$(var) app/$(var)/main.go;)

$(ODIR):
	mkdir -p $(ODIR)

build:
	$(foreach var, $(SERVICES), docker build -t $(REGISTRY)/$(PROJECT)/$(var):$(VERSION) -f ./deploy/$(var)/Dockerfile .;)

deploy:
	$(PROJECT)_version=$(VERSION) docker-compose -f deploy/docker-compose.yml --project-name "$(PROJECT)" up -d

down:
	$(PROJECT)_version=$(VERSION) docker-compose -f deploy/docker-compose.yml --project-name "$(PROJECT)" down

up:
	$(PROJECT)_version=$(VERSION) docker-compose -f deploy/docker-compose.yml --project-name "$(PROJECT)" up
