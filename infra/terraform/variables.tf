
variable "yc_service_account_key_file" {}
variable "cloud_id" {}
variable "folder_id" {}
variable "zone" {
  description = "Зона доступности"
  type        = string
  default     = "ru-central1-a"
}
variable "yc_s3_access_key" {}
variable "yc_s3_secret_key" {}

variable "yc_subnet_id" {}
variable "yc_network_id" {}

variable "yc_service_acc_id" {}
