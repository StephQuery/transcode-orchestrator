// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Modify a job template by sending a request with the job template name and
// any of the following that you wish to change: description, category, and
// queue.
type UpdateJobTemplateInput struct {
	_ struct{} `type:"structure"`

	// Accelerated transcoding can significantly speed up jobs with long, visually
	// complex content. Outputs that use this feature incur pro-tier pricing. For
	// information about feature limitations, see the AWS Elemental MediaConvert
	// User Guide.
	AccelerationSettings *AccelerationSettings `locationName:"accelerationSettings" type:"structure"`

	// The new category for the job template, if you are changing it.
	Category *string `locationName:"category" type:"string"`

	// The new description for the job template, if you are changing it.
	Description *string `locationName:"description" type:"string"`

	// Optional list of hop destinations.
	HopDestinations []HopDestination `locationName:"hopDestinations" type:"list"`

	// The name of the job template you are modifying
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// Specify the relative priority for this job. In any given queue, the service
	// begins processing the job with the highest value first. When more than one
	// job has the same priority, the service begins processing the job that you
	// submitted first. If you don't specify a priority, the service uses the default
	// value 0.
	Priority *int64 `locationName:"priority" type:"integer"`

	// The new queue for the job template, if you are changing it.
	Queue *string `locationName:"queue" type:"string"`

	// JobTemplateSettings contains all the transcode settings saved in the template
	// that will be applied to jobs created from it.
	Settings *JobTemplateSettings `locationName:"settings" type:"structure"`

	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch
	// Events. Set the interval, in seconds, between status updates. MediaConvert
	// sends an update at this interval from the time the service begins processing
	// your job to the time it completes the transcode or encounters an error.
	StatusUpdateInterval StatusUpdateInterval `locationName:"statusUpdateInterval" type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateJobTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobTemplateInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Priority != nil && *s.Priority < -50 {
		invalidParams.Add(aws.NewErrParamMinValue("Priority", -50))
	}
	if s.AccelerationSettings != nil {
		if err := s.AccelerationSettings.Validate(); err != nil {
			invalidParams.AddNested("AccelerationSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.HopDestinations != nil {
		for i, v := range s.HopDestinations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "HopDestinations", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			invalidParams.AddNested("Settings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccelerationSettings != nil {
		v := s.AccelerationSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "accelerationSettings", v, metadata)
	}
	if s.Category != nil {
		v := *s.Category

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "category", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HopDestinations != nil {
		v := s.HopDestinations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "hopDestinations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Priority != nil {
		v := *s.Priority

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "priority", protocol.Int64Value(v), metadata)
	}
	if s.Queue != nil {
		v := *s.Queue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Settings != nil {
		v := s.Settings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "settings", v, metadata)
	}
	if len(s.StatusUpdateInterval) > 0 {
		v := s.StatusUpdateInterval

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusUpdateInterval", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Successful update job template requests will return the new job template
// JSON.
type UpdateJobTemplateOutput struct {
	_ struct{} `type:"structure"`

	// A job template is a pre-made set of encoding instructions that you can use
	// to quickly create a job.
	JobTemplate *JobTemplate `locationName:"jobTemplate" type:"structure"`
}

// String returns the string representation
func (s UpdateJobTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobTemplate != nil {
		v := s.JobTemplate

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jobTemplate", v, metadata)
	}
	return nil
}

const opUpdateJobTemplate = "UpdateJobTemplate"

// UpdateJobTemplateRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Modify one of your existing job templates.
//
//    // Example sending a request using UpdateJobTemplateRequest.
//    req := client.UpdateJobTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/UpdateJobTemplate
func (c *Client) UpdateJobTemplateRequest(input *UpdateJobTemplateInput) UpdateJobTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateJobTemplate,
		HTTPMethod: "PUT",
		HTTPPath:   "/2017-08-29/jobTemplates/{name}",
	}

	if input == nil {
		input = &UpdateJobTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateJobTemplateOutput{})

	return UpdateJobTemplateRequest{Request: req, Input: input, Copy: c.UpdateJobTemplateRequest}
}

// UpdateJobTemplateRequest is the request type for the
// UpdateJobTemplate API operation.
type UpdateJobTemplateRequest struct {
	*aws.Request
	Input *UpdateJobTemplateInput
	Copy  func(*UpdateJobTemplateInput) UpdateJobTemplateRequest
}

// Send marshals and sends the UpdateJobTemplate API request.
func (r UpdateJobTemplateRequest) Send(ctx context.Context) (*UpdateJobTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateJobTemplateResponse{
		UpdateJobTemplateOutput: r.Request.Data.(*UpdateJobTemplateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateJobTemplateResponse is the response type for the
// UpdateJobTemplate API operation.
type UpdateJobTemplateResponse struct {
	*UpdateJobTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateJobTemplate request.
func (r *UpdateJobTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
