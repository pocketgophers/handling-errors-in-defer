package main

func process(ocp errorOn) (rerr error) { //!!
	err := ocp.Open()
	if err != nil {
		return err
	}
	defer func() { //!!
		err := ocp.Close() //!!
		if err != nil {    //!!
			rerr = err //!!
		} //!!
	}() //!!

	err = ocp.Process()
	if err != nil {
		return err
	}

	return nil
}
