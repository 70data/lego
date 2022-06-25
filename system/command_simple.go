package system

import (
	"os/exec"
	
	"k8s.io/klog/v2"
)

func RunCmd(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		klog.Infoln(err)
	}
}
