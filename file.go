package goutils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// RemoveOldFiles 删除给定目录下所有超过duration未更新的文件
func RemoveOldFiles(dirPath string, duration time.Duration) error {
	// 遍历目录
	return filepath.WalkDir(dirPath, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 忽略目录
		if entry.IsDir() {
			return nil
		}

		// 忽略获取文件信息失败的文件
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintln(os.Stderr, "[file] failed to get file info:", err.Error())
			return nil
		}

		// 判断文件是否需要删除
		if time.Since(info.ModTime()) < duration {
			return nil
		}

		// 删除文件, 忽略删除失败的文件
		err = os.Remove(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "[file] failed to remove file:", path, err.Error())
		}
		return nil
	})
}

// CreateAndBackupFile 创建并备份文件，返回的File对象需要手动Close
func CreateAndBackupFile(file string) (*os.File, error) {
	// 如果旧日志文件存在则重命名，重命名加上日期时间后缀，如果已经存在则覆盖
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		err = os.Rename(file, fmt.Sprintf("%s.%s", file, time.Now().Format(time.RFC3339Nano)))
		if err != nil {
			fmt.Fprintln(os.Stderr, "[file] failed to rename file:", file, err.Error())
		}
	}

	// 创建新文件，如果存在则覆盖
	return os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
}
