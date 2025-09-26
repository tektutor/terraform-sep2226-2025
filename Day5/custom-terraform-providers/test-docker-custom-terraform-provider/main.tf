terraform {
  required_providers {
    docker = {
      source = "tektutor/docker"
    }
  }
}

resource "docker_image" "nginx" {
  image_name = "bitnami/nginx:latest"
}

resource "docker_container" "ubuntu_container" {
   container_name = "c1"
   host_name = "c1"
   image_name = "tektutor/ubuntu-ansible-node:latest"
}

resource "docker_container" "rocky_container" {
   container_name = "c2"
   host_name = "c2"
   image_name = "tektutor/rocky-ansible-node:latest"
}
