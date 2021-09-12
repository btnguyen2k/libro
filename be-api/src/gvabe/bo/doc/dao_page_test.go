package doc

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/btnguyen2k/consu/reddo"
	"main/src/gvabe/bo/product"
)

const numSampleRowsPage = 100

func initSampleRowsPage(t *testing.T, testName string, dao PageDao) {
	rand.Seed(time.Now().UnixNano())
	numTopics := 1 + rand.Intn(numSampleRowsPage/10)
	topicList = make([]*Topic, numTopics)
	_tagVersion := uint64(1337)
	_app := product.NewProduct(_tagVersion, "libro", "Libro", "Libro description", true)
	for i := 0; i < numTopics; i++ {
		istr := "topic" + fmt.Sprintf("%03d", i)
		_title := "Quick start " + istr
		_icon := "default"
		_summary := "topic " + istr
		_pos := rand.Int()
		_topic := NewTopic(_tagVersion, _app, _title, _icon, _summary)
		_topic.SetPosition(_pos)
		topicList[i] = _topic
	}
	pageList = make([]*Page, numSampleRowsPage)
	topicPageCount = make(map[string]int)
	for i := 0; i < numSampleRowsPage; i++ {
		_topic := topicList[rand.Intn(numTopics)]
		istr := fmt.Sprintf("%03d", i)
		_title := "Page " + istr
		_icon := "default"
		_summary := "page summary " + istr
		_pos := rand.Int()
		_content := "page content " + istr
		_email := istr + "@libro"
		_age := float64(18 + i)
		bo := NewPage(_tagVersion, _topic, _title, _icon, _summary, _content)
		bo.SetPosition(_pos)
		bo.SetDataAttr("props.owner", "Product"+istr)
		bo.SetDataAttr("props.email", _email)
		bo.SetDataAttr("age", _age)
		if ok, err := dao.Create(bo); err != nil || !ok {
			t.Fatalf("%s failed: %#v / %s", testName+"/Create", ok, err)
		}
		pageList[i] = bo

		counter := topicPageCount[_topic.GetId()]
		topicPageCount[_topic.GetId()] = counter + 1
	}
}

func doTestPageDaoCreateGet(t *testing.T, name string, dao PageDao) {
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_app := product.NewProduct(_tagVersion, _appId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_topic := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	_topic.SetPosition(_pos)

	_content := "page one"
	bo0 := NewPage(_tagVersion, _topic, _title+"-page", _icon+"-page", _summary+"-page", _content)
	bo0.SetPosition(_pos + 1)
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
		if v1, v0 := bo1.GetTopicId(), _topic.GetId(); v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetIcon(), _icon+"-page"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetTitle(), _title+"-page"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetSummary(), _summary+"-page"; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetPosition(), _pos+1; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetContent(), _content; v1 != v0 {
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

func doTestPageDaoCreateUpdateGet(t *testing.T, name string, dao PageDao) {
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_app := product.NewProduct(_tagVersion, _appId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_topic := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	_topic.SetPosition(_pos)

	_content := "page one"
	bo0 := NewPage(_tagVersion, _topic, _title+"-page", _icon+"-page", _summary+"-page", _content)
	bo0.SetPosition(_pos + 1)
	_email := "libro@libro"
	_age := float64(35)
	bo0.SetDataAttr("props.owner", _name)
	bo0.SetDataAttr("props.email", _email)
	bo0.SetDataAttr("age", _age)
	if ok, err := dao.Create(bo0); err != nil || !ok {
		t.Fatalf("%s failed: %#v / %s", name+"/Create", ok, err)
	}

	bo0.SetAppId(_appId + "-new").SetTopicId(_topic.GetId() + "-new").
		SetIcon(_icon + "-new").SetTitle(_title + "-new").SetSummary(_summary + "-new").SetPosition(_pos + 2).SetContent(_content + "-new").
		SetTagVersion(_tagVersion + 3)
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
		if v1, v0 := bo1.GetTopicId(), _topic.GetId()+"-new"; v1 != v0 {
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
		if v1, v0 := bo1.GetPosition(), _pos+2; v1 != v0 {
			t.Fatalf("%s failed: expected %#v but received %#v", name, v0, v1)
		}
		if v1, v0 := bo1.GetContent(), _content+"-new"; v1 != v0 {
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

func doTestPageDaoCreateDelete(t *testing.T, name string, dao PageDao) {
	_tagVersion := uint64(1337)
	_appId := "libro"
	_name := "Libro"
	_desc := "Libro description"
	_isPublished := true
	_app := product.NewProduct(_tagVersion, _appId, _name, _desc, _isPublished)

	_title := "Quick start"
	_icon := "default"
	_summary := "topic one"
	_pos := rand.Intn(10242048)
	_topic := NewTopic(_tagVersion, _app, _title, _icon, _summary)
	_topic.SetPosition(_pos)

	_content := "page one"
	bo0 := NewPage(_tagVersion, _topic, _title+"-page", _icon+"-page", _summary+"-page", _content)
	bo0.SetPosition(_pos + 1)
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

func doTestPageDaoGetAll(t *testing.T, name string, dao PageDao) {
	initSampleRowsPage(t, name, dao)
	for _, topic := range topicList {
		boList, err := dao.GetAll(topic, nil, nil)
		expected := topicPageCount[topic.GetId()]
		if err != nil || len(boList) != expected {
			t.Fatalf("%s failed: expected %#v but received %#v (error %s)", name+"/GetAll", expected, len(boList), err)
		}
	}
}

func doTestPageDaoGetN(t *testing.T, name string, dao PageDao) {
	initSampleRowsPage(t, name, dao)
	for _, topic := range topicList {
		startOffset := rand.Intn(5)
		numRowsLimit := rand.Intn(10) + 1
		boList, err := dao.GetN(topic, startOffset, numRowsLimit, nil, nil)
		expected := topicPageCount[topic.GetId()]
		expected = int(math.Min(math.Max(0, float64(expected-startOffset)), float64(numRowsLimit)))
		if err != nil || len(boList) != expected {
			t.Fatalf("%s failed: expected %#v but received %#v (error %s)", name+"/"+topic.GetId(), expected, len(boList), err)
		}
	}
}
