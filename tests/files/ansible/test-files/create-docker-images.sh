#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="daxxcoin/cppjit-testrunner" docker-cppjit
docker build --tag="daxxcoin/python-testrunner" docker-python
docker build --tag="daxxcoin/go-testrunner" docker-go
