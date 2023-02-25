#!/bin/bash

# Set the name of the new branch
new_branch_name="develop"

# Check if there are any files to move
if [ $# -eq 0 ]; then
  echo "No files to move. Please provide a list of files to move."
  exit 1
fi

# Create a new branch
git checkout -b "$new_branch_name"

# Move each file to the new branch
for file in "$@"; do
  git mv "$file" .
done

# Commit the changes
git commit -m "Move files to $new_branch_name"

# Switch back to the original branch
git checkout -

echo "Files moved to branch $new_branch_name"