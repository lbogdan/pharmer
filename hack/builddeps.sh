#!/bin/bash

# https://github.com/ellisonbg/antipackage
pip install git+https://github.com/ellisonbg/antipackage.git#egg=antipackage
pip install pyyaml

go get -u github.com/jteeuwen/go-bindata/...
