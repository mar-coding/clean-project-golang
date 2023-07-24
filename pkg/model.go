package pkg

type Config struct {
	Address string   `yaml:"address" json:"address"`
	Domain  string   `yaml:"domain" json:"domain"`
	Origins []string `yaml:"origins" json:"origins"`
	Grpc    struct {
		Port            int    `yaml:"port" json:"port"`
		CertFilePath    string `yaml:"cert_file_path" json:"cert_file_path"`
		CertKeyFilePath string `yaml:"cert_key_file_path" json:"cert_key_file_path"`
	} `yaml:"grpc" json:"grpc"`
	Rest struct {
		Port            int    `yaml:"port" json:"port"`
		CertFilePath    string `yaml:"cert_file_path" json:"cert_file_path"`
		CertKeyFilePath string `yaml:"cert_key_file_path" json:"cert_key_file_path"`
	} `yaml:"rest" json:"rest"`
	Websocket struct {
		Port            int    `yaml:"port" json:"port"`
		CertFilePath    string `yaml:"cert_file_path" json:"cert_file_path"`
		CertKeyFilePath string `yaml:"cert_key_file_path" json:"cert_key_file_path"`
	} `yaml:"websocket" json:"websocket"`
	SocketFilePath string         `yaml:"socket_file_path" json:"socket_file_path"`
	Development    bool           `yaml:"development" json:"development"`
	GrpcClients    []*GrpcClient  `yaml:"grpc_clients" json:"grpc_clients"`
	Database       *Database      `yaml:"database" json:"database"`
	Broker         *Broker        `yaml:"broker" json:"broker"`
	Logger         *Logger        `yaml:"logger" json:"logger"`
	ExtraData      map[string]any `yaml:"extra_data" json:"extra_data"`
}

type GrpcClient struct {
	Name           string `yaml:"name" json:"name"`
	Address        string `yaml:"address" json:"address"`
	Port           int    `yaml:"port" json:"port"`
	SocketFilePath string `yaml:"socket_file_path" json:"socket_file_path"`
	CertCAFilePath string `yaml:"cert_ca_file_path" json:"cert_ca_file_path"`
}

type Database struct {
	Mongodb struct {
		URI          string `yaml:"uri" json:"uri"`
		DatabaseName string `yaml:"database_name" json:"database_name"`
	} `yaml:"mongodb" json:"mongodb"` // Mongodb URI address
	MySQL struct {
		URI          string `yaml:"uri" json:"uri"`
		DatabaseName string `yaml:"database_name" json:"database_name"`
	} `yaml:"mysql" json:"mysql"` // MySQL URI address
	Postgres struct {
		URI          string `yaml:"uri" json:"uri"`
		DatabaseName string `yaml:"database_name" json:"database_name"`
	} `yaml:"postgres" json:"postgres"` // Postgres URI address
	Redis struct {
		Address  string `yaml:"address" json:"address"`
		Username string `yaml:"username" json:"username"`
		Password string `yaml:"password" json:"password"`
		Database int    `yaml:"database" json:"database"`
	} `yaml:"redis" json:"redis"` // Redis URI address
	Elastic struct {
		Addresses []string `yaml:"addresses" json:"addresses"`
		Username  string   `yaml:"username" json:"username"`
		Password  string   `yaml:"password" json:"password"`
	}
}

type Broker struct {
	RabbitMQ string `yaml:"rabbitmq" json:"rabbitmq"` // RabbitMQ URI address
	Pulsar   struct {
		URI   string `yaml:"uri" json:"uri"`
		Token string `yaml:"token" json:"token"`
	} `yaml:"pulsar" json:"pulsar"`
	Nats struct {
		Address  string `yaml:"address" json:"address"`
		Username string `yaml:"username" json:"username"`
		Password string `yaml:"password" json:"password"`
	} `yaml:"nats" json:"nats"`
}

type Logger struct {
	ConsoleWriter    bool         `yaml:"console_writer" json:"console_writer"`
	ColorableConsole bool         `yaml:"colorable_console" json:"colorable_console"`
	JsonWriter       bool         `yaml:"json_writer" json:"json_writer"`
	JsonExtension    string       `yaml:"json_extension" json:"json_extension"`
	FileWriter       bool         `yaml:"file_writer" json:"file_writer"`
	LogPath          string       `yaml:"log_path" json:"log_path"`
	FileRotation     *LogRotation `yaml:"file_rotation" json:"file_rotation"`
	SentryDSN        string       `yaml:"sentry_dsn" json:"sentry_dsn"`
}

type LogRotation struct {
	MaxAge   int  `yaml:"max_age" json:"max_age"`
	FileSize int  `yaml:"file_size" json:"file_size"`
	Compress bool `yaml:"compress" json:"compress"`
}
