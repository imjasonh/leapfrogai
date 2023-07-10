package leapfrog

import "github.com/gin-gonic/gin"

type LeapfrogHandler struct {
	Prefix string
}

func (l *LeapfrogHandler) Routes(r *gin.Engine) {
	sr := r.Group(l.Prefix)
	{
		sr.GET("/config", l.Config)
	}
}
