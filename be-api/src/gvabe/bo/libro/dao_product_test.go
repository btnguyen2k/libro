package libro

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/btnguyen2k/consu/reddo"
)

const numSampleRows = 100

func initSampleRows(t *testing.T, testName string, dao ProductDao) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numSampleRows; i++ {
		istr := fmt.Sprintf("%03d", i)
		_tagVersion := uint64(1337)
		_id := istr + "@libro"
		_name := "Libro" + istr
		_desc := "Libro description"
		_isPublished := i%7 == 0
		_numTopics := i%5 + 1
		_email := istr + "@libro"
		_age := float64(18 + i)
		bo := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
		bo.SetNumTopics(_numTopics)
		bo.SetDataAttr("props.owner", "User"+istr)
		bo.SetDataAttr("props.email", _email)
		bo.SetDataAttr("age", _age)
		if ok, err := dao.Create(bo); err != nil || !ok {
			t.Fatalf("%s failed: %#v / %s", testName+"/Create", ok, err)
		}
	}
}

func doTestProductDaoCreateGet(t *testing.T, name string, dao ProductDao) {
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := 3
	_email := "libro@libro"
	_age := float64(35)
	bo0 := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	bo0.SetNumTopics(_numTopics)
	bo0.SetDataAttr("props.owner", _name)
	bo0.SetDataAttr("props.email", _email)
	bo0.SetDataAttr("age", _age)
	if ok, err := dao.Create(bo0); err != nil || !ok {
		t.Fatalf("%s failed: %#v / %s", name+"/Create", ok, err)
	}

	if bo1, err := dao.Get(_id); err != nil || bo1 == nil {
		t.Fatalf("%s failed: nil or error %s", name+"/Get("+_id+")", err)
	} else {
		if v1, v0 := bo1.GetDataAttrAsUnsafe("props.owner", reddo.TypeString), _name; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetDataAttrAsUnsafe("props.email", reddo.TypeString), _email; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetDataAttrAsUnsafe("age", reddo.TypeInt), int64(_age); v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetTagVersion(), _tagVersion; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetId(), _id; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetName(), _name; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetDescription(), _desc; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.IsPublished(), _isPublished; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetNumTopics(), _numTopics; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if t1, t0 := bo1.GetTimeCreated(), bo0.GetTimeCreated(); !t1.Equal(t0) {
			t.Fatalf("%s failed: expected %#v but received %#v", name, t0.Format(time.RFC3339Nano), t1.Format(time.RFC3339Nano))
		}
		if bo1.GetChecksum() != bo0.GetChecksum() {
			t.Fatalf("%s failed: expected %#v but received %#v", name, bo0.GetChecksum(), bo1.GetChecksum())
		}
	}
}

func doTestProductDaoCreateUpdateGet(t *testing.T, name string, dao ProductDao) {
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := 3
	_email := "libro@libro"
	_age := float64(35)
	bo0 := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	bo0.SetNumTopics(_numTopics)
	bo0.SetDataAttr("props.owner", _name)
	bo0.SetDataAttr("props.email", _email)
	bo0.SetDataAttr("age", _age)
	if ok, err := dao.Create(bo0); err != nil || !ok {
		t.Fatalf("%s failed: %#v / %s", name+"/Create", ok, err)
	}

	bo0.SetName(_name + "-new").SetDescription(_desc + "-new").SetPublished(!_isPublished).SetNumTopics(_numTopics + 2).SetTagVersion(_tagVersion + 3)
	bo0.SetDataAttr("props.owner", _name+"-new")
	bo0.SetDataAttr("props.email", _email+"-new")
	bo0.SetDataAttr("age", _age+2)
	if ok, err := dao.Update(bo0); err != nil {
		t.Fatalf("%s failed: %s", name+"/Update", err)
	} else if !ok {
		t.Fatalf("%s failed: cannot update record", name)
	}
	if bo1, err := dao.Get(_id); err != nil || bo1 == nil {
		t.Fatalf("%s failed: nil or error %s", name+"/Get("+_id+")", err)
	} else {
		if v1, v0 := bo1.GetDataAttrAsUnsafe("props.owner", reddo.TypeString), _name+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetDataAttrAsUnsafe("props.email", reddo.TypeString), _email+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetDataAttrAsUnsafe("age", reddo.TypeInt), int64(_age+2); v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetTagVersion(), _tagVersion+3; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetId(), _id; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetName(), _name+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetDescription(), _desc+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.IsPublished(), !_isPublished; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetNumTopics(), _numTopics+2; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if t1, t0 := bo1.GetTimeCreated(), bo0.GetTimeCreated(); !t1.Equal(t0) {
			t.Fatalf("%s failed: expected %#v but received %#v", name, t0.Format(time.RFC3339Nano), t1.Format(time.RFC3339Nano))
		}
		if bo1.GetChecksum() != bo0.GetChecksum() {
			t.Fatalf("%s failed: expected %#v but received %#v", name, bo0.GetChecksum(), bo1.GetChecksum())
		}
	}
}

func doTestProductDaoCreateDelete(t *testing.T, name string, dao ProductDao) {
	_tagVersion := uint64(1337)
	_id := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_numTopics := 3
	_email := "libro@libro"
	_age := float64(35)
	bo0 := NewProduct(_tagVersion, _id, _name, _desc, _isPublished)
	bo0.SetNumTopics(_numTopics)
	bo0.SetDataAttr("props.owner", _name)
	bo0.SetDataAttr("props.email", _email)
	bo0.SetDataAttr("age", _age)
	if ok, err := dao.Create(bo0); err != nil || !ok {
		t.Fatalf("%s failed: %#v / %s", name+"/Create", ok, err)
	}

	if user1, err := dao.Get(_id); err != nil || user1 == nil {
		t.Fatalf("%s failed: nil or error %s", name+"/Get("+_id+")", err)
	} else if ok, err := dao.Delete(user1); !ok || err != nil {
		t.Fatalf("%s failed: not-ok or error %s", name+"/Delete("+_id+")", err)
	}

	if user1, err := dao.Get(_id); err != nil || user1 != nil {
		t.Fatalf("%s failed: not-nil or error %s", name+"/Get("+_id+")", err)
	}
}

func doTestProductDaoGetAll(t *testing.T, name string, dao ProductDao) {
	initSampleRows(t, name, dao)
	userList, err := dao.GetAll(nil, nil)
	if err != nil || len(userList) != numSampleRows {
		t.Fatalf("%s failed: expected %#v but received %#v (error %s)", name+"/GetAll", numSampleRows, len(userList), err)
	}
}

func doTestProductDaoGetN(t *testing.T, name string, dao ProductDao) {
	initSampleRows(t, name, dao)
	userList, err := dao.GetN(3, 5, nil, nil)
	if err != nil || len(userList) != 5 {
		t.Fatalf("%s failed: expected %#v but received %#v (error %s)", name+"/GetN", 5, len(userList), err)
	}
}
