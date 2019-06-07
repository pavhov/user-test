package postgres

import (
	//"github.com/gin-gonic/gin"
	//"gopkg.in/mgo.v2"
)

type Filter struct {
	Sort  string `form:"sort"`
	Skip  int    `form:"skip"`
	Limit int    `form:"limit"`
}

//func Build(c *gin.Context, q *mgo.Query) *mgo.Query {
//	var filter Filter
//
//	err := c.BindQuery(&filter)
//
//	if err == nil {
//		if filter.Sort != "" {
//			q.Sort(filter.Sort)
//		}
//
//		if filter.Skip != 0 {
//			q.Skip(filter.Skip)
//		}
//
//		if filter.Limit != 0 {
//			q.Limit(filter.Limit)
//		}
//	}
//
//	return q
//}
