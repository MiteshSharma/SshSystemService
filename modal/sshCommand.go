package modal

import "io"

type SSHCommand struct {
	Id     string 	 `json:"id" bson:"_id"`
	Path   []string  `json:"paths" bson:"paths"`
	Env    []string  `json:"envs" bson:"envs"`
	Stdin  io.Reader `json:"-"`
	Stdout io.Writer `json:"-"`
	Stderr io.Writer `json:"-"`
}

func (sc SSHCommand) IsValid() bool {
	if (sc.Path != nil) {
		return true
	}
	return false
}