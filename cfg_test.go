package cfg

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	cfg := NewCfg()
	cfg.ParseFile("./test.kcfg")
	v := cfg.root.Child("COMMENT").value.GetString()
	t.Log(v, " ", len(v))
	t.Log("just a test", " ", len("just a test"))
	assertEqual(t, v, "just a test", "not equal,0 real:"+v)
	v = cfg.root.Childs("COMMENT")[1].value.GetString()
	assertEqual(t, v, "just a test 2", "not equal,1 real:"+v)
}
