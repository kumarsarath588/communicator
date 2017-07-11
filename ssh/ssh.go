package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

//SSHClient struct
type SSHClient struct {
	Address  string
	Port     int
	Username string
	Password string
	comm     *ssh.Client
}

//Connect ssh Connection
func (sshclient *SSHClient) Connect() error {
	config := &ssh.ClientConfig{
		User: sshclient.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshclient.Password),
		},
	}

	endpoint := fmt.Sprintf("%s:%d", sshclient.Address, sshclient.Port)
	c, err := ssh.Dial("tcp", endpoint, config)
	if err != nil {
		return err
	}
	sshclient.comm = c
	return nil
}

//ExecCommand execute command
func (sshclient *SSHClient) ExecCommand(cliCmd string) error {
	s, err := sshclient.comm.NewSession()
	defer s.Close()
	if err != nil {
		return err
	}
	err = s.Run(cliCmd)
	if err != nil {
		return err
	}
	return nil
}
