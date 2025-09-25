variable "container_name1" {
   description = "Name of the container"
   type = string
   default = "ubuntu"
}

variable "container_name2" {
   description = "Name of the container"
   type = string
   default = "rocky"
}

variable "container_count" {
   description = "Number of containers"
   type = number

   validation {
      condition = var.container_count >= 5 && var.container_count <= 10
      error_message = "The container count must be between 5 and 10 inclussive."
   }
}
