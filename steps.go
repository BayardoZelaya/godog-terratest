package main

import (
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var terraformOptions *terraform.Options
var vpcID string
var vpcCidr string
var testT *testing.T


func iDeployTheVPCTerraformModuleWithCIDR(cidr string) error {
	terraformOptions = &terraform.Options{
		TerraformDir: "terraform/vpc",
		Vars:         map[string]interface{}{"cidr": cidr},
	}
	terraform.InitAndApply(testT, terraformOptions)

	vpcID = terraform.Output(testT, terraformOptions, "vpc_id")
	vpcCidr = terraform.Output(testT, terraformOptions, "vpc_cidr")

	return nil
}

func theVPCShouldExist() error {
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		region = "us-east-1"
	}
	_, err := aws.GetVpcByIdE(testT, vpcID, region)
	println("VPC ID:", vpcID)
	return err
}

func itsCIDRBlockShouldBe(expected string) error {
	assert.Equal(testT, expected, vpcCidr)
	println("VPC CIDR:", vpcCidr)
	return nil
}

func FeatureContext(ctx *godog.ScenarioContext) {
	ctx.Step(`^I deploy the VPC Terraform module with CIDR "([^"']*)"$`, iDeployTheVPCTerraformModuleWithCIDR)
	ctx.Step(`^the VPC should exist$`, theVPCShouldExist)
	ctx.Step(`^its CIDR block should be "([^"']*)"$`, itsCIDRBlockShouldBe)
}
