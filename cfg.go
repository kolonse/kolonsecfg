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
	"errors"
	"io/ioutil"
)

// 定义关键字符号
const (
	COMMENT_B = '#'
	OBJECT_S  = '{'
	OBJECT_E  = '}'
	ARRAY_S   = '['
	ARRAY_E   = ']'
	LINE_END  = "\r\n"
)

type Cfg struct {
	root    *Node
	path    string
	content string
}

func (cfg *Cfg) readFile(path string) string {
	cfg.path = path
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(buff)
}

func (cfg *Cfg) parseComment(offset int, farther *Node) int {
	n := NewNode("#")
	i := offset + 1
	for ; i < len(cfg.content); i++ {
		// 只要碰到 \r 或者 \n 那么直接跳出
		if cfg.content[i] == '\r' || cfg.content[i] == '\n' {
			break
		}
	}
	v := NewValue(STRING, cfg.content[offset+1:i])
	n.SetValue(v)
	farther.AddChild(n)
	return i + 1
}

func (cfg *Cfg) parseObject(key string, offset int, farther *Node) int {
	return 0
}

func (cfg *Cfg) parseArray(key string, offset int, farther *Node) int {
	return 0
}

func (cfg *Cfg) parseString(offset int) (string, int) {
	i := offset
	for ; i < len(cfg.content); i++ {
		c := cfg.content[i]
		if '\r' == c || '\n' == c {
			return cfg.content[offset:i], i + 1
		} else if i == len(cfg.content)-1 {
			return cfg.content[offset : i+1], i + 1
		}
	}
	panic(errors.New("string 值不存在"))
	return "", -1
}

func (cfg *Cfg) parseKey(offset int) (string, int) {
	i := offset
	key := ""
	keyIndex := i
	for ; i < len(cfg.content); i++ {
		c := cfg.content[i]
		switch {
		case ARRAY_S == c:
			fallthrough
		case OBJECT_S == c:
			fallthrough
		case ' ' == c || '\t' == c:
			key = cfg.content[keyIndex:i]
			println(key)
			return key, i
		case '\r' == c || '\n' == c:
			panic(errors.New("关键字没有值"))
		}
	}
	panic(errors.New("配置格式应该为[key value/{}/[]]"))
	return "", i
}

func (cfg *Cfg) parseValue(key string, offset int) (*Node, int) {
	// 开始进行值处理
	i := offset
	for ; i < len(cfg.content); i++ {
		c := cfg.content[i]
		switch {
		case ARRAY_S == c: // 数组值

		case OBJECT_S == c: // 对象值

		case ' ' != c && '\t' != c: // 字符串值
			n := NewNode(key)
			value, index := cfg.parseString(i)
			v := NewValue(STRING, value)
			n.SetValue(v)
			return n, index
		}
	}
	panic(errors.New("配置错误 只有关键字没有值"))
	return nil, -1
}

func (cfg *Cfg) parseAttr(offset int, farther *Node) int {
	key, i := cfg.parseKey(offset)
	n, i := cfg.parseValue(key, i)
	farther.AddChild(n)
	return i
}

func (cfg *Cfg) ParseFile(path string) *Cfg {
	cont := cfg.readFile(path)
	return cfg.ParseByString(cont)
}

func (cfg *Cfg) ParseByString(content string) *Cfg {
	cfg.content = content
	// 遍历BUFF 对内容进行解析
	//row := 0
	for i := 0; i < len(cfg.content); {
		var offset int
		c := cfg.content[i]
		switch {
		case COMMENT_B == c:
			offset = cfg.parseComment(i, cfg.root)
		case 'a' <= c && 'z' >= c:
			fallthrough
		case 'A' <= c && 'Z' >= c:
			offset = cfg.parseAttr(i, cfg.root)
		default:
			offset = i + 1
		}
		i = offset
	}
	return cfg
}

func (cfg *Cfg) Dump() string {
	//childs := cfg.root.Childs()
	return cfg.root.Dump("")
}

func NewCfg() *Cfg {
	return &Cfg{
		root: NewNode(""),
	}
}
