# Day 1

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
