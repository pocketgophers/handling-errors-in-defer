package main

import (
	"github.com/hashicorp/go-multierror"
)

func process(ocp errorOn) (rerr error) {
	err := ocp.Open()
	if err != nil {
		return multierror.Append(rerr, err)
	}
	defer func() {
		err := ocp.Close()
		if err != nil {
			rerr = multierror.Append(rerr, err)
		}
	}()

	err = ocp.Process()
	if err != nil {
		return multierror.Append(rerr, err)
	}

	return nil
}
