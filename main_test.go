package main

import (
	"testing"

	"github.com/princeparmar/telecaller-app/controller"
	"github.com/rightjoin/fuel"
)

func TestMainMethod(t *testing.T) {
	server := fuel.NewServer()
	server.AddService(&controller.UploaderService{})
	server.Version = "1"
	server.Prefix = "api"
	server.Run()
}
