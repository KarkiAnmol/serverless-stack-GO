package handlers

import (
	"net/http"

	user "github.com/KarkiAnmol/serverless-stack-GO/pkg/users"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	// Extract the 'email' parameter from the query string of the request
	email := req.QueryStringParameters["email"]

	// Check if the 'email' parameter is provided in the query string
	if len(email) > 0 {
		// If 'email' is provided, fetch a specific user using the FetchUser function
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			// If there's an error fetching the user, return a Bad Request response
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		// If successful, return a OK response with the user data
		return apiResponse(http.StatusOK, result)
	}

	// If 'email' is not provided, fetch all users using the FetchUsers function
	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil {
		// If there's an error fetching all users, return a Bad Request response
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}

	// If successful, return an OK response with the list of users
	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
}
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
}
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
}
func UnhandleMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
