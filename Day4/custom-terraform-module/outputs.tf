
output "ubuntu_container_names" {
   value = module.create-docker-containers.ubuntu_container_names 
}

output "rocky_container_names" {
   value = module.create-docker-containers.rocky_container_names 
}

output "ubuntu_container_ips" {
   value = module.create-docker-containers.ubuntu_container_ips
}
output "rocky_container_ips" {
   value = module.create-docker-containers.rocky_container_ips 
}
