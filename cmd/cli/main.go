package main

import (
	"log"
	"os"

	"github.com/tobiaszuercher/vervet/config"
	"github.com/tobiaszuercher/vervet/pkg/cli"
	"github.com/tobiaszuercher/vervet/pkg/scanner"
)

func main() {
	dir := os.Args[1]

	cfg := config.FromEnvironment()

	s := scanner.New(cfg)

	artifacts, err := s.Find(dir)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := s.Check(artifacts); err != nil {
		log.Fatal(err.Error())
	}

	cli.List(artifacts)

	if !cli.Prompt() {
		os.Exit(0)
	}

	if err := scanner.Update(artifacts); err != nil {
		log.Fatal(err.Error())
	}
}
