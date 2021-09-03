package doc

import (
	"main/src/gvabe/bo/app"
)

var (
	appList        []*app.App
	topicList      []*Topic
	pageList       []*Page
	appTopicCount  map[string]int
	topicPageCount map[string]int
)
