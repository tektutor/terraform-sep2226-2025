---
page_title: "Resource docker_image - terraform-provider-docker"
subcategory: ""
description: |-
  Manages the lifecycle of a docker image in your docker host. It can be used to pull an existing one from a registry.
---
# Resource (docker_image)

Manages the lifecycle of a docker image in your docker host. It can be used to pull an existing one from a registry.

## Example Usage

### Basic

Finds and downloads the latest `ubuntu:precise` image but does not check
for further updates of the image

```terraform
resource "docker_image" "ubuntu" {
  name = "ubuntu:precise"
}
```
