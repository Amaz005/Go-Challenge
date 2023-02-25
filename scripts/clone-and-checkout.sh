#!/bin/bash

# Set the URL of the remote repository
remote_repo_url="https://github.com/Amaz005/Go-Challenge.git"

# Set the name of the branch to checkout
branch_name="develop"

# Clone the remote repository
git clone "$remote_repo_url"

# Change into the newly cloned repository
cd repo

# Checkout the desired branch
git checkout "$branch_name"

echo "Repository cloned and $branch_name branch checked out."
