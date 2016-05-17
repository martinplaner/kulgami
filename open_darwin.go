package main

import "os/exec"

func open(file string) {
	exec.Command("open", file).Run()
}
