// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "os/exec"

func open(file string) {
	exec.Command("xdg-open", file).Run()
}
