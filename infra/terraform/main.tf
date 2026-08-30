
terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
    }

    kubernetes = {
      source = "hashicorp/kubernetes"
    }
  }
  required_version = ">= 0.14.8"

  backend "s3" {
    endpoints = {
      s3 = "https://storage.yandexcloud.net"
    }
    bucket = "olmi-images"
    region = "ru-central1"
    key    = "k8s/terraform.tfstate"

    skip_region_validation      = true
    skip_credentials_validation = true
    skip_requesting_account_id  = true # Необходимая опция Terraform для версии 1.6.1 и старше.
    skip_s3_checksum            = true # Необходимая опция при описании бэкенда для Terraform версии 1.6.3 и старше.
  }
}

resource "yandex_kubernetes_cluster" "momo-store-cluster-learn" {
  name               = "momo-store-cluster-learn"
  network_id         = var.yc_network_id
  service_account_id = var.yc_service_acc_id

  master {
    version = "1.33"
    zonal {
      zone      = var.zone
      subnet_id = var.yc_subnet_id
    }

    public_ip = true
  }

  node_service_account_id = var.yc_service_acc_id

  lifecycle {
    # Игнорируем изменения, которые могут вызвать пересоздание кластера
    ignore_changes = [
      master[0].version,
      master[0].zonal[0].subnet_id,
      master[0].zonal[0].zone,
    ]
  }
}

resource "yandex_kubernetes_node_group" "momo-store-node-group-learn" {
  name       = "momo-store-node-group-learn"
  cluster_id = yandex_kubernetes_cluster.momo-store-cluster-learn.id
  version    = "1.33"

  instance_template {
    platform_id = "standard-v1"

    resources {
      cores       = 2
      memory      = 8
    }

    boot_disk {
      type = "network-ssd"
      size = 30
    }

    # Сетевой интерфейс:
    # нужно указать идентификатор подсети, к которой будет подключена ВМ
    network_interface {
        subnet_ids = [var.yc_subnet_id]
        nat       = true
    }
  }

  scale_policy {
    fixed_scale {
      size = 2
    }
  }

  lifecycle {
    # Позволяем обновлять всё в нодах без пересоздания node group
    ignore_changes = [
      instance_template[0].network_interface,
      instance_template[0].boot_disk,
      version
    ]
  }
}
