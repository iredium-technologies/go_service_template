# Golang REST API project template

## Description

Golang REST API project template

## Onboarding and Development Guide

### Installation
0. Depedency
  ```
  install golang -> https://golang.org/doc/install
  get govendor -> go get github.com/kardianos/govendor
  ```
1. Clone `go_service_template` to your $GOPATH:
  ```
  go get github.com/iredium_technologies/go_service_template
  ```
2. Rename and re git init go_service_template to your [project_name]
  ```
  mv $GOPATH/src/github.com/iredium_technologies/go_service_template $GOPATH/src/github.com/[your_repo]/[project_name]
  cd $GOPATH/src/github.com/[your_repo]/[project_name]
  sudo rm -rf .git
  git init
  git add .
  git commit -m 'first commit'
  ```
3. Go to [project_name] directory, then install dependency
  ```
  cd $GOPATH/src/github.com/[your_repo]/[project_name]
  make dep
  ```
4. Compile, build and run the server
  ```
  make compile
  make build
  make deploy
  ```
  or
  ```
  make all
  ```