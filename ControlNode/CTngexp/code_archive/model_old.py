from mergexp import *

net = Network('modified_topology', routing == static, experimentnetresolution == False)

# Assume n, m, and p nodes in each group
groupA = [net.node(f'CA{i}') for i in range(8)]
groupB = [net.node(f'Logger{i}') for i in range(4)]
groupC = [net.node(f'Monitor{i}') for i in range(4)]

def create_bipartite(groupA, groupB):
    for i in range(len(groupA)):
        j = i % len(groupB)
        link = net.connect([groupA[i], groupB[j]])
        link[groupA[i]].socket.addrs = ip4(f'10.0.0.{i}/16')
        link[groupB[j]].socket.addrs = ip4(f'10.0.1.{j}/16')

# Connect nodes within groupC (clique topology)
def create_clique(group):
    ip_counter = 0
    for i, node1 in enumerate(group):
        for j, node2 in enumerate(group[i+1:]):
            link = net.connect([node1, node2])
            link[node1].socket.addrs = ip4(f'10.0.2.{ip_counter}/16')
            ip_counter += 1
            link[node2].socket.addrs = ip4(f'10.0.2.{ip_counter}/16')
            ip_counter += 1

create_bipartite(groupA, groupB)
create_clique(groupC)
'''
for i, a_node in enumerate(groupA):
    c_node = groupC[i % len(groupC)]
    link = net.connect([a_node, c_node])
    # Reusing the IP addresses since we're not employing router-like functionality.
    link[a_node].socket.addrs = ip4(f'10.0.0.{i+1}/16')
    link[c_node].socket.addrs = ip4(f'10.0.2.{i+1}/16')

for i, b_node in enumerate(groupB):
    c_node = groupC[i % len(groupC)]
    link = net.connect([b_node, c_node])
    # Reusing the IP addresses since we're not employing router-like functionality.
    link[b_node].socket.addrs = ip4(f'10.0.1.{i+1}/16')
    link[c_node].socket.addrs = ip4(f'10.0.2.{i+1}/16')   
'''

# Connect every node in Group A to every node in Group C
for a_node in groupA:
    for c_node in groupC:
        link = net.connect([a_node, c_node])
        link[a_node].socket.addrs = ip4(f'10.0.0.{groupA.index(a_node)+1}/16')
        link[c_node].socket.addrs = ip4(f'10.0.2.{groupC.index(c_node)+1}/16')

# Connect every node in Group B to every node in Group C
for b_node in groupB:
    for c_node in groupC:
        link = net.connect([b_node, c_node])
        link[b_node].socket.addrs = ip4(f'10.0.1.{groupB.index(b_node)+1}/16')
        link[c_node].socket.addrs = ip4(f'10.0.2.{groupC.index(c_node)+1}/16')


experiment(net)
