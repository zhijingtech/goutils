package goutils

import (
	"os"
	"path"
	"time"
)

type StdErrLogger struct {
	file *os.File
}

func NewStdErrLogger(file string, keep time.Duration) (*StdErrLogger, error) {
	if keep <= 0 {
		keep = time.Hour * 24 * 7 // 默认保留7天
	}
	// 删除超过指定天数的日志文件
	err := RemoveOldFiles(path.Dir(file), keep)
	if err != nil {
		return nil, err
	}
	// 创建新文件，如果存在则备份
	f, err := CreateAndBackupFile(file)
	if err != nil {
		return nil, err
	}
	// 重定向标准错误输出到文件
	err = RedirectStdErrToFile(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &StdErrLogger{file: f}, nil
}

func (l *StdErrLogger) Close() error {
	if l.file == nil {
		return nil
	}
	return l.file.Close()
}
