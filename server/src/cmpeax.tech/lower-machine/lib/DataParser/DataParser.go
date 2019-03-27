package DataParser

import (
	"github.com/axgle/mahonia"
)

//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func convertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

// func convertToGBK(utfStr string) string {
// 	var enc mahonia.Encoder
// 	enc = mahonia.NewEncoder("gbk")
// 	return enc.ConvertString(utfStr)
// }

// gbk 转 utf8
func Parser(responseData string) string {
	//包一层解析结构的
	response := convertToByte(responseData, "gb18030", "utf8")
	return string(response[:len(response)])
}

//utf8 转 gbk
func ParserToGbk(responseData string) string {
	//包一层解析结构的
	var enc mahonia.Encoder
	enc = mahonia.NewEncoder("gbk")
	return enc.ConvertString(responseData)
}
