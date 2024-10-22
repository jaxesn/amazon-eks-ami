// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// By default, all vCPUs for the instance type are active when you launch an
// instance. When you configure the number of active vCPUs for the instance, it can
// help you save on licensing costs and optimize performance. The base cost of the
// instance remains unchanged.
//
// The number of active vCPUs equals the number of threads per CPU core multiplied
// by the number of cores. The instance must be in a Stopped state before you make
// changes.
//
// Some instance type options do not support this capability. For more
// information, see [Supported CPU options]in the Amazon EC2 User Guide.
//
// [Supported CPU options]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/cpu-options-supported-instances-values.html
func (c *Client) ModifyInstanceCpuOptions(ctx context.Context, params *ModifyInstanceCpuOptionsInput, optFns ...func(*Options)) (*ModifyInstanceCpuOptionsOutput, error) {
	if params == nil {
		params = &ModifyInstanceCpuOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceCpuOptions", params, optFns, c.addOperationModifyInstanceCpuOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceCpuOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceCpuOptionsInput struct {

	// The number of CPU cores to activate for the specified instance.
	//
	// This member is required.
	CoreCount *int32

	// The ID of the instance to update.
	//
	// This member is required.
	InstanceId *string

	// The number of threads to run for each CPU core.
	//
	// This member is required.
	ThreadsPerCore *int32

	// Checks whether you have the required permissions for the operation, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyInstanceCpuOptionsOutput struct {

	// The number of CPU cores that are running for the specified instance after the
	// update.
	CoreCount *int32

	// The ID of the instance that was updated.
	InstanceId *string

	// The number of threads that are running per CPU core for the specified instance
	// after the update.
	ThreadsPerCore *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstanceCpuOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceCpuOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceCpuOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyInstanceCpuOptions"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpModifyInstanceCpuOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceCpuOptions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyInstanceCpuOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyInstanceCpuOptions",
	}
}