package main

import "os/exec"

func open(file string) {
	exec.Command("xdg-open", file).Run()
}
