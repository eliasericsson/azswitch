package main

import (
    "io"
    "os"
    "os/exec"
	"strings"
	"runtime"
)

func withFilter(command string, input func(in io.WriteCloser)) []string {
    shell := os.Getenv("SHELL")
    if len(shell) == 0 {
        shell = "sh"
    }
	cmd := exec.Command(shell, "-c", command)
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("powershell", command)
	}

    cmd.Stderr = os.Stderr
    in, _ := cmd.StdinPipe()
    go func() {
        input(in)
        in.Close()
    }()
    result, _ := cmd.Output()
    return strings.Split(string(result), "\n")
}