package config

type ServerConfig struct {
	AppEnv     string
	HttpServer *HttpServerConfig
	LineConfig *LineConfig
}

type LineConfig struct {
	ChannelId     string
	ChannelSecret string
	ChannelToken  string
}
type HttpServerConfig struct {
	Address    string
	ServerName string
	Mode       string
}
