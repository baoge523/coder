
## business scene record about git

### first: how to use git rebase command
describe:
```text
i have two git-branch, and i need to rebase one to anther one
branch info:
one: release_a
anther one: dev_a

now,i want to rebase release_a all changes to dev_a
```
command:
```bash
git chekcout release_a
git pull  // pull the release_a newest data
git checkout dev_d
git pull // pull the dev_d newest data
git rebase release_a
// while the rebase, if it has conflicts, we need to resolve the conflicts
git push -f // push the data to remote with strongly

// tips: we do not neet to execute git pull --rebase, 
// if we do,that we need to resolve the all conflicts again
```




