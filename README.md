# JDiary
  个人记事本  命令行版本

### 1. 为什么做这个？
有时候在敲代码的时候，经常会由一个情景得到一些想法，对于经常使用命令行的我来说，打开一个有图形界面的软件去记录还不如直接在命令行敲。
(没错，我就是在装x (￣▽￣)／)

### 3. 如何使用？
```shell script
# 安装 

go build -a -v -o diary *.go
chmod a+x ./diary

# 使用
#    打印使用方法
        ./diary 
#    创建一个日记本 (创建一个属于Jemmy的日记本)
        ./diary create -name Jemmy
#    添加日记
        ./diary add -owner Jemmy -content "your diary content" -notes "comment for this diary"
#    显示日记(head从过去到现在的顺序打印简单内容,-simple 表示简单打印, -n 表示打印几条)
        ./diary print -owner Jemmy -type head -n 5
```