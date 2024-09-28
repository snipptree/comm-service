package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

// AWS SES Based SendEmail function

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {

	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}
	// Construct the email input for SES
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{msg.To},
		},
		Message: &types.Message{
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String(msg.Data),
				},
			},
			Subject: &types.Content{
				Data: aws.String(msg.Subject),
			},
		},
		Source: aws.String(app.Mailer.FromEmail),
	}

	// Send the email using AWS SES
	_, err = app.Mailer.SESClient.SendEmail(context.TODO(), input)
	if err != nil {
		log.Printf("failed to send email: %v", err)
	}
}

// func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
// 	type mailMessage struct {
// 		From    string `json:"from"`
// 		To      string `json:"to"`
// 		Subject string `json:"subject"`
// 		Message string `json:"message"`
// 	}

// 	var requestPaylod mailMessage
// 	err := app.readJSON(w, r, &requestPaylod)
// 	if err != nil {
// 		log.Println(err)
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	msg := Message{
// 		From:    requestPaylod.From,
// 		To:      requestPaylod.To,
// 		Subject: requestPaylod.Subject,
// 		Data:    requestPaylod.Message,
// 	}

// 	err = app.Mailer.SendSMTPMessage(msg)
// 	if err != nil {
// 		log.Println(err)
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	payload := jsonResponse{
// 		Error:   false,
// 		Message: "send to" + requestPaylod.To,
// 	}

// 	app.writeJSON(w, http.StatusOK, payload)
// }
