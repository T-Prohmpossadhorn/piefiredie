package route

import (
	"net/http"
	"piefiredie/internal/bootstrap"
	"piefiredie/internal/pathhttp"
	"piefiredie/internal/usecase/summarize"

	"github.com/gin-gonic/gin"
)

type result struct {
	Beef map[string]int32 `json:"beef"`
}

func BeefStockSummary(c *gin.Context) {
	input, err := pathhttp.Get(bootstrap.Env.DataUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ret, err := summarize.SummarizeBeefStock(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := result{
		Beef: ret,
	}

	c.JSON(http.StatusOK, result)
}
