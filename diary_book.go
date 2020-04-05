package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
	user2 "os/user"
	"time"
)

/*
* @CreateTime: 2019/12/14 17:15
* @Author: hujiaming
* @Description: 日记本
 */
const dbFile = ".diary_%s.db"
const dbBucket = "diary"
const headKey = "head"
const tailKey = "tail"

type DiaryBook struct {
	Head []byte
	Tail []byte
	//Diaries []*Diary
	DB *bolt.DB
}

func CreateNewDiaryBook(name string) *DiaryBook {
	dbFile := fmt.Sprintf(dbFile, name)
	if dbExist(dbFile) {
		fmt.Println("DiaryBook already exists")
		os.Exit(1)
	} else {

	}

	var headByte, tailByte []byte

	head := NewHead(name)
	tail := NewTail(name)
	head.ID = []byte(headKey)
	tail.ID = []byte(tailKey)

	tail.PrevID = head.ID
	head.NextID = tail.ID

	fmt.Println(head)
	fmt.Println(tail)

	db, err := bolt.Open(getDbFile(name), 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		var b *bolt.Bucket
		var err error
		b = tx.Bucket([]byte(dbBucket))
		if b == nil {
			b, err = tx.CreateBucket([]byte(dbBucket))
			if err != nil {
				return err
			}
		}

		// 添加头
		err = b.Put([]byte(headKey), head.Serialize())
		if err != nil {
			return err
		}

		// 添加尾
		err = b.Put([]byte(tailKey), tail.Serialize())
		if err != nil {
			return err
		}
		headByte = []byte(headKey)
		tailByte = []byte(tailKey)
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &DiaryBook{Head: headByte, Tail: tailByte, DB: db}
}

func GetDiaryBook(name string) *DiaryBook {
	dbFile := getDbFile(name)
	if dbExist(dbFile) == false {
		fmt.Printf("No existing DiaryBook found for %s. \nUse `diary create -owner %s` to create one first.\n", name, name)
		os.Exit(1)
	}

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	return &DiaryBook{
		Head: []byte(headKey),
		Tail: []byte(tailKey),
		DB:   db,
	}
}

func (db *DiaryBook) AddDiary(diary *Diary) {
	diary.SetID(time.Now())
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))
		diaryInDb := b.Get(diary.ID)
		// 已经存在这个key
		if diaryInDb != nil {
			return nil
		}

		tailBytes := b.Get([]byte(tailKey))
		//fmt.Println(tailKey)
		tail := DeserializeDiary(tailBytes)
		//fmt.Println("fuck")
		//fmt.Println(tail)
		preBytes := b.Get(tail.PrevID)
		//fmt.Println(preBytes)
		pre := DeserializeDiary(preBytes)
		//fmt.Println(pre)

		pre.NextID = diary.ID
		diary.NextID = tail.ID
		diary.PrevID = pre.ID
		tail.PrevID = diary.ID
		tail.Timestamp = time.Now().Unix()

		// 更新pre(因为pre的数据有变)
		err := b.Put(pre.ID, pre.Serialize())
		if err != nil {
			return err
		}
		// 保存当前diary
		err = b.Put(diary.ID, diary.Serialize())
		if err != nil {
			return err
		}
		// 保存tail
		err = b.Put(tail.ID, tail.Serialize())
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

func getDbFile(name string) string {
	user, _ := user2.Current()

	return user.HomeDir + "/" + fmt.Sprintf(dbFile, name)
}
func dbExist(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}
