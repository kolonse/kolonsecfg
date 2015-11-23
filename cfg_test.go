package cfg

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	cfg := NewCfg()
	cfg.ParseFile("./test.kcfg")
	t.Log("--------------------------------")
	t.Log(cfg.Dump())
	t.Log("--------------------------------")
	v := cfg.root.Child("#").value.GetString()
	t.Log(v, " ", len(v))
	t.Log("just a test", " ", len("just a test"))
	assertEqual(t, v, "just a test", "not equal,0 real:"+v)
	v = cfg.root.Childs("#")[1].value.GetString()
	assertEqual(t, v, "just a test 2", "not equal,1 real:"+v)
	v = cfg.root.Child("dev").value.GetString()
	t.Log(v, " ", len(v))
	assertEqual(t, v, "true", "not equal,0 real:"+v)
}
