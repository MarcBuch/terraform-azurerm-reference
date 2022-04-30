# Terraform Azure Resource Group Module

<!-- BEGIN_TF_DOCS -->

## Requirements

| Name                                                                     | Version  |
| ------------------------------------------------------------------------ | -------- |
| <a name="requirement_terraform"></a> [terraform](#requirement_terraform) | >= 1.1.8 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement_azurerm)       | >= 3.1.0 |

## Providers

| Name                                                         | Version  |
| ------------------------------------------------------------ | -------- |
| <a name="provider_azurerm"></a> [azurerm](#provider_azurerm) | >= 3.1.0 |

## Modules

No modules.

## Resources

| Name                                                                                                                          | Type     |
| ----------------------------------------------------------------------------------------------------------------------------- | -------- |
| [azurerm_resource_group.this](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/resource_group) | resource |

## Inputs

| Name                                                      | Description                                    | Type       | Default | Required |
| --------------------------------------------------------- | ---------------------------------------------- | ---------- | ------- | :------: |
| <a name="input_location"></a> [location](#input_location) | (Required) The location of the Resource Group. | `string`   | n/a     |   yes    |
| <a name="input_name"></a> [name](#input_name)             | (Required) The name of the Resource Group.     | `string`   | n/a     |   yes    |
| <a name="input_tags"></a> [tags](#input_tags)             | (Optional) The tags for the Resource Group.    | `map(any)` | `{}`    |    no    |

## Outputs

| Name                                            | Description |
| ----------------------------------------------- | ----------- |
| <a name="output_id"></a> [id](#output_id)       | n/a         |
| <a name="output_name"></a> [name](#output_name) | n/a         |

<!-- END_TF_DOCS -->
