package doc

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"main/src/gvabe/bo/app"
)

const numSampleRowsTopic = 100

func initSampleRowsTopic(t *testing.T, testName string, dao TopicDao) {
	rand.Seed(time.Now().UnixNano())
	numApps := 1 + rand.Intn(numSampleRowsTopic/10)
	appList = make([]*app.App, numApps)
	_tagVersion := uint64(1337)
	for i := 0; i < numApps; i++ {
		_appId := "libro" + fmt.Sprintf("%02d", i)
		_name := "Libro" + _appId
		_desc := "Libro description " + _appId
		_isVisible := rand.Int()%7 == 0
		_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)
		appList[i] = _app
	}
	topicList = make([]*Topic, numSampleRowsTopic)
	appTopicCount = make(map[string]int)
	for i := 0; i < numSampleRowsTopic; i++ {
		_app := appList[rand.Intn(numApps)]
		istr := fmt.Sprintf("%03d", i)
		_title := "Quick start " + istr
		_icon := "default"
		_summary := "topic " + istr
		_pos := rand.Int()
		_email := istr + "@libro"
		_age := float64(18 + i)
		bo := NewTopic(_tagVersion, _app, _title, _icon, _summary)
		bo.SetPosition(_pos)
		bo.SetDataAttr("props.owner", "App"+istr)
		bo.SetDataAttr("props.email", _email)
		bo.SetDataAttr("age", _age)
		if ok, err := dao.Create(bo); err != nil || !ok {
			t.Fatalf("%s failed: %#v / %s", testName+"/Create", ok, err)
		}
		topicList[i] = bo

		counter := appTopicCount[_app.GetId()]
		appTopicCount[_app.GetId()] = counter + 1
	}
}

func doTestTopicDaoCreateGet(t *testing.T, name string, dao TopicDao) {
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	bo0 := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	bo0.SetPosition(_pos)
	_email := "libro@libro"
	_age := float64(35)
	bo0.SetDataAttr("props.owner", _name)
	bo0.SetDataAttr("props.email", _email)
	bo0.SetDataAttr("age", _age)
	if ok, err := dao.Create(bo0); err != nil || !ok {
		t.Fatalf("%s failed: %#v / %s", name+"/Create", ok, err)
	}

	_id := bo0.GetId()
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
		if v1, v0 := bo1.GetAppId(), _app.GetId(); v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetIcon(), _icon; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetTitle(), _title; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetSummary(), _summary; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetPosition(), _pos; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if t1, t0 := bo1.GetTimeCreated(), bo0.GetTimeCreated(); !t1.Equal(t0) {
			t.Fatalf("%s failed: expected %#v but received %#v", name, t0.Format(time.RFC3339), t1.Format(time.RFC3339))
		}
		if bo1.GetChecksum() != bo0.GetChecksum() {
			t.Fatalf("%s failed: expected %#v but received %#v", name, bo0.GetChecksum(), bo1.GetChecksum())
		}
	}
}

func doTestTopicDaoCreateUpdateGet(t *testing.T, name string, dao TopicDao) {
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	bo0 := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	bo0.SetPosition(_pos)
	_email := "libro@libro"
	_age := float64(35)
	bo0.SetDataAttr("props.owner", _name)
	bo0.SetDataAttr("props.email", _email)
	bo0.SetDataAttr("age", _age)
	if ok, err := dao.Create(bo0); err != nil || !ok {
		t.Fatalf("%s failed: %#v / %s", name+"/Create", ok, err)
	}

	bo0.SetAppId(_appId + "-new").SetIcon(_icon + "-new").SetTitle(_title + "-new").SetSummary(_summary + "-new").SetPosition(_pos + 1).SetTagVersion(_tagVersion + 3)
	bo0.SetDataAttr("props.owner", _name+"-new")
	bo0.SetDataAttr("props.email", _email+"-new")
	bo0.SetDataAttr("age", _age+2)
	if ok, err := dao.Update(bo0); err != nil {
		t.Fatalf("%s failed: %s", name+"/Update", err)
	} else if !ok {
		t.Fatalf("%s failed: cannot update record", name)
	}

	_id := bo0.GetId()
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
		if v1, v0 := bo1.GetAppId(), _app.GetId()+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetIcon(), _icon+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetTitle(), _title+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetSummary(), _summary+"-new"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetPosition(), _pos+1; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if t1, t0 := bo1.GetTimeCreated(), bo0.GetTimeCreated(); !t1.Equal(t0) {
			t.Fatalf("%s failed: expected %#v but received %#v", name, t0.Format(time.RFC3339), t1.Format(time.RFC3339))
		}
		if bo1.GetChecksum() != bo0.GetChecksum() {
			t.Fatalf("%s failed: expected %#v but received %#v", name, bo0.GetChecksum(), bo1.GetChecksum())
		}
	}
}

func doTestTopicDaoCreateDelete(t *testing.T, name string, dao TopicDao) {
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isVisible := true
	_app := app.NewApp(_tagVersion, _appId, _name, _desc, _isVisible)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	bo0 := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	bo0.SetPosition(_pos)
	_email := "libro@libro"
	_age := float64(35)
	bo0.SetDataAttr("props.owner", _name)
	bo0.SetDataAttr("props.email", _email)
	bo0.SetDataAttr("age", _age)
	if ok, err := dao.Create(bo0); err != nil || !ok {
		t.Fatalf("%s failed: %#v / %s", name+"/Create", ok, err)
	}

	_id := bo0.GetId()
	if bo1, err := dao.Get(_id); err != nil || bo1 == nil {
		t.Fatalf("%s failed: nil or error %s", name+"/Get("+_id+")", err)
	} else if ok, err := dao.Delete(bo1); !ok || err != nil {
		t.Fatalf("%s failed: not-ok or error %s", name+"/Delete("+_id+")", err)
	}
	if bo1, err := dao.Get(_id); err != nil || bo1 != nil {
		t.Fatalf("%s failed: not-nil or error %s", name+"/Get("+_id+")", err)
	}
}

func doTestTopicDaoGetAll(t *testing.T, name string, dao TopicDao) {
	initSampleRowsTopic(t, name, dao)
	for _, app := range appList {
		boList, err := dao.GetAll(app, nil, nil)
		expected := appTopicCount[app.GetId()]
		if err != nil || len(boList) != expected {
			t.Fatalf("%s failed: expected %#v but received %#v (error %s)", name+"/GetAll", expected, len(boList), err)
		}
	}
}

func doTestTopicDaoGetN(t *testing.T, name string, dao TopicDao) {
	initSampleRowsTopic(t, name, dao)
	for _, app := range appList {
		startOffset := rand.Intn(5)
		numRowsLimit := rand.Intn(10) + 1
		boList, err := dao.GetN(app, startOffset, numRowsLimit, nil, nil)
		expected := appTopicCount[app.GetId()]
		expected = int(math.Min(math.Max(0, float64(expected-startOffset)), float64(numRowsLimit)))
		if err != nil || len(boList) != expected {
			// fmt.Printf("%s - %#v - %#v / %#v\n", app.GetId(), startOffset, numRowsLimit, appTopicCount)
			t.Fatalf("%s failed: expected %#v but received %#v (error %s)", name+"/"+app.GetId(), expected, len(boList), err)
		}
	}
}
