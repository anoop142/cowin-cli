#!/bin/sh
# Install script for termux

install_api(){
	 pkg i termux-api
}

install_dependency(){
	pkg i golang git imagemagick
}

install_pixterm(){
	go get -u github.com/eliukblau/pixterm/cmd/pixterm
}

install_cowin_cli(){
	go get -u github.com/anoop142/cowin-cli

}


install_dependency && install_pixterm && install_cowin_cli 
install_api
