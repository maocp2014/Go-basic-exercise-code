package main

import (
	"fmt"
	"strconv"
)

/*
题目：

让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。
例如，IPAddr{1, 2, 3, 4} 应当输出 “1.2.3.4”。
*/

// 别名方式自定义类型
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// 版本1
// func (ip IPAddr) String() string {
// 	return fmt.Sprintf("%d.%d.%d.%d", ip[0],ip[1],ip[2],ip[3])
// }

// 版本2
// 改变IPAddr的默认打印格式 
func (ipa IPAddr) String() string {
    lastIndex := len(ipa) - 1
    var s string
    for i, v := range ipa {
		s += strconv.Itoa(int(v)) // int to string
		fmt.Println(s)
        if i != lastIndex {
            s += "."
        }
    }
    return s
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
