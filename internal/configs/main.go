package configs

func CreateConfigFiles() error {
	if err := createTSConfig(); err != nil {
		return err
	}

	if err := createVitestConfigs(); err != nil {
		return err
	}

	if err := createBiome(); err != nil {
		return err
	}

	return nil

}
