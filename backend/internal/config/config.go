package config

type Config struct {
	Srv configSRV `toml:"srv"`
	DB  configDB  `toml:"db"`
}

type configDB struct {
	DBhost     string `toml:"host"`
	DBuser     string `toml:"user"`
	DBpassword string `toml:"passwor"`
	DBname     string `toml:"name"`
	DBport     int64  `toml:"port"`
}

type configSRV struct {
	Port int64 `toml:"port"`
}
