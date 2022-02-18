package config

type configuration struct {
	UrlConnection string
	UserAgents    []string
}

type Config *configuration

func Boot() Config {
	return &configuration{
		UrlConnection: getEnv("PROXY_TOR_CLIENT_CONNECTION_URL", urlConnection),
		UserAgents: getEnvAsSlice("PROXY_TOR_CLIENT_USER_AGENTS", []string{
			userAgent,
		}, "|"),
	}
}
