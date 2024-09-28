package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

const (
	webPort = "80"
)

type Config struct {
	Mailer Mail
}

func main() {
	app := Config{
		Mailer: createMail(),
	}
	log.Println("starting communication service on: ", webPort)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func createMail() Mail {
	// This is for MailHog
	// port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	// m := Mail{
	// 	Domain:      os.Getenv("MAIL_DOMAIN"),
	// 	Host:        os.Getenv("MAIL_HOST"),
	// 	Port:        port,
	// 	Username:    os.Getenv("MAIL_USERNAME"),
	// 	Password:    os.Getenv("MAIL_PASSWORD"),
	// 	Encryption:  os.Getenv("MAIL_ENCRYPTION"),
	// 	FromName:    os.Getenv("MAIL_NAME"),
	// 	FromAddress: os.Getenv("MAIL_ADDRESS"),
	// }
	// return m

	// Load AWS SDK configuration
	staticCredentials := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("REDACTED_AWS_ACCESS_KEY_ID", "REDACTED_AWS_SECRET_ACCESS_KEY", ""))
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"), config.WithCredentialsProvider(staticCredentials))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config: %v", err)
	}

	// Create SES client
	sesClient := ses.NewFromConfig(cfg)

	return Mail{
		SESClient: sesClient,
		FromName:  "SnippTree",             //os.Getenv("MAIL_NAME"),
		FromEmail: "noreply@snipptree.com", //os.Getenv("MAIL_ADDRESS")
	}
}
