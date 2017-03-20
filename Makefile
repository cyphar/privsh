# privsh: run setuid interpreted shell scripts
# Copyright (C) 2017 Aleksa Sarai
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

GO ?= go
PROJECT := github.com/cyphar/privsh
PREFIX ?= /usr
LDFLAGS ?=

# Version information.
VERSION := $(shell cat ./VERSION)
COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)
COMMIT := $(if $(shell git status --porcelain --untracked-files=no),"${COMMIT_NO}-dirty","${COMMIT_NO}")
LDFLAGS += -s -w -X main.gitCommit=${COMMIT} -X main.version=${VERSION}

# Conditional static compilation.
ifeq ($(STATIC),1)
LDFLAGS += -extldflags '-static'
GO := CGO_ENABLED=0 $(GO)
endif

privsh: $(wildcard *.go)
	$(GO) build -ldflags "$(LDFLAGS)" -o $@ $(PROJECT)

clean:
	rm -f privsh

install: privsh
	install -m4755 -D privsh $(PREFIX)/bin/privsh

uninstall:
	rm -f $(PREFIX)/bin/privsh
