package runner

import (
	"fmt"
	"runtime"
)

const (
	LINUX   = "linux"
	MACOS   = "darwin"
	WINDOWS = "windows"
)

func FlushDNS() (cmd string, err error) {
	flushCMD, goos := "", runtime.GOOS
	switch runtime.GOOS {
	case LINUX:
		// service restart nscd
		flushCMD = "CentOS: service restart nscd  Ubuntu: sudo systemctl restart nscd"
	case MACOS:
		// sudo killall -HUP mDNSResponder
		goos = fmt.Sprintf("%s(MacOS)", goos)
		flushCMD = "sudo killall -HUP mDNSResponder"
	case WINDOWS:
		// ipconfig flushdns
		flushCMD = "ipconfig flushdns"
	default:
		err = fmt.Errorf("unsupported os type: %v", runtime.GOOS)
	}

	cmd = fmt.Sprintf("1.系统类型: %s 系统架构: %s\n", goos, runtime.GOARCH)
	cmd = fmt.Sprintf("%s2.根据系统类型先修改HOST文件\nWindows: C:/Windows/system32/drivers/etc/hosts\nMacOS/Linux: /etc/hosts\n", cmd)
	cmd = fmt.Sprintf("%s3.CMD刷新DNS: %s", cmd, flushCMD)
	return
}
