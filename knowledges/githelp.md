git help

##### git fetch origin localv1.0
##### git merge origin/localv1.0 (git pull)

##### 若失败
##### git stash
##### git merge origin/localv1.0 --commit "remark"

##### 或先 add commit 本地修改之后 merge

##### 若执行git stash
##### git stash pop

##### git add . / git add filename
##### git commit -m "remark"
##### git pull origin localv1.0

##### 查看当前信息
##### git status 

##### 本地删除指定的commit
##### git reset HEAD [fileName] / git reset commit-id 

##### 撤销本地修改的文件
##### git checkout file-name 

##### 撤销某次操作，此次操作之前和之后的commit和history都会保留，并且把这次撤销
##### git revert HEAD / git revert HEAD^

##### 撤销指定的版本，撤销也会作为一次提交进行保存。
##### git revert commit-id 

##### 提交一个新的版本，将需要revert的版本的内容再反向修改回去，版本会递增，不影响之前提交的内容
##### git revert

##### 查看远程所有分支
##### git branch -a 

##### 删除远程分支
##### git push origin --delete <branchName>

##### 重命名本地分支
##### git branch -m devel develop

##### git remote show origin

##### 创建并切换到分支
##### git checkout -b dev
#####  git push --set-upstream origin dev

##### 删除本地 / 远程分支
##### git branch -d dev /  git push origin --delete <branchName>

##### 合并到master分支
##### git checkout master && git pull origin localv1.0
##### git merge origin/localv1.0
##### git push

##### git 放弃本地修改 强制更新
##### git fetch --all
##### git reset --hard origin/master

##### 让master与develop完全相同
##### git checkout master
##### git reset --hard develop // 将本地的旧分支 master 重置成 develop
##### git push origin master --force // 再推送到远程仓库


##### 删除分支
##### git branch -D develop
##### git push origin :develop (origin 后面有空格)

##### 清理分支
##### git remote prune origin

##### 自动生成model文件jdk_user
##### bee generate appcode -tables="jdk_user"  -driver=mysql -conn="root:XZhenzhuangtodb@(106.14.194.171:59027)/jdkopen" -level=1

##### 添加tag
##### git tag -a v0.1.0 -m 'init'
##### git push origin --tags/git push origin v0.1.0
