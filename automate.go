// package main

// import (
//     "fmt"
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/service/ec2"
// )

// func main() {
//     // Create a new AWS session
//     sess, err := session.NewSession(&aws.Config{
//         Region: aws.String("ap-south-1"), // Corrected AWS region format
//     })
//     if err != nil {
//         fmt.Println("Error creating session:", err)
//         return
//     }

//     // Create a new EC2 service client
//     svc := ec2.New(sess)

//     // Specify the parameters for the instance
//     params := &ec2.RunInstancesInput{
//         ImageId:      aws.String("ami-0ad37af6d0ff85233"), // Specify the AMI ID
//         InstanceType: aws.String("t2.micro"),              // Specify the instance type
//         MinCount:     aws.Int64(1),
//         MaxCount:     aws.Int64(1),
//     }

//     // Run the instance
//     resp, err := svc.RunInstances(params)
//     if err != nil {
//         fmt.Println("Error launching instance:", err)
//         return
//     }

//     // Output the instance ID
//     instanceID := *resp.Instances[0].InstanceId
//     fmt.Println("Launched instance with ID:", instanceID)
// }

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	// "time"
)

func main() {
	// Create a new AWS session

	// session.NewSession() this is a function from the AWS SDK  library for go
	// it is used to create a new session
	// it will returns two values 'sess' a new AWS session and 'err' an error(if any)
	sess, err := session.NewSession(&aws.Config{
		
		Region: aws.String("ap-south-1"), // Corrected AWS region format
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create a new EC2 service client
	svc := ec2.New(sess)

	// Specify the parameters for the instance
	params := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-0ad37af6d0ff85233"), // Specify the AMI ID
		InstanceType: aws.String("t2.micro"),              // Specify the instance type
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	// Run the instance
	resp, err := svc.RunInstances(params)
	if err != nil {
		fmt.Println("Error launching instance:", err)
		return
	}

	// Output the instance ID
	instanceID := *resp.Instances[0].InstanceId
	fmt.Println("Launched instance with ID:", instanceID)

	// Wait for the instance to be in a running state
	fmt.Println("Waiting for the instance to be in a running state...")
	err = svc.WaitUntilInstanceRunning(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	})
	if err != nil {
		fmt.Println("Error waiting for instance to be running:", err)
		return
	}

}
