output "ubuntu_container_ids" {
  value = docker_container.my_ubuntu_container[*].id
}

output "rocky_container_ids" {
  value = docker_container.my_rocky_container[*].id
}

output "ubuntu_container_names" {
   value = docker_container.my_ubuntu_container[*].name
}

output "rocky_container_names" {
   value = docker_container.my_rocky_container[*].name
}

output "ubuntu_container_ips" {
   value = docker_container.my_ubuntu_container[*].network_data[0].ip_address
}
output "rocky_container_ips" {
   value = docker_container.my_rocky_container[*].network_data[0].ip_address
}
