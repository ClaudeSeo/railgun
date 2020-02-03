package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/claudeseo/railgun/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type PlaceController struct{}

type place struct {
	PlaceID     *string `form:"placeId"`
	PlaceNumber *int    `form:"placeNumber"`
}

type placeUpdateInfo struct {
	LastUpdatedAt int `json:"last_updated_at"`
}

func validatePlace(p place) error {
	var msg string

	if p.PlaceID == nil || *p.PlaceID == "" {
		msg = "`placeId` 는 필수입니다."
	} else if p.PlaceNumber == nil || *p.PlaceNumber == 0 {
		msg = "`placeNumber` 는 필수입니다."
	} else {
		return nil
	}

	return errors.New(msg)
}

func makeHash(p place) string {
	n := strings.Join([]string{strconv.Itoa(*p.PlaceNumber), *p.PlaceID}, "|")
	h := sha1.New()
	h.Write([]byte(n))
	return hex.EncodeToString(h.Sum(nil))
}

func findByHash(h string) *placeUpdateInfo {
	r := database.GetRedis()
	result, err := r.Get(h).Result()

	if err == redis.Nil {
		return nil
	} else if err != nil {
		panic(err)
	}

	p := placeUpdateInfo{}
	json.Unmarshal([]byte(result), &p)

	return &p
}

func (ctrl PlaceController) GetPlace(c *gin.Context) {
	var p place
	c.ShouldBindQuery(&p)

	if err := validatePlace(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	h := makeHash(p)
	result := findByHash(h)

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "매장 정보가 존재하지 않습니다.",
			"status":  http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
