package main

import (
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"widgtr-backend/internal/models/user_model"
	"widgtr-backend/internal/models/widget_model"
)

func main() {
	converter := typescriptify.New().
		Add(user_model.User{}).
		Add(widget_model.Widget{}).
		Add(widget_model.TextContent{}).
		Add(widget_model.ParagraphContent{}).
		Add(widget_model.HeadingH1Content{}).
		Add(widget_model.HeadingH2Content{}).
		Add(widget_model.HeadingH3Content{}).
		Add(widget_model.HeadingH4Content{}).
		Add(widget_model.HeadingH5Content{}).
		Add(widget_model.HeadingH6Content{}).
		AddEnum(widget_model.AllWidgetTypes)
	converter.CreateInterface = true
	converter.BackupDir = ""
	err := converter.ConvertToFile("ts/models.ts")
	if err != nil {
		panic(err.Error())
	}
}
