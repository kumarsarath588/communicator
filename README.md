# communicator
golang communicator

Example:
-------
```package main

import (
	"fmt"
	"os"

	"github.com/kumarsarath588/communicator"
)

func main() {
	host := &communicator.Host{
		Address:  "10.7.111.162",
		Port:     5985,
		Username: "administrator",
		Password: "nutanix/4u",
		Kind:     "winrm",
	}
	c := communicator.New(host)
	if c == nil {
		fmt.Println("Unknown kind of communication").
		os.Exit(0)/
	}
	err := c.Connect()
	fmt.Println("Connecting to host:", host.Address)
	if err != nil {
		fmt.Println("Error Connecting Server:", host.Address, err)
		panic(err)
	}
	fmt.Println("Executing command")
	err = c.ExecCommand("ipconfig /all")
	if err != nil {
		fmt.Println("Error Executing Command:", err)
		panic(err)
	}
}```
