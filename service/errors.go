package service

import "fmt"

var (
	NotImplementedErr    = fmt.Errorf("not implemented")
	BadDiscoveryTokenErr = fmt.Errorf("bad discovery token")
	BadRequestErr        = fmt.Errorf("bad request")
)
