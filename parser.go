/*
 * privsh: run setuid interpreted shell scripts
 * Copyright (C) 2017 Aleksa Sarai
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Script stores information about a parsed privsh script. Currently privsh
// scripts are only valid if they have this precise syntax:
//
//   #!<path to privsh>
//   #?<path to script>
//
// Trailing lines or extra spaces are forbidden.
type Script struct {
	Comm string
}

// Remove blank lines.
func filterLines(lines [][]byte) [][]byte {
	var filtered [][]byte
	for _, line := range lines {
		if len(line) > 0 {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

// Parse the given io.Reader as described in the Script docstring. Returns an
// error if the parsing failed in some way.
func Parse(r io.Reader) (Script, error) {
	file, err := ioutil.ReadAll(r)
	if err != nil {
		return Script{}, err
	}

	lines := bytes.Split(file, []byte("\n"))
	lines = filterLines(lines)

	// Don't allow us to mis-parse a script that wasn't meant for us.
	if len(lines) != 2 {
		return Script{}, fmt.Errorf("invalid syntax: must only contain two lines")
	}

	// Don't allow cases where the first line isn't what we would expect.
	execName, err := executable()
	if err != nil {
		return Script{}, err
	}
	interpExpected := fmt.Sprintf("#!%s", execName)
	if !bytes.Equal(lines[0], []byte(interpExpected)) {
		return Script{}, fmt.Errorf("interpreter line is not #!<privsh>")
	}

	// Parse the script name.
	var comm string
	if !bytes.HasPrefix(lines[1], []byte("#?")) {
		return Script{}, fmt.Errorf("invalid syntax: script line doesn't start with #?")
	}
	comm = strings.TrimPrefix(string(lines[1]), "#?")

	// Don't allow spaces in the script name -- in the future I'll probably
	// implement argument splitting but for now scripts don't get any
	// arguments.
	if strings.Contains(comm, " ") {
		return Script{}, fmt.Errorf("invalid syntax: script line contains spaces")
	}

	// Don't allow relative paths. We don't trust $PATH.
	if !filepath.IsAbs(comm) {
		return Script{}, fmt.Errorf("script is not an absolute path")
	}

	return Script{
		Comm: comm,
	}, nil
}
