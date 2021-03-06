package cloudformation

import (
	"encoding/json"
)

// AWSKinesisFirehoseDeliveryStream_ProcessingConfiguration AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.ProcessingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html
type AWSKinesisFirehoseDeliveryStream_ProcessingConfiguration struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-enabled
	Enabled *Value `json:"Enabled,omitempty"`

	// Processors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-processors
	Processors []AWSKinesisFirehoseDeliveryStream_Processor `json:"Processors,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_ProcessingConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.ProcessingConfiguration"
}

func (r *AWSKinesisFirehoseDeliveryStream_ProcessingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
