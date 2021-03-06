// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// CodeDeploy provides the API operation methods for making requests to
// AWS CodeDeploy. See this package's package overview docs
// for details on the service.
//
// CodeDeploy methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type CodeDeploy struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*CodeDeploy)

// Used for custom request initialization logic
var initRequest func(*CodeDeploy, *aws.Request)

// Service information constants
const (
	ServiceName = "codedeploy" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName  // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CodeDeploy client with a config.
//
// Example:
//     // Create a CodeDeploy client from just a config.
//     svc := codedeploy.New(myConfig)
func New(config aws.Config) *CodeDeploy {
	var signingName string
	signingRegion := config.Region

	svc := &CodeDeploy{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2014-10-06",
				JSONVersion:   "1.1",
				TargetPrefix:  "CodeDeploy_20141006",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a CodeDeploy operation and runs any
// custom request initialization.
func (c *CodeDeploy) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
