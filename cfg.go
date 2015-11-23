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

func (cfg *Cfg) parseValue(offset int, farther *Node) int {
	i := offset
	key := ""
	// 读取KEY 完成与否标志
	isGetKey := false
	isValueString := true
	valueIndex := -1
	end := false
	var v *Value
	for i < len(cfg.content) {
		c := cfg.content[i]
		switch {
		case ARRAY_S == c && !isGetKey:
			// 解析数组
			isValueString = false
			fallthrough
		case OBJECT_S == c && !isGetKey:
			// 解析对象
			isValueString = false
			fallthrough
		case ' ' == c && !isGetKey:
			key = cfg.content[offset:i]
			isGetKey = true
		case (' ' == c || '\t' == c) && isGetKey: // 如果是多个 空格那么让索引加一
		case ' ' != c && isGetKey && valueIndex == -1:
			// 如果获取到 key 以后 那么 value 的值起始位置在 ' ' 后
			valueIndex = i
		case '\r' == c || '\n' == c || i == len(cfg.content)-1:
			if !isGetKey {
				key = cfg.content[offset:i]
				isGetKey = true
			}
			if isValueString && valueIndex != -1 {
				if i == len(cfg.content)-1 {
					v = NewValue(STRING, cfg.content[valueIndex:i+1])
				} else {
					v = NewValue(STRING, cfg.content[valueIndex:i])
				}
			}
			// 添加完成节点后要进行返回
			end = true
		}
		if end {
			if isGetKey && v == nil {
				panic(errors.New("只有关键字没有值"))
			} else if isGetKey && v != nil {
				// 获取到关键字和获取到值
				n := NewNode(key)
				n.SetValue(v)
				farther.AddChild(n)
			} else {
				// 当前行没有获取到任何内容
			}
			break
		} else {
			i += 1
		}
	}

	return i + 1
}

func (cfg *Cfg) ParseFile(path string) *Cfg {
	cont := cfg.readFile(path)
	return cfg.ParseString(cont)
}

func (cfg *Cfg) ParseString(content string) *Cfg {
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
			offset = cfg.parseValue(i, cfg.root)
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
