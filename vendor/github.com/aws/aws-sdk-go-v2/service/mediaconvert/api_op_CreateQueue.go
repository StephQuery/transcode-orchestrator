// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Create an on-demand queue by sending a CreateQueue request with the name
// of the queue. Create a reserved queue by sending a CreateQueue request with
// the pricing plan set to RESERVED and with values specified for the settings
// under reservationPlanSettings. When you create a reserved queue, you enter
// into a 12-month commitment to purchase the RTS that you specify. You can't
// cancel this commitment.
type CreateQueueInput struct {
	_ struct{} `type:"structure"`

	// Optional. A description of the queue that you are creating.
	Description *string `locationName:"description" type:"string"`

	// The name of the queue that you are creating.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// Specifies whether the pricing plan for the queue is on-demand or reserved.
	// For on-demand, you pay per minute, billed in increments of .01 minute. For
	// reserved, you pay for the transcoding capacity of the entire queue, regardless
	// of how much or how little you use it. Reserved pricing requires a 12-month
	// commitment. When you use the API to create a queue, the default is on-demand.
	PricingPlan PricingPlan `locationName:"pricingPlan" type:"string" enum:"true"`

	// Details about the pricing plan for your reserved queue. Required for reserved
	// queues and not applicable to on-demand queues.
	ReservationPlanSettings *ReservationPlanSettings `locationName:"reservationPlanSettings" type:"structure"`

	// Initial state of the queue. If you create a paused queue, then jobs in that
	// queue won't begin.
	Status QueueStatus `locationName:"status" type:"string" enum:"true"`

	// The tags that you want to add to the resource. You can tag resources with
	// a key-value pair or with only a key.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateQueueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateQueueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateQueueInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.ReservationPlanSettings != nil {
		if err := s.ReservationPlanSettings.Validate(); err != nil {
			invalidParams.AddNested("ReservationPlanSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateQueueInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.PricingPlan) > 0 {
		v := s.PricingPlan

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pricingPlan", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ReservationPlanSettings != nil {
		v := s.ReservationPlanSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "reservationPlanSettings", v, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Successful create queue requests return the name of the queue that you just
// created and information about it.
type CreateQueueOutput struct {
	_ struct{} `type:"structure"`

	// You can use queues to manage the resources that are available to your AWS
	// account for running multiple transcoding jobs at the same time. If you don't
	// specify a queue, the service sends all jobs through the default queue. For
	// more information, see https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html.
	Queue *Queue `locationName:"queue" type:"structure"`
}

// String returns the string representation
func (s CreateQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateQueueOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Queue != nil {
		v := s.Queue

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "queue", v, metadata)
	}
	return nil
}

const opCreateQueue = "CreateQueue"

// CreateQueueRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Create a new transcoding queue. For information about queues, see Working
// With Queues in the User Guide at https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html
//
//    // Example sending a request using CreateQueueRequest.
//    req := client.CreateQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/CreateQueue
func (c *Client) CreateQueueRequest(input *CreateQueueInput) CreateQueueRequest {
	op := &aws.Operation{
		Name:       opCreateQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/2017-08-29/queues",
	}

	if input == nil {
		input = &CreateQueueInput{}
	}

	req := c.newRequest(op, input, &CreateQueueOutput{})

	return CreateQueueRequest{Request: req, Input: input, Copy: c.CreateQueueRequest}
}

// CreateQueueRequest is the request type for the
// CreateQueue API operation.
type CreateQueueRequest struct {
	*aws.Request
	Input *CreateQueueInput
	Copy  func(*CreateQueueInput) CreateQueueRequest
}

// Send marshals and sends the CreateQueue API request.
func (r CreateQueueRequest) Send(ctx context.Context) (*CreateQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateQueueResponse{
		CreateQueueOutput: r.Request.Data.(*CreateQueueOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateQueueResponse is the response type for the
// CreateQueue API operation.
type CreateQueueResponse struct {
	*CreateQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateQueue request.
func (r *CreateQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
