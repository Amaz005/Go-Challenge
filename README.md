# Challenges and Interview

## Golang basic
- you can find the result of this challenge in `main.go` find using
```sh
go run main.go
```

## Git Basic

1. How do I create a new branch in Git?
- By using `git branch <new-branch>` or `git checkout -b ＜new-branch＞`
2. How do I switch branches in Git?
- Using `git checkout <branch>`
3. How do I push changes to a remote repository?
- First you need to to run <b>git add</b> to stage the change in your local repo
- Then run <b>git commit -m "your-comment"</b> to commit the change in your local repo
- Then push the commit from local repo to your remote repository by run <b>git push -u remote-name branch-name</b> where remote-name is the nickname the local repo uses for the remote repositories and branch-name is the name of the branch to push to the repository
4. Write a script that will take a list of files and move them to a new branch.
- Run the script with a list of files as arguments, for example ./move-files-to-branch.sh file1.txt file2.txt.
- The script will create a new branch called develop, move the specified files to it, commit the changes, and switch back to the original branch.
5. Write a script that will clone a remote repository and checkout a specific branch.
- Run the script by running ./clone-and-checkout.sh.
6. Write a script that will compare two different branches and list the differences in their contents
7. Create a script that will list all commits made by a specific user
- Run the script by running ./list-commits-by-user.sh