terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.6.2"
    }
  }
}

provider "docker" {
  # Configuration options
} 

resource "docker_container" "nginx_container_1" {
  name = "nginx4"
  image = "bitnami/nginx:1.20"
}

resource "docker_container" "nginx_container_2" {
  name = "nginx2"
  image = "bitnami/nginx:latest"
}


resource "docker_container" "nginx_container_3" {
  name = "nginx3"
  image = "bitnami/nginx:latest"
}

