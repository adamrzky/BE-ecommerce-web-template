package controllers

import (
	"BE-ecommerce-web-template/utils/resp"
	"BE-ecommerce-web-template/utils/token"

	"github.com/gin-gonic/gin"
)

type DummyController struct {
}

func NewDummyController() *DummyController {
	return &DummyController{}
}

func (controller *DummyController) MyClaims(c *gin.Context) {
	// panggil token.ExtractClaims() function untuk mendapatkan properti dari JWTClaims
	// beberapa fields dari JWTClaims yang bisa digunakan: `ID`, `Username`, `Role`
	claims, err := token.ExtractClaims(c)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	// response model, for example purpose
	var response struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	// example of how we utilize every claims field
	response.ID = claims.ID
	response.Username = claims.Username
	response.Role = claims.Role

	resp.NewResponseSuccess(c, response, "claims received")
}

func (controller *DummyController) AdminAndDev(c *gin.Context) {
	resp.NewResponseSuccess(c, nil, "Welcome admin and dev!")
}
