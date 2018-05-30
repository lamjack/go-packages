package main

import (
    "fmt"
    "regexp"
    "bytes"
)

func main() {
    // 将字符串中的特殊字符转为其转义格式
    // fmt.Println(regexp.QuoteMeta("(?P:Hello) [a-z]"))

    // match()
    // createRegexp()
    find()
}

// 判断是否存在匹配的子串
func match() {
    pattern := "H.* "

    // []bytes
    fmt.Println(regexp.Match(pattern, []byte("Hello World!")))

    // io.RuneReader
    r := bytes.NewReader([]byte("Hello World"))
    fmt.Println(regexp.MatchReader(pattern, r))

    // string
    fmt.Println(regexp.MatchString(pattern, "Hello World"))
}

// 创建Regexp对象，Regexp对象可以在任意文本上执行需要的操作
func createRegexp() {
    var (
        err error
        reg *regexp.Regexp
    )

    // Compile 用来解析正则表达式是否合法，如果合法，则返回一个 Regexp 对象
    reg, err = regexp.Compile(`\w+`)
    fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)

    // CompilePOSIX 的作用和 Compile 一样
    // 不同的是，CompilePOSIX 使用 POSIX 语法，
    // 同时，它采用最左最长方式搜索，
    // 而 Compile 采用最左最短方式搜索
    // POSIX 语法不支持 Perl 的语法格式：\d、\D、\s、\S、\w、\W
    reg, err = regexp.CompilePOSIX(`[[:word:]]+`)
    fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err)

    // MustCompile 的作用和 Compile 一样
    // 不同的是，当正则表达式 str 不合法时，MustCompile 会抛出异常
    // 而 Compile 仅返回一个 error 值
    reg = regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindString("Hello World"))

    // MustCompilePOSIX 的作用和 CompilePOSIX 一样
    // 不同的是，当正则表达式 str 不合法时，MustCompilePOSIX 会抛出异常
    // 而 CompilePOSIX 仅返回一个 error 值
    reg = regexp.MustCompilePOSIX(`[[:word:]].+ `)
    fmt.Printf("%q\n", reg.FindString("Hello World!"))
}

func find() {
    r := bytes.NewReader([]byte("Hello World!"))
    reg := regexp.MustCompile(`\w+`)

    // 在 []byte 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
    fmt.Printf("%q\n", reg.Find([]byte("Hello World!")))

    // 在 string 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
    fmt.Println(reg.FindString("Hello World!"))

    // 在 []byte 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
    // 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
    // ["Hello" "World"]
    fmt.Printf("%q\n", reg.FindAll([]byte("Hello World!"), -1))

    // 在 string 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
    // 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
    // ["Hello" "World"]
    fmt.Printf("%q\n", reg.FindAllString("Hello World!", -1))

    // 在 []byte 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
    // [0 5]
    fmt.Println(reg.FindIndex([]byte("Hello World!")))

    // 在 string 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
    // [0 5]
    fmt.Println(reg.FindStringIndex("Hello World!"))

    // 在 io.RuneReader 中查找 re 中编译好的正则表达式，并返回第一个匹配的位置
    // [0 5]
    fmt.Println(reg.FindReaderIndex(r))

    // 在 string 中查找 re 中编译好的正则表达式，并返回所有匹配的位置
    // 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
    // [[0 5] [6 11]]
    fmt.Println(reg.FindAllStringIndex("Hello World!", -1))

    // 返回第一个匹配到的结果及其分组内容（结果以 b 的切片形式返回）。
    // 返回值中的第 0 个元素是整个正则表达式的匹配结果，后续元素是各个分组的
    // 匹配内容，分组顺序按照“(”的出现次序而定。
    // ["Hello"]
    fmt.Printf("%q\n", reg.FindSubmatch([]byte("Hello World!")))

    fmt.Printf("%q\n", reg.FindStringSubmatch("Hello World!"))
}
