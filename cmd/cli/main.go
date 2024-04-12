package main

import (
	"github.com/tobiaszuercher/vervet/config"
	"github.com/tobiaszuercher/vervet/pkg/cli"
	"github.com/tobiaszuercher/vervet/pkg/scanner"
)

func main() {
	dir := "D:\\git\\zuehlke\\platformplane\\platformplane"

	cfg := config.FromEnvironment()

	s := scanner.New(cfg)

	artifacts, err := s.Find(dir)

	cli.ArtifactTable(artifacts)

	check, err := s.Check(artifacts)

	cli.ArtifactTable(artifacts)

	if err != nil {
		return
	}

	_ = check

	if err != nil {
		return
	}

}
