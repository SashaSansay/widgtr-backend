package widget_model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
	"widgtr-backend/internal/models"
)

type WidgetType string

const (
	Paragraph WidgetType = "paragraph"
	Page                 = "page"
	HeadingH1            = "heading_h1"
	HeadingH2            = "heading_h2"
	HeadingH3            = "heading_h3"
	HeadingH4            = "heading_h4"
	HeadingH5            = "heading_h5"
	HeadingH6            = "heading_h6"
)

const CollectionName = "widgets"

var AllWidgetTypes = []struct {
	Value  WidgetType
	TSName string
}{
	{Paragraph, strings.ToUpper(string(Paragraph))},
	{Page, strings.ToUpper(Page)},
	{HeadingH1, strings.ToUpper(HeadingH1)},
	{HeadingH2, strings.ToUpper(HeadingH2)},
	{HeadingH3, strings.ToUpper(HeadingH3)},
	{HeadingH4, strings.ToUpper(HeadingH4)},
	{HeadingH5, strings.ToUpper(HeadingH5)},
	{HeadingH6, strings.ToUpper(HeadingH6)},
}

type TextContent struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}

type ParagraphContent struct {
	TextContent `json:"text_content"`
}
type HeadingH1Content struct {
	TextContent `json:"text_content"`
}
type HeadingH2Content struct {
	TextContent `json:"text_content"`
}
type HeadingH3Content struct {
	TextContent `json:"text_content"`
}
type HeadingH4Content struct {
	TextContent `json:"text_content"`
}
type HeadingH5Content struct {
	TextContent `json:"text_content"`
}
type HeadingH6Content struct {
	TextContent `json:"text_content"`
}

type Widget struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	CreatedBy  primitive.ObjectID   `json:"created_by"`
	CreatedAt  primitive.DateTime   `json:"created_at"`
	WidgetType WidgetType           `json:"type"`
	Children   []primitive.ObjectID `json:"children"`
	ParentID   *primitive.ObjectID  `json:"parentId"`
	Parent     *Widget              `json:"parent,omitempty"`
	Paragraph  *ParagraphContent    `json:"paragraph,omitempty"`
	HeadingH1  *HeadingH1Content    `json:"heading_h1,omitempty"`
	HeadingH2  *HeadingH2Content    `json:"heading_h2,omitempty"`
	HeadingH3  *HeadingH3Content    `json:"heading_h3,omitempty"`
	HeadingH4  *HeadingH4Content    `json:"heading_h4,omitempty"`
	HeadingH5  *HeadingH5Content    `json:"heading_h5,omitempty"`
	HeadingH6  *HeadingH6Content    `json:"heading_h6,omitempty"`
}

func NewWidget(widgetIn Widget) *Widget {
	id := widgetIn.ID
	if widgetIn.ID.IsZero() {
		id = primitive.NewObjectID()
	}
	createdAt := primitive.NewDateTimeFromTime(time.Now())
	children := widgetIn.Children
	if children == nil {
		children = []primitive.ObjectID{}
	}
	return &Widget{
		ID:         id,
		CreatedBy:  widgetIn.CreatedBy,
		CreatedAt:  createdAt,
		WidgetType: widgetIn.WidgetType,
		Children:   children,
		Paragraph:  widgetIn.Paragraph,
		HeadingH1:  widgetIn.HeadingH1,
		HeadingH2:  widgetIn.HeadingH2,
		HeadingH3:  widgetIn.HeadingH3,
		HeadingH4:  widgetIn.HeadingH4,
		HeadingH5:  widgetIn.HeadingH5,
		HeadingH6:  widgetIn.HeadingH6,
	}
}

func (w *Widget) InsertIntoDB() (*Widget, error) {
	_, err := models.DB.Collection(CollectionName).InsertOne(context.TODO(), w)

	if err != nil {
		return nil, err
	}

	return w, nil
}

func GetFromDB(id string) (*Widget, error) {
	var widget Widget
	err := models.DB.Collection(CollectionName).FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&widget)

	if err != nil {
		return nil, err
	}

	return &widget, nil
}
