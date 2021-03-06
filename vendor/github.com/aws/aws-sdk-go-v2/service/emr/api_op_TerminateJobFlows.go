// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Input to the TerminateJobFlows operation.
type TerminateJobFlowsInput struct {
	_ struct{} `type:"structure"`

	// A list of job flows to be shutdown.
	//
	// JobFlowIds is a required field
	JobFlowIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s TerminateJobFlowsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateJobFlowsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateJobFlowsInput"}

	if s.JobFlowIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobFlowIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TerminateJobFlowsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TerminateJobFlowsOutput) String() string {
	return awsutil.Prettify(s)
}

const opTerminateJobFlows = "TerminateJobFlows"

// TerminateJobFlowsRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// TerminateJobFlows shuts a list of clusters (job flows) down. When a job flow
// is shut down, any step not yet completed is canceled and the EC2 instances
// on which the cluster is running are stopped. Any log files not already saved
// are uploaded to Amazon S3 if a LogUri was specified when the cluster was
// created.
//
// The maximum number of clusters allowed is 10. The call to TerminateJobFlows
// is asynchronous. Depending on the configuration of the cluster, it may take
// up to 1-5 minutes for the cluster to completely terminate and release allocated
// resources, such as Amazon EC2 instances.
//
//    // Example sending a request using TerminateJobFlowsRequest.
//    req := client.TerminateJobFlowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/TerminateJobFlows
func (c *Client) TerminateJobFlowsRequest(input *TerminateJobFlowsInput) TerminateJobFlowsRequest {
	op := &aws.Operation{
		Name:       opTerminateJobFlows,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateJobFlowsInput{}
	}

	req := c.newRequest(op, input, &TerminateJobFlowsOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return TerminateJobFlowsRequest{Request: req, Input: input, Copy: c.TerminateJobFlowsRequest}
}

// TerminateJobFlowsRequest is the request type for the
// TerminateJobFlows API operation.
type TerminateJobFlowsRequest struct {
	*aws.Request
	Input *TerminateJobFlowsInput
	Copy  func(*TerminateJobFlowsInput) TerminateJobFlowsRequest
}

// Send marshals and sends the TerminateJobFlows API request.
func (r TerminateJobFlowsRequest) Send(ctx context.Context) (*TerminateJobFlowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateJobFlowsResponse{
		TerminateJobFlowsOutput: r.Request.Data.(*TerminateJobFlowsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateJobFlowsResponse is the response type for the
// TerminateJobFlows API operation.
type TerminateJobFlowsResponse struct {
	*TerminateJobFlowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateJobFlows request.
func (r *TerminateJobFlowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
