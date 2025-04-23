package main

import (
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestFeatures(t *testing.T) {
	testT = t

	defer func() {
		if terraformOptions != nil {
			terraform.Destroy(testT, terraformOptions)
			terraformOptions = nil
		}
	}()

	suite := godog.TestSuite{
		Name:                "go-bdd",
		ScenarioInitializer: FeatureContext,
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{"features"},
			Randomize: time.Now().UTC().UnixNano(),
			Output:    colors.Colored(os.Stdout),
		},
	}
	if suite.Run() != 0 {
		t.Fail()
	}
}
