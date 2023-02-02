#!/bin/sh -l

echo "Argument one: $1 Argument two: $2"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT