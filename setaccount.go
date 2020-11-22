package main

import (
	"os"
	"os/exec"
	"strings"
	"log"
)

func setaccount(id []string) {
	account := strings.Join(id, "")
	var acc = strings.Split(account, ": ")
	_, err := exec.Command("az", "account", "set", "-s", acc[0]).Output()
	if err != nil {
		log.Fatal(err)
	}
	Success(os.Stderr, "%+v", acc[1])
}