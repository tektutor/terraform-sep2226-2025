data "docker_image" "tektutor_ansible_ubuntu_image" {
  name = "tektutor/ubuntu-ansible-node:latest"
}

data "docker_image" "tektutor_ansible_rocky_image" {
  name = "tektutor/rocky-ansible-node:latest"
}

resource "docker_container" "my_ubuntu_container1" {
  image = data.docker_image.tektutor_ansible_ubuntu_image.name
  name =  var.container_name1

  ports {
     internal = "22"
     external = "2001"
  }

  ports {
     internal = "80"
     external = "8001"
  }
  
  depends_on = [
     data.docker_image.tektutor_ansible_ubuntu_image
  ]
}

resource "docker_container" "my_rocky_container1" {
  image = data.docker_image.tektutor_ansible_rocky_image.name
  name =  var.container_name2

  ports {
     internal = "22"
     external = "2002"
  }

  ports {
     internal = "80"
     external = "8002"
  }

  depends_on = [
     data.docker_image.tektutor_ansible_rocky_image
  ]
}

resource "null_resource" "invoke_ansible_playbook" {

  triggers = {
     always_run = timestamp()
  }


  provisioner "local-exec" {
    environment = {
      ANSIBLE_CONFIG = "${path.module}/ansible.cfg"
    }
    command = "ansible-playbook install-nginx-playbook.yml"
  }

  depends_on = [
     docker_container.my_ubuntu_container1,
     docker_container.my_rocky_container1,
  ]
}
