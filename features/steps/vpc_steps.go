package main

import (
	"os"
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
	vpcCidr = terraform.Output(testT, terraformOptions, "cidr_block")
	return nil
}

func theVPCShouldExist() error {
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		region = "us-east-2"
	}
	_, err := aws.GetVpcByIdE(testT, vpcID, region)
	return err
}

func itsCIDRBlockShouldBe(expected string) error {
	assert.Equal(testT, expected, vpcCidr)
	return nil
}

func FeatureContext(ctx *godog.ScenarioContext) {
	ctx.Step(`^Given I deploy the VPC Terraform module with CIDR "([^"']*)"$`, iDeployTheVPCTerraformModuleWithCIDR)
	ctx.Step(`^Then the VPC should exist$`, theVPCShouldExist)
	ctx.Step(`^And its CIDR block should be "([^"']*)"$`, itsCIDRBlockShouldBe)
	ctx.AfterScenario(func(*godog.Scenario, error) {
		terraform.Destroy(testT, terraformOptions)
	})
}
