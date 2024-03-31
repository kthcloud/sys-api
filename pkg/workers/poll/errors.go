package poll

import "strings"

type FailedTaskErr struct {
	Hosts []string
}

func (e *FailedTaskErr) Error() string {
	return "task failed for hosts " + strings.Join(e.Hosts, ", ")
}

func NewFailedTaskErr(hosts []string) error {
	return &FailedTaskErr{Hosts: hosts}
}

var ()
