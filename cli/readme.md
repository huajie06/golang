# Golang examples

## Use the go module system
Steps are following
- `go mod init <module name>`
- then in the code, `import <model name>/<package name>`

## The package structure need to pay attention to
This is what i have.
```sh
.
├── cmd
│   ├── arch.go
│   └── uname.go
├── go.mod
├── go.sum
├── main.go
└── readme.md
```
`main.go` import as, so the package name needs to be the same as the folder name
```go
import (
	"coreutils/cmd"
	"fmt"
)
```


Recreate some of the [GNU utilities](https://github.com/coreutils/coreutils)

## A
- arch

## B
- ~b2sum~
- ~base32~
- base64
- basename
- basenc

## C
- cat
- chcon
- chgrp
- chmod
- chown
- chroot
- cksum
- comm
- coreutils
- cp
- csplit
- cut

## D
- date
- dd
- df
- dir
- dircolors
- dirname
- du

## E
- echo
- env
- expand
- expr

## F
- factor
- false
- fmt
- fold

## G
- groups

## H
- head
- hostid
- hostname

## I
- id
- install

## J
- join
- kill
- link
- ln
- logname
- ls
- md5sum
- mkdir
- mkfifo
- mknod
- mktemp
- mv
- nice
- nl
- nohup
- nproc
- numfmt
- od
- paste
- pathchk
- pinky
- pr
- printenv
- printf
- ptx
- pwd
- readlink
- realpath
- rm
- rmdir
- runcon
- seq
- sha1sum
- sha224sum
- sha256sum
- sha384sum
- sha512sum
- shred
- shuf
- sleep
- sort
- split
- stat
- stdbuf
- stty
- sum
- sync
- tac
- tail
- tee
- test
- timeout
- touch
- tr
- true
- truncate
- tsort
- tty
- uname
- unexpand
- uniq
- unlink
- uptime
- users
- vdir
- wc
- who
- whoami
- yes

