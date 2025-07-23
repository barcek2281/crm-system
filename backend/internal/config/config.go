package config

type Config struct {
	Srv configSRV `toml:"srv"`
	DB  configDB  `toml:"db"`
	S3  configS3  `toml:"s3"`
}

type configDB struct {
	DBhost     string `toml:"host"`
	DBuser     string `toml:"user"`
	DBpassword string `toml:"password"`
	DBname     string `toml:"name"`
	DBport     int64  `toml:"port"`
}

type configSRV struct {
	Port      int64  `toml:"port"`
	SecretJws string `toml:"secret_jws"`
}

type configS3 struct {
	Host      string `toml:"host"`
	Port      int64  `toml:"port"`
	PublicUrl string `toml:"public_uri"`
	User      string `toml:"root_user"`
	Password  string `toml:"root_password"`
	SSL       bool   `toml:"ssl"`
}
