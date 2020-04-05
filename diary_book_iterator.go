package main

import (
	"github.com/boltdb/bolt"
	"log"
)

/*
* @CreateTime: 2019/12/14 18:13
* @Author: hujiaming
* @Description:
 */

type DiaryBookIterator struct {
	CurrentKey []byte
	DB         *bolt.DB
}

// Next 从过去到现在迭代
func (i *DiaryBookIterator) Next() *Diary {
	var diary *Diary

	err := i.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))
		encodedDiary := b.Get(i.CurrentKey)
		if len(encodedDiary) == 0{
			return nil
		}
		diary = DeserializeDiary(encodedDiary)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}


	i.CurrentKey = diary.NextID
	return diary
}

// Prev 从现在向过去迭代
func (i *DiaryBookIterator) Prev() *Diary {
	var diary *Diary

	err := i.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))
		encodedDiary := b.Get(i.CurrentKey)
		if len(encodedDiary) == 0 {
			return nil
		}
		diary = DeserializeDiary(encodedDiary)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	//if diary == nil {
	//	return nil
	//}
	i.CurrentKey = diary.PrevID
	return diary
}
