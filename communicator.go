package communicator

import (
	"github.com/kumarsarath588/communicator/ssh"
	"github.com/kumarsarath588/communicator/winrm"
)

//Host host details
type Host struct {
	Address  string
	Port     int
	Username string
	Password string
	Kind     string
}

//Client interface
type Client interface {
	Connect() error
	ExecCommand(cliCmd string) error
}

//New client creation
func New(h *Host) Client {
	if h.Kind == "ssh" {
		client := &ssh.SSHClient{
			Address:  h.Address,
			Username: h.Username,
			Password: h.Password,
			Port:     h.Port,
		}
		return client
	} else if h.Kind == "winrm" {
		client := &winrm.WINClient{
			Address:  h.Address,
			Username: h.Username,
			Password: h.Password,
			Port:     h.Port,
		}
		return client
	}
	return nil
}
