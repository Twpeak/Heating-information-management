package main

import (
	"fmt"
	"regexp"
)

func main() {
	rule := ""
	if _,err := regexp.CompilePOSIX(rule); err != nil{
		fmt.Printf("您的规则中的正则表达式%s不合法,err：%v",rule,err)
	}
}
