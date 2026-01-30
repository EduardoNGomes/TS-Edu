package env

func CreateEnvConfig(dir string) error {
	if err := createEnvDIR(dir); err != nil {
		return err
	}
	if err := createEnvFile(); err != nil {
		return err
	}

	return nil
}
