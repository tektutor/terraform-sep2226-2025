terraform {
  required_providers {
     local = {
       source = "hashicorp/local"
     }
  }
}

provider "local" {}


resource "local_file" "myfile" {
   filename = "${path.module}/message-${terraform.workspace}.txt"
   content  = "Hello from Terraform Workspace - ${terraform.workspace}"
}
