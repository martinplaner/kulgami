package main

import "os/exec"

func open(file string) {
	exec.Command("start", file).Run()
}
