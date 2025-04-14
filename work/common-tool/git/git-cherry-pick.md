## cherry-pick
[cherry-pick](https://git-scm.com/docs/git-cherry-pick)
将已经提交的变更，应用到某个其他分支中

例如：
将指定commitId的变更应用到当前分支中
```bash
git checkout branch_a
git cherry-pick commitId
# 如果有冲突，需要先解决冲突，如果没有冲突，就可以直接合入

```

### git cherry-pick (--continue | --skip | --abort | --quit) 

在 `git cherry-pick` 的上下文中，这四个参数的含义如下：
1. **--continue**：在解决冲突后，继续进行 cherry-pick 操作。
2. **--skip**：跳过当前的 commit，继续 cherry-pick 下一个 commit。
3. **--abort**：取消当前的 cherry-pick 操作，恢复到 cherry-pick 开始前的状态。
4. **--quit**：退出 cherry-pick 操作，但保留当前的合并状态，以便后续继续处理。