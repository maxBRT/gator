package config

func (c *Config) SetUser(username string) error {
	// Set the username in the config
	configData, err := ReadConfig()
	if err != nil {
		return err
	}
	if configData.DBURL == "" {
		configData.DBURL = db_url
	} else {
		c.DBURL = configData.DBURL
	}
	c.USERNAME = username
	writeConfig(c)
	return nil
}
