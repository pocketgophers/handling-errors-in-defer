package main

import (
	"errors"
	"fmt"
	"log"
)

type errorOn struct {
	open    error
	close   error
	process error
}

func (e errorOn) Open() error {
	log.Println("Open()")
	return e.open
}

func (e errorOn) Close() error {
	log.Println("Close()")
	return e.close
}

func (e errorOn) Process() error {
	log.Println("Process()")
	return e.process
}

func main() {
	log.SetFlags(log.Lshortfile)

	ocp := errorOn{
		open:    errors.New("open error"),
		close:   errors.New("close error"),
		process: errors.New("process error"),
	}

	fmt.Println("--- ignored ---")
	err := ignored(ocp)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("--- []errors ---")
	errs := withErrors(ocp)
	for _, err := range errs {
		log.Println(err)
	}

	fmt.Println("--- multiErr ---")
	err = withMultiErr(ocp)
	if err != nil {
		log.Println(err)
	}
}

func ignored(ocp errorOn) error {
	err := ocp.Open()
	if err != nil {
		return err
	}
	defer ocp.Close()

	err = ocp.Process()
	if err != nil {
		return err
	}

	return nil
}

func withErrors(ocp errorOn) (errs []error) {
	err := ocp.Open()
	if err != nil {
		return append(errs, err)
	}
	defer func() {
		err := ocp.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}()

	err = ocp.Process()
	if err != nil {
		return append(errs, err)
	}

	return nil
}

type multiErr struct {
	errs []error
}

func (me multiErr) Error() string {
	return fmt.Sprint(me.errs)
}

func (me *multiErr) append(errs ...error) error {
	me.errs = append(me.errs, errs...)
	return errors.New("return value for multiErr must be set in the first deferred function")
}

func (me *multiErr) orNil() error {
	if len(me.errs) > 0 {
		return me
	}
	return nil
}

func withMultiErr(ocp errorOn) (rerr error) {
	var errs multiErr
	defer func() {
		rerr = errs.orNil()
	}()

	err := ocp.Open()
	if err != nil {
		return errs.append(err)

	}
	defer func() {
		err := ocp.Close()
		if err != nil {
			errs.append(err)
		}
	}()

	err = ocp.Process()
	if err != nil {
		return errs.append(err)
	}

	return nil
}
