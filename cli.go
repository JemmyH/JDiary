package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	user2 "os/user"
	"time"
)

/*
* @CreateTime: 2019/12/14 19:07
* @Author: hujiaming
* @Description:
 */

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  create 创建一个日记本，并指定Owner\n      -owner NAME  创建一个日记本")
	//fmt.Println("  add -owner OWNER -content CONTENT -notes NOTES - 保存一篇日记到OWNER的日记本中")
	fmt.Println("  add 添加日记\n      -owner OWNER 日记创建者\n      -content CONTENT 日记内容\n      -notes NOTES 日记额外信息(可选)")
	fmt.Println("  print 显示日记\n      -owner OWNER 日记所有者\n      -type TYPE 显示类型(head 从旧到新,tail 从新到旧)\n      -simple 是否简要显示\n      -n 显示几条日记(默认7条)")
}
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
func (cli *CLI) Run() {
	cli.validateArgs()
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printCmd := flag.NewFlagSet("print", flag.ExitOnError)

	createName := createCmd.String("owner", "", "the owner of this DiaryBook")
	addName := addCmd.String("owner", "", "the owner of this DiaryBook. Default is the current user.")
	addContent := addCmd.String("content", "", "the words you want to write to your DiaryBook")
	addNotes := addCmd.String("notes", "", "the additional information about this diary")
	printSimple := printCmd.Bool("simple", false, "whether use simple mode to print diary")
	printName := printCmd.String("owner", "", "the owner of this DiaryBook")
	printNum := printCmd.Int("n", 7, "the max number of diary you want to print. Default is 7.")
	printType := printCmd.String("type", "tail", "the way to print. 'head' or 'tail' is needed")

	switch os.Args[1] {
	case "create":
		err := createCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "add":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := printCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	}

	if createCmd.Parsed() {
		name := ""
		if *createName == "" {
			user, err := user2.Current()
			if err != nil {
				fmt.Println("get current user failed")
				os.Exit(1)
			}
			name = user.Username
		} else {
			name = *createName
		}
		createDiaryBook(name)
	}

	if addCmd.Parsed() {
		if *addName == "" {
			fmt.Println("please specify a user")
			os.Exit(1)
		}
		if *addContent == "" {
			fmt.Println("you cannot write an empty content to the diary")
			os.Exit(1)
		}
		addDiary(*addName, *addContent, *addNotes)
	}

	if printCmd.Parsed() {
		printLimitDiary(*printName, *printType, *printNum, *printSimple)
	}
}

func createDiaryBook(name string) {
	CreateNewDiaryBook(name)
}
func addDiary(name, content, notes string) {
	db := GetDiaryBook(name)
	diary := &Diary{
		Content:   EncryptString(content),
		Notes:     EncryptString(notes),
		Timestamp: time.Now().Unix(),
	}
	db.AddDiary(diary)
}

func printLimitDiary(name, types string, num int, simple bool) {
	switch types {
	case "head":
		// 打印从head开始的前num条
		db := GetDiaryBook(name)
		dbi := DiaryBookIterator{
			CurrentKey: []byte("head"),
			DB:         db.DB,
		}
		for num > 0 {
			diary := dbi.Next()
			if diary == nil {
				break
			}
			if simple {
				fmt.Println(diary.SimpleString())
			} else {
				fmt.Println(diary)
			}
			if diary.NextID == nil {
				break
			}
			num--
		}
	case "tail":
		db := GetDiaryBook(name)
		dbi := DiaryBookIterator{
			CurrentKey: []byte("tail"),
			DB:         db.DB,
		}
		for num > 0 {
			diary := dbi.Prev()
			if diary == nil {
				break
			}
			if simple {
				fmt.Println(diary.SimpleString())
			} else {
				fmt.Println(diary)
			}
			if diary.PrevID == nil {
				break
			}
			num--
		}
	case "day":
		// TODO: 某一天以及以后几天
	}
}
