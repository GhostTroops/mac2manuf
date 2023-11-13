package netparse

import (
	_ "embed"
	"strings"
)

var d = map[string]string{}

//go:embed manuf
var manuf string

// 合并、去重下面所有厂商信息
// https://standards-oui.ieee.org/oui/oui.txt
// https://mac2vendor.com/download/ieee-oui-database.txt
// https://mac2vendor.com/download/vendorMacs.prop
// https://www.wireshark.org/download/automated/data/manuf
// 通过mac地址计算厂家
func init() {
	r := strings.Split(manuf, "\n")
	for _, x := range r {
		l := strings.Split(x, "\t")
		if len(l) >= 2 {
			var s1 string = l[1]
			if len(l) > 2 {
				s1 = strings.Join(l[2:], ",")
				s1 = strings.Replace(s1, " ", "", -1)
				s1 = strings.Replace(s1, "\t", "", -1)
				if -1 == strings.Index(s1, l[1]) {
					s1 = l[1] + "," + s1
				}
			}

			d[strings.ToUpper(l[0])] = s1
		}
	}
}

// 根据mac前8位搜索厂商信息
// format: xx:xx:xx
func Search(mac string) string {
	if 8 > len(mac) {
		return ""
	}
	if s, ok := d[strings.ToUpper(mac[0:8])]; ok {
		return s
	}
	return ""
}
