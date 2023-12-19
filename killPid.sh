#!/bin/bash

# Check if the PID file exists
PID_FILE="pid"

if [ -f "$PID_FILE" ]; then
    # Read the PID from the file
    PID=$(cat $PID_FILE)

    # Check if the process with this PID exists and kill it
    if ps -p $PID > /dev/null
    then
       echo "Killing process with PID $PID"
       kill $PID
    else
       echo "No process found with PID $PID"
    fi

    # Optionally, remove the PID file
    # rm $PID_FILE
else
    echo "PID file does not exist."
fi