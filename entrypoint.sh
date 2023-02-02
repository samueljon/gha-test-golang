#!/bin/sh -l

echo "Argument one: $1 Argument two: $2 Argument 3: $3"
./gha-test-golang -name $1 -timeout $2 -namespace $3

time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT