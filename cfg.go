/**
配置解析文件 配置文件格式符号说明:
# 表示注释
 关键字下方需要有一个 tab 用来表示以下的属性都属于 关键字下的列表
key:{} 表示关键字key下有一组属性是key:value的列表
key:[] 表示关键字 key的属性是一个数组
"" 表示一串字符串
*/
package kolonsecfg

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

func (cfg *Cfg) parseComment(content string, offset int, farther *Node) int {
	n := NewNode("#")
	i := offset + 1
	for ; i < len(content); i++ {
		// 只要碰到 \r 或者 \n 那么直接跳出
		if content[i] == '\r' || content[i] == '\n' {
			break
		}
	}
	v := NewValue(STRING, content[offset+1:i])
	n.SetValue(v)
	farther.AddChild(n)
	return i + 1
}

func (cfg *Cfg) parseObject(content string, key string, offset int, farther *Node) int {
	stack := 1
	i := offset
	for ; stack != 0 && i < len(content); i++ {
		if content[i] == '{' {
			stack++
		} else if content[i] == '}' {
			stack--
		}
	}
	if stack != 0 {
		panic(errors.New("{ 不匹配 } 数目"))
	} else {
		childString := content[offset:i]
		cfg.parseByString(childString, farther)
	}
	return i + 1
}

func (cfg *Cfg) parseArray(content string, key string, offset int, farther *Node) int {
	stack := 1
	i := offset
	for ; stack != 0 && i < len(content); i++ {
		if content[i] == '[' {
			stack++
		} else if content[i] == ']' {
			stack--
		}
	}
	if stack != 0 {
		panic(errors.New("[ 不匹配 ] 数目"))
	} else {
		childString := content[offset:i]
		cfg.parseByString(childString, farther)
	}
	return i + 1
}

func (cfg *Cfg) parseString(content string, offset int) (string, int) {
	i := offset
	for ; i < len(content); i++ {
		c := content[i]
		if '\r' == c || '\n' == c {
			return content[offset:i], i + 1
		} else if i == len(content)-1 {
			return content[offset : i+1], i + 1
		}
	}
	panic(errors.New("string 值不存在"))
	return "", -1
}

func (cfg *Cfg) parseKey(content string, offset int) (string, int) {
	i := offset
	key := ""
	keyIndex := i
	for ; i < len(content); i++ {
		c := content[i]
		switch {
		case ARRAY_S == c:
			fallthrough
		case OBJECT_S == c:
			fallthrough
		case ' ' == c || '\t' == c:
			key = content[keyIndex:i]
			return key, i
		case '\r' == c || '\n' == c:
			panic(errors.New("关键字没有值"))
		}
	}
	panic(errors.New("配置格式应该为[key value/{}/[]]"))
	return "", i
}

func (cfg *Cfg) parseValue(content string, key string, offset int) (*Node, int) {
	// 开始进行值处理
	i := offset
	for ; i < len(content); i++ {
		c := content[i]
		switch {
		case ARRAY_S == c: // 数组值
			n := NewNode(key)
			index := cfg.parseArray(content, key, i+1, n)
			return n, index
		case OBJECT_S == c: // 对象值
			n := NewNode(key)
			index := cfg.parseObject(content, key, i+1, n)
			return n, index
		case ' ' != c && '\t' != c: // 字符串值
			n := NewNode(key)
			value, index := cfg.parseString(content, i)
			v := NewValue(STRING, value)
			n.SetValue(v)
			return n, index
		}
	}
	panic(errors.New("配置错误 只有关键字没有值"))
	return nil, -1
}

func (cfg *Cfg) parseAttr(content string, offset int, farther *Node) int {
	key, i := cfg.parseKey(content, offset)
	n, i := cfg.parseValue(content, key, i)
	farther.AddChild(n)
	return i
}

func (cfg *Cfg) parseByString(content string, farther *Node) {
	// 遍历BUFF 对内容进行解析
	//row := 0
	for i := 0; i < len(content); {
		var offset int
		c := content[i]
		switch {
		case COMMENT_B == c:
			offset = cfg.parseComment(content, i, farther)
		case 'a' <= c && 'z' >= c:
			fallthrough
		case 'A' <= c && 'Z' >= c:
			offset = cfg.parseAttr(content, i, farther)
		default:
			offset = i + 1
		}
		i = offset
	}
}

func (cfg *Cfg) ParseFile(path string) *Cfg {
	content := cfg.readFile(path)
	return cfg.ParseByString(content)
}

func (cfg *Cfg) ParseByString(content string) *Cfg {
	cfg.content = content
	cfg.parseByString(content, cfg.root)
	return cfg
}

func (cfg *Cfg) Dump() string {
	//childs := cfg.root.Childs()
	return cfg.root.Dump("")
}

func (cfg *Cfg) Value(key string) *Node {
	kr := NewKeyRoute()
	kr.parse(key)
	find := func(root *Node, k string) *Node {
		return root.Child(k)
	}
	n := cfg.root.Child(kr.next())
	for !kr.end() {
		n = find(n, kr.next())
	}
	return n
}

func (cfg *Cfg) Values(key string) []*Node {
	kr := NewKeyRoute()
	kr.parse(key)
	find := func(root []*Node, k string) []*Node {
		var ret []*Node
		for _, r := range root {
			ns := r.Childs(k)
			if ns != nil {
				ret = append(ret, ns...)
			}
		}
		return ret
	}
	n := cfg.root.Childs(kr.next())
	for !kr.end() {
		n = find(n, kr.next())
	}
	return n
}

func NewCfg() *Cfg {
	return &Cfg{
		root: NewNode(""),
	}
}
