<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/774ad9be-8078-4a2f-9291-36af0c0acbf3" /># Day 2

## Info - Ansible role
<pre>
- is a way we can write reusable code
- ansible roles can't be executed like playbooks directly
- ansible roles is similar to DLL(Dynamic Link Library) - it has reusable code but can't be executed directly
- just like dll can be invoked from application, ansible roles can be invoked from ansible playbooks
- the same ansible role can be used from multiple playbooks
- ansible roles following a recommended directory structure
- it looks like a playbook but it is not a playbook
- using ansible-galaxy one can download and use read-made ansible roles from galaxy.ansible.com portal
- we could also develop our custom ansible role using ansible-galaxy tool
</pre>

## Lab - Developing an ansible role to install nginx,configure web root folder and deploy custom web page
```
cd ~/terraform-sep2226-2025
git pull
cd Day2/ansible/role
ansible-galaxy init nginx
tree nginx

cp scripts/default nginx/files
cp scripts/nginx.conf nginx/files
cp scripts/index.html.j2 nginx/templates
cp scripts/nginx-vars.yml nginx/defaults
cp scripts/nginx-vars.yml nginx/vars
cp scripts/restart* nginx/handlers
cp scripts/*.yml nginx/tasks
tree nginx

ansible-playbook install-nginx-playbook.yml
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/5781c876-7474-4a3f-a471-e9813fc48e14" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/e24c4385-d8bb-4278-94ed-cd56318714d7" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3d3f175d-bbfc-4988-b9cb-8c9c25553636" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/7ffa38d2-0ef9-4f65-8d30-76b9cf1b00d0" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3c747e8b-1d16-4ff5-b5fd-044f3e1ef076" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/8ab0b3a5-8f2d-4a4e-bfd8-94368a3ceb93" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b16ae53a-2fa6-4591-b537-477bad25d6ea" />

## Lab - Installing AWX 

#### Let's install minikube
```
curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64

minikube config set cpus 4
minikube config set memory 12288
minikube start --driver=docker

# Download kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x ./kubectl
sudo mv kubectl /usr/local/bin

docker ps -a
minikube status
kubectl get nodes
```
Expected output
![image](img1.png)
![image](img2.png)

#### Let's install 
```
# Clone the awx operator to install Ansible Tower within minikube
git clone https://github.com/ansible/awx-operator.git
cd awx-operator
git checkout tags/2.7.2
export VERSION=2.7.2

# Install make
sudo apt install make -y
make deploy
```

#### Check if the AWX required pods are running
```
kubectl get pods -n awx -w
```

#### Troubleshooting pods crash and AWX Dashboard login failure

Create a file named kustomization.yaml with below code
<pre>
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: awx
resources:
  - github.com/ansible/awx-operator/config/default?ref=2.19.1

images:
  - name: quay.io/ansible/awx-operator
    newTag: 2.19.1
</pre>

Apply
```
kubectl apply -k .
```

#### Let's create a nodeport service for AWX
Create a file awx.yml with the below code
<pre>
---
apiVersion: awx.ansible.com/v1beta1
kind: AWX
metadata:
  name: awx-tower
spec:
  service_type: nodeport
</pre>

Let's create the ansible tower instance
```
kubectl config set-context --current --namespace=awx
kubectl apply -f awx.yml
kubectl logs -f deployments/awx-operator-controller-manager -c awx-manager
kubectl get svc -l "app.kubernetes.io/managed-by=awx-operator"
minikube ip
kubectl get svc -n awx
```
You may access the ansible tower webconsole
```
http://192.168.49.2:30181
```
![image](img3.png)

Retrieve the password
```
kubectl get secret awx-tower-admin-password -n awx -o jsonpath='{.data.password}' | base64 -d; echo
```

Login credentials
<pre>
username - admin
password - 
</pre>

Once you login, you will get the below page
![image](img4.png)

## Lab - Creating a Project in Ansible AWX

Navigate to Ansible Automation Platform on your RPS Lab machine - chrome web browser
![image](img5.png)

On the menu that appears on the left side, Navigate to Resources --> Projects
![image](img6.png)

Click "Add"
![image](img7.png)
<pre>
Under the Name, type "TekTutor Training Repository"
Under the Source Code Type, select "Git"
Under the Source Control url, type "https://github.com/tektutor/terraform-sep2226-2025.git"
Under the Source Control Branch/Tag/Commit, type "main"
Under the Options, enable the check box that says "Update Revision on Launch"
</pre>
![image](img8.png)
Click "Save"
![image](img9.png)
![image](img10.png)
Click "Successful"
![image](img11.png)
![image](img12.png)
![image](img13.png)

## Lab - Creating Inventory in Ansible Automation Platform(AWX)

Navigate to Ansible Automation Platform(AWX)
![image](img14.png)

Click Resources --> Inventories
![image](img15.png)
Click Add
Select the first option "Add Inventory
![image](img16.png)
![image](img17.png)
![image](img18.png)
Click "Save"
![image](img19.png)
Click the Tab named "Hosts" within the Inventory you saved just now
![image](img20.png)
Click "Add"
![image](img21.png)
![image](img22.png)
![image](img23.png)
Click "Save"
![image](img24.png)
![image](img25.png)
click Add to create other ansible nodes on the similar fashion
![image](img26.png)
![image](img27.png)
Click "Save"
![image](img28.png)
click Add to create other ansible nodes on the similar fashion
![image](img29.png)
Click "Add"
![image](img30.png)
![image](img31.png)
Click "Save"
![image](img32.png)

Repeat the procedure to add "Rocky2"
![image](img33.png)
![image](img34.png)
![image](img35.png)

To verify if all the hosts(ansible nodes) added to the inventory are reachable to Ansible Tower, Click on your inventory and move to the Hosts tab
![image](img36.png)
Click "Run command"
![image](img37.png)
Under the Module, choose "ping"
![image](img38.png)
![image](img39.png)
Click "Next"
![image](img40.png)
Click "Next"
![image](img41.png)
Select "RPS Private Key" we saved
Click "Next"
![image](img42.png)
Click "Launch"
![image](img43.png)
![image](img44.png)
![image](img45.png)


## Lab - Creating Credentials to store the Private key 
Navigate to Ansible Tower Dashboard
![image](img46.png)

Click Resources --> Credentials
![image](img47.png)
Click "Add"
![image](img48.png)
![image](img49.png)
Select "Machine" Credential Type
![image](img50.png)
Open your RPS Cloud Machine Terminal, type "cat /home/rps/.ssh/id_ed25519"
![image](img51.png)
Copy the private key including the Begin and End as shown below
![image](img52.png)
Paste the private key you copied under the "SSH Private Key" field (Remove extra space)
![image](img53.png)
Scroll down to save
![image](img54.png)
Click Save
![image](img55.png)

## Lab - Creating Job Template to invoke a playbook from Ansible Tower
Navigate to Ansible Tower Dashboard
![image](img56.png)

Click "Resources->Templates"
![image](img57.png)
Click "Add"
![image](img58.png)
Select "Add Job Template"
![image](img59.png)
<pre>
Under the Name, type "Install nginx playbook"
Click Search in Inventory and select "Docker Inventory" that we created
</pre>
![image](img60.png)
![image](img61.png)

Click Search in Project and Select "TekTutor Training Repository"
![image](img62.png)
![image](img63.png)
![image](img64.png)
Under the Playbook, select "Day2/ansible/after-refactoring/install-nginx-playbook.yml"
![image](img65.png)
Under Credential, click search and select "RPS private key file"
![image](img66.png)
![image](img67.png)
![image](img68.png)
Scroll down and click "Save"
![image](img69.png)


To run the playbook, click "Launch" Button
![image](img70.png)
![image](img71.png)
![image](img72.png)
![image](img73.png)
![image](img74.png)
![image](img75.png)


## Lab - Installing Golang in Ubuntu
```
cd ~
wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz
tar xvf go1.25.1.linux-amd64.tar.gz
```

Edit your /home/student/.bashrc and append the below at the end of the file
```
export PATH=$PATH:/home/student/go/bin
export GOROOT=/home/student/go
export GOPATH=/home/student/go/bin
```

To apply the exported variables 
```
source ~/.bashrc
```

Now you may verify the golang version
```
go version
```

It will report
<pre>
go version go1.25.1 linux/amd64  
</pre>


## Info - Golang Overview
<pre>
- golang is developed by Google using C programming language
- golang is a compiled programming language
- golang syntax resembles very close to C programming language
- golang has 25 keywords
- golang only supports for loop
- just like C/C++/C# main function is the entry-point function ( the very function that will invoked when you run a go application )
- golang supports pointers but memory is managed by garbage collector unlike C/C++
- golang doesn't support class
- golang only supports functions
- using golang one can develop a new compiler/interpreter, a game, console based application, graphical application that runs on your local machine, can develop mobile applications, can develop AI/ML applications, web applications, etc.,
- using golang one can develop REST API, Microservices, etc.,
- is case-sensitive
- statically typed programming language
- performance wise, it is faster than most compiled languages, definitely more faster than interpretted and scripting languages
- even compilation is done faster in go lang for bulky applications
- Some of the popular tools developed in Golang
  - Terraform
  - Docker
  - Kubernetes
  - Openshift  
</pre>

## Info - Golang keywords
<pre>
- break
- default
- func
- interface
- select
- case
- defer
- go
- map
- struct
- chan
- else
- goto
- package
- switch
- const
- fallthrough
- if
- range
- type
- continue
- for
- import
- return
- var
</pre>

## Lab - Write your first Hello World in golang
Create a file named hello.go with the below content
<pre>
package main

import "fmt"

func main() {
  fmt.Println ("Hello World !")
}
</pre>

Run your application
```
go run ./hello.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/79a6528d-7b9f-4c1a-a470-51be85f04f2e" />

## Lab - Accepting user inputs

Create a file named userinputs.go with the below content
<pre>
package main

import "fmt"

func main() {
	//The below line declares variables named x and y and initializes them with value 0
	x := 0

	//This is yet another way we can declare variables in go lang
	var y int

	//This is an ansignment
	y = 0

	fmt.Print("Enter your first integer input :" )
	fmt.Scanf("%d", &x)

	fmt.Print("Enter your second integer input :" )
	fmt.Scanf("%d", &y)

	fmt.Println("Value of x :", x )
	fmt.Println("Value of y :", y )

	var temp string
	fmt.Println("Press any key to exit ...")
	fmt.Scanln(&temp)
}  
</pre>

Run it
```
go run ./userinputs.go
```

## Lab - If else
Create file named if-else.go with the below code
```
package main

import "fmt"

func main() {
	x := 100

	if x % 2 == 0 {
	    fmt.Println(x, "is an even number")
	} else {
            fmt.Println(x, "is an odd number")
	}
}
```

Run it
```
go run ./if-else.go
```

## Lab - Golang arrays

Create a file named arrays.go with the below content
<pre>
package main

import "fmt"

func main() {

	//We have declared an array of integers of size 5
	//so we can store upto 5 integer values into this array
	//go lang array size is fixed
	//array index starts from 0
	//valid array index range is 0 to 4, total 5 values
	var arr [5]int

	//let's assign valaues into the array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	//arr[5] = 60 This will report array index out of bounds erro

	fmt.Println("Array elements are ...")
	fmt.Println(arr)

	count := len(arr)
	fmt.Println("Length of array :", count)

	//Modifying values stored in an array
	arr[3] = 25

	fmt.Println("Array elements are ...")
	for i := 0; i < count; i++ {
            fmt.Printf("%d\t", arr[i])
	}
	fmt.Println()
}	
</pre>

Run it 
```
go run ./array.go
```

## Lab - Golang error-handling

Create a file named error-handling.go with below code
<pre>
package main

import (
  "fmt"
  "os"
  "os/user"
  "github.com/sttk/stringcase"
)

func main() {
    u, err := user.Current()

    if err != nil {
       fmt.Println("Cannot get current user:", err)
       os.Exit(1)
    }

    fmt.Printf("Hello %s, welcome !\n", stringcase.PascalCase(u.Username) )
}	
</pre>

Run it
```
ls
go mod init main
cat go.mod
go mod tidy
cat go.mod
go run ./error-handling.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/156d9430-086e-4a62-9fbf-b1e187f35193" />

## Lab - Golang user-defined functions

Create a file named functions.go with the below code
<pre>
package main

import "fmt"

func yetAnotherFunction() {
     fmt.Println("Yet Another Function invoked")
}

func main() {
	fmt.Println( sayHello("Golang") )
	fmt.Println( sayHello("World") )
	yetAnotherFunction()
}

//This function accepts a string input argument and returns a string output
func sayHello( msg string ) string {
  return "Hello, " + msg + " !"
}

/* function overloading is not supported in golang
func sayHello() string {
   return "Hello World !"
}
*/


</pre>

Run it
```
go run ./functions.go
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a4d62eb0-fc09-4475-ae6b-ebe79f6f0bd3" />

## Lab - Golang user-defined function that returns multiple values
Create a file named function-with-multiple-returns.go with below code
<pre>
package main

import "fmt"

func myFunction() (int,int) {
  return 10, 20
}

func main() {
   x, y := myFunction() // := is a short form of declaring a new variable and initialized with some value

   fmt.Println( "Value of x is ", x )
   fmt.Println( "Value of y is ", y )
}	
</pre>

Run it
```
go run ./functions-with-multiple-returns.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c089e5c2-84ca-4903-84d1-a02d9cd67000" />

## Lab - Golang loops

Create a file named loops.go with below code
<pre>
package main

import "fmt"

func main() {

	//Declares a count variable of type int and assigns a value 5
	count := 5

	//similar to while loop
	for count > 0 {
		fmt.Println("Before decrementing ", count)

		count--  //equivalent to count = count - 1
		//golang doesn't support pre-decrement or pre-increment unlike C or C++

		fmt.Println("After decrementing ", count )
	}
	fmt.Println("Value of count is ", count, " after for loop")

	//Resetting already declared variable count with value 0
	count = 0

	fmt.Println()
	//Regular for loop
	for count=1; count<10; count++ {
           fmt.Printf( "%d\t", count )
	}
	fmt.Println()

	count = 0

	//similar to do while - infinite loop
	for {

		fmt.Printf("Inside for loop %d\n", count )
		count++

		if count > 3 {
		   break
		}
	}
	fmt.Println("Control reached outside infinite for loop")
}
</pre>

Run it
```
go run ./loops.go
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/1229fbf2-d47f-4022-8580-a5f50f6f5993" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/dfd996b9-1bec-4621-a15c-0052d5b0b9f3" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b706203b-6e84-4cc0-b654-80444d24b7d8" />

## Lab - Golang map

Create a file named map.go with below code
<pre>
package main

import "fmt"

func main() {

	//key,value can be of different data type
	toolsPath := map[string]string {
		"java_home": "/usr/lib/jdk11",
		"mvn_home" : "/usr/share/maven",
	}

	fmt.Println("Java home directory ", toolsPath["java_home"])

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key,value := range toolsPath {
          fmt.Println(key,value)
	}

	//delete a key-vaule pair from an existing map
	delete(toolsPath, "go_home")
	fmt.Println(toolsPath)
}	
</pre>

Run it
```
go run ./map.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b483c611-d89f-4889-8f4c-c57cd5b18c96" />

## Lab - Golang pointers
Create a filed named pointer.go with the below code
<pre>
package main

import "fmt"

func sayHello ( msgPtr *string ) {

	//Dereferencing - the values stored at address pointed by msgPtr will be printed here
	fmt.Println( "Inside sayHello funciton", *msgPtr )

	//Here the address pointed by msgPtr pointer will be printed
	fmt.Println("Address pointed by msgPtr is", msgPtr )

	//Print the address of msgPtr
	fmt.Println("Address of msgPtr is", &msgPtr )

	//The values stored at the address pointed by msgPtr is assigned to tmp string
	tmp := *msgPtr

	//We are changing the value stored at address pointed by msgPtr pointer
	*msgPtr = tmp + " Golang" + " !"

	fmt.Println("Inside sayHello before returning ", *msgPtr)

}

func main() {
   //declares a string variable name str with value "Hello"
   msg := "Hello"

   fmt.Println("Message before calling sayHello function is ", msg )
   fmt.Println("Address of msg string is ", &msg )

   sayHello( &msg )

   fmt.Println("Message after calling sayHello function is ", msg )
}
</pre>

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c4234890-f981-4cf5-9899-853abbffa2f1" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/82e848e5-53f9-4520-b3f7-2594337c9ea3" />

## Lab - Golang struct with Methods
Create a file named struct.go with the below code
<pre>
package main

import "fmt"

type Rectangle struct {
  length int
  width  int
}

//Method - these are functions that can only be invoked via struct variables
//Area is member function of Rectangle structure
//Area function takes zero arguments and it return an integer
func ( rect Rectangle ) Area() int {
   area := rect.length * rect.width
   return area
}

func ( rect Rectangle ) GetLength() int {
   return rect.length
}

func ( rect Rectangle ) GetWidth() int {
   return rect.width
}

func main() {
     rectangle := Rectangle {
        length: 100,
	width : 200,
     }

     fmt.Printf("Length of rectangle : %d\n", rectangle.GetLength() )
     fmt.Printf("Width  of rectangle : %d\n", rectangle.GetWidth () )
     fmt.Printf("Area of rectangle   : %d\n", rectangle.Area() )
}	
</pre>

Run it
```
go run ./struct.go
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/19215565-faff-43ed-bd0c-3be81cb6e353" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/fb4ba20a-baa3-42a9-9691-429e86671234" />

## Lab - Golang switch case

Create a file named switch-case.go with the below code
<pre>
package main

import "fmt"

func main() {

	var direction string

	fmt.Println("Possible values are east,west,south,north")

	fmt.Print("Enter some direction :")
	fmt.Scanln(&direction)

	switch direction {
	case "east":
	     fmt.Println("You entered direction ", direction)
	case "west":
	     fmt.Println("You entered direction ", direction)
	case "south":
	     fmt.Println("You entered direction ", direction)
	case "north":
	     fmt.Println("You entered direction ", direction)
	default:
	     fmt.Println("Invalid direction", "possible values are east, west, north, south")

	}
}	
</pre>

Run it
```
go run ./switch-case.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/e849d231-23ad-4ea5-b93c-dd67eb2d8f07" />
