#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Please provide the number of monitor directories to remove."
    exit 1
fi

# Get the input value
input=$1

# Check if input is a number
if ! [[ $input =~ ^[0-9]+$ ]]; then
    echo "The input must be a positive number."
    exit 1
fi

# Loop and execute rm -rv
for (( i=1; i<=input; i++ ))
do
    echo "rm -rv Monitor$i"
    # Uncomment the next line to actually execute the command
    rm -rv "Monitor$i"
done
