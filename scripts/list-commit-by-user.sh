#!/bin/bash

# Set the name of the user to filter by
user_name="amaz005"

# Set the name of the repository to search in
repo_name="Go-Challenge"

# Change into the repository directory
cd "$repo_name"

# List all commits made by the user
git log --author="$user_name"

echo "Commits made by $user_name in $repo_name:"
