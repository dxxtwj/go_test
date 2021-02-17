package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is    kkk@qq.com
email2 is ddd@abc.com.cn
`

func main()  {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// 会在text中找符合MustCompile规则的表达式
	//match := re.FindAllString(text, -1)// 第二个参数说明我们要找多少个这样的规则匹配，-1是所有
	match := re.FindAllStringSubmatch(text, -1)// 子匹配，可以获得加括号提取出来的内容
	for _, m := range match {
		fmt.Println(m)
	}
}
