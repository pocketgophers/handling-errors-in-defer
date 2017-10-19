package main

func process(ocp errorOn) (errs []error) {
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
