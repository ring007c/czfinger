package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func IsFile(path string) bool {
	return !IsDir(path)
}
func ReadLineFile(path string) (error, []string) {
	lineSlice := make([]string, 0)

	// 判断所给路径是否为文件
	if !IsFile(path) {
		return os.ErrNotExist, nil
	}
	// 打开文件
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err, nil
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		lineSlice = append(lineSlice, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err, nil
			}
		}
	}
	return nil, lineSlice

}
