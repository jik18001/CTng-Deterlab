# CTng-Deterlab

This repo contains a step by step guide to setup up the experiments on Deterlab.

## Software list:
- **OS:** Ubuntu 22.04
- **g++:** 11.4.0
- **ansible:** 2.10.8
- **python:** 3.10.6
- **jinja2:** 3.0.3
- **golang:** 1.19.9

## Requirements:
- All the yml files should be in the /project/CTngexp folder on the XDC VM (Ubuntu 22.04)
- go1.19.9.linux-amd64.tar.gz needs to be the in the /project folder
- the repo used is up to date
 ```
go get github.com/jik18001/CTngV2@latest
```
## Step 1: Config generation and setup
### Step 1.1: Generate the Node List for Automation
- Run the following command:
```
python3 genini.py
```
### Step 1.2: Streamline Config Generation (mainly crypto related)
- Edit `gen_test.go`.
- Run the test:
```
go test
```
## Step 2: Make sure we have the correct version of ansible-playbook (2.10.8) and Jinja2 (3.03) 
#### Run the following commands on both the XDC VM and the Control Node VM
### command to check version: 
```
ansible --version 
```
```
pip show Jinja2
```
### command to install pip if not present
```
sudo apt update
```
```
sudo apt install python3-pip --fix-missing
```
### command to install Jinja2 if not present
```
pip install Jinja2 == 3.03
```
## Step 3: Setup the Control Node and all other nodes
#### Run the following command on XDC VM
```
ansible-playbook -i control.ini ssh.yml
```
## Step 4: Environment Setup and Build the experiment
#### Run these commands on the XDC VM
```
ansible-playbook -i control.ini ssh.yml
```
```
ansible all -i inv.ini -m ping
```
```
ansible-playbook -i inv.ini gpp.yml
```
```
ansible-playbook -i inv.ini go.yml
```
```
ansible-playbook -i inv.ini distribute.yml
 ```
## Step 5: Check Topology and Run Tests
- check ansible.cfg
- In the `CTng/Topology` folder, verify `topo_test.go` to ensure the number matches the intended test setup.
- In the `CTng/DMLCG` folder, execute:
```
go test
```
- In the `CTngexp` folder:
- Distribute the new DMLCG config:
  ```
  ansible-playbook -i inv.ini topology.yml
  ```
- Distribute other server settings (non-crypto related, exp variables):
  ```
  ansible-playbook -i inv.ini dist.yml
  ```

## Additional deterlab checks
## Part 1: Experiment Node resolution
Ensure your `/etc/resolv.conf` file is configured properly. It should look something like this:
```
search infra.ctng32.cliquetest.jik18001 mergev1-xdc.svc.cluster.local svc.cluster.local cluster.local mgmt.mdr.mergetb.net
nameserver 172.30.0.1
nameserver 10.2.0.10
option ndots:5
```
## Part 2: Check Bandwidth

### Installation
First, install `iperf3` on the nodes:
```
sudo apt -y install iperf3
```

### Running the Test
1. Start `iperf3` in server mode on one of the nodes (e.g., Monitor1):

    ```
    iperf3 -s
    ```

2. On another node, initiate the client to connect to the server:

    ```
    iperf3 -c Monitor1
    ```


