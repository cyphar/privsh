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

func run(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	// Check permissions before we do anything.
	perm := fi.Mode()

	// If the script is not world-readable then we have to be silent about
	// what specific error was hit. Otherwise we might accidentally start
	// dumping file information that is not safe.
	if perm&0444 != 0444 {
		silent = true
	}
	// The script _must_ be executable.
	if perm&0111 != 0111 {
		return fmt.Errorf("privsh: script not executable")
	}

	setuid := perm&os.ModeSetuid == os.ModeSetuid
	setgid := perm&os.ModeSetgid == os.ModeSetgid

	uid, gid, err := owner(fi)
	if err != nil {
		return err
	}

	fh, err := os.Open(path)
	if err != nil {
		return err
	}

	script, err := Parse(fh)
	if err != nil {
		return err
	}

	// First, drop groups.
	if err := setgroups(nil); err != nil {
		return err
	}

	// Always set group before user.
	if setgid {
		if err := setegid(gid); err != nil {
			return err
		}
	} else {
		// We are a setuid/setgid binary so our non-effective gid is the
		// (unprivileged) user which triggered this script.
		if err := setegid(os.Getgid()); err != nil {
			return err
		}
	}

	// Set user.
	if setuid {
		if err := seteuid(uid); err != nil {
			return err
		}
	} else {
		// We are a setuid/setgid binary so our non-effective uid is the
		// (unprivileged) user which triggered this script.
		if err := seteuid(os.Getuid()); err != nil {
			return err
		}
	}

	// TODO: We probably need to clean up os.Environ.
	return execve(script.Comm, []string{script.Comm}, os.Environ())
}
