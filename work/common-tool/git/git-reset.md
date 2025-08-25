## 本地committed后，如何退回到工作区

查看当前本地的提交状态:
> git status

回退本地的最新提交; 将提交的信息回退到工作区
> git reset --soft HEAD~1

回退本地的最新提交；不要最新的提交，丢弃最新的提交  -- 慎用
> git reset --hard HEAD~1