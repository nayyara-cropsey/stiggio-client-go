# Stigg.io Go Client

Go Client for Stigg GraphQL API.

This client is auto-generated from the Stigg GraphQL Schema Definition
Language([SDL](https://studio.apollographql.com/public/Stigg-API/variant/Production/schema/reference)) available via
Apollo GraphQL explorer. It is stored in the file [schema.graphql](schema/schema.graphql).

It uses the [gqlgenc](https://github.com/Yamashou/gqlgenc) client generator Go library to generated the client.

## Motivation

The Stigg GO SDK currently available from Stigg does not capture some of the GraphQL queries provided via the public
API. This is an alternative client that expands the queries available for use with the SDK.

This library should be used as an alternative to Go SDK for Stigg.

## Getting started

1. Install

```
go get github.com/nayyara-cropsey/stiggio-client-go
```

2. Retrieve your server API key from [Stigg console](https://app.stigg.io/account/settings)
3. Init your Go Client

```
import github.com/nayyara-cropsey/stiggio-client-go

....

apiKey := "XXXXXXXXX"
client := stigg.NewStiggClient(apiKey, nil, nil)
```
