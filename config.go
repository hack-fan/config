package config

// Load config to `dst` struct pointer from shell env variables and docker secrets.
func Load(dst interface{}) error {
	return LoadEnvAndDockerSecret(dst)
}

// MustLoad load config to `dst` struct pointer from shell env variables and docker secrets.
// It panic when an error occurs.
func MustLoad(dst interface{}) {
	err := Load(dst)
	if err != nil {
		panic(err)
	}
}

// LoadEnvAndSecret load config to `dst` struct pointer from shell env variables and container secrets.
func LoadEnvAndSecret(dst interface{}, secretPath string) error {
	l := loader{
		Env:    true,
		Secret: true,
		Path:   secretPath,
	}
	return l.load(dst)
}

// LoadEnvAndDockerSecret load config to `dst` struct pointer from shell env variables and docker secrets.
func LoadEnvAndDockerSecret(dst interface{}) error {
	return LoadEnvAndSecret(dst, "/run/secrets")
}

// LoadEnvAndKubernetesSecret load config to `dst` struct pointer from shell env variables and kubernetes secrets.
func LoadEnvAndKubernetesSecret(dst interface{}) error {
	return LoadEnvAndSecret(dst, "/etc/secret-volume")
}

// LoadEnv load config to `dst` struct pointer from shell env variables only
func LoadEnv(dst interface{}) error {
	l := loader{
		Env:    true,
		Secret: false,
	}
	return l.load(dst)
}
