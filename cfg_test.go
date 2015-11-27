package kolonsecfg

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	cfg := NewCfg()
	cfg.ParseFile("./test.kcfg")
	t.Log("--------------------------------")
	t.Log(cfg.Dump())
	t.Log("--------------------------------")
	v := cfg.Child("#").GetString()
	t.Log(v, " ", len(v))
	t.Log("just a test", " ", len("just a test"))
	assertEqual(t, v, "just a test", "not equal,0 real:"+v)
	v = cfg.Childs("#")[1].GetString()
	assertEqual(t, v, "just a test 2", "not equal,1 real:"+v)
	v = cfg.Child("dev").GetString()
	t.Log(v, " ", len(v))
	assertEqual(t, v, "true", "not equal,0 real:"+v)
	assertEqual(t, cfg.Child("jszhou2.tt").GetString(), "haha", "not equal")
	assertEqual(t, cfg.Child("jszhou2.woqu.dddd").GetString(), "ddd", "not equal")
	assertEqual(t, cfg.Childs("jszhou2.tt")[1].GetString(), "tdfsdfd", "not equal")
	assertEqual(t, cfg.Childs("jszhou2.tt")[0].GetString(), "haha", "not equal")
}
