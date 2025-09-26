---
layout: ""
page_title: "Provider: Docker"
description: |-
  The Docker provider is used to interact with Docker resources, like containers, images, etc.
---

# Docker Provider

The Docker provider is used to interact with Docker containers and images.
It uses the Docker API to manage the lifecycle of Docker containers. Because
the Docker provider uses the Docker API, it is immediately compatible not
only with single server Docker but Swarm and any additional Docker-compatible
API hosts.

Use the navigation to the left to read about the available resources.

## Example Usage

Terraform 0.13 and later:

```terraform
terraform {
  required_providers {
    docker = {
      source  = "tektutor/docker"
    }
  }
}

provider "docker" {
}

# Pulls the image
resource "docker_image" "ubuntu" {
  name = "ubuntu:latest"
}

# Create a container
resource "docker_container" "foo" {
  image_name = docker_image.ubuntu.image_id
}
```

Required:

- `image_name` (String) Docker Image name 

Optional:
