/**
配置解析文件 配置文件格式符号说明:
# 表示注释
 关键字下方需要有一个 tab 用来表示以下的属性都属于 关键字下的列表
key:{} 表示关键字key下有一组属性是key:value的列表
key:[] 表示关键字 key的属性是一个数组
"" 表示一串字符串
*/
package cfg

import (
	"io/ioutil"
)

// 定义关键字符号
const (
	COMMENT_B = '#'
	OBJECT_S  = '{'
	OBJECT_E  = '}'
	ARRAY_S   = '['
	ARRAY_E   = ']'
	KEY_E     = ':'
	LINE_END  = "\r\n"
)

type Cfg struct {
	root    *Node
	path    string
	content string
}

func (cfg *Cfg) readFile(path string) {
	cfg.path = path
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	cfg.content = string(buff)
}

func (cfg *Cfg) parseComment(offset int) (*Node, int) {
	n := NewNode("COMMENT")
	for i := offset + 1; i < len(cfg.content); i++ {
		//		if
	}
	return n, 0
}

func (cfg *Cfg) Parse(path string) *Cfg {
	cfg.readFile(path)
	// 遍历BUFF 对内容进行解析
	for i := 0; i < len(cfg.content); {
		var n *Node
		//		var offset int
		switch cfg.content[i] {
		case COMMENT_B:

		}
		cfg.root.AddChild(n)
	}
	return cfg
}

func NewCfg() *Cfg {
	return &Cfg{
		root: NewNode(""),
	}
}
