package controller

import (
	"template-instance-operator/pkg/controller/kscinstance"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kscinstance.Add)
}
