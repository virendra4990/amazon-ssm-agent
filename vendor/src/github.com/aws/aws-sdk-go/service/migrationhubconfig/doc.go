// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package migrationhubconfig provides the client and types for making API
// requests to AWS Migration Hub Config.
//
// The AWS Migration Hub home region APIs are available specifically for working
// with your Migration Hub home region. You can use these APIs to determine
// a home region, as well as to create and work with controls that describe
// the home region.
//
// You can use these APIs within your home region only. If you call these APIs
// from outside your home region, your calls are rejected, except for the ability
// to register your agents and connectors.
//
// You must call GetHomeRegion at least once before you call any other AWS Application
// Discovery Service and AWS Migration Hub APIs, to obtain the account's Migration
// Hub home region.
//
// The StartDataCollection API call in AWS Application Discovery Service allows
// your agents and connectors to begin collecting data that flows directly into
// the home region, and it will prevent you from enabling data collection information
// to be sent outside the home region.
//
// For specific API usage, see the sections that follow in this AWS Migration
// Hub Home Region API reference.
//
// The Migration Hub Home Region APIs do not support AWS Organizations.
//
// See https://docs.aws.amazon.com/goto/WebAPI/migrationhub-config-2019-06-30 for more information on this service.
//
// See migrationhubconfig package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/migrationhubconfig/
//
// Using the Client
//
// To contact AWS Migration Hub Config with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Migration Hub Config client MigrationHubConfig for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/migrationhubconfig/#New
package migrationhubconfig
