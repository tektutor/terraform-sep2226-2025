# ansible-docker-dynamic-inventory

You can use the dynamic_inventory.py in the place of ansible static inventory.

For instance to execute an Ansible ad-hoc command, you may try the below command
ansible -i ./dynamic_inventory all -m ping

As we have ansible.cfg as part of this project folder, it already takes care of pointing the dynamic_inventory.py 
to be used as inventory in the place of a static inventory, hence you can execute ansible commands as below

ansible all -m ping<br>
ansible-playbook ping.yml
