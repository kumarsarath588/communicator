package winrm

import (
	//"bytes"
	"io"
	"os"

	"github.com/masterzen/winrm"
)

//WINClient is windows client struct
type WINClient struct {
	Address  string
	Port     int
	Username string
	Password string
	comm     *winrm.Client
}

//Connect Winrm Connection
func (winclient *WINClient) Connect() error {
	endpoint := &winrm.Endpoint{
		Host:     winclient.Address,
		Port:     winclient.Port,
		HTTPS:    false,
		Insecure: true,
	}
	c, err := winrm.NewClient(endpoint, winclient.Username, winclient.Password)
	shell, err := c.CreateShell()
	if err != nil {
		return err
	}
	err = shell.Close()
	if err != nil {
		return err
	}
	winclient.comm = c
	return nil
}

//ExecCommand execute command
func (winclient *WINClient) ExecCommand(cliCmd string) error {

	shell, err := winclient.comm.CreateShell()
	if err != nil {
		panic(err)
	}
	defer shell.Close()
	var cmd *winrm.Command
	cmd, err = shell.Execute(cliCmd)
	if err != nil {
		return err
	}
	//stdin := bytes.NewBufferString(cliCmd)
	//go io.Copy(cmd.Stdin, stdin)
	go io.Copy(os.Stdout, cmd.Stdout)
	go io.Copy(os.Stderr, cmd.Stderr)
	cmd.Wait()
	return nil
}
