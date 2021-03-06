// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elbv2

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// ELBV2 provides the API operation methods for making requests to
// Elastic Load Balancing. See this package's package overview docs
// for details on the service.
//
// ELBV2 methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ELBV2 struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*ELBV2)

// Used for custom request initialization logic
var initRequest func(*ELBV2, *aws.Request)

// Service information constants
const (
	ServiceName = "elasticloadbalancing" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName            // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ELBV2 client with a config.
//
// Example:
//     // Create a ELBV2 client from just a config.
//     svc := elbv2.New(myConfig)
func New(config aws.Config) *ELBV2 {
	var signingName string
	signingRegion := config.Region

	svc := &ELBV2{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2015-12-01",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a ELBV2 operation and runs any
// custom request initialization.
func (c *ELBV2) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
