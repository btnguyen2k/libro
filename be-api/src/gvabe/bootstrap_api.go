package gvabe

import (
	"reflect"
	"regexp"
	"strings"

	"main/src/gvabe/bo/user"
	"main/src/itineris"
)

/*
Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)
  - api-handler function must have the following signature: func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
*/
// TODO change this function to implement application's business logic
func initApiHandlers(router *itineris.ApiRouter) {
	router.SetHandler("info", apiInfo)
	router.SetHandler("login", apiLogin)
	router.SetHandler("verifyLoginToken", apiVerifyLoginToken)
	router.SetHandler("systemInfo", apiSystemInfo)

	router.SetHandler("adminGetStats", apiAdminGetStats)
	router.SetHandler("adminGetProductList", apiAdminGetProductList)
	router.SetHandler("adminAddProduct", apiAdminAddProduct)
	router.SetHandler("adminGetProduct", apiAdminGetProduct)
	router.SetHandler("adminUpdateProduct", apiAdminUpdateProduct)
	router.SetHandler("adminDeleteProduct", apiAdminDeleteProduct)
	router.SetHandler("adminMapDomain", apiAdminMapDomain)
	router.SetHandler("adminUnmapDomain", apiAdminUnmapDomain)
	router.SetHandler("adminGetProductTopics", apiAdminGetProductTopics)
	router.SetHandler("adminAddProductTopic", apiAdminAddProductTopic)
	// router.SetHandler("adminGetTopic", apiAdminGetTopic)
	// router.SetHandler("adminUpdateTopic", apiAdminUpdateTopic)
	// router.SetHandler("adminDeleteTopic", apiAdminDeleteTopic)

	// router.SetHandler("myFeed", apiMyFeed)
	// router.SetHandler("myBlog", apiMyBlog)
	// router.SetHandler("createBlogPost", apiCreateBlogPost)
	// router.SetHandler("getBlogPost", apiGetBlogPost)
	// router.SetHandler("updateBlogPost", apiUpdateBlogPost)
	// router.SetHandler("deleteBlogPost", apiDeleteBlogPost)
	//
	// router.SetHandler("getUserVoteForPost", apiGetUserVoteForPost)
	// router.SetHandler("voteForPost", apiVoteForPost)
}

/*------------------------------ shared variables and functions ------------------------------*/

var (
	// those APIs will not need authentication.
	// "false" means client, however, needs to send app-id along with the API call
	// "true" means the API is free for public call
	publicApis = map[string]bool{
		"login":            false,
		"info":             true,
		"getApp":           false,
		"verifyLoginToken": true,
		"loginChannelList": true,
	}
)

// available since template-v0.2.0
func _currentUserFromContext(ctx *itineris.ApiContext) (*SessionClaims, *user.User, error) {
	sessClaims, ok := ctx.GetContextValue(ctxFieldSession).(*SessionClaims)
	if !ok || sessClaims == nil {
		return nil, nil, nil
	}
	user, err := userDao.Get(sessClaims.UserId)
	return sessClaims, user, err
}

// available since template-v0.2.0
func _extractParam(params *itineris.ApiParams, paramName string, typ reflect.Type, defValue interface{}, regexp *regexp.Regexp) interface{} {
	v, _ := params.GetParamAsType(paramName, typ)
	if v == nil {
		v = defValue
	}
	if v != nil {
		if _, ok := v.(string); ok {
			v = strings.TrimSpace(v.(string))
			if regexp != nil && !regexp.Match([]byte(v.(string))) {
				return nil
			}
		}
	}
	return v
}


// /*
// apiUpdateBlogPost handles API call "updateBlogPost"
//
// @available since template-v0.2.0
// */
// func apiUpdateBlogPost(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
// 	_, user, err := _currentUserFromContext(ctx)
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	resultNoPermission := itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("current user has no permission to edit this post")
// 	if user == nil {
// 		return resultNoPermission
// 	}
// 	id := _extractParam(params, "id", reddo.TypeString, "", nil)
// 	blogPost, err := blogPostDaov2.Get(id.(string))
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	if blogPost == nil {
// 		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("post not found")
// 	}
// 	if blogPost.GetOwnerId() != user.GetId() {
// 		return resultNoPermission
// 	}
// 	isPublic := _extractParam(params, "is_public", reddo.TypeBool, false, nil)
// 	title := _extractParam(params, "title", reddo.TypeString, "", nil)
// 	if title == "" {
// 		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("title is empty")
// 	}
// 	content := _extractParam(params, "content", reddo.TypeString, "", nil)
// 	if content == "" {
// 		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("content is empty")
// 	}
// 	blogPost.SetPublic(isPublic.(bool)).SetTitle(title.(string)).SetContent(content.(string))
// 	ok, err := blogPostDaov2.Update(blogPost)
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	if !ok {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage("cannot update blog post")
// 	}
// 	return itineris.NewApiResult(itineris.StatusOk)
// }

// /*
// apiGetUserVoteForPost handles API call "getUserVoteForPost"
//
// @available since template-v0.2.0
// */
// func apiGetUserVoteForPost(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
// 	_, user, err := _currentUserFromContext(ctx)
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	if user == nil {
// 		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("current user not found")
// 	}
// 	postId := _extractParam(params, "postId", reddo.TypeString, "", nil).(string)
// 	vote, err := blogVoteDaov2.GetUserVoteForTarget(user, postId)
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	value := 0
// 	if vote != nil {
// 		value = vote.GetValue()
// 	}
// 	return itineris.NewApiResult(itineris.StatusOk).SetData(value)
// }

// /*
// apiVoteForPost handles API call "voteForPost"
//
// @available since template-v0.2.0
// */
// func apiVoteForPost(ctx *itineris.ApiContext, _ *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
// 	_, user, err := _currentUserFromContext(ctx)
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	if user == nil {
// 		return itineris.NewApiResult(itineris.StatusErrorClient).SetMessage("current user not found")
// 	}
// 	value := _extractParam(params, "vote", reddo.TypeInt, 0, nil).(int64)
// 	if value == 0 {
// 		return itineris.NewApiResult(itineris.StatusOk).SetData(map[string]interface{}{"vote": false})
// 	}
// 	postId := _extractParam(params, "postId", reddo.TypeString, "", nil).(string)
// 	blogPost, err := blogPostDaov2.Get(postId)
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	if blogPost == nil {
// 		return itineris.NewApiResult(itineris.StatusNotFound).SetMessage("post not found")
// 	}
// 	if !blogPost.IsPublic() && blogPost.GetOwnerId() != user.GetId() {
// 		return itineris.NewApiResult(itineris.StatusNoPermission).SetMessage("current user has no permission to existingVote for this post")
// 	}
// 	if value > 1 {
// 		value = 1
// 	} else if value < -1 {
// 		value = -1
// 	}
// 	existingVote, err := blogVoteDaov2.GetUserVoteForTarget(user, postId)
// 	if err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	log.Printf("Existing vote: %#v\n", existingVote)
// 	newVote := blog.NewBlogVote(goapi.AppVersionNumber, user, blogPost.GetId(), int(value))
// 	if existingVote == nil {
// 		// new vote
// 		if value > 0 {
// 			blogPost.IncNumVotesUp(1)
// 		} else {
// 			blogPost.IncNumVotesDown(1)
// 		}
// 		if _, err := blogVoteDaov2.Create(newVote); err != nil {
// 			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 		}
// 	} else {
// 		newVote.SetId(existingVote.GetId())
// 		if existingVote.GetValue() == newVote.GetValue() {
// 			// cancel existing vote
// 			newVote.SetValue(0)
// 			if value > 0 {
// 				blogPost.IncNumVotesUp(-1)
// 			} else {
// 				blogPost.IncNumVotesDown(-1)
// 			}
// 		} else {
// 			// chance existing vote
// 			if value > 0 {
// 				blogPost.IncNumVotesUp(1)
// 				if existingVote.GetValue() != 0 {
// 					blogPost.IncNumVotesDown(-1)
// 				}
// 			} else {
// 				if existingVote.GetValue() != 0 {
// 					blogPost.IncNumVotesUp(-1)
// 				}
// 				blogPost.IncNumVotesDown(1)
// 			}
// 		}
// 		log.Printf("New vote: %#v\n", newVote)
// 		if _, err := blogVoteDaov2.Update(newVote); err != nil {
// 			return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 		}
// 	}
// 	if _, err := blogPostDaov2.Update(blogPost); err != nil {
// 		return itineris.NewApiResult(itineris.StatusErrorServer).SetMessage(err.Error())
// 	}
// 	return itineris.NewApiResult(itineris.StatusOk).SetData(map[string]interface{}{
// 		"vote": true, "value": newVote.GetValue(), "num_votes_up": blogPost.GetNumVotesUp(), "num_votes_down": blogPost.GetNumVotesDown(),
// 	})
// }
