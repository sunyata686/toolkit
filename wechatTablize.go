package toolkit

import (
	"fmt"
)

// （为微信）二维字符串数组对齐，不足补空格,默认第一行为表头；
// 如果str2d行宽度不一致，报错。
func WechatStr2dTablizeRender(str2d [][]string) (string, error) {
	if err := checkStr2dRowLenAligned(str2d); err != nil {
		return "", err
	}

	str2d = wechatArr2dFillAndAlign(str2d)
	maxWith := getWechatTablizeRowFullWidth(str2d)
	map1 := getWechatCharWidthMap()
	var s string
	for i := 0; i < len(str2d); i++ {
		if i == 1 {
			for n := float32(0); n < maxWith/map1["-"]*0.95; n++ {
				s += "-"
			}
			s += "\n"
		}
		for j := 0; j < len(str2d[i]); j++ {
			if j == 0 {
				s = s + "  " + str2d[i][j]
			} else {
				s = s + "  |  " + str2d[i][j]
			}

		}
		s += "\n"
	}
	return s, nil
}

// 检查str2d宽度是否一致
func checkStr2dRowLenAligned(str2d [][]string) error {
	if len(str2d) == 0 {
		return nil
	}
	rowLen := len(str2d[0])
	for i, row := range str2d {
		if len(row) != rowLen {
			return fmt.Errorf("二维数组列宽度不一致： RowLength %d excepted, but rowLength %d at rowIndex %d", rowLen, len(row), i)
		}
	}

	return nil
}

// 获取wehcat各字符的宽度map
func getWechatCharWidthMap() map[string]float32 {
	//初始数据：1080P 100%
	//	2k      asus    msi(125)   msi(150)(比例无关)
	//数字  11      8.0     11         13.2
	//文字  18      13.6    18          21.5
	//空格  5       4        5          6
	// M  18        14
	m := make(map[string]float32)
	//m[" "] = 4.0
	m[" "] = 3.75
	m["!"] = 4.0
	m["！"] = 14.0
	m["~"] = 10.0
	m["·"] = 3.0
	m["`"] = 4.0
	m["@"] = 14.0
	m["#"] = 9.0
	m["$"] = 8.0
	m["%"] = 12.0
	m["^"] = 10.0
	m["&"] = 12.0
	m["*"] = 6.0
	m["("] = 5.0
	m[")"] = 5.0
	m["（"] = 14.0
	m["）"] = 14.0
	m["-"] = 6.0
	m["—"] = 15.0
	m["_"] = 6.0
	m["="] = 10.0
	m["+"] = 10.0
	m["["] = 5.0
	m["]"] = 5.0
	m["\\"] = 6.0
	m["{"] = 5.0
	m["}"] = 5.0
	m["|"] = 4.0
	m["【"] = 14.0
	m["】"] = 14.0
	m["、"] = 14.0
	m["；"] = 14.0
	m["‘"] = 9.0
	m["："] = 14.0
	m["“"] = 14.0
	m[";"] = 3.0
	m["'"] = 4.0
	m[":"] = 3.0
	m["\""] = 6.0
	m[","] = 3.0
	m["."] = 3.0
	m["/"] = 6.0
	m["<"] = 10.0
	m[">"] = 10.0
	m["?"] = 7.0
	m["，"] = 14.0
	m["。"] = 14.0
	m["《"] = 14.0
	m["》"] = 14.0
	m["？"] = 14.0

	//数字部分
	for i := 0; i < 10; i++ {
		m[fmt.Sprintf("%d", i)] = 8.3
	}
	//13.6 /18 = 0.75
	scale := float32(0.96)
	//字母部分
	m["a"] = 8.0 * scale
	m["b"] = 9.0 * scale
	m["c"] = 7.0 * scale
	m["d"] = 9.0 * scale
	m["e"] = 8.0 * scale
	m["f"] = 5.0 * scale
	m["g"] = 9.0 * scale
	m["h"] = 9.0 * scale
	m["i"] = 4.0 * scale
	m["j"] = 4.0 * scale
	m["k"] = 8.0 * scale
	m["l"] = 4.0 * scale
	m["m"] = 13.0 * scale
	m["n"] = 9.0 * scale
	m["o"] = 9.0 * scale
	m["p"] = 9.0 * scale
	m["q"] = 9.0 * scale
	m["r"] = 5.0 * scale
	m["s"] = 6.0 * scale
	m["t"] = 5.0 * scale
	m["u"] = 9.0 * scale
	m["v"] = 7.0 * scale
	m["w"] = 11.0 * scale
	m["x"] = 7.0 * scale
	m["y"] = 7.0 * scale
	m["z"] = 7.0 * scale

	m["A"] = 10.0 * scale
	m["B"] = 9.0 * scale
	m["C"] = 9.0 * scale
	m["D"] = 11.0 * scale
	m["E"] = 8.0 * scale
	m["F"] = 7.0 * scale
	m["G"] = 10.0 * scale
	m["H"] = 11.0 * scale
	m["I"] = 4.0 * scale
	m["J"] = 6.0 * scale
	m["K"] = 9.0 * scale
	m["L"] = 7.0 * scale
	m["M"] = 14.0 * scale
	m["N"] = 11.0 * scale
	m["O"] = 11.0 * scale
	m["P"] = 9.0 * scale
	m["Q"] = 11.0 * scale
	m["R"] = 9.0 * scale
	m["S"] = 8.0 * scale
	m["T"] = 8.0 * scale
	m["U"] = 10.0 * scale
	m["V"] = 9.0 * scale
	m["W"] = 14.0 * scale
	m["X"] = 9.0 * scale
	m["Y"] = 8.0 * scale
	m["Z"] = 9.0 * scale
	//log.Println("开始检查map是否有错误")
	//for k, v := range m {
	//	if strings.Contains(k, " ") {
	//		log.Println(k, " : ", v)
	//	}
	//}
	return m
}

// 获取单个字符在wechat中的宽度
func getWechatCharWidth(s string) float32 {
	m := getWechatCharWidthMap()
	w, ok := m[s]
	if ok {
		return w
	} else {
		return 13.6
	}
}

// 获取字符串的在wehchat中的宽度
func getWechatStrWidth(s string) float32 {
	p := new(float32)
	for _, ch := range s {
		*p = *p + getWechatCharWidth(fmt.Sprintf("%c", ch))
	}

	return *p
}

// 获取二维字符串数组str2d每列在wehcat中的最大宽度
func getWechatTablizeColMaxWidth(str2d [][]string) []float32 {
	colNum := len(str2d[0])
	resArr := []float32{}
	P_resArr := &resArr
	for i := 0; i < colNum; i++ {
		resArr = append(resArr, 0.0)
	}

	for _, row := range str2d {
		for j, str := range row {
			if getWechatStrWidth(str) > (*P_resArr)[j] {
				(*P_resArr)[j] = getWechatStrWidth(str)
			}
		}
	}
	//log.Println(resArr)
	return resArr
}

// str2d每一列在微信中的每一列的最大宽度求和，作为行物理宽度
func getWechatTablizeRowFullWidth(str2d [][]string) float32 {
	fs := getWechatTablizeColMaxWidth(str2d)
	f := new(float32)
	m := getWechatCharWidthMap()
	for _, v := range fs {
		*f = *f + v
	}
	//加上缩进宽度
	plus := float32(len(str2d[0])-1) * (m[" "]*4 + m["1"])
	*f = *f + plus
	//fmt.Println("总像素宽度：", *f)
	return *f
}

// 填充str2d每一格到列最大宽度
func wechatArr2dFillAndAlign(str2d [][]string) [][]string {
	m := getWechatCharWidthMap()
	mws := getWechatTablizeColMaxWidth(str2d)
	str2d2 := str2d
	for i := 0; i < len(str2d2); i++ {
		for j := 0; j < len(str2d2[i]); j++ {
			d := mws[j] - getWechatStrWidth(str2d2[i][j])
			for k := 0; float32(k)*m[" "] < d; k++ {
				str2d2[i][j] += " "
			}
		}
	}
	//打印
	//fmt.Println("=============打印改动后微信宽度===========")
	//for i := 0; i < len(str2d2); i++ {
	//	for j := 0; j < len(str2d2[i]); j++ {
	//		fmt.Printf("第%v行第%v列宽度： %v", i, j, getWechatStrWidth(str2d2[i][j]))
	//	}
	//	fmt.Println()
	//}
	return str2d2
}
