package repository

import (
	"strings"
	"github.com/hashicorp/aws-sdk-go/gen/ec2"
	"github.com/hashicorp/aws-sdk-go/aws"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"os"
)

type EC2InstanceRepository struct {
	key string
	secret string
}

func NewEC2InstanceRepository() *EC2InstanceRepository {
	ec2Service:= &EC2InstanceRepository{}
	ec2Service.key = os.Getenv("AWS_ACCESS_KEY_ID")
	ec2Service.secret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	return ec2Service;
}

func (ecr EC2InstanceRepository) GetDetails() ([]modal.InstanceDetail, error) {
	creds := aws.Creds(ecr.key, ecr.secret, "")
	client := ec2.New(creds, os.Getenv("AWS_REGION"), nil)

	// Only grab instances that are running or just started
	filters := []ec2.Filter{
		ec2.Filter{
			aws.String("instance-state-name"),
			[]string{"running", "pending"},
		},
	}

	request := ec2.DescribeInstancesRequest{Filters: filters}
	result, err := client.DescribeInstances(&request)

	if err != nil {
		return nil, err
	}

	instanceDetails := make([]modal.InstanceDetail, len(result.Reservations))
	index := 0
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			var name aws.StringValue
			tags := instance.Tags
			for _, tag := range tags {
				if (strings.EqualFold(*tag.Key, "name")) {
					name = tag.Value;
				}
			}
			instanceDetail := modal.NewInstanceDetail(*name, *instance.PublicIPAddress);
			instanceDetails[index] = instanceDetail;
			index++;
		}
	}

	return instanceDetails, nil
}
