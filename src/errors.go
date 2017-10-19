package main

//!+ process
func process(ocp errorOn) (errs []error) { //!!
	//!- process
	err := ocp.Open()
	if err != nil {
		return append(errs, err) //!!
	}
	//!+ defer
	defer func() {
		err := ocp.Close()
		if err != nil {
			errs = append(errs, err) //!!
		}
	}()
	//!- defer

	err = ocp.Process()
	if err != nil {
		return append(errs, err) //!!
	}

	return nil
}
