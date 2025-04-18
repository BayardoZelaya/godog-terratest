Feature: AWS VPC Module

  Scenario: Create a VPC with a given CIDR block
    Given I deploy the VPC Terraform module with CIDR "10.0.0.0/16"
    Then the VPC should exist
    And its CIDR block should be "10.0.0.0/16"