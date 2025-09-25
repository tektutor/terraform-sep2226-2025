# Day 4

## Lab - Terraform Workspace
<pre>
cd ~/terraform-sep-2025
git pull
cd Day4/terraform-workspace
terraform init
ls -lha
tree .terraform

# Creates a new workspace named dev  
terraform workspace new dev
ls -lha
tree terraform.tfstate.d

# Creates a new workspace named stage  
terraform workspace new stage
ls -lha
tree terraform.tfstate.d

# Creates a new workspace named prod  
terraform workspace new prod
ls -lha
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select dev
terraform apply --auto-approve
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select stage
terraform apply --auto-approve
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select prod
terraform apply --auto-approve
tree terraform.tfstate.d

ls -1
</pre>
