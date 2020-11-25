package main

import "github.com/topxeq/tk"

// convert Gotx file to Go source file.
func main() {
	fileT := tk.Trim(tk.GetSwitchWithDefaultValue(tk.GetOSArgsShort(), "-file=", ""))

	if fileT == "" {
		tk.Pl("not enough parameters")
		tk.Exit()
	}

	fcT := tk.LoadStringFromFile(fileT)

	if tk.IsErrStr(fcT) {
		tk.Pl("failed to load file content: %v", tk.GetErrStr(fcT))
		tk.Exit()
	}

	contentT := tk.RegReplaceX(fcT, `(\s+)(pl\()`, "${1}tk.Pl(")
	contentT = tk.RegReplaceX(contentT, `((\s+)|(\())(isErrStr\()`, "${1}tk.IsErrStr(")
	contentT = tk.RegReplaceX(contentT, `((\s+)|(\())(getSwitch\()`, "${1}tk.GetSwitchWithDefaultValue(")
	contentT = tk.RegReplaceX(contentT, `((\s+)|(\())(trim\()`, "${1}tk.Trim(")
	contentT = tk.RegReplaceX(contentT, `((\s+)|(\())(getErrStr\()`, "${1}tk.GetErrStr(")
	contentT = tk.RegReplaceX(contentT, `(\s+)(exit\()`, "${1}tk.Exit(")
	contentT = tk.RegReplaceX(contentT, `tk.GetOSArgsShort()`, "tk.GetOSArgsShort()")

	rs := tk.SaveStringToFile(contentT, fileT+".go")

	if tk.IsErrStr(rs) {
		tk.Pl("failed to save file content: %v", tk.GetErrStr(rs))
		tk.Exit()
	}

	tk.Pl("Done.")
}
