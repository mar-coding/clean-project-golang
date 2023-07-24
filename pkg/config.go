package pkg

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/credentials"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

const (
	yamlExt      = ".yaml"
	shortYamlExt = ".yml"
	jsonExt      = ".json"
)

// New load config file from configPath and return Config object
//
//	example:
//		config, err := New("./config.yaml")
func New(configPath string) (*Config, error) {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.Join(errors.New("service config: failed to read config file "), err)
	}

	switch filepath.Ext(configPath) {
	case yamlExt, shortYamlExt:
		return unmarshalYaml(b)
	case jsonExt:
		return unmarshalJson(b)
	}

	return nil, errors.New("service config: config path is invalid")
}

// LoadGrpcServerCredentials create transport credential for grpc server for TLS handshake
func (c *Config) LoadGrpcServerCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(c.Grpc.CertFilePath, c.Grpc.CertKeyFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("service config: load x509 key pair got error %s", err.Error()))
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

// LoadGrpcClientCredentials create transport credential for grpc client to access server side for TLS handshake
func (c *Config) LoadGrpcClientCredentials(client *GrpcClient) (credentials.TransportCredentials, error) {
	if client == nil {
		return nil, errors.New("service config: client is nil")
	}

	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(client.CertCAFilePath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, errors.New("service config: failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func unmarshalYaml(payload []byte) (*Config, error) {
	cfg := new(Config)
	if err := yaml.Unmarshal(payload, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func unmarshalJson(payload []byte) (*Config, error) {
	cfg := new(Config)
	if err := json.Unmarshal(payload, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
