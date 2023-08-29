// Code generated by hertz generator.

package feed

import (
	"context"
	"simple-douyin/api/biz/client"
	mw "simple-douyin/api/biz/middleware"
	bizFeed "simple-douyin/api/biz/model/feed"
	"simple-douyin/kitex_gen/feed"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	apiLog "github.com/sirupsen/logrus"
)

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var bizReq bizFeed.FeedRequest
	resp := new(bizFeed.FeedResponse)
	err = c.BindAndValidate(&bizReq)
	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57002
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
	}

	var userId = new(int64)
	apiLog.Info("Token: ", bizReq.GetToken())
	// 这里对token进行校验

	// 这里token是nil时，在路由时鉴权会报错
	//loggedClaims, exist := c.Get("JWT_PAYLOAD")
	//if !exist {
	//	resp.StatusCode = 57002
	//	if resp.StatusMsg == nil {
	//		resp.StatusMsg = new(string)
	//	}
	//	*resp.StatusMsg = "Unauthorized"
	//	c.JSON(consts.StatusOK, resp)
	//	return
	//}
	//userId = int64(loggedClaims.(jwt.MapClaims)[mw.JwtMiddleware.IdentityKey].(float64))

	if bizReq.Token != nil {
		_, err := mw.JwtMiddleware.ParseTokenString(*bizReq.Token)
		if err != nil {
			apiLog.Info(err)
			resp.StatusCode = 57001
			if resp.StatusMsg == nil {
				resp.StatusMsg = new(string)
			}
			*resp.StatusMsg = "Unauthorized"
			c.JSON(consts.StatusBadRequest, resp)
			return
		}
		// 用户token失效了也能用feed
		_, err = mw.JwtMiddleware.CheckIfTokenExpire(ctx, c)
		if err != nil {
			apiLog.Info(err)
			resp.StatusCode = 0
			if resp.StatusMsg == nil {
				resp.StatusMsg = new(string)
			}
			*resp.StatusMsg = "token expired"
			c.JSON(consts.StatusOK, resp)
		}
		claims, err := mw.JwtMiddleware.GetClaimsFromJWT(ctx, c)
		if err != nil {
			apiLog.Info(err)
			resp.StatusCode = 57001
			if resp.StatusMsg == nil {
				resp.StatusMsg = new(string)
			}
			*resp.StatusMsg = "Unauthorized"
			c.JSON(consts.StatusBadRequest, resp)
			return
		}
		*userId = int64(claims[mw.IdentityKey].(float64))
	}

	apiLog.Info("userId: ", userId)
	apiLog.Info("latestTime: ", bizReq.GetLatestTime())
	req := feed.FeedRequest{
		UserId:     userId,
		LatestTime: bizReq.LatestTime,
	}

	resp, err = client.Feed(ctx, &req)

	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57002
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
