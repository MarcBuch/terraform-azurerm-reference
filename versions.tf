terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.1.0"
    }
  }
  backend "azurerm" {

  }
  required_version = ">= 1.1.8"
}
