package toolkit

import (
	"regexp"
	"strings"
)

// 储存中文数字到阿拉伯数字的映射
var numMap = make(map[string]string)

func init() {
	numMap["一"] = "1"
	numMap["二"] = "2"
	numMap["三"] = "3"
	numMap["四"] = "4"
	numMap["五"] = "5"
	numMap["六"] = "6"
	numMap["七"] = "7"
	numMap["八"] = "8"
	numMap["九"] = "9"
}

// StringPurgerForMatch 去除标点空格，英文字母转为小写，中文数字转化为阿拉伯数字
func StringPurgerForMatch(txt string) string {
	//去掉标点符号
	regp := regexp.MustCompile(`[\pP]*`)
	txt = regp.ReplaceAllString(txt, "")
	//去掉空格
	txt = strings.Replace(txt, " ", "", -1)
	//全部转为小写
	txt = strings.ToLower(txt)
	//中文数字转换为阿拉伯数字
	for k, v := range numMap {
		txt = strings.Replace(txt, k, v, -1)
	}
	return txt
}
