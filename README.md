# golang-cep
[![Go Report Card](https://goreportcard.com/badge/erozario/golang-cep)](https://goreportcard.com/report/erozario/golang-cep)
[![Docker Build Status](https://img.shields.io/docker/build/eduardoagrj/golang-cep.svg)](https://hub.docker.com/r/eduardoagrj/golang-cep/builds/)

Dockerized API for address search based on the informed CEP

## Requirements

- Docker version 17.05 or later.

## Getting started

`docker run -it -p 8080:8080 --name server -d eduardoagrj/golang-cep`

## Build

`docker build -t eduardoagrj/golang-cep .`

## Usage

`curl -XGET http://localhost:8080/addresses/06815250`
