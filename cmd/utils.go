/*
	Copyright © 2020 Sankarsan Kampa

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// change the current working directory to `path`
func Cd(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Println(err.Error())
		os.Exit(126)
	}
}

// checks whether a directory exists in the `path`
func IsDir(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// create `newpath` as a symbolic link to `oldpath`.
func Ln(oldpath, newpath string) {
	if err := os.Symlink(oldpath, newpath); err != nil {
		fmt.Println(err.Error())
		os.Exit(126)
	}
}

// remove the `path` and any children it contains
func MkDir(path string) {
	if err := os.Mkdir(path, 0755); err != nil {
		fmt.Println(err.Error())
		os.Exit(126)
	}
}

// remove the `path` and any children it contains
func RmRf(path string) {
	if err := os.RemoveAll(path); err != nil {
		fmt.Println(err.Error())
		os.Exit(126)
	}
}

// convert the relative `path` to an absolute path
func Resolve(path string) string {
	abspath, err := filepath.Abs(path);
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(126)
	}
	return abspath
}

// run a shell command and stream the output to standard stream
func Exec(command string, args ...string) {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// run a `yarn` command, and fall back to `npm` if yarn fails
func ExecNode(args ...string) {
	cmd := exec.Command("yarn", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		cmd = exec.Command("npm", args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
