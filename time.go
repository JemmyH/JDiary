package main

import "time"

/*
* @CreateTime: 2019/12/14 17:42
* @Author: hujiaming
* @Description:
 */

func TimeStrSub(source string, stepDay int) string {
	s, _ := time.Parse("20160102150405", source)
	return s.Add(time.Hour * time.Duration(24*stepDay)).Format("20160102150405")
}
