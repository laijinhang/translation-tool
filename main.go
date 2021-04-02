package main

import (
	"bufio"
	"fmt"
	"os"
	"translation-tool/translation"
)

func main() {
	var text string
	for {
		fmt.Print("请输入：")
		fmt.Scan()
		text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		text = text[:len(text)-1]
		fmt.Println("翻译结果：")
		fmt.Println("----------------------------------")
		res := translation.Translation(text)
		for i, _ := range res {
			fmt.Printf("%s翻译结果>> %s\n\n", res[i].Name, res[i].Results)
		}
		fmt.Println("----------------------------------")
	}
}
