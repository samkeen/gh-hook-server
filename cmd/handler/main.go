package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/samkeen/github-webhook-serverless/pkg/ghpayloads"
	"github.com/samkeen/github-webhook-serverless/pkg/templatization"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var slackWebhookUrl = os.Getenv("SLACK_WEBHOOK_URL")

const WATCH_EVENT = "watch"
const REPOSITORY_EVENT = "repository"

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	var requestPayload = request.Body
	fmt.Println(requestPayload)

	ghEvent, ok := request.Headers["X-GitHub-Event"]
	if ok {
		fmt.Printf("Github Event Type: %s\n", ghEvent)
	} else {
		fmt.Println("X-GitHub-Event header was not found")
	}

	if ghEvent == WATCH_EVENT {
		var starredEvent ghpayloads.StarredEventPayload
		if err := json.Unmarshal([]byte(requestPayload), &starredEvent); err != nil {
			fmt.Printf("There was an error unmarsheling the github event payload JSON: %s", err)
		} else {
			fmt.Printf("action: %s, Repo Name: %s\n", starredEvent.Action, starredEvent.Repository.Name)
			sendStargazeEventToSlack(starredEvent.Repository.Name,
				strconv.Itoa(starredEvent.Repository.StargazersCount),
				starredEvent.Sender.Login,
				starredEvent.Sender.HTMLURL)
		}
	} else if ghEvent == REPOSITORY_EVENT {
		var repositoryEvent ghpayloads.RepositoryEventPayload
		if err := json.Unmarshal([]byte(requestPayload), &repositoryEvent); err != nil {
			fmt.Printf("There was an error unmarsheling the github event payload JSON: %s", err)
		} else {
			fmt.Printf("action: %s, Repo Name: %d\n", repositoryEvent.Action, repositoryEvent.Repository.Name)
			sendRepoEventToSlack(
				repositoryEvent.Repository.Name,
				repositoryEvent.Action,
				repositoryEvent.Sender.Login,
				repositoryEvent.Sender.HTMLURL)
		}
	} else {
		fmt.Printf("Unrecognized Event: %s\n", ghEvent)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}

func sendStargazeEventToSlack(repo string, stars string, username string, url string) {

	var tmpl = templatization.GetSlackMessageTemplate("newStargazer")
	var renderedTemplate = templatization.ExecuteTemplate(
		tmpl,
		templatization.Payload{repo, stars, username, url})

	var postPayload = map[string]string{"text": renderedTemplate}
	postPayloadJson, _ := json.Marshal(postPayload)

	fmt.Printf("Payload to post to slack: %s\n", postPayloadJson)

	resp, err := http.Post(
		slackWebhookUrl,
		"application/json; charset=utf-8",
		bytes.NewBuffer([]byte(postPayloadJson)))
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)
	fmt.Printf("Response  Body: %s\n", bodyString)

}

func sendRepoEventToSlack(repo string, action string, username string, url string) {



}
