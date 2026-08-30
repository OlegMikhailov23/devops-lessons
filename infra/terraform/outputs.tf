
data "yandex_kubernetes_cluster" "momo-store-cluster-learn" {
  cluster_id = yandex_kubernetes_cluster.momo-store-cluster-learn.id
}

output "endpoint" {
  value = data.yandex_kubernetes_cluster.momo-store-cluster-learn.master[0].external_v4_endpoint
}
