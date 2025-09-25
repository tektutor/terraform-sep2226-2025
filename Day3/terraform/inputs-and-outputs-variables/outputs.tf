output "container1_id" {
  value = docker_container.my_ubuntu_container1.id
}

output "container2_id" {
  value = docker_container.my_rocky_container1.id
}

output "container_name1" {
   value = docker_container.my_ubuntu_container1.name
}

output "container_name2" {
   value = docker_container.my_rocky_container1.name
}

output "container1_ip" {
   value = docker_container.my_ubuntu_container1.network_data[0].ip_address
}
output "container2_ip" {
   value = docker_container.my_rocky_container1.network_data[0].ip_address
}
