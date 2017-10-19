package main

func process(ocp errorOn) error {
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
