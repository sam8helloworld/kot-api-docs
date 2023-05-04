package openapi

import (
	"context"

	openapiclient "github.com/sam8helloworld/kot-api-docs"
)

func configuration() *openapiclient.Configuration {
	return &openapiclient.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: openapiclient.ServerConfigurations{
			{
				URL:         "http://localhost:8001",
				Description: "Mock server (uses example data)",
			},
		},
		OperationServers: map[string]openapiclient.ServerConfigurations{},
	}
}

func contextWithAuthorization() context.Context {
	bearer := "8j9f7v4893y58rvt7nyfq2893n75tr78937n83"
	return context.WithValue(context.Background(), openapiclient.ContextAccessToken, bearer)
}
