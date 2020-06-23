// +build !noguigiu

package main

import (
	"reflect"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/sqweek/dialog"
	"github.com/topxeq/tk"
)

func loadFont() {
	fonts := giu.Context.IO().Fonts()

	rangeVarT := tk.GetVar("FontRange")

	ranges := imgui.NewGlyphRanges()

	builder := imgui.NewFontGlyphRangesBuilder()

	if rangeVarT == nil {
		builder.AddRanges(fonts.GlyphRangesDefault())
	} else {
		rangeStrT := rangeVarT.(string)
		if rangeStrT == "" {
			builder.AddRanges(fonts.GlyphRangesChineseSimplifiedCommon())
			builder.AddText(specialCharsG)
		} else if tk.StartsWith(rangeStrT, "COMMON") {
			builder.AddRanges(fonts.GlyphRangesChineseSimplifiedCommon())
			builder.AddText(specialCharsG + rangeStrT[6:])
		} else if rangeStrT == "FULL" {
			builder.AddRanges(fonts.GlyphRangesChineseFull())
			builder.AddText(specialCharsG)
		} else {
			builder.AddText(rangeStrT)
		}
	}

	builder.BuildRanges(ranges)

	fontPath := "c:/Windows/Fonts/simhei.ttf"

	if tk.Contains(tk.GetOSName(), "rwin") {
		fontPath = "/Library/Fonts/Microsoft/SimHei.ttf"
	}

	fontVarT := tk.GetVar("Font") // "c:/Windows/Fonts/simsun.ttc"

	if fontVarT != nil {
		fontPath = fontVarT.(string)
	}

	fontSizeStrT := "16"

	fontSizeVarT := tk.GetVar("FontSize")

	if fontSizeVarT != nil {
		fontSizeStrT = fontSizeVarT.(string)
	}

	fontSizeT := tk.StrToIntWithDefaultValue(fontSizeStrT, 16)

	// fonts.AddFontFromFileTTF(fontPath, 14)
	fonts.AddFontFromFileTTFV(fontPath, float32(fontSizeT), imgui.DefaultFontConfig, ranges.Data())
}

// var vgInch = float64(vg.Inch)

func getConfirmGUI(titleA string, messageA string) bool {
	return dialog.Message("%v", messageA).Title(titleA).YesNo()
}

func simpleInfo(titleA string, messageA string) {
	dialog.Message("%v", messageA).Title(titleA).Info()
}

func simpleError(titleA string, messageA string) {
	dialog.Message("%v", messageA).Title(titleA).Error()
}

// filename, err := dialog.File().Filter("XML files", "xml").Title("Export to XML").Save()
func selectFileToSaveGUI(titleA string, filterNameA string, filterTypeA string) string {
	fileNameT, errT := dialog.File().Filter(filterNameA, filterTypeA).Title(titleA).Save()

	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	return fileNameT
}

// fileNameT, errT := dialog.File().Filter("Mp3 audio file", "mp3").Load()
func selectFileGUI(titleA string, filterNameA string, filterTypeA string) string {
	fileNameT, errT := dialog.File().Filter(filterNameA, filterTypeA).Title(titleA).Load()

	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	return fileNameT
}

// directory, err := dialog.Directory().Title("Load images").Browse()
func selectDirectoryGUI(titleA string) string {
	directoryT, errT := dialog.Directory().Title(titleA).Browse()

	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	return directoryT
}

var specialCharsG = "ɪɔː辑ˌɡɜɔŋæʌʃɛəʒɪɑɒθʊː"

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

		"EditFile": reflect.ValueOf(editFile),
		// "LoopWindow": reflect.ValueOf(loopWindow),

		"LayoutP": reflect.ValueOf(giu.Layout{}),

		"Layout": reflect.ValueOf((*giu.Layout)(nil)),
		"Widget": reflect.ValueOf((*giu.Widget)(nil)),
	}

	GotxSymbols["giu"] = GotxSymbols["github.com/AllenDang/giu"]
}
