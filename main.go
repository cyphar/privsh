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
	"fmt"
	"os"
)

var silent = false

// TODO: Actually use these.
const (
	gitCommit = ""
	version   = ""
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: privsh <script>")
		os.Exit(1)
	}

	// TODO: Need to switch this to a proper cli library.
	path := os.Args[1]
	if err := run(path); err != nil {
		if !silent {
			fmt.Fprintf(os.Stderr, "privsh: %v\n", err)
		}
		os.Exit(1)
	}
}
