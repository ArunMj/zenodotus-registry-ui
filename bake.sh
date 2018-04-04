#!/bin/bash

set -e

function clean_all(){
    rm -rf build
    rm -rf dist
    mkdir -p build/src
    mkdir -p dist
}

function build_ui(){
    echo -e "\n~~~~~~~~~~~~~~~~~~ build ui ~~~~~~~~~~~~~~~~~~~"
    cd ui
    npm run build
    echo -e "\n~~~~~ copy static assets to build dir ~~~~~~~~~"
    cd -
    mv ui/build build/src/web
}


function dist_binary(){
    cp -r backend/* build/src/
    cd build/src
    echo -e "n~~~~~~ embed static assets to binary ~~~~~~~~~~~"
    rice embed-go -v -i .
    echo -e  "\n~~~~~~~~~~~~~~~~ build binary ~~~~~~~~~~~~~~~~~~"
    go build -o ../../dist/zenodotus .
    cd -
}

clean_all
build_ui
dist_binary