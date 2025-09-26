# Day 5

## Lab - Implementing a custom Terraform Provider using Golang

You need to create a folder
```
mkdir -p ~/go/bin
touch ~/.terraformrc
```

Paste the below code in your ~/.terraformrc file
<pre>
provider_installation {	
    dev_overrides {
    	"registry.terraform.io/tektutor/file" = "/home/student/go/bin",
    }
    direct{}
}  
</pre>

Then you may proceed with the below instructions
```
cd ~/terraform-sep2226-2025
git pull
cd Day5/custom-terraform-providers/
tree
cd file

go mod init github.com/tektutor/terraform-provider-file
go mod tidy
ls -l
go build
ls -l
go install
ls -l
ls -l ~/go/bin
```

## Lab - Using our custom teraform file provider in our Terraform project
```
cd ~/terraform-sep2226-2025
git pull
cd Day5/custom-terraform-providers/test-file-custom-terraform-provider
ls -l
terraform plan
terraform apply --auto-approve
ls -l
cat terraform.tfstate
cat myfile.txt
```
