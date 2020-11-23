package main

import (
	"io"
	"log"
	"fmt"
	"os/exec"
	"encoding/json"
)

type Account struct {
	Id			string	`json:"id"`
	IsDefault	bool	`json:"isDefault"`
	Name		string	`json:"name"`
}

func main() {	
	filtered := withFilter("fzf", func(in io.WriteCloser) {
		list, err := exec.Command("az", "account", "list").Output()
		if err != nil {
			log.Fatal(err)
		}

		var account []Account
		if err := json.Unmarshal([]byte(list), &account); err != nil {
			panic(err)
		}

		for _, account := range account {
			fmt.Fprintf(in, "%s: %s\n", account.Id, account.Name)
		}
	})
	setaccount(filtered)
}