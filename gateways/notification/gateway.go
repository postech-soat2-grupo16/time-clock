package notification

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
	"time-clock/entities"
)

type Gateway struct {
	notificationTopic string
	notification      *sns.SNS
}

func NewGateway(notificationClient *sns.SNS) *Gateway {
	if notificationClient == nil {
		return NewGatewayMock()
	}
	return &Gateway{
		notificationTopic: os.Getenv("NOTIFICATION_TOPIC"),
		notification:      notificationClient,
	}
}

func (g *Gateway) ClientSubscriber(user *entities.User) error {

	topicArn := g.notificationTopic
	protocol := "email"
	endpoint := user.Email

	filterPolicy := map[string]interface{}{
		"target": []string{user.Email},
	}
	filterPolicyJSON, err := json.Marshal(filterPolicy)
	if err != nil {
		fmt.Println("Error marshaling filter policy:", err)
		return err
	}

	fmt.Printf("Criando params para o user %d\n", user.ID)
	params := &sns.SubscribeInput{
		TopicArn:   aws.String(topicArn),
		Protocol:   aws.String(protocol),
		Endpoint:   aws.String(endpoint),
		Attributes: map[string]*string{"FilterPolicy": aws.String(string(filterPolicyJSON))},
	}

	fmt.Printf("Enviando Subscription de USER para o TOPICO --- %d\n", user.ID)
	resp, err := g.notification.Subscribe(params)
	if err != nil {
		fmt.Printf("Erro ao inscrever o user %d, err: %v\n", user.ID, err)
		return nil
	}

	fmt.Printf("Subscription de USER para %d enviada com sucesso: %v\n", user.ID, resp.String())

	return nil
}

func (g *Gateway) SendNotification(texto string, user *entities.User) error {

	notificationMessage := fmt.Sprintf("Relatorio do usuario %d, nome %s, registration %s",
		user.ID, user.Name, user.Registration)
	fmt.Printf("Sending message: %s\n", notificationMessage)

	//Build message
	message := &sns.PublishInput{
		TopicArn: &g.notificationTopic,
		Subject:  aws.String(fmt.Sprintf("Relatorio de ponto %s", user.Registration)),
		Message:  &notificationMessage,
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"target": {
				DataType:    aws.String("String"),
				StringValue: aws.String(user.Email),
			},
		}}

	fmt.Printf("Enviando Notificação de relatorio ID %d\n", user.ID)
	messageResult, err := g.notification.Publish(message)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem para a fila:", err)
		return nil
	}
	fmt.Printf("Notificação de relatorio para user %d enviada com sucesso: %v\n", user.ID, messageResult)

	return nil
}

func NewGatewayMock() *Gateway {
	return nil
}

type GatewayMock struct {
}

func (g GatewayMock) ClientSubscriber(user *entities.User) error {
	return nil
}

func (g GatewayMock) SendNotification(texto string, user *entities.User) error {
	return nil
}
