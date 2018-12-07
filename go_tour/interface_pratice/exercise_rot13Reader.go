/*
题目：
一个常见模式是 io.Reader 包裹另一个 io.Reader，然后通过某种形式修改数据流。
例如，gzip.NewReader 函数接受 io.Reader（压缩的数据流）并且返回同样实现了 io.Reader 的 *gzip.Reader（解压缩后的数据流）。

编写一个实现了 io.Reader 的 rot13Reader， 并从一个 io.Reader 读取， 利用 rot13 代换密码对数据流进行修改。
已经帮你构造了 rot13Reader 类型。 通过实现 Read 方法使其匹配 io.Reader。

提示：
ROT13是过去在古罗马开发的**凯撒加密**的一种变体。套用ROT13到一段文字上仅仅只需要检查字元字母顺序并取代它在13位之后的对应字母，
有需要超过时则重新绕回26英文字母开头即可[2]。A换成N、B换成O、依此类推到M换成Z，然后序列反转：N换成A、O换成B、最后Z换成M。
只有这些出现在英文字母里头的字元受影响；数字、符号、空白字元以及所有其他字元都不变。 因为只有在英文字母表里头只有26个，并且26=2×13，
ROT13函数是它自己的逆反： 
*/ 

package main

import (
	"io"
	"os"
	"strings"
)

// 封装io.Reader
type rot13Reader struct {
	r io.Reader
}

// 转换byte  前进13位/后退13位
func rot13(b byte) byte {
    switch {
    case 'A' <= b && b <= 'M':
        b = b + 13
    case 'M' < b && b <= 'Z':
        b = b - 13
    case 'a' <= b && b <= 'm':
        b = b + 13
    case 'm' < b && b <= 'z':
        b = b - 13
    }
    return b
}

// 实现接口Read方法
func (mr rot13Reader) Read(b []byte) (int, error) {
    n, e := mr.r.Read(b)
    for i := 0; i < n; i++ {
        b[i] = rot13(b[i])
    }
    return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}