# Project Commands Documentation

# Initialize a new Git repository. This creates a new .git directory.
git init

# Check the configured remote repository for your project. Initially, this will show nothing.
git remote -v

# Add a remote repository. This command links your local repository to a remote server, allowing you to push and pull changes.
git remote add origin https://github.com/ahsan03h/I200966_Quiz2.git

# Create a new branch named 'dev' and switch to it immediately.
git checkout -b dev

# Add all changes in the working directory to the staging area, making them ready for a commit.
git add .

# Commit the staged changes to the repository with a message describing what was done.
git commit -m "Add Go source file"

# Switch to the 'main' branch. If 'main' does not exist locally, you might need to fetch it from the remote or create it.
git checkout main

# Merge the 'dev' branch into the 'main' branch. This command may prompt you to resolve conflicts if they exist.
git merge dev -m "Merge dev branch"

# Push the changes from the 'dev' branch to the remote repository. This updates the remote 'dev' branch with your local changes.
git push origin dev

# List all branches in your repository. This command shows both local and remote branches.
git branch

# Switch back to the 'dev' branch. Useful if you need to make more changes or commits to 'dev' after merging.
git checkout dev

# Show the status of the working directory and staging area. It lets you see which changes have been staged, which haven't, and which files aren't being tracked by Git.
git status

# Push the changes from the 'main' branch to the remote repository. This command is incorrectly ordered here; it should ideally be used after commits to 'main'.
git push origin main

# Display the commit logs. This command shows the commit history, allowing you to see the details of each commit.
git log
