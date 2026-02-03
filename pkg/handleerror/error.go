package handleerror

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	ErrNotFound = errors.New("未找到数据")
	// 其他自定义错误类型...
)

func JsonError(c *gin.Context, msg interface{}) {
	c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": msg})
}

func HandleError(c *gin.Context, err error) bool {
	if err != nil {
		//logrus.WithError(err).Error("gin context http handler error")
		JsonError(c, err.Error())
		return true
	}
	return false
}

func WsHandleError(ws *websocket.Conn, err error) bool {
	if err != nil {
		dt := time.Now().Add(time.Second)
		if writeErr := ws.WriteControl(websocket.CloseMessage, []byte(err.Error()), dt); writeErr != nil {
		}
		return true
	}
	return false
}
