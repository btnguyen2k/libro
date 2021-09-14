package respicite

import (
	"testing"
)

func doTestM2mDao_GetNotExist(t *testing.T, testName string, dao M2oDao) {
	if result, err := dao.Get("not-exist"); result != nil || err != nil {
		t.Fatalf("%s failed: expected (nil, nil) but received (%#v, %T)", testName, result, err)
	}
}

func doTestM2mDao_SetGet(t *testing.T, testName string, dao M2oDao) {
	src := "email@domain.com"
	dest := "user"

	if result, err := dao.Set(src, dest); err != nil || !result {
		t.Fatalf("%s failed: %#v / %s", testName+"/Set", result, err)
	}

	if result, err := dao.Get(src); result == nil || err != nil {
		t.Fatalf("%s failed: %T / %s", testName+"/Get", result, err)
	} else if result.Src != src || result.Dest != dest {
		t.Fatalf("%s failed: expected (%#v, %#v) but received (%#v, %#v)", testName+"/Get", src, dest, result.Src, result.Dest)
	}
}

func doTestM2mDao_SetDuplicated(t *testing.T, testName string, dao M2oDao) {
	src := "email@domain.com"
	dest := "user"

	if result, err := dao.Set(src, dest); err != nil || !result {
		t.Fatalf("%s failed: %#v / %s", testName+"/Set{1}", result, err)
	}

	if result, err := dao.Set(src, dest); err != ErrDuplicated || !result {
		t.Fatalf("%s failed: %#v / %s", testName+"/Set{2}", result, err)
	}

	if result, err := dao.Set(src, dest+"-new"); err != ErrDuplicated || !result {
		t.Fatalf("%s failed: %#v / %s", testName+"/Set{2}", result, err)
	}
}

func doTestM2mDao_SetRemove(t *testing.T, testName string, dao M2oDao) {
	src := "email@domain.com"
	dest := "user"

	if result, err := dao.Set(src, dest); err != nil || !result {
		t.Fatalf("%s failed: %#v / %s", testName+"/Set", result, err)
	}

	if result, err := dao.Get(src); result == nil || err != nil {
		t.Fatalf("%s failed: %T / %s", testName+"/Get", result, err)
	} else if result.Src != src || result.Dest != dest {
		t.Fatalf("%s failed: expected (%#v, %#v) but received (%#v, %#v)", testName+"/Get", src, dest, result.Src, result.Dest)
	}

	if result, err := dao.Remove(src, dest); err != nil || !result {
		t.Fatalf("%s failed: %T / %s", testName+"/Remove", result, err)
	}

	if result, err := dao.Get(src); result != nil || err != nil {
		t.Fatalf("%s failed: expected (nil, nil) but received (%#v, %T)", testName+"/Get", result, err)
	}
}

func doTestM2mDao_RemoveNotExist(t *testing.T, testName string, dao M2oDao) {
	if result, err := dao.Remove("Src-not-exist", "Dest-not-exist"); err != nil || result {
		t.Fatalf("%s failed: %T / %s", testName, result, err)
	}
}

func doTestM2mDao_SetRget(t *testing.T, testName string, dao M2oDao) {
	src1 := "email1@domain.com"
	src2 := "email2@domain.com"
	dest := "user"

	if result, err := dao.Set(src1, dest); err != nil || !result {
		t.Fatalf("%s failed: %#v / %s", testName+"/Set{1}", result, err)
	}
	if result, err := dao.Set(src2, dest); err != nil || !result {
		t.Fatalf("%s failed: %#v / %s", testName+"/Set{2}", result, err)
	}

	if result, err := dao.Rget(dest); result == nil || err != nil {
		t.Fatalf("%s failed: %T / %s", testName+"/Rget", result, err)
	} else if len(result) != 2 {
		t.Fatalf("%s failed: expected 2 items but received %#v", testName+"/Get", len(result))
	} else if result[0].Src != src1 || result[0].Dest != dest || result[1].Src != src2 || result[1].Dest != dest {
		t.Fatalf("%s failed: expected [{%s:%s}, {%s:%s}] but received [{%s:%s}, {%s:%s}]", testName+"/Get",
			src1, dest, src2, dest, result[0].Src, result[0].Dest, result[1].Src, result[1].Dest)
	}
}
