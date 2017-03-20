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
	"path/filepath"
	"syscall"
)

func execve(comm string, argv []string, envv []string) error {
	return syscall.Exec(comm, argv, envv)
}

func executable() (string, error) {
	return filepath.EvalSymlinks("/proc/self/exe")
}

func owner(fi os.FileInfo) (int, int, error) {
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok || stat == nil {
		return -1, -1, fmt.Errorf("could not cast stat_t")
	}
	return int(stat.Uid), int(stat.Gid), nil
}

func setegid(gid int) error {
	return syscall.Setresgid(gid, gid, gid)
}

func seteuid(uid int) error {
	return syscall.Setresuid(uid, uid, uid)
}

func setgroups(groups []int) error {
	return syscall.Setgroups(groups)
}
