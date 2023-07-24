package commands

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func init() {
	//rootCmd.AddCommand(runCmd)
}

//var runCmd = &cobra.Command{
//	Use:   "run",
//	Short: "run feedback service",
//	RunE: func(cmd *cobra.Command, args []string) error {
//cfg, err := config.NewConfig(configPath)
//if err != nil {
//	cmd.PrintErrln(err)
//}
//
//grpcServer, err := server.NewGrpcServer(cfg, info, errHandler, os.Getenv("PUBLIC_SECRET"), os.Getenv("PRIVATE_SECRET"), os.Getenv("GOOGLE_CAPTCHA_SECRET_KEY"), zapper, cfg.Development)
//if err != nil {
//	cmd.PrintErrln(err)
//}
//
//httpServer, err := server.NewHttpServer(cmd.Context(), cfg.Address, strconv.Itoa(cfg.Rest.Port), strconv.Itoa(cfg.Grpc.Port), cfg.Rest.CertFilePath, cfg.Rest.CertKeyFilePath, cfg.Development, cfg.Origins...)
//if err != nil {
//	cmd.PrintErrln(err)
//}
//
//client := client.NewClient(zapper, cfg.GrpcClients...)
//
//private, err := loadRSAKey()
//if err != nil {
//	cmd.PrintErrln(err)
//}
//
//elasticsearch, err := elastic.NewElasticSearch(cmd.Context(), cfg.Database.Elastic.Addresses, cfg.Database.Elastic.Username, cfg.Database.Elastic.Password, cfg.Development)
//if err != nil {
//	cmd.PrintErrln(err)
//}
//
//application, err := app.New(cmd.Context(),
//	grpcServer, httpServer,
//	client,
//	cfg, info, errHandler, zapper, jwt, redisClient, private, mailgun.New(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_PRIVATEKEY"), os.Getenv("MAILGUN_SENDER_EMAIL"), os.Getenv("MAILGUN_BASE_URL"), errHandler), elasticsearch)
//if err != nil {
//	cmd.PrintErrln(err)
//}
//
//application.Run(cmd.Context())
//return nil
//},
//}

func loadRSAKey() (*rsa.PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(os.Getenv("RSA_PRIVATE"))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("private key is invalid")
	}

	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		return nil, fmt.Errorf("unknown key type %q, want %q", got, want)
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return private, nil
}
