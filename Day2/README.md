# Day 2

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
ansible init nginx
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
Under the Source Control url, type "https://github.com/tektutor/terraform-may-2025.git"
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
