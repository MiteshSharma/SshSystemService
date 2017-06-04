package domain

import (
	"net"
	"os/user"
	"io/ioutil"
	"golang.org/x/crypto/ssh"
	"github.com/MiteshSharma/SshSystemSetup/modal"
)

type SshClient interface {
	RunCommand(instance modal.InstanceDetail, cmd *modal.SSHCommand) error
}

func GetSshClient() SshClient {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.ssh/id_rsa"
	config := &ssh.ClientConfig{
		User: "msharma",
		Auth: []ssh.AuthMethod{
			PublicKeyFile(file),
		},
		HostKeyCallback: HostKeyCallbackFunc,
	}
	return &SSHMainClient{config};
}

func HostKeyCallbackFunc(hostname string, remote net.Addr, key ssh.PublicKey) error {
	return nil
}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}