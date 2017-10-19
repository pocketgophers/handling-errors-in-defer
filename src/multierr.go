package main

import (
	"errors"
	"fmt"
)

//!+ multiErr
type multiErr struct {
	Errors []error
}

func (me multiErr) Error() string {
	return fmt.Sprint(me.Errors)
}

//!- multiErr

//!+ append
func (me *multiErr) Append(errs ...error) error {
	me.Errors = append(me.Errors, errs...)
	return errors.New("return value for multiErr must be set in the first deferred function")
}

//!- append

//!+ return
func (me *multiErr) Return(rerr *error) {
	if len(me.Errors) > 0 {
		*rerr = me
	}
}

//!- return

func process(ocp errorOn) (rerr error) { //!!
	var errs multiErr        //!!
	defer errs.Return(&rerr) //!!

	err := ocp.Open()
	if err != nil {
		return errs.Append(err) //!!
	}
	defer func() {
		err := ocp.Close()
		if err != nil {
			errs.Append(err) //!!
		}
	}()

	err = ocp.Process()
	if err != nil {
		return errs.Append(err) //!!
	}

	return nil
}
