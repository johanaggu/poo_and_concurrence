package main

import "fmt"

//Software send notifications with interface "INotificationFactory"
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

//SMS Sender
type SMSNotification struct{}

func (n *SMSNotification) SendNotification() {
	fmt.Println("Send message with SMS...")
}

func (n *SMSNotification) GetSender() ISender {
	return &SMSNotificationSender{}
}

type SMSNotificationSender struct{}

func (s *SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (s *SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// EMAIL Sender
type EmailNotification struct{}

func (e *EmailNotification) SendNotification() {
	fmt.Println("Send message with EMAIL")
}

func (e *EmailNotification) GetSender() ISender {
	return &EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (e *EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (e *EmailNotificationSender) GetSenderChannel() string {
	return "Mailer to Go"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "EMAIL":
		return &EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("No notification type")
	}
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	method := f.GetSender().GetSenderMethod()
	fmt.Println(method)
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	sendNotification(emailFactory)
}
