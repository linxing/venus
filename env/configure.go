package env

func Configure() error {

	// config database
	if err := configDatabase(); err != nil {
		return err
	}

	// config redis
	if err := configRedis(); err != nil {
		return err
	}

	// config timezone
	if err := timeLocation(); err != nil {
		return err
	}

	// config services
	if err := services(); err != nil {
		//return err
		return nil
	}

	return nil
}
