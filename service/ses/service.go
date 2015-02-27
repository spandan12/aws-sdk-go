package ses

import (
	"github.com/awslabs/aws-sdk-go/aws/protocol/query"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
)

// SES is a client for Amazon SES.
type SES struct {
	*aws.Service
}

type SESConfig struct {
	*aws.Config
}

// New returns a new SES client.
func New(config *SESConfig) *SES {
	if config == nil {
		config = &SESConfig{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "email",
		APIVersion:  "2010-12-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)

	return &SES{service}
}