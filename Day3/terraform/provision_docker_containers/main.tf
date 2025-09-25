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
