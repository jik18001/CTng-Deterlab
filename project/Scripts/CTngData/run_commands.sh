#!/bin/bash

# Check if the user has provided a number of iterations
if [ $# -eq 0 ]; then
    echo "Please enter the number of iterations:"
    read iterations
else
    iterations=$1
fi

# Check if iterations is a number
re='^[0-9]+$'
if ! [[ $iterations =~ $re ]] ; then
   echo "Error: Not a number" >&2; exit 1
fi

# Loop for the specified number of iterations
for (( i=1; i<=iterations; i++ ))
do
    echo "Iteration $i"

    # Navigate to the CTngexp/Topology directory
    cd ../CTngexp/Topology || exit

    # Run the go test
    go test

    # Navigate back to the CTngexp directory
    cd ..

    # Run ansible playbooks
    ansible-playbook -i inv.ini topo.yml
    ansible-playbook -i inv.ini playbook.yml
    ansible-playbook -i inv.ini collect.yml

    # Navigate to the CTngData directory
    cd ../CTngData || exit

    # Run the Python script
    python3 computes.py
done
