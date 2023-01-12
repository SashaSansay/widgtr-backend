package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"widgtr-backend/internal/constants"
	"widgtr-backend/internal/models/user_model"
	"widgtr-backend/internal/models/widget_model"
	"widgtr-backend/internal/utils/reponse_handlers"
)

type CreatePageIn struct {
	ParentID primitive.ObjectID `json:"parent_id"`
}

func CreatePageHandler(c *gin.Context) {
	var input CreatePageIn

	if err := c.ShouldBindJSON(&input); err != nil {
		reponse_handlers.Handle400(c, err)
		return
	}

	user, _ := c.Get(constants.MetaUserKey)

	w := widget_model.Widget{
		WidgetType: widget_model.Page,
		Children:   nil,
		CreatedBy:  user.(*user_model.User).ID,
	}

	widget, err := widget_model.NewWidget(w).InsertIntoDB()

	if err != nil {
		reponse_handlers.Handle400(c, err)
		return
	}

	reponse_handlers.Handle201(c, widget)
}
