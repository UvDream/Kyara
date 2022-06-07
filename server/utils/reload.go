package utils

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

// Reload 刷新程序
func Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New("系统不支持")
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
