package main

import (
	"github.com/AllenDang/giu"
	"reflect"
)

func init() {
	GotxSymbols["github.com/AllenDang/giu"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewMasterWindow":         reflect.ValueOf(giu.NewMasterWindow),
		"SingleWindow":            reflect.ValueOf(giu.SingleWindow),
		"Window":                  reflect.ValueOf(giu.Window),
		"SingleWindowWithMenuBar": reflect.ValueOf(giu.SingleWindowWithMenuBar),
		"WindowV":                 reflect.ValueOf(giu.WindowV),

		"MasterWindowFlagsNotResizable": reflect.ValueOf(giu.MasterWindowFlagsNotResizable),
		"MasterWindowFlagsMaximized":    reflect.ValueOf(giu.MasterWindowFlagsMaximized),
		"MasterWindowFlagsFloating":     reflect.ValueOf(giu.MasterWindowFlagsFloating),

		// "Layout":          reflect.ValueOf(giu.Layout),

		"NewTextureFromRgba": reflect.ValueOf(giu.NewTextureFromRgba),

		"Label":                  reflect.ValueOf(giu.Label),
		"Line":                   reflect.ValueOf(giu.Line),
		"Button":                 reflect.ValueOf(giu.Button),
		"InvisibleButton":        reflect.ValueOf(giu.InvisibleButton),
		"ImageButton":            reflect.ValueOf(giu.ImageButton),
		"InputTextMultiline":     reflect.ValueOf(giu.InputTextMultiline),
		"Checkbox":               reflect.ValueOf(giu.Checkbox),
		"RadioButton":            reflect.ValueOf(giu.RadioButton),
		"Child":                  reflect.ValueOf(giu.Child),
		"ComboCustom":            reflect.ValueOf(giu.ComboCustom),
		"Combo":                  reflect.ValueOf(giu.Combo),
		"ContextMenu":            reflect.ValueOf(giu.ContextMenu),
		"Group":                  reflect.ValueOf(giu.Group),
		"Image":                  reflect.ValueOf(giu.Image),
		"ImageWithFile":          reflect.ValueOf(giu.ImageWithFile),
		"ImageWithUrl":           reflect.ValueOf(giu.ImageWithUrl),
		"InputText":              reflect.ValueOf(giu.InputText),
		"InputTextV":             reflect.ValueOf(giu.InputTextV),
		"InputTextFlagsPassword": reflect.ValueOf(giu.InputTextFlagsPassword),
		"InputInt":               reflect.ValueOf(giu.InputInt),
		"InputFloat":             reflect.ValueOf(giu.InputFloat),
		"MainMenuBar":            reflect.ValueOf(giu.MainMenuBar),
		"MenuBar":                reflect.ValueOf(giu.MenuBar),
		"MenuItem":               reflect.ValueOf(giu.MenuItem),
		"PopupModal":             reflect.ValueOf(giu.PopupModal),
		"OpenPopup":              reflect.ValueOf(giu.OpenPopup),
		"CloseCurrentPopup":      reflect.ValueOf(giu.CloseCurrentPopup),
		"ProgressBar":            reflect.ValueOf(giu.ProgressBar),
		"Separator":              reflect.ValueOf(giu.Separator),
		"SliderInt":              reflect.ValueOf(giu.SliderInt),
		"SliderFloat":            reflect.ValueOf(giu.SliderFloat),
		"HSplitter":              reflect.ValueOf(giu.HSplitter),
		"VSplitter":              reflect.ValueOf(giu.VSplitter),
		"TabItem":                reflect.ValueOf(giu.TabItem),
		"TabBar":                 reflect.ValueOf(giu.TabBar),
		"Row":                    reflect.ValueOf(giu.Row),
		"Table":                  reflect.ValueOf(giu.Table),
		"FastTable":              reflect.ValueOf(giu.FastTable),
		"Tooltip":                reflect.ValueOf(giu.Tooltip),
		"TreeNode":               reflect.ValueOf(giu.TreeNode),
		"Spacing":                reflect.ValueOf(giu.Spacing),
		"Custom":                 reflect.ValueOf(giu.Custom),
		"Condition":              reflect.ValueOf(giu.Condition),
		"ListBox":                reflect.ValueOf(giu.ListBox),
		"DatePicker":             reflect.ValueOf(giu.DatePicker),
		"Dummy":                  reflect.ValueOf(giu.Dummy),
		// "Widget":             reflect.ValueOf(giu.Widget),

		"PrepareMessageBox": reflect.ValueOf(giu.PrepareMsgbox),
		"MessageBox":        reflect.ValueOf(giu.Msgbox),

		"LoadFont": reflect.ValueOf(loadFont),

		"GetConfirm": reflect.ValueOf(getConfirmGUI),

		"SimpleInfo":      reflect.ValueOf(simpleInfo),
		"SimpleError":     reflect.ValueOf(simpleError),
		"SelectFile":      reflect.ValueOf(selectFileGUI),
		"SelectSaveFile":  reflect.ValueOf(selectFileToSaveGUI),
		"SelectDirectory": reflect.ValueOf(selectDirectoryGUI),

		"EditFile":   reflect.ValueOf(editFile),
		"LoopWindow": reflect.ValueOf(loopWindow),

		"LayoutP": reflect.ValueOf(giu.Layout{}),

		"Layout": reflect.ValueOf((*giu.Layout)(nil)),
		"Widget": reflect.ValueOf((*giu.Widget)(nil)),
	}

	GotxSymbols["giu"] = GotxSymbols["github.com/AllenDang/giu"]
}
