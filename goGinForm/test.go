package main

/*
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func TopicUrl(f1 validator.FieldLevel) bool {
	return true //返回true表示验证成功
}

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok { //类型断言
		v.RegisterValidation("topicurl", TopicUrl) //注册调用tag和自定义验证器
	}
	v1 := router.Group("/v1/topics")
	{
		v1.GET("/:topic_id", GetTopicDetail)
	}
}

type Topic struct {
	TopicID         int    `form:"id" json:"id"`
	TopicTitle      string `form:"title" json:"title" binding:"min=4,max=20"`                  //长度4到20之间
	TopicShortTitle string `form:"stitle" json:"stitle" binding:"required,nefield=ToipcTitle"` //required非空,nfield不能等于TopicTitle字段
	UserIP          string `form:"ip" json:"ip" binding:"ipv4"`
	TopicScore      int    `form:"score" json:"score" binding:"omitempty,gt=5"` //omitempty要么不传，传的话就要大于5
}
type Topic struct {
	TopicID         int    `form:"id" json:"id"`
	TopicTitle      string `form:"title" json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `form:"stitle" json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string `form:"ip" json:"ip" binding:"ipv4"`
	TopicScore      int    `form:"score" json:"score" binding:"omitempty,gt=5"`
	TopicUrl        string `form:"url" json:"url" binding:"topicurl"` //绑定自定义的topicurl验证器
}

type Topics struct {
	TopicList []*Topic `form:"topiclist" json:"topiclist" binding:"gt=0,lt=3,dive"` //dive表示进入列表或数组里面的字段去验证，这里是验证上面的Topic，dive要写在后面，不然如果lt写在dive后面就是验证里面的数据的lt>3，而不是列表的lt>3
	Size      int      `form:"size" json:"size"`
	Age       int      `form:"age" json:"age" binding:"required,oneof=1 2 3"` //age只能是1 2 3中一个
}
*/
