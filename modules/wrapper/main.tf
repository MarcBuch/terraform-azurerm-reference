module "wrapper" {
  source = "../tf-azure-resource-group"

  name     = var.input.name
  location = var.input.location
}
