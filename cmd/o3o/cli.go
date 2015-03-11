package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mozillazg/go-o3o"
)

const repoURL string = "https://github.com/mozillazg/go-o3o"

func notfound() {
	fail, _ := o3o.O3O("摊手").(string)
	fail += " 没有这个情绪类别哦。。\n要不帮我加上？欢迎 Fork & PR 项目地址: "
	fail += repoURL
	fmt.Println(fail)
	return
}

func main() {
	args := os.Args
	key := ""
	if len(args) >= 2 {
		key = args[1]
	}

	if key == "" {
		fmt.Println(o3o.O3O("*"))
		return
	}
	if key == "ls" || key == "list" {
		fmt.Println(strings.Join(o3o.O3O("").([]string), ", "))
		return
	}

	yan := o3o.O3O(key)
	if yan == nil || yan == "" {
		notfound()
		return
	}
	fmt.Println(yan)
}
