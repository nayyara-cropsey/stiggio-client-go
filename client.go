// Package stigg contains client code for Stigg GraphQL APi
// This file is modified from https://github.com/stiggio/api-client-go/blob/master/client.go official Go SDK for Stigg

package stigg

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/hashicorp/go-retryablehttp"
)

const RetryCount = 3

func NewStiggClient(apiKey string, httpClient *http.Client, baseUrl *string) StiggClient {
	if baseUrl == nil {
		defaultStiggBaseUrl := "https://api.stigg.io/graphql"
		baseUrl = &defaultStiggBaseUrl
	}
	if httpClient == nil {
		retryClient := retryablehttp.NewClient()
		retryClient.RetryMax = RetryCount
		retryClient.Logger = nil

		httpClient = retryClient.StandardClient()
		httpClient.Timeout = time.Second * 30
	}

	return NewClient(httpClient, *baseUrl, &clientv2.Options{ParseDataAlongWithErrors: true},
		func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res any, next clientv2.RequestInterceptorFunc) error {
			req.Header.Set("x-api-key", fmt.Sprintf(apiKey))
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("x-graphql-operation-name", gqlInfo.Request.OperationName)

			return next(ctx, req, gqlInfo, res)
		})
}
