package main

import "fmt"
import s "strings"
import "strconv"
/*
1.字符串函数
2.字符串格式化
3.数字转换
4.时间函数
5.JSON转换
*/
var p = fmt.Printf


func main() {
	p("Contains:  %v\n", s.Contains("test", "es"))
	p("Count:     %v\n", s.Count("test", "t"))
	p("HasPrefix: %v\n", s.HasPrefix("test", "te"))
	p("HasSuffix: %v\n", s.HasSuffix("test", "st"))
	p("Index:     %v\n", s.Index("test", "e"))
	p("Join:      %v\n", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    %v\n", s.Repeat("a", 5))
	p("Replace:   %v\n", s.Replace("foo", "o", "0", -1))
	p("Replace:   %v\n", s.Replace("foo", "o", "0", 1))
	p("Split:     %v\n", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   %v\n", s.ToLower("TEST"))
	p("ToUpper:   %v\n", s.ToUpper("test"))
	// 虽然不是 strings 的一部分，但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制。
	p("Len: %v\n", len("hello"))
	p("Char: %v\n", "hello"[1])

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// 你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// 要左对齐，使用 - 标志。
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// 你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。
	// 这是基本的右对齐宽度表示。
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// 要左对齐，和数字一样，使用 - 标志。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	// 到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。
	// Sprintf 则格式化并返回一个字符串而不带任何输出。
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)


	// 内置的 strconv 包提供了数字解析功能。
	// 使用 ParseFloat 解析浮点数，这里的 64 表示表示解析的数的位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// 在使用 ParseInt 解析整形数时，例子中的参数 0 表示自动推断字符串所表示的数字的进制。
	// 64 表示返回的整形数是以 64 位存储的。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// ParseInt 会自动识别出十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	// ParseUint 也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// Atoi 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}