provider "yandex" {
  service_account_key_file = var.yc_service_account_key_file
  cloud_id  = var.cloud_id
  folder_id = var.folder_id
  zone      = var.zone
}

data "yandex_client_config" "client" {}

provider "kubernetes" {
  host                   = yandex_kubernetes_cluster.k8s-cluster.master[0].external_v4_endpoint
  cluster_ca_certificate = yandex_kubernetes_cluster.k8s-cluster.master[0].cluster_ca_certificate
  token                  = data.yandex_client_config.client.iam_token
}
