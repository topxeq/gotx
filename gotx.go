package main

import (
	"bufio"
	"bytes"

	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sync"
	"time"

	"fmt"

	// "runtime"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"

	_ "github.com/denisenkom/go-mssqldb"
	// "github.com/badgraph-io/badger"
	// _ "github.com/godror/godror"

	"image"
	"image/color"
	"image/draw"
	"image/png"

	// "github.com/fogleman/gg"
	// "github.com/topxeq/imagetk"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"

	// "gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	// "github.com/topxeq/sqltk"

	// "github.com/beevik/etree"

	// GUI related start
	"github.com/sqweek/dialog"
	// GUI related end

	"github.com/topxeq/tk"

	// GUI related start
	"github.com/AllenDang/giu"
	// g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	// GUI related end
)

// Non GUI related

var versionG = "0.996a"

var verboseG = false

var variableG = make(map[string]interface{})

var ygVMG *interp.Interpreter = nil

var varMutexG sync.Mutex

func ygEval(strA string) string {
	// if ygVMG == nil {
	// 	return tk.GenerateErrorString("VM nil")
	// }

	vmT := interp.New(interp.Options{})

	vmT.Use(stdlib.Symbols)

	GotxSymbols["builtin"] = map[string]reflect.Value{
		"eval":             reflect.ValueOf(ygEval),
		"printfln":         reflect.ValueOf(tk.Pl),
		"fprintf":          reflect.ValueOf(fmt.Fprintf),
		"pl":               reflect.ValueOf(tk.Pl),
		"pln":              reflect.ValueOf(fmt.Println),
		"plv":              reflect.ValueOf(tk.Plv),
		"plerr":            reflect.ValueOf(tk.PlErr),
		"exit":             reflect.ValueOf(tk.Exit),
		"setValue":         reflect.ValueOf(tk.SetValue),
		"getValue":         reflect.ValueOf(tk.GetValue),
		"bitXor":           reflect.ValueOf(tk.BitXor),
		"setVar":           reflect.ValueOf(tk.SetVar),
		"getVar":           reflect.ValueOf(tk.GetVar),
		"checkError":       reflect.ValueOf(tk.CheckError),
		"checkErrorString": reflect.ValueOf(tk.CheckErrorString),
		"getInput":         reflect.ValueOf(tk.GetUserInput),
		"getInputf":        reflect.ValueOf(tk.GetInputf),
		"run":              reflect.ValueOf(runFile),
		"typeOf":           reflect.ValueOf(tk.TypeOfValue),
		"typeOfReflect":    reflect.ValueOf(tk.TypeOfValueReflect),
		"remove":           reflect.ValueOf(tk.RemoveItemsInArray),
		"runScript":        reflect.ValueOf(runScript),
		"getClipText":      reflect.ValueOf(tk.GetClipText),
		"setClipText":      reflect.ValueOf(tk.SetClipText),
	}

	vmT.Use(GotxSymbols)

	_, errT := vmT.Eval(`import(. "builtin")`)
	if errT != nil {
		return tk.GenerateErrorStringF("failed to run init routine(%v): %v", "init", errT)

	}

	result, errT := vmT.Eval(strA)
	if errT != nil {
		return tk.GenerateErrorStringF("failed to run script: %v", errT)
	}

	return tk.Spr("%v", result)
}

func panicIt(valueA interface{}) {
	panic(valueA)
}

func runScript(codeA string, modeA string, argsA ...string) interface{} {
	if modeA == "" || modeA == "0" || modeA == "yg" {
		vmT := interp.New(interp.Options{})

		vmT.Use(stdlib.Symbols)

		GotxSymbols["builtin"] = map[string]reflect.Value{
			"eval":             reflect.ValueOf(ygEval),
			"printfln":         reflect.ValueOf(tk.Pl),
			"fprintf":          reflect.ValueOf(fmt.Fprintf),
			"pl":               reflect.ValueOf(tk.Pl),
			"pln":              reflect.ValueOf(fmt.Println),
			"plv":              reflect.ValueOf(tk.Plv),
			"plerr":            reflect.ValueOf(tk.PlErr),
			"exit":             reflect.ValueOf(tk.Exit),
			"setValue":         reflect.ValueOf(tk.SetValue),
			"getValue":         reflect.ValueOf(tk.GetValue),
			"bitXor":           reflect.ValueOf(tk.BitXor),
			"setVar":           reflect.ValueOf(tk.SetVar),
			"getVar":           reflect.ValueOf(tk.GetVar),
			"checkError":       reflect.ValueOf(tk.CheckError),
			"checkErrorString": reflect.ValueOf(tk.CheckErrorString),
			"getInput":         reflect.ValueOf(tk.GetUserInput),
			"getInputf":        reflect.ValueOf(tk.GetInputf),
			"run":              reflect.ValueOf(runFile),
			"typeOf":           reflect.ValueOf(tk.TypeOfValue),
			"typeOfReflect":    reflect.ValueOf(tk.TypeOfValueReflect),
			"remove":           reflect.ValueOf(tk.RemoveItemsInArray),
			"runScript":        reflect.ValueOf(runScript),
			"getClipText":      reflect.ValueOf(tk.GetClipText),
			"setClipText":      reflect.ValueOf(tk.SetClipText),
		}

		vmT.Use(GotxSymbols)

		_, errT := ygVMG.Eval(`import(. "builtin")`)
		if errT != nil {
			return tk.GenerateErrorStringF("failed to run init routine(%v): %v", "init", errT)

		}

		result, errT := vmT.Eval(codeA)
		if errT != nil {
			return tk.GenerateErrorStringF("failed to run script: %v", errT)
		}

		if result.IsValid() {
			return tk.Spr("%v", result)
		}

		return nil
	} else {
		return systemCmd("gotx", append([]string{codeA}, argsA...)...)
	}

}

func systemCmd(cmdA string, argsA ...string) string {
	var out bytes.Buffer

	cmd := exec.Command(cmdA, argsA...)

	cmd.Stdout = &out
	errT := cmd.Run()
	if errT != nil {
		return tk.GenerateErrorStringF("failed: %v", errT)
	}

	rStrT := tk.Trim(out.String())

	return rStrT
}

// full version related start
func newRGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

func newNRGBAFromHex(strA string) color.NRGBA {
	r, g, b, a := tk.ParseHexColor(strA)

	return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func newRGBAFromHex(strA string) color.RGBA {
	r, g, b, a := tk.ParseHexColor(strA)

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func newPlotXY(xA, yA float64) plotter.XY {
	return plotter.XY{X: xA, Y: yA}
}

func loadRGBAFromImage(imageA image.Image) (*image.RGBA, error) {
	switch imageT := imageA.(type) {
	case *image.RGBA:
		return imageT, nil
	default:
		rgba := image.NewRGBA(imageT.Bounds())
		draw.Draw(rgba, imageT.Bounds(), imageT, image.Pt(0, 0), draw.Src)
		return rgba, nil
	}

}

func LoadPlotImage(p *plot.Plot, w vg.Length, h vg.Length) (*image.RGBA, error) {

	var bufT bytes.Buffer

	writerT, errT := p.WriterTo(w, h, "png")

	if errT != nil {
		return nil, errT
	}

	_, errT = writerT.WriteTo(&bufT)

	if errT != nil {
		return nil, errT
	}

	readerT := bytes.NewReader(bufT.Bytes())

	// defer readerT.Close()

	// imgFile, err := os.Open(imgPath)
	// if err != nil {
	// 	return nil, err
	// }
	// defer imgFile.Close()

	img, err := png.Decode(readerT)
	if err != nil {
		return nil, err
	}

	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg, nil
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
		return rgba, nil
	}
}

func showHelp() {
	tk.Pl("Gotx by TopXeQ V%v\n", versionG)

	tk.Pl("Usage: gotx [-v|-h] test.gt\n")
	tk.Pl("or just gotx without arguments to start REPL instead.\n")

}

// full version related start

func runInteractive() int {
	var following bool
	var source string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if following {
			source += "\n"
			fmt.Print("  ")
		} else {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}

		source += scanner.Text()
		if source == "" {
			continue
		}

		if source == "quit()" {
			break
		}

		result, err := ygVMG.Eval(source)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			following = false
			source = ""
			continue
		}

		if result.IsValid() {
			fmt.Println(result)
		}

		following = false
		source = ""
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "ReadString error:", err)
			return 12
		}
	}

	return 0
}

// Non GUI related end

// GUI related start

// full version related start

var specialCharsG = "ɪɔː辑ˌɡɜɔŋæʌʃɛəʒɪɑɒθʊː"

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

// full version related end

// full version related start

var vgInch = float64(vg.Inch)

// full version related end

// full version related start
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

var (
	editorG            imgui.TextEditor
	errMarkersG        imgui.ErrorMarkers
	editFileNameG      string
	editFileCleanFlagG string
	editSecureCodeG    string
	editArgsG          string
)

func editorLoad() {
	if editorG.IsTextChanged() {
		editFileCleanFlagG = "*"
	}

	if editFileCleanFlagG != "" {
		rs := getConfirmGUI("Please confirm", "File modified, load another file anyway?")

		if rs == false {
			return
		}
	}

	fileNameNewT := selectFileGUI("Select the file to open...", "All files", "*")

	if tk.IsErrorString(fileNameNewT) {
		if tk.EndsWith(fileNameNewT, "Cancelled") {
			giu.Msgbox("Info", tk.Spr("Action cancelled by user"))
			return
		}

		giu.Msgbox("Error", tk.Spr("Failed to select file: %v", tk.GetErrorString(fileNameNewT)))
		return
	}

	fcT := tk.LoadStringFromFile(fileNameNewT)

	if tk.IsErrorString(fcT) {
		giu.Msgbox("Error", tk.Spr("Failed to load file content: %v", tk.GetErrorString(fileNameNewT)))
		return
	}

	editFileNameG = fileNameNewT
	editorG.SetText(fcT)
	editFileCleanFlagG = ""

}

func editorSaveAs() {
	fileNameNewT := selectFileToSaveGUI("Select the file to save...", "All file", "*")

	if tk.IsErrorString(fileNameNewT) {
		if tk.EndsWith(fileNameNewT, "Cancelled") {
			giu.Msgbox("Info", tk.Spr("Action cancelled by user"))
			return
		}

		giu.Msgbox("Error", tk.Spr("Failed to select file: %v", tk.GetErrorString(fileNameNewT)))
		return
	}

	editFileNameG = fileNameNewT

	rs := true
	// if tk.IfFileExists(editFileNameG) {
	// 	rs = getConfirmGUI("请再次确认", "文件已存在，是否覆盖?")
	// }

	if rs == true {
		rs1 := tk.SaveStringToFile(editorG.GetText(), editFileNameG)

		if rs1 != "" {
			giu.Msgbox("Error", tk.Spr("Failed to save: %v", rs))
			return
		}

		giu.Msgbox("Info", tk.Spr("File saved to: %v", editFileNameG))

		editFileCleanFlagG = ""
	}

}

func editorSave() {
	if editFileNameG == "" {
		editorSaveAs()

		return
	}

	rs := false

	if tk.IfFileExists(editFileNameG) {
		rs = getConfirmGUI("Please confirm", "The file already exists, confirm to overwrite?")
	}

	if rs == true {
		rs1 := tk.SaveStringToFile(editorG.GetText(), editFileNameG)

		if rs1 != "" {
			giu.Msgbox("Error", tk.Spr("Failed to save: %v", rs))
			return
		}

		giu.Msgbox("Info", tk.Spr("File saved to file: %v", editFileNameG))

		editFileCleanFlagG = ""
	}

}

func editEncrypt() {
	imgui.CloseCurrentPopup()

	sourceT := editorG.GetText()

	encStrT := tk.EncryptStringByTXDEF(sourceT, editSecureCodeG)

	if tk.IsErrorString(encStrT) {
		simpleError("Error", tk.Spr("failed to encrypt content: %v", tk.GetErrorString(encStrT)))
		return
	}

	editorG.SetText("//TXDEF#" + encStrT)
	editFileCleanFlagG = "*"

	editSecureCodeG = ""
}

func editEncryptClick() {
	giu.OpenPopup("Please enter:##EncryptInputSecureCode")
}

func editDecrypt() {
	imgui.CloseCurrentPopup()

	sourceT := tk.Trim(editorG.GetText())

	encStrT := tk.DecryptStringByTXDEF(sourceT, editSecureCodeG)

	if tk.IsErrorString(encStrT) {
		simpleError("Error", tk.Spr("failed to decrypt content: %v", tk.GetErrorString(encStrT)))
		return
	}

	editorG.SetText(encStrT)
	editFileCleanFlagG = "*"
	editSecureCodeG = ""

}

func editDecryptClick() {
	giu.OpenPopup("Please enter:##DecryptInputSecureCode")
}

func editRun() {
	imgui.CloseCurrentPopup()

	runScript(editorG.GetText(), "", editArgsG)
}

func editRunClick() {
	giu.OpenPopup("Please enter:##RunInputArgs")
}

func onButtonCloseClick() {
	os.Exit(0)
}

func loopWindow(windowA *giu.MasterWindow, loopA func()) {
	// wnd := giu.NewMasterWindow("Gotx Editor", 800, 600, 0, loadFont)

	windowA.Main(loopA)

}

func editorLoop() {
	giu.SingleWindow("Gotx Editor", giu.Layout{
		giu.Label(editFileNameG + editFileCleanFlagG),
		giu.Dummy(30, 0),
		giu.Line(
			giu.Button("Load", editorLoad),
			giu.Button("Save", editorSave),
			giu.Button("Save As...", editorSaveAs),
			giu.Button("Check", func() {

				// sourceT := editorG.GetText()

				// parser.EnableErrorVerbose()
				// _, errT := parser.ParseSrc(sourceT)
				// // tk.Plv(stmts)

				// e, ok := errT.(*parser.Error)

				// if ok {
				// 	errMarkersG.Clear()
				// 	errMarkersG.Insert(e.Pos.Line, tk.Spr("[col: %v, size: %v] %v", e.Pos.Column, errMarkersG.Size(), e.Error()))

				// 	editorG.SetErrorMarkers(errMarkersG)

				// } else if errT != nil {
				// 	giu.Msgbox("Error", tk.Spr("%#v", errT))
				// } else {
				// 	giu.Msgbox("Info", "Syntax check passed.")
				// }

			}),
			giu.Button("Encrypt", editEncryptClick),
			giu.Button("Decrypt", editDecryptClick),
			giu.Button("Run", editRunClick),
			giu.Button("Close", onButtonCloseClick),
			// giu.Button("Get Text", func() {
			// 	if editorG.HasSelection() {
			// 		fmt.Println(editorG.GetSelectedText())
			// 	} else {
			// 		fmt.Println(editorG.GetText())
			// 	}

			// 	column, line := editorG.GetCursorPos()
			// 	fmt.Println("Cursor pos:", column, line)

			// 	column, line = editorG.GetSelectionStart()
			// 	fmt.Println("Selection start:", column, line)

			// 	fmt.Println("Current line is", editorG.GetCurrentLineText())
			// }),
			// giu.Button("Set Text", func() {
			// 	editorG.SetText("Set text")
			// 	editFileNameG = "Set text"
			// }),
			// giu.Button("Set Error Marker", func() {
			// 	errMarkersG.Clear()
			// 	errMarkersG.Insert(1, "Error message")
			// 	fmt.Println("ErrMarkers Size:", errMarkersG.Size())

			// 	editorG.SetErrorMarkers(errMarkersG)
			// }),
		),
		giu.PopupModal("Please enter:##EncryptInputSecureCode", giu.Layout{
			giu.Line(
				giu.Label("Secure code"),
				giu.InputTextV("", 40, &editSecureCodeG, giu.InputTextFlagsPassword, nil, nil),
			),
			giu.Line(
				giu.Button("Ok", editEncrypt),
				giu.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		giu.PopupModal("Please enter:##DecryptInputSecureCode", giu.Layout{
			giu.Line(
				giu.Label("Secure code"),
				giu.InputTextV("", 40, &editSecureCodeG, giu.InputTextFlagsPassword, nil, nil),
			),
			giu.Line(
				giu.Button("Ok", editDecrypt),
				giu.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		giu.PopupModal("Please enter:##RunInputArgs", giu.Layout{
			giu.Line(
				giu.Label("Arguments to pass to VM"),
				giu.InputText("", 80, &editArgsG),
			),
			giu.Line(
				giu.Button("Ok", editRun),
				giu.Button("Cancel", func() { imgui.CloseCurrentPopup() }),
			),
		}),
		giu.Custom(func() {
			editorG.Render("Hello", imgui.Vec2{X: 0, Y: 0}, true)
			if giu.IsItemHovered() {
				if editorG.IsTextChanged() {
					editFileCleanFlagG = "*"
				}
			}
		}),
		giu.PrepareMsgbox(),
	})
}

func editFile(fileNameA string) {
	var fcT string

	if fileNameA == "" {
		editFileNameG = ""

		fcT = ""

		editFileCleanFlagG = "*"
	} else {
		editFileNameG = fileNameA

		fcT = tk.LoadStringFromFile(editFileNameG)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load file %v: %v", editFileNameG, tk.GetErrorString(fcT))
			return

		}

		editFileCleanFlagG = ""

	}

	errMarkersG = imgui.NewErrorMarkers()

	editorG = imgui.NewTextEditor()

	editorG.SetShowWhitespaces(false)
	editorG.SetTabSize(2)
	editorG.SetText(fcT)
	editorG.SetLanguageDefinitionC()

	// setVar("Font", "c:/Windows/Fonts/simsun.ttc")
	tk.SetVar("FontRange", "COMMON")
	tk.SetVar("FontSize", "15")

	tk.SetVar("Font", "c:/Windows/Fonts/simhei.ttf")

	wnd := giu.NewMasterWindow("Gotx Editor", 800, 600, 0, loadFont)
	// tk.Pl("%T", wnd)
	wnd.Main(editorLoop)

}

// full version related end

// GUI related end

func runFile(argsA ...string) interface{} {
	lenT := len(argsA)

	// full version related start
	// GUI related start

	if lenT < 1 {
		rs := selectFileGUI("Please select file to run...", "All files", "*")

		if tk.IsErrorString(rs) {
			return tk.Errf("Failed to load file: %v", tk.GetErrorString(rs))
		}

		fcT := tk.LoadStringFromFile(rs)

		if tk.IsErrorString(fcT) {
			return tk.Errf("Invalid file content: %v", tk.GetErrorString(fcT))
		}

		return runScript(fcT, "")

	}
	// GUI related end
	// full version related end

	if lenT < 1 {
		return nil
	}

	fcT := tk.LoadStringFromFile(argsA[0])

	if tk.IsErrorString(fcT) {
		return tk.Errf("Invalid file content: %v", tk.GetErrorString(fcT))
	}

	return runScript(fcT, "", argsA[1:]...)
}

var GotxSymbols = map[string]map[string]reflect.Value{}

func initYGVM() {
	if ygVMG == nil {
		ygVMG = interp.New(interp.Options{})

		ygVMG.Use(stdlib.Symbols)

		GotxSymbols["builtin"] = map[string]reflect.Value{
			"eval":             reflect.ValueOf(ygEval),
			"printfln":         reflect.ValueOf(tk.Pl),
			"fprintf":          reflect.ValueOf(fmt.Fprintf),
			"pl":               reflect.ValueOf(tk.Pl),
			"pln":              reflect.ValueOf(fmt.Println),
			"plv":              reflect.ValueOf(tk.Plv),
			"plerr":            reflect.ValueOf(tk.PlErr),
			"exit":             reflect.ValueOf(tk.Exit),
			"setValue":         reflect.ValueOf(tk.SetValue),
			"getValue":         reflect.ValueOf(tk.GetValue),
			"bitXor":           reflect.ValueOf(tk.BitXor),
			"setVar":           reflect.ValueOf(tk.SetVar),
			"getVar":           reflect.ValueOf(tk.GetVar),
			"checkError":       reflect.ValueOf(tk.CheckError),
			"checkErrorString": reflect.ValueOf(tk.CheckErrorString),
			"getInput":         reflect.ValueOf(tk.GetUserInput),
			"getInputf":        reflect.ValueOf(tk.GetInputf),
			"run":              reflect.ValueOf(runFile),
			"typeOf":           reflect.ValueOf(tk.TypeOfValue),
			"typeOfReflect":    reflect.ValueOf(tk.TypeOfValueReflect),
			"remove":           reflect.ValueOf(tk.RemoveItemsInArray),
			"runScript":        reflect.ValueOf(runScript),
			"getClipText":      reflect.ValueOf(tk.GetClipText),
			"setClipText":      reflect.ValueOf(tk.SetClipText),
		}

		ygVMG.Use(GotxSymbols)

		_, errT := ygVMG.Eval(`import(. "builtin")

		var argsG = os.Args[1:]
		`)
		if errT != nil {
			tk.Pl("failed to run init routine(%v): %v", "init", errT)

			return

		}

	}
}

func downloadStringFromSSH(sshA string, filePathA string) string {
	aryT := tk.Split(sshA, ":")

	basePathT, errT := tk.EnsureBasePath("gotx")

	if errT != nil {
		return tk.GenerateErrorStringF("failed to find base path: %v", errT)
	}

	if len(aryT) != 5 {
		aryT = tk.Split(tk.LoadStringFromFile(tk.JoinPath(basePathT, "ssh.cfg"))+filePathA, ":")

		if len(aryT) != 5 {
			return tk.ErrStrF("invalid ssh config: %v", "")
		}

	}

	clientT, errT := tk.NewSSHClient(aryT[0], tk.StrToIntWithDefaultValue(aryT[1], 22), aryT[2], aryT[3])

	if errT != nil {
		return tk.ErrToStrF("failed to create SSH client:", errT)
	}

	tmpPathT := tk.JoinPath(basePathT, "tmp")

	errT = tk.EnsureMakeDirsE(tmpPathT)

	if errT != nil {
		return tk.ErrToStrF("failed to create tmp dir:", errT)
	}

	tmpFileT, errT := tk.CreateTempFile(tmpPathT, "")

	if errT != nil {
		return tk.ErrToStrF("failed to create tmp dir:", errT)
	}

	defer os.Remove(tmpFileT)

	errT = clientT.Download(aryT[4], tmpFileT)

	if errT != nil {
		return tk.ErrToStrF("failed to download file:", errT)
	}

	fcT := tk.LoadStringFromFile(tmpFileT)

	return fcT
}

func getCfgString(fileNameA string) string {
	basePathT, errT := tk.EnsureBasePath("gotx")

	if errT == nil {
		cfgPathT := tk.JoinPath(basePathT, fileNameA)

		cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

		if !tk.IsErrorString(cfgStrT) {
			return cfgStrT
		}

		return tk.ErrStrF("failed to get config string: %v", tk.GetErrorString(cfgStrT))

	}

	return tk.ErrStrF("failed to get config string")
}

func main() {
	// var errT error
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
		}
	}()

	test()

	rand.Seed(time.Now().Unix())

	argsT := os.Args

	if tk.IfSwitchExistsWhole(argsT, "-version") {
		tk.Pl("Gotx by TopXeQ V%v", versionG)
		return
	}

	if tk.IfSwitchExistsWhole(argsT, "-h") {
		showHelp()
		return
	}

	scriptsT := tk.GetAllParameters(argsT)[1:]

	lenT := len(scriptsT)

	// GUI related start

	// full version related start
	if tk.IfSwitchExistsWhole(argsT, "-edit") {
		if lenT < 1 {
			editFile("")
		} else {
			editFile(scriptsT[0])
		}

		return
	}
	// full version related end

	// GUI related end

	if lenT < 1 {

		autoPathT := filepath.Join(tk.GetApplicationPath(), "auto.gt")

		if tk.IfFileExists(autoPathT) {
			scriptsT = []string{autoPathT}
		} else {
			initYGVM()

			runInteractive()

			// tk.Pl("not enough parameters")

			return
		}

	}

	encryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-encrypt=", "")

	if encryptCodeT != "" {
		for i, v := range scriptsT {
			fcT := tk.LoadStringFromFile(v)

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load file [%v] %v: %v", i, v, tk.GetErrorString(fcT))
				continue
			}

			encStrT := tk.EncryptStringByTXDEF(fcT, encryptCodeT)

			if tk.IsErrorString(encStrT) {
				tk.Pl("failed to encrypt content [%v] %v: %v", i, v, tk.GetErrorString(encStrT))
				continue
			}

			rsT := tk.SaveStringToFile("//TXDEF#"+encStrT, v+"e")

			if tk.IsErrorString(rsT) {
				tk.Pl("failed to encrypt file [%v] %v: %v", i, v, tk.GetErrorString(rsT))
				continue
			}
		}

		return
	}

	decryptCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrypt=", "")

	if decryptCodeT != "" {
		for i, v := range scriptsT {
			fcT := tk.LoadStringFromFile(v)

			if tk.IsErrorString(fcT) {
				tk.Pl("failed to load file [%v] %v: %v", i, v, tk.GetErrorString(fcT))
				continue
			}

			decStrT := tk.DecryptStringByTXDEF(fcT, decryptCodeT)

			if tk.IsErrorString(decStrT) {
				tk.Pl("failed to decrypt content [%v] %v: %v", i, v, tk.GetErrorString(decStrT))
				continue
			}

			rsT := tk.SaveStringToFile(decStrT, v+"d")

			if tk.IsErrorString(rsT) {
				tk.Pl("failed to decrypt file [%v] %v: %v", i, v, tk.GetErrorString(rsT))
				continue
			}
		}

		return
	}

	decryptRunCodeT := tk.GetSwitchWithDefaultValue(argsT, "-decrun=", "")

	if !tk.IfSwitchExistsWhole(argsT, "-m") {
		scriptsT = scriptsT[0:1]
	}

	// if verboseG {
	// 	tk.Pl("currenttid: %v", tk.GetCurrentThreadID())
	// }

	ifExampleT := tk.IfSwitchExistsWhole(argsT, "-example")
	ifGoPathT := tk.IfSwitchExistsWhole(argsT, "-gopath")
	ifLocalT := tk.IfSwitchExistsWhole(argsT, "-local")
	ifRemoteT := tk.IfSwitchExistsWhole(argsT, "-remote")
	ifCloudT := tk.IfSwitchExistsWhole(argsT, "-cloud")
	ifViewT := tk.IfSwitchExistsWhole(argsT, "-view")
	sshT := tk.GetSwitchWithDefaultValue(argsT, "-ssh=", "")
	ifShowResultT := tk.IfSwitchExistsWhole(argsT, "-showResult")
	verboseG = tk.IfSwitchExistsWhole(argsT, "-verbose")

	scriptT := scriptsT[0]

	var fcT string

	if ifExampleT {
		if !tk.EndsWith(scriptT, ".gt") {
			scriptT += ".gt"
		}
		fcT = tk.DownloadPageUTF8("https://gitee.com/topxeq/gotx/raw/master/scripts/"+scriptT, nil, "", 30)
	} else if ifRemoteT {
		fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
	} else if ifCloudT {
		if (!tk.EndsWith(scriptT, ".gt")) && (!tk.EndsWith(scriptT, ".yg")) && (!tk.EndsWith(scriptT, ".go")) {
			scriptT += ".gt"
		}

		basePathT, errT := tk.EnsureBasePath("gotx")

		gotT := false

		if errT == nil {
			cfgPathT := tk.JoinPath(basePathT, "cloud.cfg")

			cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

			if !tk.IsErrorString(cfgStrT) {
				fcT = tk.DownloadPageUTF8(cfgStrT+scriptT, nil, "", 30)

				gotT = true
			}

		}

		if !gotT {
			fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
		}
	} else if sshT != "" {
		if (!tk.EndsWith(scriptT, ".gt")) && (!tk.EndsWith(scriptT, ".yg")) && (!tk.EndsWith(scriptT, ".go")) {
			scriptT += ".gox"
		}

		fcT = downloadStringFromSSH(sshT, scriptT)

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to get script from SSH: %v", tk.GetErrorString(fcT))
			return

		}
	} else if ifGoPathT {
		if (!tk.EndsWith(scriptT, ".gt")) && (!tk.EndsWith(scriptT, ".yg")) && (!tk.EndsWith(scriptT, ".go")) {
			scriptT += ".gt"
		}

		fcT = tk.LoadStringFromFile(filepath.Join(tk.GetEnv("GOPATH"), "src", "github.com", "topxeq", "gotx", "scripts", scriptT))
	} else if ifLocalT {
		if (!tk.EndsWith(scriptT, ".gt")) && (!tk.EndsWith(scriptT, ".yg")) && (!tk.EndsWith(scriptT, ".go")) {
			scriptT += ".gt"
		}

		localPathT := getCfgString("localScriptPath.cfg")

		if tk.IsErrorString(localPathT) {
			tk.Pl("failed to get local path: %v", tk.GetErrorString(localPathT))

			return
		}

		fcT = tk.LoadStringFromFile(filepath.Join(localPathT, scriptT))
	} else {
		fcT = tk.LoadStringFromFile(scriptT)
	}

	if tk.IsErrorString(fcT) {
		tk.Pl("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))

		return
	}

	if tk.StartsWith(fcT, "//TXDEF#") {
		if decryptRunCodeT == "" {
			tk.Prf("Password: ")
			decryptRunCodeT = tk.Trim(tk.GetInputBufferedScan())

			// fcT = fcT[8:]
		}
	}

	if decryptRunCodeT != "" {
		fcT = tk.DecryptStringByTXDEF(fcT, decryptRunCodeT)
	}

	if ifViewT {
		tk.Pl("%v", fcT)

		return
	}

	if true {
		initYGVM()
		// i := interp.New(interp.Options{})

		// _, errT := ygVMG.Eval(`import(. "builtin")`)
		// if errT != nil {
		// 	tk.Pl("failed to run init routine(%v): %v", "init", errT)

		// 	if p, ok := errT.(interp.Panic); ok {
		// 		tk.Pl("%v", string(p.Stack))
		// 	}

		// 	return

		// }

		result, errT := ygVMG.Eval(fcT)
		if errT != nil {
			tk.Pl("failed to run script(%v): %v", scriptT, errT)

			if p, ok := errT.(interp.Panic); ok {
				tk.Pl("%v", string(p.Stack))
			}

		}

		if ifShowResultT {
			tk.Pl("%v", result)
		}

		return
	}
	// full version related end
}

func test() {

}
