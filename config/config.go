package config

type Config struct {
	Server   *Server
	Database *Database
}

func Load() (*Config, error) {
	c := &Config{
		Database: &Database{},
		Server:   &Server{},
	}
	c.loadVars()

	return c, nil
}

func (c *Config) loadVars() {
	c.Database.Load()
	c.Server.Load()
}