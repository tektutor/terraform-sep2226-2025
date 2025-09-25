data "docker_image" "tektutor_ansible_ubuntu_image" {
  name = "tektutor/ubuntu-ansible-node:latest"
}

data "docker_image" "tektutor_ansible_rocky_image" {
  name = "tektutor/rocky-ansible-node:latest"
}

resource "docker_container" "my_ubuntu_container" {
  count =  var.container_count 

  image = data.docker_image.tektutor_ansible_ubuntu_image.name
  name =  "${var.container_name1}-${count.index+1}"
  hostname =  "${var.container_name1}-${count.index+1}"

  ports {
     internal = "22"
     external = "2${format("%03d",count.index+1)}"
  }

  ports {
     internal = "80"
     external = "8${format("%03d",count.index+1)}"
  }
  
  depends_on = [
     data.docker_image.tektutor_ansible_ubuntu_image
  ]
}

resource "docker_container" "my_rocky_container" {
  count =  var.container_count 

  image = data.docker_image.tektutor_ansible_rocky_image.name
  name =  "${var.container_name2}-${count.index+1}"
  hostname =  "${var.container_name2}-${count.index+1}"

  ports {
     internal = "22"
     external = "3${format("%03d",count.index+1)}"
  }

  ports {
     internal = "80"
     external = "9${format("%03d",count.index+1)}"
  }

  depends_on = [
     data.docker_image.tektutor_ansible_rocky_image
  ]
}

resource "null_resource" "ubuntu_remote_connection" {
  count = var.container_count

  triggers = {
     always_run = timestamp()
  }

  connection {
     type = "ssh"
     user = "root"
     port = "2${format("%03d",count.index+1)}"
     host = "localhost"
     private_key = file("~/.ssh/id_ed25519")
     timeout = "15s" # Wait upto 15s before attempting SSH
     agent  = false
  }

  provisioner "remote-exec" {
     inline = [
         "hostname",
         "hostname -i",
     ]
  }

  depends_on = [
     docker_container.my_ubuntu_container,
  ]
}

resource "null_resource" "rocky_remote_connection" {
  count = var.container_count

  triggers = {
     always_run = timestamp()
  }

  connection {
     type = "ssh"
     user = "root"
     port = "3${format("%03d",count.index+1)}"
     host = "localhost"
     private_key = file("~/.ssh/id_ed25519")
     timeout = "15s" # Wait upto 15s before attempting SSH
     agent  = false
  }

  provisioner "remote-exec" {
     inline = [
         "hostname",
         "hostname -i",
     ]
  }

  depends_on = [
     docker_container.my_rocky_container,
  ]
}
