# JDiary
  个人记事本  命令行版本

### 1. 为什么做这个？
有时候在敲代码的时候，经常会由一个情景得到一些想法，对于经常使用命令行的我来说，打开一个有图形界面的软件去记录还不如直接在命令行敲。

(没错，我就是在装x (￣▽￣)／)

### 2. 如何使用？
```shell
# 安装 
git clone https://github.com/JemmyH/JDiary
cd JDiary
go build -a -v -o diary *.go
chmod a+x ./diary

# 使用
Usage:
  create 创建一个日记本,并指定Owner
      -owner NAME  创建一个日记本
  add 添加日记
      -owner OWNER 日记创建者
      -content CONTENT 日记内容
      -notes NOTES 日记额外信息(可选)
  print 显示日记
      -owner OWNER 日记所有者
      -type TYPE 显示类型(head 从旧到新,tail 从新到旧)
      -simple 是否简要显示
      -n 显示几条日记(默认7条)
```
### 3. TODO
- [ ] 使用密码操作
- [ ] 指定删除哪一条
- [ ] 指定显示某天到某天