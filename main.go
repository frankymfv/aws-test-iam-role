package main

import (
	"aws_test_iam_role/awsv1"
	"aws_test_iam_role/awsv2"
	myconfig "aws_test_iam_role/config"
	"fmt"
)

type Manager struct{}

func main() {
	// Example usage

	awscfg := myconfig.LoadCfg()
	fmt.Printf("datacfg: %+v\n", awscfg)

	fmt.Printf("========= aws v22 =========\n\n\n")
	awsv2.QueryS3v2(awscfg)

	fmt.Printf("========= aws v1 =========\n\n")
	awsv1.QueryS3withSetup(awscfg)
	awsv1.QueryS3withInitS3Config(awscfg)
	awsv1.SendMessageToQueue(awscfg, "hello world 222")

}
