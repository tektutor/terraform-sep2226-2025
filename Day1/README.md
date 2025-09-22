<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/7fc02cb2-f5a5-4fe9-9702-90d432c091c2" /># Day 1

## Lab - Installing Ansible core in Ubuntu
```
sudo apt update && apt install -y ansible
```

## Info - Configuration Management Tools
<pre>
- are helpful in automating system administrative tasks
- i.e if you have a machine with some OS pre-installed, we can install/upgrade softwares on those servers
- the servers managed by Configuration Management Tools are generally referred as Nodes
- the Nodes can be 
  - an Unix/Linux/Mac Server
  - an AWS ec2 instance, or an Azure Virtual Machine
  - a on-prem bare-metal server or a virtual machine in your data-center
  - a Network Switch/Router
  - a Windows Server
- examples
  - Puppet
  - Chef
  - Salt/SaltStack
  - Ansible
</pre>

## Info - Puppet Overview
<pre>
- this is one of the oldest Configuration Management tool
- came around year 2005
- uses Puppet's proprietrary language for automation
- follows client/server architecture
- the machines that are managed by Puppet are called Puppet Nodes
- Puppet Nodes can be a
  - Windows Server
  - Unix/Linux/Mac Server
  - Network Switch/Routers
  - Raspberry Pi
- DSL used is Puppet's proprietary language
- DSL means Domain Specific Language - i.e the language in which the automation code is written
- Software requirements on the Puppet Nodes
  - Puppet Agent must be installed
- Puppet automation code is written on your laptop/desktop
- Puppet is where the Puppet automation scripts are maintained
- the Puppet agents running on the Puppet nodes, connects to the server periodically and then they poll to check if there are
  any new scripts, update on the existing scripts available, when they see new scripts or updated scripts, it is the 
  responsibility of the Puppet agent to pull the automation scripts onto the Puppet Node, run it and updates the status back to
  the Puppet Server
  - this architecture is called Pull architecture
  - comes in 2 flavours
    1. Puppet opensource
    2. Puppet Enterprise ( licensed, commercial product )
- drawbacks
  - installation is time consuming
  - learning curve is steep as we need learn the Puppet's proprietary language in order to master Puppet
</pre>

## Info - Chef Overview
<pre>
- Chef also follows similar architecture as Puppet
  - client/server architecture
  - pull based architecture
- the servers managed by Chef are referred as Chef Nodes
  - Chef Nodes might be Windows Server, Unix/Linux/Mac server, ec2 instance or azure VMs, Network switches/routers
- uses Ruby as the DSL to write automation code
- installation and learning is very difficult
- comes in 2 flavours
  1. Chef opensource variant
  2. Chef Enterprise
</pre>

## Info - Ansible
<pre>
- Ansible is agentless
- developed in Python by Michael Deehan
- Michael Deehan was a former employee of Red Hat
- Michael Deehan started a company named Ansible Inc
- Ansible Inc made Ansible core an open source product
- uses a simple architecture as compared to Puppet/Chef
- there is no need to install proprietary tools on Ansible nodes 
- software requirements on Ansible nodes are
  - Windows Servers
    - WinRM must be supported
    - PowerShell should be installed
  - Unix/Linux/Mac Server
    - SSH Server
    - Python should be installed
- YAML is the DSL used in Ansible
- comes in 3 flavours
  - Ansible Core 
    - open source
    - supports only command-line 
  - AWX 
    - open source
    - supports web console
    - developed on top of Ansible core
  - Ansible Automation Platform 
    - formerly called Ansible Tower - Red Hat licensed Enterprise product
    - supports web console
    - developed on top of AWX open source product
    - comes with world-wide support from Red Hat ( an IBM company )
</pre>


## Info - Ansible Inventory
<pre>
- captures connectivity details to Unix, Linux, Windows ansible nodes
- two types
  1. Static Inventory
  2. Dynamic Inventory
</pre>

## Lab - Create a custom ubuntu ansible node container image
```
cd ~/terraform-sep2226-2025
git pull
cd Day1/ansible/CustomAnsibleNodeDockerImages/ubuntu
cat Dockerfile

ssh-keygen
cp ~/.ssh/id_ed25519.pub authorized_keys

docker build -t tektutor/ubuntu-ansible-node:latest .

docker images | grep tektutor
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/95258d1d-d7ae-48b8-ad3a-1e2c74999b49" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/581f14e3-b0fa-4dd1-bbc3-96eff14da3ff" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3d3658c5-8f6d-4304-bbd3-fa6d325eccbf" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/147aa580-8027-4431-a0d1-3b18173dc6f6" />

## Lab - Let's create couple of ubuntu ansible node containers using our custom docker image
```
docker images
docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:latest
```
<img width="2848" height="1390" alt="image" src="https://github.com/user-attachments/assets/a808ad47-e578-4e34-89f1-02656b1a143a" />

List and see if the ubuntu1 and ubuntu2 containers are running
```
docker ps
```
<img width="1942" height="745" alt="image" src="https://github.com/user-attachments/assets/44eba12f-9715-4bc6-9b63-43e8b082cd52" />

## Lab - Running ansible ad-hoc command
Check if you are able to SSH into ubuntu1 and ubuntu2 ansible contaible nodes
```
ssh -p 2001 root@localhost
ls
hostname
hostname -i
exit

ssh -p 2002 root@localhost
ls
hostname
hostname -i
exit
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9e7a7157-8963-4f0d-a0d2-c0bd231c6880" />

Run the ansible-hoc command to ping the ansible nodes
```
cd ~/terraform-sep2226-2025
git pull
cd Day1/ansible
cat inventory
ansible -i inventory all -m ping
ansible -i inventory ubuntu1 -m ping
ansible -i inventory ubuntu2 -m ping
ansible -i inventory all -m shell -a "hostname -i"
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/18850abd-9e94-482c-ad6a-fdc20e2734c5" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a3e0700b-58db-48d3-99f2-9c9d6767b5e8" />

## Lab - Running your first ansible playbook
```
cd ~/terraform-sep2226-2025
git pull
cd Day1/ansible
cat inventory
cat ping-playbook.yml
ansible-playbook -i inventory ping-playbook.yml 
ansible-playbook -i inventory ping-playbook.yml --limit=ubuntu1
ansible-playbook -i inventory ping-playbook.yml --limit=ubuntu2
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/6cf6394b-a64a-4389-8a5c-2e4a0a71ed48" />

## Lab - Running install nginx playbook
```
cd ~/terraform-sep2226-2025
git pull
cd Day1/ansible/un-refactored
cat inventory
cat install-nginx-playbook.yml
ansible-playbook -i inventory install-nginx-playbook.yml

curl http://localhost:8001
curl http://localhost:8002
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3c017405-af3e-401c-9444-daf1bda1ff99" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/ae2824d4-4713-43e0-848e-795380366849" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/711bb639-4008-4d02-b7bd-a261b596d0b8" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/8570a5de-998d-4284-9779-75f1c6160c6a" />
