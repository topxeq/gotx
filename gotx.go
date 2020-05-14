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
	"strings"
	"sync"
	"time"

	"errors"
	"fmt"

	// "runtime"

	"net"
	"strconv"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"

	"github.com/atotto/clipboard"

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

var versionG = "0.985a"

var verboseG = false

var variableG = make(map[string]interface{})

var ygVMG *interp.Interpreter = nil

var varMutexG sync.Mutex

func exit() {
	defer func() {
		if r := recover(); r != nil {
			tk.Printfln("发生异常，错误信息：%v", r)

			return
		}
	}()

	os.Exit(1)
}

func getVar(nameA string) interface{} {
	varMutexG.Lock()
	rs, ok := variableG[nameA]
	varMutexG.Unlock()

	if !ok {
		tk.GenerateErrorString("no key")
	}
	return rs
}

func setVar(nameA string, valueA interface{}) {
	varMutexG.Lock()
	variableG[nameA] = valueA
	varMutexG.Unlock()
}

func ygEval(strA string) string {
	// if ygVMG == nil {
	// 	return tk.GenerateErrorString("VM nil")
	// }

	vmT := interp.New(interp.Options{})

	result, errT := vmT.Eval(strA)
	if errT != nil {
		return tk.GenerateErrorStringF("failed to run script: %v", errT)
	}

	return tk.Spr("%v", result)
}

func getClipText() string {
	textT, errT := clipboard.ReadAll()
	if errT != nil {
		return tk.GenerateErrorStringF("could not get text from clipboard: %v", errT.Error())
	}

	return textT
}

func setClipText(textA string) {
	clipboard.WriteAll(textA)
}

func panicIt(valueA interface{}) {
	panic(valueA)
}

func checkErrorFunc(errA error, funcA func()) {
	if errA != nil {
		tk.PlErr(errA)

		if funcA != nil {
			funcA()
		}

		os.Exit(1)
	}

}

func checkError(errA error, funcsA ...(func())) {
	if errA != nil {
		tk.PlErr(errA)

		if funcsA != nil {
			for _, v := range funcsA {
				v()
			}
		}

		os.Exit(1)
	}

}

func checkErrorString(strA string, funcA func()) {
	if tk.IsErrorString(strA) {
		tk.PlErrString(strA)

		if funcA != nil {
			funcA()
		}

		os.Exit(1)
	}

}

func newSSHClient(hostA string, portA int, userA string, passA string) (*goph.Client, error) {
	authT := goph.Password(passA)

	clientT := &goph.Client{
		Addr: hostA,
		Port: portA,
		User: userA,
		Auth: authT,
	}

	errT := goph.Conn(clientT, &ssh.ClientConfig{
		User:    clientT.User,
		Auth:    clientT.Auth,
		Timeout: 20 * time.Second,
		HostKeyCallback: func(host string, remote net.Addr, key ssh.PublicKey) error {
			return nil
			// hostFound, err := goph.CheckKnownHost(host, remote, key, "")

			// if hostFound && err != nil {
			// 	return err
			// }

			// if hostFound && err == nil {
			// 	return nil
			// }

			// return goph.AddKnownHost(host, remote, key, "")
		},
	})

	// clientT, errT := goph.NewConn(userA, hostA, authT, func(host string, remote net.Addr, key ssh.PublicKey) error {

	// 	hostFound, err := goph.CheckKnownHost(host, remote, key, "")

	// 	if hostFound && err != nil {
	// 		return err
	// 	}

	// 	if hostFound && err == nil {
	// 		return nil
	// 	}

	// 	return goph.AddKnownHost(host, remote, key, "")
	// })

	return clientT, errT
}

func remove(aryA []interface{}, startA int, endA int) []interface{} {
	if startA < 0 || startA >= len(aryA) {
		tk.Pl("Runtime error: %v", "index out of range")
		exit()
	}

	if endA < 0 || endA >= len(aryA) {
		tk.Pl("Runtime error: %v", "index out of range")
		exit()
	}

	return append(aryA[:startA], aryA[endA+1:]...)
	// if idxT == 0 {
	// 	return ayrA[idxT + 1:]
	// }

	// if idxT == len(aryA) - 1 {
	// 	return ayrA[0:len(aryA) - 1]
	// }

	// return append(aryA[:idxA], aryA[idxA+1:]...)

}

func toStringFromRuneSlice(sliceA []rune) string {
	return string(sliceA)
}

// toInt converts all reflect.Value-s into int.
func toInt(vA interface{}) int {
	v := reflect.ValueOf(&vA)
	i, _ := tryToInt(v)
	return i
}

// tryToInt attempts to convert a value to an int.
// If it cannot (in the case of a non-numeric string, a struct, etc.)
// it returns 0 and an error.
func tryToInt(v reflect.Value) (int, error) {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float64, reflect.Float32:
		return int(v.Float()), nil
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return int(v.Int()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.String:
		s := v.String()
		var i int64
		var err error
		if strings.HasPrefix(s, "0x") {
			i, err = strconv.ParseInt(s, 16, 64)
		} else {
			i, err = strconv.ParseInt(s, 10, 64)
		}
		if err == nil {
			return int(i), nil
		}
	}
	return 0, errors.New("couldn't convert to integer")
}

func getUint64Value(v reflect.Value) uint16 {
	tk.Pl("%x", v.Interface())

	var p *uint16

	p = (v.Interface().(*uint16))

	return *p
}

func runScript(codeA string, modeA string, argsA ...string) interface{} {

	if modeA == "" || modeA == "0" || modeA == "ql" {
		// vmT := qlang.New()

		// if argsA != nil && len(argsA) > 0 {
		// 	vmT.SetVar("argsG", argsA)
		// }

		// retG = notFoundG

		// errT := vmT.SafeEval(codeA)

		// if errT != nil {
		// 	return errT.Error()
		// }

		// // if retG != notFoundG {
		// // 	fmt.Println(retG)
		// // }

		// // rs, _ := vmT.GetVar("outG")

		// // if !ok {
		// // 	return ""
		// // }

		// return retG
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

func typeOfValue(vA interface{}) string {
	return fmt.Sprintf("%T", vA)
}

func typeOfValueReflect(vA interface{}) string {
	rs := reflect.TypeOf(vA)
	return rs.String()
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

func setValue(p interface{}, v interface{}) {
	// tk.Pl("%#v", reflect.TypeOf(p).Kind())
	// p = v

	srcRef := reflect.ValueOf(v)
	vp := reflect.ValueOf(p)
	vp.Elem().Set(srcRef)
}

func getValue(p interface{}) interface{} {
	vp := reflect.Indirect(reflect.ValueOf(p))
	return vp.Interface()
}

func bitXor(p interface{}, v interface{}) interface{} {
	switch p.(type) {
	case int:
		return p.(int) ^ v.(int)
	case int64:
		return p.(int64) ^ v.(int64)
	case int32:
		return p.(int32) ^ v.(int32)
	case int16:
		return p.(int16) ^ v.(int16)
	case int8:
		return p.(int8) ^ v.(int8)
	case uint64:
		return p.(uint64) ^ v.(uint64)
	case uint32:
		return p.(uint32) ^ v.(uint32)
	case uint16:
		return p.(uint16) ^ v.(uint16)
	case uint8:
		return p.(uint8) ^ v.(uint8)
	case uint:
		return p.(uint) ^ v.(uint)
	}

	return 0
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

		fmt.Println(result)

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

func loadFont() {
	fonts := giu.Context.IO().Fonts()

	rangeVarT := getVar("FontRange")

	ranges := imgui.NewGlyphRanges()

	builder := imgui.NewFontGlyphRangesBuilder()

	if rangeVarT == nil {
		builder.AddRanges(fonts.GlyphRangesDefault())
	} else {
		rangeStrT := rangeVarT.(string)
		if rangeStrT == "" || tk.StartsWith(rangeStrT, "COMMON") {
			builder.AddRanges(fonts.GlyphRangesChineseSimplifiedCommon())
			builder.AddText("辑" + rangeStrT[6:])
		} else if rangeStrT == "FULL" {
			builder.AddRanges(fonts.GlyphRangesChineseFull())
		} else {
			builder.AddText(rangeStrT)
		}
	}

	builder.BuildRanges(ranges)

	fontPath := "c:/Windows/Fonts/simhei.ttf"

	if tk.Contains(tk.GetOSName(), "rwin") {
		fontPath = "/Library/Fonts/Microsoft/SimHei.ttf"
	}

	fontVarT := getVar("Font") // "c:/Windows/Fonts/simsun.ttc"

	if fontVarT != nil {
		fontPath = fontVarT.(string)
	}

	fontSizeStrT := "16"

	fontSizeVarT := getVar("FontSize")

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
	exit()
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
	setVar("FontRange", "COMMON")
	setVar("FontSize", "15")

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
			"exit":             reflect.ValueOf(exit),
			"setValue":         reflect.ValueOf(setValue),
			"getValue":         reflect.ValueOf(getValue),
			"bitXor":           reflect.ValueOf(bitXor),
			"setVar":           reflect.ValueOf(setVar),
			"getVar":           reflect.ValueOf(getVar),
			"checkError":       reflect.ValueOf(checkError),
			"checkErrorString": reflect.ValueOf(checkErrorString),
			"getInput":         reflect.ValueOf(tk.GetUserInput),
			"getInputf":        reflect.ValueOf(tk.GetInputf),
			"newSSHClient":     reflect.ValueOf(newSSHClient),
			"run":              reflect.ValueOf(runFile),
			"typeOf":           reflect.ValueOf(typeOfValueReflect),
			"remove":           reflect.ValueOf(remove),
			"runScript":        reflect.ValueOf(runScript),
			"getClipText":      reflect.ValueOf(getClipText),
			"setClipText":      reflect.ValueOf(setClipText),
		}

		GotxSymbols["native"] = GotxSymbols["builtin"]

		ygVMG.Use(GotxSymbols)

	}
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
	ifRemoteT := tk.IfSwitchExistsWhole(argsT, "-remote")
	ifCloudT := tk.IfSwitchExistsWhole(argsT, "-cloud")
	ifViewT := tk.IfSwitchExistsWhole(argsT, "-view")
	ifShowResultT := tk.IfSwitchExistsWhole(argsT, "-showResult")
	verboseG = tk.IfSwitchExistsWhole(argsT, "-verbose")

	for _, scriptT := range scriptsT {
		var fcT string

		if ifExampleT {
			if !tk.EndsWith(scriptT, ".gt") {
				scriptT += ".gt"
			}
			fcT = tk.DownloadPageUTF8("https://gitee.com/topxeq/gotx/raw/master/scripts/"+scriptT, nil, "", 30)
		} else if ifRemoteT {
			fcT = tk.DownloadPageUTF8(scriptT, nil, "", 30)
		} else if ifCloudT {
			if !tk.EndsWith(scriptT, ".gt") {
				scriptT += ".gt"
			}
			fcT = tk.DownloadPageUTF8("http://scripts.frenchfriend.net/xaf/scripts/"+scriptT, nil, "", 30)
		} else if ifGoPathT {
			if !tk.EndsWith(scriptT, ".gt") {
				scriptT += ".gt"
			}

			fcT = tk.LoadStringFromFile(filepath.Join(tk.GetEnv("GOPATH"), "src", "github.com", "topxeq", "gotx", "scripts", scriptT))
		} else {
			fcT = tk.LoadStringFromFile(scriptT)
		}

		if tk.IsErrorString(fcT) {
			tk.Pl("failed to load script from %v: %v", scriptT, tk.GetErrorString(fcT))

			continue
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
}

func test() {
	// return

	// p, _ := plot.New()

	// p.Title.Text = "a"

	// tk.Pl("p: %#v", p)

	// typeT := reflect.TypeOf(p)

	// m := 1
	// kind := 2
	// name := "aa"

	// fmt.Printf("1m: %#v, obj: %#v, kind: %v, %v, Name: %v\n", m, typeT, kind, m, name)
	// lenT := typeT.NumMethod()

	// fmt.Printf("typeT: %#v, methodNum: %#v\n", typeT, lenT)
	// for i := 0; i < lenT; i++ {
	// 	fmt.Printf("m %v: %#v, method: %#v\n", i, typeT.Method(i), typeT.Method(i).Name)

	// }

}
