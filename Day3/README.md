# Day 3

## Info - Terraform Overview
<pre>
- is one of the Infrastructure as a code tool (Iac)
- it is cloud newtral, hence this Iac tool can be used in any public cloud environment like AWS, Azure, GCE, etc.,
- this also called used to provision infrastructure on your on-prem data-centres
- it helps you provision containers, manage images, provision virtual machines locally or on public cloud, etc.,
- it can be used provision storage cluster, etc.,
- it can be used to provision eks, aks, ROSA, ARO on public cloud
- unlike the AWS cloudformation,  Terraform works on any environment and any cloud
- it comes in 2 flavours
  1. Terraform core ( command-line only - opensource and free )
  2. Terraform Enterprise ( Web console and it is a paid tool )
</pre>

## Info - Terraform High Level Architecture
![Terraform](terraform-architecture-diagram.png)

## Lab - Checking the Terraform version
```
terraform --version
```
<img width="1760" height="422" alt="image" src="https://github.com/user-attachments/assets/7c9dc5b9-074c-4803-93f8-d2a07d00b2e6" />

## Info - Terraform Providers
<pre>
- Terraform depends on Providers to provision resources
- For example
  - In order to provision an ec2 instance in AWS, Terraform depends on a provider called AWS ( registry.terraform.io )
  - IN order to provision an azure VM in Azure portal, Terraform depends on a provider called Azure
  - as long as there is a provider, Terraform can provision resources on that environment
  - In case, to provision a particular type of resource within your organization and there is no read-made provider, you can
    develop your own provider in Golang using Terraform Provider SDK
  - Providers supports two types of objects/resources
    1. Resources
       - If you wish to Provision ec2 instances using Terraform, then you will define a resource block expressing your expected state
       - Terrafrom can Create, Replace, Update and Delete the resources managed by Terraform
    2. DataSources ( already existing resources - these objects will be treated by Terraform as a read-only resource )
       - these resources are not managed by Terraform
       - they are managed outside Terraform
       - Terraform can refer and use it the HCL (Hashicorp Configuration Language - Terraform's proprietary language )
       - IN case to provision certain resource you declarative terraform script(manifest) file depends on already existing resource
         then, we call them as DataSources or Data block
</pre>

## Info - Terraform Resources
<pre>
- Each Terraform Provider supports one to many Resources and one to many Datasources
- For instance, the docker provider supports the following resources
  - docker_image
  - docker_container
</pre>

## Lab - Downloading docker image using Terraform
Create a file named main.tf with the below code
<pre>
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

resource "docker_image" "nginx_docker_image" {
    name = "bitnami/nginx:latest"
}  
</pre>

Then let's download the docker provider using the below command
```
terraform init
tree .terraform
terraform plan
terraform apply
cat terraform.tfsate
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/8c19f2a2-d6cc-4c16-91ff-ede9619bff7d" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/199b680e-ddcd-45f1-b6b9-9aced8aad60a" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b4a5a745-cc46-4397-8b83-a2fb9ab0da92" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/24fb6f90-fc8e-4950-93a1-db931c224ceb" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/031e5f7e-db91-48d2-91a8-3b74999c6231" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/bd6a1af6-e155-404d-99ee-8a9ef6d6a000" />

## Lab - Provisioning container using Terraform

Create a separate folder for each exercise
```
cd ~
mkdir provision-docker-containers
cd provision-docker-containers
touch providers.tf main.tf
tree
```

Create a file named providers.tf with the below code
<pre>
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
</pre>

Create a file named main.tf with the below code
<pre>
data "docker_image" "bitnami_nginx_image" {
   name = "bitnami/nginx:latest"
}

resource "docker_container" "my_nginx1_container" {
   name = "nginx_container_1"
   image = data.docker_image.bitnami_nginx_image.name
}

resource "docker_container" "my_nginx2_container" {
   name = "nginx_container_2"
   image = data.docker_image.bitnami_nginx_image.name
}  
</pre>

You may download the providers and do the terraform provisioning
```
terraform init
terraform plan
terraform apply --auto-approve
terraform show
docker images  | grep nginx
docker ps
```
