package main

import (
	"os"
	"log"
	"encoding/json"
	"os/exec"
)

type Account struct {
	Id			string	`json:"id"`
	IsDefault	bool	`json:"isDefault"`
	Name		string	`json:"name"`
}

func listaccounts() {
	list, err := exec.Command("az", "account", "list").Output()
	if err != nil {
        log.Fatal(err)
	}

	var account []Account

	if err := json.Unmarshal([]byte(list), &account); err != nil {
        panic(err)
	}
	
	Success(os.Stderr, "%+v", account[0])
	Warning(os.Stderr, "%+v", account[1])
	Error(os.Stderr, "%+v", account[2])
	Error(os.Stderr, "%+v", account[3])
}