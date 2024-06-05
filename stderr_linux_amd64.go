package main

import (
	"os"
	"syscall"
)

func RedirectStdErrToFile(f *os.File) error {
	return syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd())) // 将stderr重定向到文件，比如捕获panic
}
