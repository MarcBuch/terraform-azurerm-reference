variable "name" {
  description = "(Required) The name of the Resource Group."
  type        = string
}

variable "location" {
  description = "(Required) The location of the Resource Group."
  type        = string
}

variable "tags" {
  description = "(Optional) The tags for the Resource Group."
  type        = map(any)
  default     = {}
}
