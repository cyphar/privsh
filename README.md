## `privsh` ##
[![Release](https://img.shields.io/github/release/cyphar/privsh.svg)](CHANGELOG.md)
[![License](https://img.shields.io/github/license/cyphar/privsh.svg)](COPYING)

A fairly dodgy interpreter which will allow you to run setuid shell scripts.
Due to various race conditions, setuid shell scripts don't actually run as
setuid on Linux (on BSDs they do because of the `/dev/fd` interface resolving
the race conditions mentioned).

As a result if you wanted to hack together a shell script and make it setuid,
you were out of luck. Until now.

### License ###

`privsh` is licensed under the terms of the GNU GPLv3 or later.

```
privsh: run setuid interpreted shell scripts
Copyright (C) 2017 Aleksa Sarai

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```
