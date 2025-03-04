package main

import (
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func main() {
	for {
		cmd := exec.Command("tasklist")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, err := cmd.Output()
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		if strings.Contains(string(out), "RobloxPlayerBeta.exe") {
			killCmd := exec.Command("taskkill", "/IM", "RobloxPlayerBeta.exe", "/F")
			killCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			killCmd.Run()
		}

		time.Sleep(5 * time.Second)
	}
}
