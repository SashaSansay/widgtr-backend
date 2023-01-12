package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"widgtr-backend/internal/constants"
	"widgtr-backend/internal/models"
	"widgtr-backend/internal/models/user_model"
	"widgtr-backend/internal/models/widget_model"
	"widgtr-backend/internal/utils/reponse_handlers"
)

func FindAllWidgetsHandler(c *gin.Context) {
	widgets := []widget_model.Widget{}
	cursor, err := models.DB.Collection("widgets").Find(context.TODO(), bson.D{})
	if err != nil {
		reponse_handlers.Handle500(c, err)
		return
	}
	if err = cursor.All(context.TODO(), &widgets); err != nil {
		reponse_handlers.Handle500(c, err)
		return
	}

	reponse_handlers.Handle200(c, widgets)
}

type CreateWidgetIn struct {
	Type      widget_model.WidgetType       `json:"type" binding:"required"`
	ParentID  *primitive.ObjectID           `json:"parent_id"`
	Paragraph widget_model.ParagraphContent `json:"paragraph" binding:"dive"`
	HeadingH1 widget_model.HeadingH6Content `json:"heading_h1" binding:"dive"`
	HeadingH2 widget_model.HeadingH6Content `json:"heading_h2" binding:"dive"`
	HeadingH3 widget_model.HeadingH6Content `json:"heading_h3" binding:"dive"`
	HeadingH4 widget_model.HeadingH6Content `json:"heading_h4" binding:"dive"`
	HeadingH5 widget_model.HeadingH6Content `json:"heading_h5" binding:"dive"`
	HeadingH6 widget_model.HeadingH6Content `json:"heading_h6" binding:"dive"`
}

func CreateWidgetHandler(c *gin.Context) {
	var input CreateWidgetIn

	if err := c.ShouldBindJSON(&input); err != nil {
		reponse_handlers.Handle400(c, err)
		return
	}

	user, _ := c.Get(constants.MetaUserKey)

	w := widget_model.Widget{
		WidgetType: input.Type,
		Children:   nil,
		Paragraph:  &input.Paragraph,
		CreatedBy:  user.(*user_model.User).ID,
	}

	widget, err := widget_model.NewWidget(w).InsertIntoDB()

	if err != nil {
		reponse_handlers.Handle400(c, err)
		return
	}

	reponse_handlers.Handle201(c, widget)
}
