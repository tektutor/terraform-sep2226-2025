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
  image = "bitnami/nginx:latest"
  name  = "nginx1"
  hostname  = "nginx1"
}

resource "docker_container" "nginx_container_2" {
  image = "bitnami/nginx:latest"
  name  = "nginx2"
  hostname  = "nginx2"
}


resource "docker_container" "nginx_container_3" {
  image = "bitnami/nginx:latest"
  name  = "nginx3"
  hostname  = "nginx3"
}

