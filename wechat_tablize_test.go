package toolkit

import (
	"fmt"
	"testing"
)

func TestWechatStr2dTablizeRender(t *testing.T) {
	t.Helper()
	str2dA := [][]string{{"Code", "节目名称", "内容类型", "内容方"}, {"jaiojfiojaoifjoasjf", "打工", "综艺", "爱奇艺"}, {"adfagryd1232asdffsa", "大地飞歌", "少儿", "咪咕常规"}}
	outA, err := WechatStr2dTablizeRender(str2dA)
	fmt.Printf("【input】: \n%v\n【output】: \n%v\n【(returned) error】: %v\n", str2dA, outA, err)
}
