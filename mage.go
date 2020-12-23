// +build mage

package main

import (
	magesh "github.com/magefile/mage/sh"
)

func Run() error {
	args := []string{
		"run",
		"client/main.go",
	}

	return magesh.RunV("go", args...)
}
