package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"reflect"
	"unsafe"
)

func init() {
	GotxSymbols["github.com/ying32/govcl/vcl"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ActionFromInst":                      reflect.ValueOf(vcl.ActionFromInst),
		"ActionFromObj":                       reflect.ValueOf(vcl.ActionFromObj),
		"ActionFromUnsafePointer":             reflect.ValueOf(vcl.ActionFromUnsafePointer),
		"ActionListFromInst":                  reflect.ValueOf(vcl.ActionListFromInst),
		"ActionListFromObj":                   reflect.ValueOf(vcl.ActionListFromObj),
		"ActionListFromUnsafePointer":         reflect.ValueOf(vcl.ActionListFromUnsafePointer),
		"Application":                         reflect.ValueOf(&vcl.Application).Elem(),
		"ApplicationFromInst":                 reflect.ValueOf(vcl.ApplicationFromInst),
		"ApplicationFromObj":                  reflect.ValueOf(vcl.ApplicationFromObj),
		"ApplicationFromUnsafePointer":        reflect.ValueOf(vcl.ApplicationFromUnsafePointer),
		"AsAction":                            reflect.ValueOf(vcl.AsAction),
		"AsActionList":                        reflect.ValueOf(vcl.AsActionList),
		"AsApplication":                       reflect.ValueOf(vcl.AsApplication),
		"AsBevel":                             reflect.ValueOf(vcl.AsBevel),
		"AsBitBtn":                            reflect.ValueOf(vcl.AsBitBtn),
		"AsBitmap":                            reflect.ValueOf(vcl.AsBitmap),
		"AsBoundLabel":                        reflect.ValueOf(vcl.AsBoundLabel),
		"AsBrush":                             reflect.ValueOf(vcl.AsBrush),
		"AsButton":                            reflect.ValueOf(vcl.AsButton),
		"AsCanvas":                            reflect.ValueOf(vcl.AsCanvas),
		"AsCheckBox":                          reflect.ValueOf(vcl.AsCheckBox),
		"AsCheckListBox":                      reflect.ValueOf(vcl.AsCheckListBox),
		"AsClipboard":                         reflect.ValueOf(vcl.AsClipboard),
		"AsCollection":                        reflect.ValueOf(vcl.AsCollection),
		"AsCollectionItem":                    reflect.ValueOf(vcl.AsCollectionItem),
		"AsColorBox":                          reflect.ValueOf(vcl.AsColorBox),
		"AsColorDialog":                       reflect.ValueOf(vcl.AsColorDialog),
		"AsColorListBox":                      reflect.ValueOf(vcl.AsColorListBox),
		"AsComboBox":                          reflect.ValueOf(vcl.AsComboBox),
		"AsComboBoxEx":                        reflect.ValueOf(vcl.AsComboBoxEx),
		"AsComboExItem":                       reflect.ValueOf(vcl.AsComboExItem),
		"AsComboExItems":                      reflect.ValueOf(vcl.AsComboExItems),
		"AsComponent":                         reflect.ValueOf(vcl.AsComponent),
		"AsControl":                           reflect.ValueOf(vcl.AsControl),
		"AsControlScrollBar":                  reflect.ValueOf(vcl.AsControlScrollBar),
		"AsCoolBand":                          reflect.ValueOf(vcl.AsCoolBand),
		"AsCoolBands":                         reflect.ValueOf(vcl.AsCoolBands),
		"AsCoolBar":                           reflect.ValueOf(vcl.AsCoolBar),
		"AsDateTimePicker":                    reflect.ValueOf(vcl.AsDateTimePicker),
		"AsDragDockObject":                    reflect.ValueOf(vcl.AsDragDockObject),
		"AsDragObject":                        reflect.ValueOf(vcl.AsDragObject),
		"AsDrawGrid":                          reflect.ValueOf(vcl.AsDrawGrid),
		"AsEdit":                              reflect.ValueOf(vcl.AsEdit),
		"AsException":                         reflect.ValueOf(vcl.AsException),
		"AsFindDialog":                        reflect.ValueOf(vcl.AsFindDialog),
		"AsFlowPanel":                         reflect.ValueOf(vcl.AsFlowPanel),
		"AsFont":                              reflect.ValueOf(vcl.AsFont),
		"AsFontDialog":                        reflect.ValueOf(vcl.AsFontDialog),
		"AsForm":                              reflect.ValueOf(vcl.AsForm),
		"AsFrame":                             reflect.ValueOf(vcl.AsFrame),
		"AsGIFImage":                          reflect.ValueOf(vcl.AsGIFImage),
		"AsGauge":                             reflect.ValueOf(vcl.AsGauge),
		"AsGraphic":                           reflect.ValueOf(vcl.AsGraphic),
		"AsGroupBox":                          reflect.ValueOf(vcl.AsGroupBox),
		"AsHeaderControl":                     reflect.ValueOf(vcl.AsHeaderControl),
		"AsHeaderSection":                     reflect.ValueOf(vcl.AsHeaderSection),
		"AsHeaderSections":                    reflect.ValueOf(vcl.AsHeaderSections),
		"AsIcon":                              reflect.ValueOf(vcl.AsIcon),
		"AsIconOptions":                       reflect.ValueOf(vcl.AsIconOptions),
		"AsImage":                             reflect.ValueOf(vcl.AsImage),
		"AsImageButton":                       reflect.ValueOf(vcl.AsImageButton),
		"AsImageList":                         reflect.ValueOf(vcl.AsImageList),
		"AsIniFile":                           reflect.ValueOf(vcl.AsIniFile),
		"AsJPEGImage":                         reflect.ValueOf(vcl.AsJPEGImage),
		"AsLabel":                             reflect.ValueOf(vcl.AsLabel),
		"AsLabeledEdit":                       reflect.ValueOf(vcl.AsLabeledEdit),
		"AsLinkLabel":                         reflect.ValueOf(vcl.AsLinkLabel),
		"AsList":                              reflect.ValueOf(vcl.AsList),
		"AsListBox":                           reflect.ValueOf(vcl.AsListBox),
		"AsListColumn":                        reflect.ValueOf(vcl.AsListColumn),
		"AsListColumns":                       reflect.ValueOf(vcl.AsListColumns),
		"AsListItem":                          reflect.ValueOf(vcl.AsListItem),
		"AsListItems":                         reflect.ValueOf(vcl.AsListItems),
		"AsListView":                          reflect.ValueOf(vcl.AsListView),
		"AsMainMenu":                          reflect.ValueOf(vcl.AsMainMenu),
		"AsMargins":                           reflect.ValueOf(vcl.AsMargins),
		"AsMaskEdit":                          reflect.ValueOf(vcl.AsMaskEdit),
		"AsMemo":                              reflect.ValueOf(vcl.AsMemo),
		"AsMemoryStream":                      reflect.ValueOf(vcl.AsMemoryStream),
		"AsMenuItem":                          reflect.ValueOf(vcl.AsMenuItem),
		"AsMiniWebview":                       reflect.ValueOf(vcl.AsMiniWebview),
		"AsMonitor":                           reflect.ValueOf(vcl.AsMonitor),
		"AsMonthCalendar":                     reflect.ValueOf(vcl.AsMonthCalendar),
		"AsMouse":                             reflect.ValueOf(vcl.AsMouse),
		"AsObject":                            reflect.ValueOf(vcl.AsObject),
		"AsOpenDialog":                        reflect.ValueOf(vcl.AsOpenDialog),
		"AsOpenPictureDialog":                 reflect.ValueOf(vcl.AsOpenPictureDialog),
		"AsPageControl":                       reflect.ValueOf(vcl.AsPageControl),
		"AsPageSetupDialog":                   reflect.ValueOf(vcl.AsPageSetupDialog),
		"AsPaintBox":                          reflect.ValueOf(vcl.AsPaintBox),
		"AsPanel":                             reflect.ValueOf(vcl.AsPanel),
		"AsPen":                               reflect.ValueOf(vcl.AsPen),
		"AsPicture":                           reflect.ValueOf(vcl.AsPicture),
		"AsPngImage":                          reflect.ValueOf(vcl.AsPngImage),
		"AsPopupMenu":                         reflect.ValueOf(vcl.AsPopupMenu),
		"AsPrintDialog":                       reflect.ValueOf(vcl.AsPrintDialog),
		"AsPrinter":                           reflect.ValueOf(vcl.AsPrinter),
		"AsPrinterSetupDialog":                reflect.ValueOf(vcl.AsPrinterSetupDialog),
		"AsProgressBar":                       reflect.ValueOf(vcl.AsProgressBar),
		"AsRadioButton":                       reflect.ValueOf(vcl.AsRadioButton),
		"AsRadioGroup":                        reflect.ValueOf(vcl.AsRadioGroup),
		"AsRegistry":                          reflect.ValueOf(vcl.AsRegistry),
		"AsReplaceDialog":                     reflect.ValueOf(vcl.AsReplaceDialog),
		"AsSaveDialog":                        reflect.ValueOf(vcl.AsSaveDialog),
		"AsSavePictureDialog":                 reflect.ValueOf(vcl.AsSavePictureDialog),
		"AsScreen":                            reflect.ValueOf(vcl.AsScreen),
		"AsScrollBar":                         reflect.ValueOf(vcl.AsScrollBar),
		"AsScrollBox":                         reflect.ValueOf(vcl.AsScrollBox),
		"AsShape":                             reflect.ValueOf(vcl.AsShape),
		"AsSizeConstraints":                   reflect.ValueOf(vcl.AsSizeConstraints),
		"AsSpeedButton":                       reflect.ValueOf(vcl.AsSpeedButton),
		"AsSpinEdit":                          reflect.ValueOf(vcl.AsSpinEdit),
		"AsSplitter":                          reflect.ValueOf(vcl.AsSplitter),
		"AsStaticText":                        reflect.ValueOf(vcl.AsStaticText),
		"AsStatusBar":                         reflect.ValueOf(vcl.AsStatusBar),
		"AsStatusPanel":                       reflect.ValueOf(vcl.AsStatusPanel),
		"AsStatusPanels":                      reflect.ValueOf(vcl.AsStatusPanels),
		"AsStringGrid":                        reflect.ValueOf(vcl.AsStringGrid),
		"AsStringList":                        reflect.ValueOf(vcl.AsStringList),
		"AsStrings":                           reflect.ValueOf(vcl.AsStrings),
		"AsTabSheet":                          reflect.ValueOf(vcl.AsTabSheet),
		"AsTaskDialog":                        reflect.ValueOf(vcl.AsTaskDialog),
		"AsTaskDialogBaseButtonItem":          reflect.ValueOf(vcl.AsTaskDialogBaseButtonItem),
		"AsTaskDialogButtonItem":              reflect.ValueOf(vcl.AsTaskDialogButtonItem),
		"AsTaskDialogButtons":                 reflect.ValueOf(vcl.AsTaskDialogButtons),
		"AsTaskDialogRadioButtonItem":         reflect.ValueOf(vcl.AsTaskDialogRadioButtonItem),
		"AsTimer":                             reflect.ValueOf(vcl.AsTimer),
		"AsToolBar":                           reflect.ValueOf(vcl.AsToolBar),
		"AsToolButton":                        reflect.ValueOf(vcl.AsToolButton),
		"AsTrackBar":                          reflect.ValueOf(vcl.AsTrackBar),
		"AsTrayIcon":                          reflect.ValueOf(vcl.AsTrayIcon),
		"AsTreeNode":                          reflect.ValueOf(vcl.AsTreeNode),
		"AsTreeNodes":                         reflect.ValueOf(vcl.AsTreeNodes),
		"AsTreeView":                          reflect.ValueOf(vcl.AsTreeView),
		"AsUpDown":                            reflect.ValueOf(vcl.AsUpDown),
		"AsValueListEditor":                   reflect.ValueOf(vcl.AsValueListEditor),
		"AsWinControl":                        reflect.ValueOf(vcl.AsWinControl),
		"AsXButton":                           reflect.ValueOf(vcl.AsXButton),
		"BevelFromInst":                       reflect.ValueOf(vcl.BevelFromInst),
		"BevelFromObj":                        reflect.ValueOf(vcl.BevelFromObj),
		"BevelFromUnsafePointer":              reflect.ValueOf(vcl.BevelFromUnsafePointer),
		"BitBtnFromInst":                      reflect.ValueOf(vcl.BitBtnFromInst),
		"BitBtnFromObj":                       reflect.ValueOf(vcl.BitBtnFromObj),
		"BitBtnFromUnsafePointer":             reflect.ValueOf(vcl.BitBtnFromUnsafePointer),
		"BitmapFromInst":                      reflect.ValueOf(vcl.BitmapFromInst),
		"BitmapFromObj":                       reflect.ValueOf(vcl.BitmapFromObj),
		"BitmapFromUnsafePointer":             reflect.ValueOf(vcl.BitmapFromUnsafePointer),
		"BoundLabelFromInst":                  reflect.ValueOf(vcl.BoundLabelFromInst),
		"BoundLabelFromObj":                   reflect.ValueOf(vcl.BoundLabelFromObj),
		"BoundLabelFromUnsafePointer":         reflect.ValueOf(vcl.BoundLabelFromUnsafePointer),
		"BrushFromInst":                       reflect.ValueOf(vcl.BrushFromInst),
		"BrushFromObj":                        reflect.ValueOf(vcl.BrushFromObj),
		"BrushFromUnsafePointer":              reflect.ValueOf(vcl.BrushFromUnsafePointer),
		"ButtonFromInst":                      reflect.ValueOf(vcl.ButtonFromInst),
		"ButtonFromObj":                       reflect.ValueOf(vcl.ButtonFromObj),
		"ButtonFromUnsafePointer":             reflect.ValueOf(vcl.ButtonFromUnsafePointer),
		"CanvasFromInst":                      reflect.ValueOf(vcl.CanvasFromInst),
		"CanvasFromObj":                       reflect.ValueOf(vcl.CanvasFromObj),
		"CanvasFromUnsafePointer":             reflect.ValueOf(vcl.CanvasFromUnsafePointer),
		"CheckBoxFromInst":                    reflect.ValueOf(vcl.CheckBoxFromInst),
		"CheckBoxFromObj":                     reflect.ValueOf(vcl.CheckBoxFromObj),
		"CheckBoxFromUnsafePointer":           reflect.ValueOf(vcl.CheckBoxFromUnsafePointer),
		"CheckListBoxFromInst":                reflect.ValueOf(vcl.CheckListBoxFromInst),
		"CheckListBoxFromObj":                 reflect.ValueOf(vcl.CheckListBoxFromObj),
		"CheckListBoxFromUnsafePointer":       reflect.ValueOf(vcl.CheckListBoxFromUnsafePointer),
		"CheckPtr":                            reflect.ValueOf(vcl.CheckPtr),
		"Clipboard":                           reflect.ValueOf(&vcl.Clipboard).Elem(),
		"ClipboardFromInst":                   reflect.ValueOf(vcl.ClipboardFromInst),
		"ClipboardFromObj":                    reflect.ValueOf(vcl.ClipboardFromObj),
		"ClipboardFromUnsafePointer":          reflect.ValueOf(vcl.ClipboardFromUnsafePointer),
		"CollectionFromInst":                  reflect.ValueOf(vcl.CollectionFromInst),
		"CollectionFromObj":                   reflect.ValueOf(vcl.CollectionFromObj),
		"CollectionFromUnsafePointer":         reflect.ValueOf(vcl.CollectionFromUnsafePointer),
		"CollectionItemFromInst":              reflect.ValueOf(vcl.CollectionItemFromInst),
		"CollectionItemFromObj":               reflect.ValueOf(vcl.CollectionItemFromObj),
		"CollectionItemFromUnsafePointer":     reflect.ValueOf(vcl.CollectionItemFromUnsafePointer),
		"ColorBoxFromInst":                    reflect.ValueOf(vcl.ColorBoxFromInst),
		"ColorBoxFromObj":                     reflect.ValueOf(vcl.ColorBoxFromObj),
		"ColorBoxFromUnsafePointer":           reflect.ValueOf(vcl.ColorBoxFromUnsafePointer),
		"ColorDialogFromInst":                 reflect.ValueOf(vcl.ColorDialogFromInst),
		"ColorDialogFromObj":                  reflect.ValueOf(vcl.ColorDialogFromObj),
		"ColorDialogFromUnsafePointer":        reflect.ValueOf(vcl.ColorDialogFromUnsafePointer),
		"ColorListBoxFromInst":                reflect.ValueOf(vcl.ColorListBoxFromInst),
		"ColorListBoxFromObj":                 reflect.ValueOf(vcl.ColorListBoxFromObj),
		"ColorListBoxFromUnsafePointer":       reflect.ValueOf(vcl.ColorListBoxFromUnsafePointer),
		"ComboBoxExFromInst":                  reflect.ValueOf(vcl.ComboBoxExFromInst),
		"ComboBoxExFromObj":                   reflect.ValueOf(vcl.ComboBoxExFromObj),
		"ComboBoxExFromUnsafePointer":         reflect.ValueOf(vcl.ComboBoxExFromUnsafePointer),
		"ComboBoxFromInst":                    reflect.ValueOf(vcl.ComboBoxFromInst),
		"ComboBoxFromObj":                     reflect.ValueOf(vcl.ComboBoxFromObj),
		"ComboBoxFromUnsafePointer":           reflect.ValueOf(vcl.ComboBoxFromUnsafePointer),
		"ComboExItemFromInst":                 reflect.ValueOf(vcl.ComboExItemFromInst),
		"ComboExItemFromObj":                  reflect.ValueOf(vcl.ComboExItemFromObj),
		"ComboExItemFromUnsafePointer":        reflect.ValueOf(vcl.ComboExItemFromUnsafePointer),
		"ComboExItemsFromInst":                reflect.ValueOf(vcl.ComboExItemsFromInst),
		"ComboExItemsFromObj":                 reflect.ValueOf(vcl.ComboExItemsFromObj),
		"ComboExItemsFromUnsafePointer":       reflect.ValueOf(vcl.ComboExItemsFromUnsafePointer),
		"ComponentFromInst":                   reflect.ValueOf(vcl.ComponentFromInst),
		"ComponentFromObj":                    reflect.ValueOf(vcl.ComponentFromObj),
		"ComponentFromUnsafePointer":          reflect.ValueOf(vcl.ComponentFromUnsafePointer),
		"ControlFromInst":                     reflect.ValueOf(vcl.ControlFromInst),
		"ControlFromObj":                      reflect.ValueOf(vcl.ControlFromObj),
		"ControlFromUnsafePointer":            reflect.ValueOf(vcl.ControlFromUnsafePointer),
		"ControlScrollBarFromInst":            reflect.ValueOf(vcl.ControlScrollBarFromInst),
		"ControlScrollBarFromObj":             reflect.ValueOf(vcl.ControlScrollBarFromObj),
		"ControlScrollBarFromUnsafePointer":   reflect.ValueOf(vcl.ControlScrollBarFromUnsafePointer),
		"CoolBandFromInst":                    reflect.ValueOf(vcl.CoolBandFromInst),
		"CoolBandFromObj":                     reflect.ValueOf(vcl.CoolBandFromObj),
		"CoolBandFromUnsafePointer":           reflect.ValueOf(vcl.CoolBandFromUnsafePointer),
		"CoolBandsFromInst":                   reflect.ValueOf(vcl.CoolBandsFromInst),
		"CoolBandsFromObj":                    reflect.ValueOf(vcl.CoolBandsFromObj),
		"CoolBandsFromUnsafePointer":          reflect.ValueOf(vcl.CoolBandsFromUnsafePointer),
		"CoolBarFromInst":                     reflect.ValueOf(vcl.CoolBarFromInst),
		"CoolBarFromObj":                      reflect.ValueOf(vcl.CoolBarFromObj),
		"CoolBarFromUnsafePointer":            reflect.ValueOf(vcl.CoolBarFromUnsafePointer),
		"CreateResForm":                       reflect.ValueOf(vcl.CreateResForm),
		"CreateResFrame":                      reflect.ValueOf(vcl.CreateResFrame),
		"DateTimePickerFromInst":              reflect.ValueOf(vcl.DateTimePickerFromInst),
		"DateTimePickerFromObj":               reflect.ValueOf(vcl.DateTimePickerFromObj),
		"DateTimePickerFromUnsafePointer":     reflect.ValueOf(vcl.DateTimePickerFromUnsafePointer),
		"DragDockObjectFromInst":              reflect.ValueOf(vcl.DragDockObjectFromInst),
		"DragDockObjectFromObj":               reflect.ValueOf(vcl.DragDockObjectFromObj),
		"DragDockObjectFromUnsafePointer":     reflect.ValueOf(vcl.DragDockObjectFromUnsafePointer),
		"DragObjectFromInst":                  reflect.ValueOf(vcl.DragObjectFromInst),
		"DragObjectFromObj":                   reflect.ValueOf(vcl.DragObjectFromObj),
		"DragObjectFromUnsafePointer":         reflect.ValueOf(vcl.DragObjectFromUnsafePointer),
		"DrawGridFromInst":                    reflect.ValueOf(vcl.DrawGridFromInst),
		"DrawGridFromObj":                     reflect.ValueOf(vcl.DrawGridFromObj),
		"DrawGridFromUnsafePointer":           reflect.ValueOf(vcl.DrawGridFromUnsafePointer),
		"EditFromInst":                        reflect.ValueOf(vcl.EditFromInst),
		"EditFromObj":                         reflect.ValueOf(vcl.EditFromObj),
		"EditFromUnsafePointer":               reflect.ValueOf(vcl.EditFromUnsafePointer),
		"EqualsObject":                        reflect.ValueOf(vcl.EqualsObject),
		"ExceptionClass":                      reflect.ValueOf(vcl.ExceptionClass),
		"ExceptionFromInst":                   reflect.ValueOf(vcl.ExceptionFromInst),
		"ExceptionFromObj":                    reflect.ValueOf(vcl.ExceptionFromObj),
		"ExceptionFromUnsafePointer":          reflect.ValueOf(vcl.ExceptionFromUnsafePointer),
		"FindDialogFromInst":                  reflect.ValueOf(vcl.FindDialogFromInst),
		"FindDialogFromObj":                   reflect.ValueOf(vcl.FindDialogFromObj),
		"FindDialogFromUnsafePointer":         reflect.ValueOf(vcl.FindDialogFromUnsafePointer),
		"FlowPanelFromInst":                   reflect.ValueOf(vcl.FlowPanelFromInst),
		"FlowPanelFromObj":                    reflect.ValueOf(vcl.FlowPanelFromObj),
		"FlowPanelFromUnsafePointer":          reflect.ValueOf(vcl.FlowPanelFromUnsafePointer),
		"FontDialogFromInst":                  reflect.ValueOf(vcl.FontDialogFromInst),
		"FontDialogFromObj":                   reflect.ValueOf(vcl.FontDialogFromObj),
		"FontDialogFromUnsafePointer":         reflect.ValueOf(vcl.FontDialogFromUnsafePointer),
		"FontFromInst":                        reflect.ValueOf(vcl.FontFromInst),
		"FontFromObj":                         reflect.ValueOf(vcl.FontFromObj),
		"FontFromUnsafePointer":               reflect.ValueOf(vcl.FontFromUnsafePointer),
		"FormFromInst":                        reflect.ValueOf(vcl.FormFromInst),
		"FormFromObj":                         reflect.ValueOf(vcl.FormFromObj),
		"FormFromUnsafePointer":               reflect.ValueOf(vcl.FormFromUnsafePointer),
		"FrameFromInst":                       reflect.ValueOf(vcl.FrameFromInst),
		"FrameFromObj":                        reflect.ValueOf(vcl.FrameFromObj),
		"FrameFromUnsafePointer":              reflect.ValueOf(vcl.FrameFromUnsafePointer),
		"GIFImageFromInst":                    reflect.ValueOf(vcl.GIFImageFromInst),
		"GIFImageFromObj":                     reflect.ValueOf(vcl.GIFImageFromObj),
		"GIFImageFromUnsafePointer":           reflect.ValueOf(vcl.GIFImageFromUnsafePointer),
		"GaugeFromInst":                       reflect.ValueOf(vcl.GaugeFromInst),
		"GaugeFromObj":                        reflect.ValueOf(vcl.GaugeFromObj),
		"GaugeFromUnsafePointer":              reflect.ValueOf(vcl.GaugeFromUnsafePointer),
		"GraphicFromInst":                     reflect.ValueOf(vcl.GraphicFromInst),
		"GraphicFromObj":                      reflect.ValueOf(vcl.GraphicFromObj),
		"GraphicFromUnsafePointer":            reflect.ValueOf(vcl.GraphicFromUnsafePointer),
		"GroupBoxFromInst":                    reflect.ValueOf(vcl.GroupBoxFromInst),
		"GroupBoxFromObj":                     reflect.ValueOf(vcl.GroupBoxFromObj),
		"GroupBoxFromUnsafePointer":           reflect.ValueOf(vcl.GroupBoxFromUnsafePointer),
		"HandleToPlatformHandle":              reflect.ValueOf(vcl.HandleToPlatformHandle),
		"HeaderControlFromInst":               reflect.ValueOf(vcl.HeaderControlFromInst),
		"HeaderControlFromObj":                reflect.ValueOf(vcl.HeaderControlFromObj),
		"HeaderControlFromUnsafePointer":      reflect.ValueOf(vcl.HeaderControlFromUnsafePointer),
		"HeaderSectionFromInst":               reflect.ValueOf(vcl.HeaderSectionFromInst),
		"HeaderSectionFromObj":                reflect.ValueOf(vcl.HeaderSectionFromObj),
		"HeaderSectionFromUnsafePointer":      reflect.ValueOf(vcl.HeaderSectionFromUnsafePointer),
		"HeaderSectionsFromInst":              reflect.ValueOf(vcl.HeaderSectionsFromInst),
		"HeaderSectionsFromObj":               reflect.ValueOf(vcl.HeaderSectionsFromObj),
		"HeaderSectionsFromUnsafePointer":     reflect.ValueOf(vcl.HeaderSectionsFromUnsafePointer),
		"IconFromInst":                        reflect.ValueOf(vcl.IconFromInst),
		"IconFromObj":                         reflect.ValueOf(vcl.IconFromObj),
		"IconFromUnsafePointer":               reflect.ValueOf(vcl.IconFromUnsafePointer),
		"IconOptionsFromInst":                 reflect.ValueOf(vcl.IconOptionsFromInst),
		"IconOptionsFromObj":                  reflect.ValueOf(vcl.IconOptionsFromObj),
		"IconOptionsFromUnsafePointer":        reflect.ValueOf(vcl.IconOptionsFromUnsafePointer),
		"ImageButtonFromInst":                 reflect.ValueOf(vcl.ImageButtonFromInst),
		"ImageButtonFromObj":                  reflect.ValueOf(vcl.ImageButtonFromObj),
		"ImageButtonFromUnsafePointer":        reflect.ValueOf(vcl.ImageButtonFromUnsafePointer),
		"ImageFromInst":                       reflect.ValueOf(vcl.ImageFromInst),
		"ImageFromObj":                        reflect.ValueOf(vcl.ImageFromObj),
		"ImageFromUnsafePointer":              reflect.ValueOf(vcl.ImageFromUnsafePointer),
		"ImageListFromInst":                   reflect.ValueOf(vcl.ImageListFromInst),
		"ImageListFromObj":                    reflect.ValueOf(vcl.ImageListFromObj),
		"ImageListFromUnsafePointer":          reflect.ValueOf(vcl.ImageListFromUnsafePointer),
		"IniFileFromInst":                     reflect.ValueOf(vcl.IniFileFromInst),
		"IniFileFromObj":                      reflect.ValueOf(vcl.IniFileFromObj),
		"IniFileFromUnsafePointer":            reflect.ValueOf(vcl.IniFileFromUnsafePointer),
		"InputBox":                            reflect.ValueOf(vcl.InputBox),
		"InputQuery":                          reflect.ValueOf(vcl.InputQuery),
		"JPEGImageFromInst":                   reflect.ValueOf(vcl.JPEGImageFromInst),
		"JPEGImageFromObj":                    reflect.ValueOf(vcl.JPEGImageFromObj),
		"JPEGImageFromUnsafePointer":          reflect.ValueOf(vcl.JPEGImageFromUnsafePointer),
		"LabelFromInst":                       reflect.ValueOf(vcl.LabelFromInst),
		"LabelFromObj":                        reflect.ValueOf(vcl.LabelFromObj),
		"LabelFromUnsafePointer":              reflect.ValueOf(vcl.LabelFromUnsafePointer),
		"LabeledEditFromInst":                 reflect.ValueOf(vcl.LabeledEditFromInst),
		"LabeledEditFromObj":                  reflect.ValueOf(vcl.LabeledEditFromObj),
		"LabeledEditFromUnsafePointer":        reflect.ValueOf(vcl.LabeledEditFromUnsafePointer),
		"LclLoaded":                           reflect.ValueOf(vcl.LclLoaded),
		"LinkLabelFromInst":                   reflect.ValueOf(vcl.LinkLabelFromInst),
		"LinkLabelFromObj":                    reflect.ValueOf(vcl.LinkLabelFromObj),
		"LinkLabelFromUnsafePointer":          reflect.ValueOf(vcl.LinkLabelFromUnsafePointer),
		"ListBoxFromInst":                     reflect.ValueOf(vcl.ListBoxFromInst),
		"ListBoxFromObj":                      reflect.ValueOf(vcl.ListBoxFromObj),
		"ListBoxFromUnsafePointer":            reflect.ValueOf(vcl.ListBoxFromUnsafePointer),
		"ListColumnFromInst":                  reflect.ValueOf(vcl.ListColumnFromInst),
		"ListColumnFromObj":                   reflect.ValueOf(vcl.ListColumnFromObj),
		"ListColumnFromUnsafePointer":         reflect.ValueOf(vcl.ListColumnFromUnsafePointer),
		"ListColumnsFromInst":                 reflect.ValueOf(vcl.ListColumnsFromInst),
		"ListColumnsFromObj":                  reflect.ValueOf(vcl.ListColumnsFromObj),
		"ListColumnsFromUnsafePointer":        reflect.ValueOf(vcl.ListColumnsFromUnsafePointer),
		"ListFromInst":                        reflect.ValueOf(vcl.ListFromInst),
		"ListFromObj":                         reflect.ValueOf(vcl.ListFromObj),
		"ListFromUnsafePointer":               reflect.ValueOf(vcl.ListFromUnsafePointer),
		"ListItemFromInst":                    reflect.ValueOf(vcl.ListItemFromInst),
		"ListItemFromObj":                     reflect.ValueOf(vcl.ListItemFromObj),
		"ListItemFromUnsafePointer":           reflect.ValueOf(vcl.ListItemFromUnsafePointer),
		"ListItemsFromInst":                   reflect.ValueOf(vcl.ListItemsFromInst),
		"ListItemsFromObj":                    reflect.ValueOf(vcl.ListItemsFromObj),
		"ListItemsFromUnsafePointer":          reflect.ValueOf(vcl.ListItemsFromUnsafePointer),
		"ListViewFromInst":                    reflect.ValueOf(vcl.ListViewFromInst),
		"ListViewFromObj":                     reflect.ValueOf(vcl.ListViewFromObj),
		"ListViewFromUnsafePointer":           reflect.ValueOf(vcl.ListViewFromUnsafePointer),
		"MainMenuFromInst":                    reflect.ValueOf(vcl.MainMenuFromInst),
		"MainMenuFromObj":                     reflect.ValueOf(vcl.MainMenuFromObj),
		"MainMenuFromUnsafePointer":           reflect.ValueOf(vcl.MainMenuFromUnsafePointer),
		"MarginsFromInst":                     reflect.ValueOf(vcl.MarginsFromInst),
		"MarginsFromObj":                      reflect.ValueOf(vcl.MarginsFromObj),
		"MarginsFromUnsafePointer":            reflect.ValueOf(vcl.MarginsFromUnsafePointer),
		"MaskEditFromInst":                    reflect.ValueOf(vcl.MaskEditFromInst),
		"MaskEditFromObj":                     reflect.ValueOf(vcl.MaskEditFromObj),
		"MaskEditFromUnsafePointer":           reflect.ValueOf(vcl.MaskEditFromUnsafePointer),
		"MemoFromInst":                        reflect.ValueOf(vcl.MemoFromInst),
		"MemoFromObj":                         reflect.ValueOf(vcl.MemoFromObj),
		"MemoFromUnsafePointer":               reflect.ValueOf(vcl.MemoFromUnsafePointer),
		"MemoryStreamFromInst":                reflect.ValueOf(vcl.MemoryStreamFromInst),
		"MemoryStreamFromObj":                 reflect.ValueOf(vcl.MemoryStreamFromObj),
		"MemoryStreamFromUnsafePointer":       reflect.ValueOf(vcl.MemoryStreamFromUnsafePointer),
		"MenuItemFromInst":                    reflect.ValueOf(vcl.MenuItemFromInst),
		"MenuItemFromObj":                     reflect.ValueOf(vcl.MenuItemFromObj),
		"MenuItemFromUnsafePointer":           reflect.ValueOf(vcl.MenuItemFromUnsafePointer),
		"MessageDlg":                          reflect.ValueOf(vcl.MessageDlg),
		"MiniWebviewFromInst":                 reflect.ValueOf(vcl.MiniWebviewFromInst),
		"MiniWebviewFromObj":                  reflect.ValueOf(vcl.MiniWebviewFromObj),
		"MiniWebviewFromUnsafePointer":        reflect.ValueOf(vcl.MiniWebviewFromUnsafePointer),
		"MonitorFromInst":                     reflect.ValueOf(vcl.MonitorFromInst),
		"MonitorFromObj":                      reflect.ValueOf(vcl.MonitorFromObj),
		"MonitorFromUnsafePointer":            reflect.ValueOf(vcl.MonitorFromUnsafePointer),
		"MonthCalendarFromInst":               reflect.ValueOf(vcl.MonthCalendarFromInst),
		"MonthCalendarFromObj":                reflect.ValueOf(vcl.MonthCalendarFromObj),
		"MonthCalendarFromUnsafePointer":      reflect.ValueOf(vcl.MonthCalendarFromUnsafePointer),
		"Mouse":                               reflect.ValueOf(&vcl.Mouse).Elem(),
		"MouseFromInst":                       reflect.ValueOf(vcl.MouseFromInst),
		"MouseFromObj":                        reflect.ValueOf(vcl.MouseFromObj),
		"MouseFromUnsafePointer":              reflect.ValueOf(vcl.MouseFromUnsafePointer),
		"NewAction":                           reflect.ValueOf(vcl.NewAction),
		"NewActionList":                       reflect.ValueOf(vcl.NewActionList),
		"NewApplication":                      reflect.ValueOf(vcl.NewApplication),
		"NewBevel":                            reflect.ValueOf(vcl.NewBevel),
		"NewBitBtn":                           reflect.ValueOf(vcl.NewBitBtn),
		"NewBitmap":                           reflect.ValueOf(vcl.NewBitmap),
		"NewBoundLabel":                       reflect.ValueOf(vcl.NewBoundLabel),
		"NewBrush":                            reflect.ValueOf(vcl.NewBrush),
		"NewButton":                           reflect.ValueOf(vcl.NewButton),
		"NewCanvas":                           reflect.ValueOf(vcl.NewCanvas),
		"NewCheckBox":                         reflect.ValueOf(vcl.NewCheckBox),
		"NewCheckListBox":                     reflect.ValueOf(vcl.NewCheckListBox),
		"NewClipboard":                        reflect.ValueOf(vcl.NewClipboard),
		"NewCollection":                       reflect.ValueOf(vcl.NewCollection),
		"NewCollectionItem":                   reflect.ValueOf(vcl.NewCollectionItem),
		"NewColorBox":                         reflect.ValueOf(vcl.NewColorBox),
		"NewColorDialog":                      reflect.ValueOf(vcl.NewColorDialog),
		"NewColorListBox":                     reflect.ValueOf(vcl.NewColorListBox),
		"NewComboBox":                         reflect.ValueOf(vcl.NewComboBox),
		"NewComboBoxEx":                       reflect.ValueOf(vcl.NewComboBoxEx),
		"NewComponent":                        reflect.ValueOf(vcl.NewComponent),
		"NewControl":                          reflect.ValueOf(vcl.NewControl),
		"NewCoolBand":                         reflect.ValueOf(vcl.NewCoolBand),
		"NewCoolBands":                        reflect.ValueOf(vcl.NewCoolBands),
		"NewCoolBar":                          reflect.ValueOf(vcl.NewCoolBar),
		"NewDateTimePicker":                   reflect.ValueOf(vcl.NewDateTimePicker),
		"NewDragDockObject":                   reflect.ValueOf(vcl.NewDragDockObject),
		"NewDragObject":                       reflect.ValueOf(vcl.NewDragObject),
		"NewDrawGrid":                         reflect.ValueOf(vcl.NewDrawGrid),
		"NewEdit":                             reflect.ValueOf(vcl.NewEdit),
		"NewFindDialog":                       reflect.ValueOf(vcl.NewFindDialog),
		"NewFlowPanel":                        reflect.ValueOf(vcl.NewFlowPanel),
		"NewFont":                             reflect.ValueOf(vcl.NewFont),
		"NewFontDialog":                       reflect.ValueOf(vcl.NewFontDialog),
		"NewForm":                             reflect.ValueOf(vcl.NewForm),
		"NewFrame":                            reflect.ValueOf(vcl.NewFrame),
		"NewGIFImage":                         reflect.ValueOf(vcl.NewGIFImage),
		"NewGauge":                            reflect.ValueOf(vcl.NewGauge),
		"NewGraphic":                          reflect.ValueOf(vcl.NewGraphic),
		"NewGroupBox":                         reflect.ValueOf(vcl.NewGroupBox),
		"NewHeaderControl":                    reflect.ValueOf(vcl.NewHeaderControl),
		"NewHeaderSection":                    reflect.ValueOf(vcl.NewHeaderSection),
		"NewHeaderSections":                   reflect.ValueOf(vcl.NewHeaderSections),
		"NewIcon":                             reflect.ValueOf(vcl.NewIcon),
		"NewImage":                            reflect.ValueOf(vcl.NewImage),
		"NewImageButton":                      reflect.ValueOf(vcl.NewImageButton),
		"NewImageList":                        reflect.ValueOf(vcl.NewImageList),
		"NewIniFile":                          reflect.ValueOf(vcl.NewIniFile),
		"NewJPEGImage":                        reflect.ValueOf(vcl.NewJPEGImage),
		"NewLabel":                            reflect.ValueOf(vcl.NewLabel),
		"NewLabeledEdit":                      reflect.ValueOf(vcl.NewLabeledEdit),
		"NewLinkLabel":                        reflect.ValueOf(vcl.NewLinkLabel),
		"NewList":                             reflect.ValueOf(vcl.NewList),
		"NewListBox":                          reflect.ValueOf(vcl.NewListBox),
		"NewListColumn":                       reflect.ValueOf(vcl.NewListColumn),
		"NewListColumns":                      reflect.ValueOf(vcl.NewListColumns),
		"NewListItem":                         reflect.ValueOf(vcl.NewListItem),
		"NewListItems":                        reflect.ValueOf(vcl.NewListItems),
		"NewListView":                         reflect.ValueOf(vcl.NewListView),
		"NewMainMenu":                         reflect.ValueOf(vcl.NewMainMenu),
		"NewMaskEdit":                         reflect.ValueOf(vcl.NewMaskEdit),
		"NewMemo":                             reflect.ValueOf(vcl.NewMemo),
		"NewMemoryStream":                     reflect.ValueOf(vcl.NewMemoryStream),
		"NewMemoryStreamFromBytes":            reflect.ValueOf(vcl.NewMemoryStreamFromBytes),
		"NewMenuItem":                         reflect.ValueOf(vcl.NewMenuItem),
		"NewMiniWebview":                      reflect.ValueOf(vcl.NewMiniWebview),
		"NewMonitor":                          reflect.ValueOf(vcl.NewMonitor),
		"NewMonthCalendar":                    reflect.ValueOf(vcl.NewMonthCalendar),
		"NewMouse":                            reflect.ValueOf(vcl.NewMouse),
		"NewObject":                           reflect.ValueOf(vcl.NewObject),
		"NewOpenDialog":                       reflect.ValueOf(vcl.NewOpenDialog),
		"NewOpenPictureDialog":                reflect.ValueOf(vcl.NewOpenPictureDialog),
		"NewPageControl":                      reflect.ValueOf(vcl.NewPageControl),
		"NewPageSetupDialog":                  reflect.ValueOf(vcl.NewPageSetupDialog),
		"NewPaintBox":                         reflect.ValueOf(vcl.NewPaintBox),
		"NewPanel":                            reflect.ValueOf(vcl.NewPanel),
		"NewPen":                              reflect.ValueOf(vcl.NewPen),
		"NewPicture":                          reflect.ValueOf(vcl.NewPicture),
		"NewPngImage":                         reflect.ValueOf(vcl.NewPngImage),
		"NewPopupMenu":                        reflect.ValueOf(vcl.NewPopupMenu),
		"NewPrintDialog":                      reflect.ValueOf(vcl.NewPrintDialog),
		"NewPrinter":                          reflect.ValueOf(vcl.NewPrinter),
		"NewPrinterSetupDialog":               reflect.ValueOf(vcl.NewPrinterSetupDialog),
		"NewProgressBar":                      reflect.ValueOf(vcl.NewProgressBar),
		"NewRadioButton":                      reflect.ValueOf(vcl.NewRadioButton),
		"NewRadioGroup":                       reflect.ValueOf(vcl.NewRadioGroup),
		"NewRegistry":                         reflect.ValueOf(vcl.NewRegistry),
		"NewRegistryAllAccess":                reflect.ValueOf(vcl.NewRegistryAllAccess),
		"NewReplaceDialog":                    reflect.ValueOf(vcl.NewReplaceDialog),
		"NewSaveDialog":                       reflect.ValueOf(vcl.NewSaveDialog),
		"NewSavePictureDialog":                reflect.ValueOf(vcl.NewSavePictureDialog),
		"NewScreen":                           reflect.ValueOf(vcl.NewScreen),
		"NewScrollBar":                        reflect.ValueOf(vcl.NewScrollBar),
		"NewScrollBox":                        reflect.ValueOf(vcl.NewScrollBox),
		"NewShape":                            reflect.ValueOf(vcl.NewShape),
		"NewSpeedButton":                      reflect.ValueOf(vcl.NewSpeedButton),
		"NewSpinEdit":                         reflect.ValueOf(vcl.NewSpinEdit),
		"NewSplitter":                         reflect.ValueOf(vcl.NewSplitter),
		"NewStaticText":                       reflect.ValueOf(vcl.NewStaticText),
		"NewStatusBar":                        reflect.ValueOf(vcl.NewStatusBar),
		"NewStatusPanel":                      reflect.ValueOf(vcl.NewStatusPanel),
		"NewStatusPanels":                     reflect.ValueOf(vcl.NewStatusPanels),
		"NewStringGrid":                       reflect.ValueOf(vcl.NewStringGrid),
		"NewStringList":                       reflect.ValueOf(vcl.NewStringList),
		"NewStrings":                          reflect.ValueOf(vcl.NewStrings),
		"NewTabSheet":                         reflect.ValueOf(vcl.NewTabSheet),
		"NewTaskDialog":                       reflect.ValueOf(vcl.NewTaskDialog),
		"NewTaskDialogBaseButtonItem":         reflect.ValueOf(vcl.NewTaskDialogBaseButtonItem),
		"NewTaskDialogButtonItem":             reflect.ValueOf(vcl.NewTaskDialogButtonItem),
		"NewTaskDialogRadioButtonItem":        reflect.ValueOf(vcl.NewTaskDialogRadioButtonItem),
		"NewTimer":                            reflect.ValueOf(vcl.NewTimer),
		"NewToolBar":                          reflect.ValueOf(vcl.NewToolBar),
		"NewToolButton":                       reflect.ValueOf(vcl.NewToolButton),
		"NewTrackBar":                         reflect.ValueOf(vcl.NewTrackBar),
		"NewTrayIcon":                         reflect.ValueOf(vcl.NewTrayIcon),
		"NewTreeNode":                         reflect.ValueOf(vcl.NewTreeNode),
		"NewTreeNodes":                        reflect.ValueOf(vcl.NewTreeNodes),
		"NewTreeView":                         reflect.ValueOf(vcl.NewTreeView),
		"NewUpDown":                           reflect.ValueOf(vcl.NewUpDown),
		"NewValueListEditor":                  reflect.ValueOf(vcl.NewValueListEditor),
		"NewWinControl":                       reflect.ValueOf(vcl.NewWinControl),
		"NewXButton":                          reflect.ValueOf(vcl.NewXButton),
		"ObjectFromInst":                      reflect.ValueOf(vcl.ObjectFromInst),
		"ObjectFromObj":                       reflect.ValueOf(vcl.ObjectFromObj),
		"ObjectFromUnsafePointer":             reflect.ValueOf(vcl.ObjectFromUnsafePointer),
		"OpenDialogFromInst":                  reflect.ValueOf(vcl.OpenDialogFromInst),
		"OpenDialogFromObj":                   reflect.ValueOf(vcl.OpenDialogFromObj),
		"OpenDialogFromUnsafePointer":         reflect.ValueOf(vcl.OpenDialogFromUnsafePointer),
		"OpenPictureDialogFromInst":           reflect.ValueOf(vcl.OpenPictureDialogFromInst),
		"OpenPictureDialogFromObj":            reflect.ValueOf(vcl.OpenPictureDialogFromObj),
		"OpenPictureDialogFromUnsafePointer":  reflect.ValueOf(vcl.OpenPictureDialogFromUnsafePointer),
		"PageControlFromInst":                 reflect.ValueOf(vcl.PageControlFromInst),
		"PageControlFromObj":                  reflect.ValueOf(vcl.PageControlFromObj),
		"PageControlFromUnsafePointer":        reflect.ValueOf(vcl.PageControlFromUnsafePointer),
		"PageSetupDialogFromInst":             reflect.ValueOf(vcl.PageSetupDialogFromInst),
		"PageSetupDialogFromObj":              reflect.ValueOf(vcl.PageSetupDialogFromObj),
		"PageSetupDialogFromUnsafePointer":    reflect.ValueOf(vcl.PageSetupDialogFromUnsafePointer),
		"PaintBoxFromInst":                    reflect.ValueOf(vcl.PaintBoxFromInst),
		"PaintBoxFromObj":                     reflect.ValueOf(vcl.PaintBoxFromObj),
		"PaintBoxFromUnsafePointer":           reflect.ValueOf(vcl.PaintBoxFromUnsafePointer),
		"PanelFromInst":                       reflect.ValueOf(vcl.PanelFromInst),
		"PanelFromObj":                        reflect.ValueOf(vcl.PanelFromObj),
		"PanelFromUnsafePointer":              reflect.ValueOf(vcl.PanelFromUnsafePointer),
		"PenFromInst":                         reflect.ValueOf(vcl.PenFromInst),
		"PenFromObj":                          reflect.ValueOf(vcl.PenFromObj),
		"PenFromUnsafePointer":                reflect.ValueOf(vcl.PenFromUnsafePointer),
		"PictureFromInst":                     reflect.ValueOf(vcl.PictureFromInst),
		"PictureFromObj":                      reflect.ValueOf(vcl.PictureFromObj),
		"PictureFromUnsafePointer":            reflect.ValueOf(vcl.PictureFromUnsafePointer),
		"PngImageFromInst":                    reflect.ValueOf(vcl.PngImageFromInst),
		"PngImageFromObj":                     reflect.ValueOf(vcl.PngImageFromObj),
		"PngImageFromUnsafePointer":           reflect.ValueOf(vcl.PngImageFromUnsafePointer),
		"PopupMenuFromInst":                   reflect.ValueOf(vcl.PopupMenuFromInst),
		"PopupMenuFromObj":                    reflect.ValueOf(vcl.PopupMenuFromObj),
		"PopupMenuFromUnsafePointer":          reflect.ValueOf(vcl.PopupMenuFromUnsafePointer),
		"PrintDialogFromInst":                 reflect.ValueOf(vcl.PrintDialogFromInst),
		"PrintDialogFromObj":                  reflect.ValueOf(vcl.PrintDialogFromObj),
		"PrintDialogFromUnsafePointer":        reflect.ValueOf(vcl.PrintDialogFromUnsafePointer),
		"Printer":                             reflect.ValueOf(&vcl.Printer).Elem(),
		"PrinterFromInst":                     reflect.ValueOf(vcl.PrinterFromInst),
		"PrinterFromObj":                      reflect.ValueOf(vcl.PrinterFromObj),
		"PrinterFromUnsafePointer":            reflect.ValueOf(vcl.PrinterFromUnsafePointer),
		"PrinterSetupDialogFromInst":          reflect.ValueOf(vcl.PrinterSetupDialogFromInst),
		"PrinterSetupDialogFromObj":           reflect.ValueOf(vcl.PrinterSetupDialogFromObj),
		"PrinterSetupDialogFromUnsafePointer": reflect.ValueOf(vcl.PrinterSetupDialogFromUnsafePointer),
		"ProgressBarFromInst":                 reflect.ValueOf(vcl.ProgressBarFromInst),
		"ProgressBarFromObj":                  reflect.ValueOf(vcl.ProgressBarFromObj),
		"ProgressBarFromUnsafePointer":        reflect.ValueOf(vcl.ProgressBarFromUnsafePointer),
		"RadioButtonFromInst":                 reflect.ValueOf(vcl.RadioButtonFromInst),
		"RadioButtonFromObj":                  reflect.ValueOf(vcl.RadioButtonFromObj),
		"RadioButtonFromUnsafePointer":        reflect.ValueOf(vcl.RadioButtonFromUnsafePointer),
		"RadioGroupFromInst":                  reflect.ValueOf(vcl.RadioGroupFromInst),
		"RadioGroupFromObj":                   reflect.ValueOf(vcl.RadioGroupFromObj),
		"RadioGroupFromUnsafePointer":         reflect.ValueOf(vcl.RadioGroupFromUnsafePointer),
		"RegisterExtEventCallback":            reflect.ValueOf(vcl.RegisterExtEventCallback),
		"RegisterFormResource":                reflect.ValueOf(vcl.RegisterFormResource),
		"RegistryFromInst":                    reflect.ValueOf(vcl.RegistryFromInst),
		"RegistryFromObj":                     reflect.ValueOf(vcl.RegistryFromObj),
		"RegistryFromUnsafePointer":           reflect.ValueOf(vcl.RegistryFromUnsafePointer),
		"ReplaceDialogFromInst":               reflect.ValueOf(vcl.ReplaceDialogFromInst),
		"ReplaceDialogFromObj":                reflect.ValueOf(vcl.ReplaceDialogFromObj),
		"ReplaceDialogFromUnsafePointer":      reflect.ValueOf(vcl.ReplaceDialogFromUnsafePointer),
		"RunApp":                              reflect.ValueOf(vcl.RunApp),
		"SaveDialogFromInst":                  reflect.ValueOf(vcl.SaveDialogFromInst),
		"SaveDialogFromObj":                   reflect.ValueOf(vcl.SaveDialogFromObj),
		"SaveDialogFromUnsafePointer":         reflect.ValueOf(vcl.SaveDialogFromUnsafePointer),
		"SavePictureDialogFromInst":           reflect.ValueOf(vcl.SavePictureDialogFromInst),
		"SavePictureDialogFromObj":            reflect.ValueOf(vcl.SavePictureDialogFromObj),
		"SavePictureDialogFromUnsafePointer":  reflect.ValueOf(vcl.SavePictureDialogFromUnsafePointer),
		"Screen":                              reflect.ValueOf(&vcl.Screen).Elem(),
		"ScreenFromInst":                      reflect.ValueOf(vcl.ScreenFromInst),
		"ScreenFromObj":                       reflect.ValueOf(vcl.ScreenFromObj),
		"ScreenFromUnsafePointer":             reflect.ValueOf(vcl.ScreenFromUnsafePointer),
		"ScrollBarFromInst":                   reflect.ValueOf(vcl.ScrollBarFromInst),
		"ScrollBarFromObj":                    reflect.ValueOf(vcl.ScrollBarFromObj),
		"ScrollBarFromUnsafePointer":          reflect.ValueOf(vcl.ScrollBarFromUnsafePointer),
		"ScrollBoxFromInst":                   reflect.ValueOf(vcl.ScrollBoxFromInst),
		"ScrollBoxFromObj":                    reflect.ValueOf(vcl.ScrollBoxFromObj),
		"ScrollBoxFromUnsafePointer":          reflect.ValueOf(vcl.ScrollBoxFromUnsafePointer),
		"SelectDirectory1":                    reflect.ValueOf(vcl.SelectDirectory1),
		"SelectDirectory2":                    reflect.ValueOf(vcl.SelectDirectory2),
		"SelectDirectory3":                    reflect.ValueOf(vcl.SelectDirectory3),
		"SetClipboard":                        reflect.ValueOf(vcl.SetClipboard),
		"ShapeFromInst":                       reflect.ValueOf(vcl.ShapeFromInst),
		"ShapeFromObj":                        reflect.ValueOf(vcl.ShapeFromObj),
		"ShapeFromUnsafePointer":              reflect.ValueOf(vcl.ShapeFromUnsafePointer),
		"ShowMessage":                         reflect.ValueOf(vcl.ShowMessage),
		"ShowMessageFmt":                      reflect.ValueOf(vcl.ShowMessageFmt),
		"SizeConstraintsFromInst":             reflect.ValueOf(vcl.SizeConstraintsFromInst),
		"SizeConstraintsFromObj":              reflect.ValueOf(vcl.SizeConstraintsFromObj),
		"SizeConstraintsFromUnsafePointer":    reflect.ValueOf(vcl.SizeConstraintsFromUnsafePointer),
		"SpeedButtonFromInst":                 reflect.ValueOf(vcl.SpeedButtonFromInst),
		"SpeedButtonFromObj":                  reflect.ValueOf(vcl.SpeedButtonFromObj),
		"SpeedButtonFromUnsafePointer":        reflect.ValueOf(vcl.SpeedButtonFromUnsafePointer),
		"SpinEditFromInst":                    reflect.ValueOf(vcl.SpinEditFromInst),
		"SpinEditFromObj":                     reflect.ValueOf(vcl.SpinEditFromObj),
		"SpinEditFromUnsafePointer":           reflect.ValueOf(vcl.SpinEditFromUnsafePointer),
		"SplitterFromInst":                    reflect.ValueOf(vcl.SplitterFromInst),
		"SplitterFromObj":                     reflect.ValueOf(vcl.SplitterFromObj),
		"SplitterFromUnsafePointer":           reflect.ValueOf(vcl.SplitterFromUnsafePointer),
		"StaticTextFromInst":                  reflect.ValueOf(vcl.StaticTextFromInst),
		"StaticTextFromObj":                   reflect.ValueOf(vcl.StaticTextFromObj),
		"StaticTextFromUnsafePointer":         reflect.ValueOf(vcl.StaticTextFromUnsafePointer),
		"StatusBarFromInst":                   reflect.ValueOf(vcl.StatusBarFromInst),
		"StatusBarFromObj":                    reflect.ValueOf(vcl.StatusBarFromObj),
		"StatusBarFromUnsafePointer":          reflect.ValueOf(vcl.StatusBarFromUnsafePointer),
		"StatusPanelFromInst":                 reflect.ValueOf(vcl.StatusPanelFromInst),
		"StatusPanelFromObj":                  reflect.ValueOf(vcl.StatusPanelFromObj),
		"StatusPanelFromUnsafePointer":        reflect.ValueOf(vcl.StatusPanelFromUnsafePointer),
		"StatusPanelsFromInst":                reflect.ValueOf(vcl.StatusPanelsFromInst),
		"StatusPanelsFromObj":                 reflect.ValueOf(vcl.StatusPanelsFromObj),
		"StatusPanelsFromUnsafePointer":       reflect.ValueOf(vcl.StatusPanelsFromUnsafePointer),
		"StringGridFromInst":                  reflect.ValueOf(vcl.StringGridFromInst),
		"StringGridFromObj":                   reflect.ValueOf(vcl.StringGridFromObj),
		"StringGridFromUnsafePointer":         reflect.ValueOf(vcl.StringGridFromUnsafePointer),
		"StringListFromInst":                  reflect.ValueOf(vcl.StringListFromInst),
		"StringListFromObj":                   reflect.ValueOf(vcl.StringListFromObj),
		"StringListFromUnsafePointer":         reflect.ValueOf(vcl.StringListFromUnsafePointer),
		"StringsFromInst":                     reflect.ValueOf(vcl.StringsFromInst),
		"StringsFromObj":                      reflect.ValueOf(vcl.StringsFromObj),
		"StringsFromUnsafePointer":            reflect.ValueOf(vcl.StringsFromUnsafePointer),
		"TActionClass":                        reflect.ValueOf(vcl.TActionClass),
		"TActionListClass":                    reflect.ValueOf(vcl.TActionListClass),
		"TApplicationClass":                   reflect.ValueOf(vcl.TApplicationClass),
		"TBevelClass":                         reflect.ValueOf(vcl.TBevelClass),
		"TBitBtnClass":                        reflect.ValueOf(vcl.TBitBtnClass),
		"TBitmapClass":                        reflect.ValueOf(vcl.TBitmapClass),
		"TBoundLabelClass":                    reflect.ValueOf(vcl.TBoundLabelClass),
		"TBrushClass":                         reflect.ValueOf(vcl.TBrushClass),
		"TButtonClass":                        reflect.ValueOf(vcl.TButtonClass),
		"TCanvasClass":                        reflect.ValueOf(vcl.TCanvasClass),
		"TCheckBoxClass":                      reflect.ValueOf(vcl.TCheckBoxClass),
		"TCheckListBoxClass":                  reflect.ValueOf(vcl.TCheckListBoxClass),
		"TClipboardClass":                     reflect.ValueOf(vcl.TClipboardClass),
		"TCollectionClass":                    reflect.ValueOf(vcl.TCollectionClass),
		"TCollectionItemClass":                reflect.ValueOf(vcl.TCollectionItemClass),
		"TColorBoxClass":                      reflect.ValueOf(vcl.TColorBoxClass),
		"TColorDialogClass":                   reflect.ValueOf(vcl.TColorDialogClass),
		"TColorListBoxClass":                  reflect.ValueOf(vcl.TColorListBoxClass),
		"TComboBoxClass":                      reflect.ValueOf(vcl.TComboBoxClass),
		"TComboBoxExClass":                    reflect.ValueOf(vcl.TComboBoxExClass),
		"TComboExItemClass":                   reflect.ValueOf(vcl.TComboExItemClass),
		"TComboExItemsClass":                  reflect.ValueOf(vcl.TComboExItemsClass),
		"TComponentClass":                     reflect.ValueOf(vcl.TComponentClass),
		"TControlClass":                       reflect.ValueOf(vcl.TControlClass),
		"TControlScrollBarClass":              reflect.ValueOf(vcl.TControlScrollBarClass),
		"TCoolBandClass":                      reflect.ValueOf(vcl.TCoolBandClass),
		"TCoolBandsClass":                     reflect.ValueOf(vcl.TCoolBandsClass),
		"TCoolBarClass":                       reflect.ValueOf(vcl.TCoolBarClass),
		"TDateTimePickerClass":                reflect.ValueOf(vcl.TDateTimePickerClass),
		"TDragDockObjectClass":                reflect.ValueOf(vcl.TDragDockObjectClass),
		"TDragObjectClass":                    reflect.ValueOf(vcl.TDragObjectClass),
		"TDrawGridClass":                      reflect.ValueOf(vcl.TDrawGridClass),
		"TEditClass":                          reflect.ValueOf(vcl.TEditClass),
		"TFindDialogClass":                    reflect.ValueOf(vcl.TFindDialogClass),
		"TFlowPanelClass":                     reflect.ValueOf(vcl.TFlowPanelClass),
		"TFontClass":                          reflect.ValueOf(vcl.TFontClass),
		"TFontDialogClass":                    reflect.ValueOf(vcl.TFontDialogClass),
		"TFormClass":                          reflect.ValueOf(vcl.TFormClass),
		"TFrameClass":                         reflect.ValueOf(vcl.TFrameClass),
		"TGIFImageClass":                      reflect.ValueOf(vcl.TGIFImageClass),
		"TGaugeClass":                         reflect.ValueOf(vcl.TGaugeClass),
		"TGraphicClass":                       reflect.ValueOf(vcl.TGraphicClass),
		"TGroupBoxClass":                      reflect.ValueOf(vcl.TGroupBoxClass),
		"THeaderControlClass":                 reflect.ValueOf(vcl.THeaderControlClass),
		"THeaderSectionClass":                 reflect.ValueOf(vcl.THeaderSectionClass),
		"THeaderSectionsClass":                reflect.ValueOf(vcl.THeaderSectionsClass),
		"TIconClass":                          reflect.ValueOf(vcl.TIconClass),
		"TIconOptionsClass":                   reflect.ValueOf(vcl.TIconOptionsClass),
		"TImageButtonClass":                   reflect.ValueOf(vcl.TImageButtonClass),
		"TImageClass":                         reflect.ValueOf(vcl.TImageClass),
		"TImageListClass":                     reflect.ValueOf(vcl.TImageListClass),
		"TIniFileClass":                       reflect.ValueOf(vcl.TIniFileClass),
		"TJPEGImageClass":                     reflect.ValueOf(vcl.TJPEGImageClass),
		"TLabelClass":                         reflect.ValueOf(vcl.TLabelClass),
		"TLabeledEditClass":                   reflect.ValueOf(vcl.TLabeledEditClass),
		"TLinkLabelClass":                     reflect.ValueOf(vcl.TLinkLabelClass),
		"TListBoxClass":                       reflect.ValueOf(vcl.TListBoxClass),
		"TListClass":                          reflect.ValueOf(vcl.TListClass),
		"TListColumnClass":                    reflect.ValueOf(vcl.TListColumnClass),
		"TListColumnsClass":                   reflect.ValueOf(vcl.TListColumnsClass),
		"TListItemClass":                      reflect.ValueOf(vcl.TListItemClass),
		"TListItemsClass":                     reflect.ValueOf(vcl.TListItemsClass),
		"TListViewClass":                      reflect.ValueOf(vcl.TListViewClass),
		"TMainMenuClass":                      reflect.ValueOf(vcl.TMainMenuClass),
		"TMarginsClass":                       reflect.ValueOf(vcl.TMarginsClass),
		"TMaskEditClass":                      reflect.ValueOf(vcl.TMaskEditClass),
		"TMemoClass":                          reflect.ValueOf(vcl.TMemoClass),
		"TMemoryStreamClass":                  reflect.ValueOf(vcl.TMemoryStreamClass),
		"TMenuItemClass":                      reflect.ValueOf(vcl.TMenuItemClass),
		"TMiniWebviewClass":                   reflect.ValueOf(vcl.TMiniWebviewClass),
		"TMonitorClass":                       reflect.ValueOf(vcl.TMonitorClass),
		"TMonthCalendarClass":                 reflect.ValueOf(vcl.TMonthCalendarClass),
		"TMouseClass":                         reflect.ValueOf(vcl.TMouseClass),
		"TObjectClass":                        reflect.ValueOf(vcl.TObjectClass),
		"TOpenDialogClass":                    reflect.ValueOf(vcl.TOpenDialogClass),
		"TOpenPictureDialogClass":             reflect.ValueOf(vcl.TOpenPictureDialogClass),
		"TPageControlClass":                   reflect.ValueOf(vcl.TPageControlClass),
		"TPageSetupDialogClass":               reflect.ValueOf(vcl.TPageSetupDialogClass),
		"TPaintBoxClass":                      reflect.ValueOf(vcl.TPaintBoxClass),
		"TPanelClass":                         reflect.ValueOf(vcl.TPanelClass),
		"TPenClass":                           reflect.ValueOf(vcl.TPenClass),
		"TPictureClass":                       reflect.ValueOf(vcl.TPictureClass),
		"TPngImageClass":                      reflect.ValueOf(vcl.TPngImageClass),
		"TPopupMenuClass":                     reflect.ValueOf(vcl.TPopupMenuClass),
		"TPrintDialogClass":                   reflect.ValueOf(vcl.TPrintDialogClass),
		"TPrinterClass":                       reflect.ValueOf(vcl.TPrinterClass),
		"TPrinterSetupDialogClass":            reflect.ValueOf(vcl.TPrinterSetupDialogClass),
		"TProgressBarClass":                   reflect.ValueOf(vcl.TProgressBarClass),
		"TRadioButtonClass":                   reflect.ValueOf(vcl.TRadioButtonClass),
		"TRadioGroupClass":                    reflect.ValueOf(vcl.TRadioGroupClass),
		"TRegistryClass":                      reflect.ValueOf(vcl.TRegistryClass),
		"TReplaceDialogClass":                 reflect.ValueOf(vcl.TReplaceDialogClass),
		"TSaveDialogClass":                    reflect.ValueOf(vcl.TSaveDialogClass),
		"TSavePictureDialogClass":             reflect.ValueOf(vcl.TSavePictureDialogClass),
		"TScreenClass":                        reflect.ValueOf(vcl.TScreenClass),
		"TScrollBarClass":                     reflect.ValueOf(vcl.TScrollBarClass),
		"TScrollBoxClass":                     reflect.ValueOf(vcl.TScrollBoxClass),
		"TShapeClass":                         reflect.ValueOf(vcl.TShapeClass),
		"TSizeConstraintsClass":               reflect.ValueOf(vcl.TSizeConstraintsClass),
		"TSpeedButtonClass":                   reflect.ValueOf(vcl.TSpeedButtonClass),
		"TSpinEditClass":                      reflect.ValueOf(vcl.TSpinEditClass),
		"TSplitterClass":                      reflect.ValueOf(vcl.TSplitterClass),
		"TStaticTextClass":                    reflect.ValueOf(vcl.TStaticTextClass),
		"TStatusBarClass":                     reflect.ValueOf(vcl.TStatusBarClass),
		"TStatusPanelClass":                   reflect.ValueOf(vcl.TStatusPanelClass),
		"TStatusPanelsClass":                  reflect.ValueOf(vcl.TStatusPanelsClass),
		"TStringGridClass":                    reflect.ValueOf(vcl.TStringGridClass),
		"TStringListClass":                    reflect.ValueOf(vcl.TStringListClass),
		"TStringsClass":                       reflect.ValueOf(vcl.TStringsClass),
		"TTabSheetClass":                      reflect.ValueOf(vcl.TTabSheetClass),
		"TTaskDialogBaseButtonItemClass":      reflect.ValueOf(vcl.TTaskDialogBaseButtonItemClass),
		"TTaskDialogButtonItemClass":          reflect.ValueOf(vcl.TTaskDialogButtonItemClass),
		"TTaskDialogButtonsClass":             reflect.ValueOf(vcl.TTaskDialogButtonsClass),
		"TTaskDialogClass":                    reflect.ValueOf(vcl.TTaskDialogClass),
		"TTaskDialogRadioButtonItemClass":     reflect.ValueOf(vcl.TTaskDialogRadioButtonItemClass),
		"TTimerClass":                         reflect.ValueOf(vcl.TTimerClass),
		"TToolBarClass":                       reflect.ValueOf(vcl.TToolBarClass),
		"TToolButtonClass":                    reflect.ValueOf(vcl.TToolButtonClass),
		"TTrackBarClass":                      reflect.ValueOf(vcl.TTrackBarClass),
		"TTrayIconClass":                      reflect.ValueOf(vcl.TTrayIconClass),
		"TTreeNodeClass":                      reflect.ValueOf(vcl.TTreeNodeClass),
		"TTreeNodesClass":                     reflect.ValueOf(vcl.TTreeNodesClass),
		"TTreeViewClass":                      reflect.ValueOf(vcl.TTreeViewClass),
		"TUpDownClass":                        reflect.ValueOf(vcl.TUpDownClass),
		"TValueListEditorClass":               reflect.ValueOf(vcl.TValueListEditorClass),
		"TWinControlClass":                    reflect.ValueOf(vcl.TWinControlClass),
		"TXButtonClass":                       reflect.ValueOf(vcl.TXButtonClass),
		"TabSheetFromInst":                    reflect.ValueOf(vcl.TabSheetFromInst),
		"TabSheetFromObj":                     reflect.ValueOf(vcl.TabSheetFromObj),
		"TabSheetFromUnsafePointer":           reflect.ValueOf(vcl.TabSheetFromUnsafePointer),
		"TaskDialogBaseButtonItemFromInst":    reflect.ValueOf(vcl.TaskDialogBaseButtonItemFromInst),
		"TaskDialogBaseButtonItemFromObj":     reflect.ValueOf(vcl.TaskDialogBaseButtonItemFromObj),
		"TaskDialogBaseButtonItemFromUnsafePointer":  reflect.ValueOf(vcl.TaskDialogBaseButtonItemFromUnsafePointer),
		"TaskDialogButtonItemFromInst":               reflect.ValueOf(vcl.TaskDialogButtonItemFromInst),
		"TaskDialogButtonItemFromObj":                reflect.ValueOf(vcl.TaskDialogButtonItemFromObj),
		"TaskDialogButtonItemFromUnsafePointer":      reflect.ValueOf(vcl.TaskDialogButtonItemFromUnsafePointer),
		"TaskDialogButtonsFromInst":                  reflect.ValueOf(vcl.TaskDialogButtonsFromInst),
		"TaskDialogButtonsFromObj":                   reflect.ValueOf(vcl.TaskDialogButtonsFromObj),
		"TaskDialogButtonsFromUnsafePointer":         reflect.ValueOf(vcl.TaskDialogButtonsFromUnsafePointer),
		"TaskDialogFromInst":                         reflect.ValueOf(vcl.TaskDialogFromInst),
		"TaskDialogFromObj":                          reflect.ValueOf(vcl.TaskDialogFromObj),
		"TaskDialogFromUnsafePointer":                reflect.ValueOf(vcl.TaskDialogFromUnsafePointer),
		"TaskDialogRadioButtonItemFromInst":          reflect.ValueOf(vcl.TaskDialogRadioButtonItemFromInst),
		"TaskDialogRadioButtonItemFromObj":           reflect.ValueOf(vcl.TaskDialogRadioButtonItemFromObj),
		"TaskDialogRadioButtonItemFromUnsafePointer": reflect.ValueOf(vcl.TaskDialogRadioButtonItemFromUnsafePointer),
		"ThreadSync":                       reflect.ValueOf(vcl.ThreadSync),
		"TimerFromInst":                    reflect.ValueOf(vcl.TimerFromInst),
		"TimerFromObj":                     reflect.ValueOf(vcl.TimerFromObj),
		"TimerFromUnsafePointer":           reflect.ValueOf(vcl.TimerFromUnsafePointer),
		"ToolBarFromInst":                  reflect.ValueOf(vcl.ToolBarFromInst),
		"ToolBarFromObj":                   reflect.ValueOf(vcl.ToolBarFromObj),
		"ToolBarFromUnsafePointer":         reflect.ValueOf(vcl.ToolBarFromUnsafePointer),
		"ToolButtonFromInst":               reflect.ValueOf(vcl.ToolButtonFromInst),
		"ToolButtonFromObj":                reflect.ValueOf(vcl.ToolButtonFromObj),
		"ToolButtonFromUnsafePointer":      reflect.ValueOf(vcl.ToolButtonFromUnsafePointer),
		"TrackBarFromInst":                 reflect.ValueOf(vcl.TrackBarFromInst),
		"TrackBarFromObj":                  reflect.ValueOf(vcl.TrackBarFromObj),
		"TrackBarFromUnsafePointer":        reflect.ValueOf(vcl.TrackBarFromUnsafePointer),
		"TrayIconFromInst":                 reflect.ValueOf(vcl.TrayIconFromInst),
		"TrayIconFromObj":                  reflect.ValueOf(vcl.TrayIconFromObj),
		"TrayIconFromUnsafePointer":        reflect.ValueOf(vcl.TrayIconFromUnsafePointer),
		"TreeNodeFromInst":                 reflect.ValueOf(vcl.TreeNodeFromInst),
		"TreeNodeFromObj":                  reflect.ValueOf(vcl.TreeNodeFromObj),
		"TreeNodeFromUnsafePointer":        reflect.ValueOf(vcl.TreeNodeFromUnsafePointer),
		"TreeNodesFromInst":                reflect.ValueOf(vcl.TreeNodesFromInst),
		"TreeNodesFromObj":                 reflect.ValueOf(vcl.TreeNodesFromObj),
		"TreeNodesFromUnsafePointer":       reflect.ValueOf(vcl.TreeNodesFromUnsafePointer),
		"TreeViewFromInst":                 reflect.ValueOf(vcl.TreeViewFromInst),
		"TreeViewFromObj":                  reflect.ValueOf(vcl.TreeViewFromObj),
		"TreeViewFromUnsafePointer":        reflect.ValueOf(vcl.TreeViewFromUnsafePointer),
		"UpDownFromInst":                   reflect.ValueOf(vcl.UpDownFromInst),
		"UpDownFromObj":                    reflect.ValueOf(vcl.UpDownFromObj),
		"UpDownFromUnsafePointer":          reflect.ValueOf(vcl.UpDownFromUnsafePointer),
		"ValueListEditorFromInst":          reflect.ValueOf(vcl.ValueListEditorFromInst),
		"ValueListEditorFromObj":           reflect.ValueOf(vcl.ValueListEditorFromObj),
		"ValueListEditorFromUnsafePointer": reflect.ValueOf(vcl.ValueListEditorFromUnsafePointer),
		"WinControlFromInst":               reflect.ValueOf(vcl.WinControlFromInst),
		"WinControlFromObj":                reflect.ValueOf(vcl.WinControlFromObj),
		"WinControlFromUnsafePointer":      reflect.ValueOf(vcl.WinControlFromUnsafePointer),
		"XButtonFromInst":                  reflect.ValueOf(vcl.XButtonFromInst),
		"XButtonFromObj":                   reflect.ValueOf(vcl.XButtonFromObj),
		"XButtonFromUnsafePointer":         reflect.ValueOf(vcl.XButtonFromUnsafePointer),

		// type definitions
		"Exception":                         reflect.ValueOf((*vcl.Exception)(nil)),
		"IComponent":                        reflect.ValueOf((*vcl.IComponent)(nil)),
		"IControl":                          reflect.ValueOf((*vcl.IControl)(nil)),
		"IObject":                           reflect.ValueOf((*vcl.IObject)(nil)),
		"IWinControl":                       reflect.ValueOf((*vcl.IWinControl)(nil)),
		"TAction":                           reflect.ValueOf((*vcl.TAction)(nil)),
		"TActionList":                       reflect.ValueOf((*vcl.TActionList)(nil)),
		"TAlignPositionEvent":               reflect.ValueOf((*vcl.TAlignPositionEvent)(nil)),
		"TApplication":                      reflect.ValueOf((*vcl.TApplication)(nil)),
		"TAs":                               reflect.ValueOf((*vcl.TAs)(nil)),
		"TBevel":                            reflect.ValueOf((*vcl.TBevel)(nil)),
		"TBitBtn":                           reflect.ValueOf((*vcl.TBitBtn)(nil)),
		"TBitmap":                           reflect.ValueOf((*vcl.TBitmap)(nil)),
		"TBoundLabel":                       reflect.ValueOf((*vcl.TBoundLabel)(nil)),
		"TBrush":                            reflect.ValueOf((*vcl.TBrush)(nil)),
		"TButton":                           reflect.ValueOf((*vcl.TButton)(nil)),
		"TCanvas":                           reflect.ValueOf((*vcl.TCanvas)(nil)),
		"TCheckBox":                         reflect.ValueOf((*vcl.TCheckBox)(nil)),
		"TCheckListBox":                     reflect.ValueOf((*vcl.TCheckListBox)(nil)),
		"TClipboard":                        reflect.ValueOf((*vcl.TClipboard)(nil)),
		"TCloseEvent":                       reflect.ValueOf((*vcl.TCloseEvent)(nil)),
		"TCloseQueryEvent":                  reflect.ValueOf((*vcl.TCloseQueryEvent)(nil)),
		"TCollection":                       reflect.ValueOf((*vcl.TCollection)(nil)),
		"TCollectionItem":                   reflect.ValueOf((*vcl.TCollectionItem)(nil)),
		"TColorBox":                         reflect.ValueOf((*vcl.TColorBox)(nil)),
		"TColorDialog":                      reflect.ValueOf((*vcl.TColorDialog)(nil)),
		"TColorListBox":                     reflect.ValueOf((*vcl.TColorListBox)(nil)),
		"TComboBox":                         reflect.ValueOf((*vcl.TComboBox)(nil)),
		"TComboBoxEx":                       reflect.ValueOf((*vcl.TComboBoxEx)(nil)),
		"TComboExItem":                      reflect.ValueOf((*vcl.TComboExItem)(nil)),
		"TComboExItems":                     reflect.ValueOf((*vcl.TComboExItems)(nil)),
		"TComponent":                        reflect.ValueOf((*vcl.TComponent)(nil)),
		"TConstrainedResizeEvent":           reflect.ValueOf((*vcl.TConstrainedResizeEvent)(nil)),
		"TContextPopupEvent":                reflect.ValueOf((*vcl.TContextPopupEvent)(nil)),
		"TControl":                          reflect.ValueOf((*vcl.TControl)(nil)),
		"TControlScrollBar":                 reflect.ValueOf((*vcl.TControlScrollBar)(nil)),
		"TCoolBand":                         reflect.ValueOf((*vcl.TCoolBand)(nil)),
		"TCoolBands":                        reflect.ValueOf((*vcl.TCoolBands)(nil)),
		"TCoolBar":                          reflect.ValueOf((*vcl.TCoolBar)(nil)),
		"TCreatingListErrorEvent":           reflect.ValueOf((*vcl.TCreatingListErrorEvent)(nil)),
		"TCustomSectionNotifyEvent":         reflect.ValueOf((*vcl.TCustomSectionNotifyEvent)(nil)),
		"TDateTimePicker":                   reflect.ValueOf((*vcl.TDateTimePicker)(nil)),
		"TDockDropEvent":                    reflect.ValueOf((*vcl.TDockDropEvent)(nil)),
		"TDockOverEvent":                    reflect.ValueOf((*vcl.TDockOverEvent)(nil)),
		"TDragDockObject":                   reflect.ValueOf((*vcl.TDragDockObject)(nil)),
		"TDragDropEvent":                    reflect.ValueOf((*vcl.TDragDropEvent)(nil)),
		"TDragObject":                       reflect.ValueOf((*vcl.TDragObject)(nil)),
		"TDragOverEvent":                    reflect.ValueOf((*vcl.TDragOverEvent)(nil)),
		"TDrawCellEvent":                    reflect.ValueOf((*vcl.TDrawCellEvent)(nil)),
		"TDrawGrid":                         reflect.ValueOf((*vcl.TDrawGrid)(nil)),
		"TDrawItemEvent":                    reflect.ValueOf((*vcl.TDrawItemEvent)(nil)),
		"TDrawSectionEvent":                 reflect.ValueOf((*vcl.TDrawSectionEvent)(nil)),
		"TDropFilesEvent":                   reflect.ValueOf((*vcl.TDropFilesEvent)(nil)),
		"TEdit":                             reflect.ValueOf((*vcl.TEdit)(nil)),
		"TEndDragEvent":                     reflect.ValueOf((*vcl.TEndDragEvent)(nil)),
		"TExceptionEvent":                   reflect.ValueOf((*vcl.TExceptionEvent)(nil)),
		"TExtEventCallback":                 reflect.ValueOf((*vcl.TExtEventCallback)(nil)),
		"TFindDialog":                       reflect.ValueOf((*vcl.TFindDialog)(nil)),
		"TFixedCellClickEvent":              reflect.ValueOf((*vcl.TFixedCellClickEvent)(nil)),
		"TFlowPanel":                        reflect.ValueOf((*vcl.TFlowPanel)(nil)),
		"TFont":                             reflect.ValueOf((*vcl.TFont)(nil)),
		"TFontDialog":                       reflect.ValueOf((*vcl.TFontDialog)(nil)),
		"TForm":                             reflect.ValueOf((*vcl.TForm)(nil)),
		"TFrame":                            reflect.ValueOf((*vcl.TFrame)(nil)),
		"TGIFImage":                         reflect.ValueOf((*vcl.TGIFImage)(nil)),
		"TGauge":                            reflect.ValueOf((*vcl.TGauge)(nil)),
		"TGetEditEvent":                     reflect.ValueOf((*vcl.TGetEditEvent)(nil)),
		"TGetSiteInfoEvent":                 reflect.ValueOf((*vcl.TGetSiteInfoEvent)(nil)),
		"TGraphic":                          reflect.ValueOf((*vcl.TGraphic)(nil)),
		"TGroupBox":                         reflect.ValueOf((*vcl.TGroupBox)(nil)),
		"THeaderControl":                    reflect.ValueOf((*vcl.THeaderControl)(nil)),
		"THeaderSection":                    reflect.ValueOf((*vcl.THeaderSection)(nil)),
		"THeaderSections":                   reflect.ValueOf((*vcl.THeaderSections)(nil)),
		"THelpEvent":                        reflect.ValueOf((*vcl.THelpEvent)(nil)),
		"TIcon":                             reflect.ValueOf((*vcl.TIcon)(nil)),
		"TIconOptions":                      reflect.ValueOf((*vcl.TIconOptions)(nil)),
		"TImage":                            reflect.ValueOf((*vcl.TImage)(nil)),
		"TImageButton":                      reflect.ValueOf((*vcl.TImageButton)(nil)),
		"TImageList":                        reflect.ValueOf((*vcl.TImageList)(nil)),
		"TIniFile":                          reflect.ValueOf((*vcl.TIniFile)(nil)),
		"TIs":                               reflect.ValueOf((*vcl.TIs)(nil)),
		"TJPEGImage":                        reflect.ValueOf((*vcl.TJPEGImage)(nil)),
		"TKeyEvent":                         reflect.ValueOf((*vcl.TKeyEvent)(nil)),
		"TKeyPressEvent":                    reflect.ValueOf((*vcl.TKeyPressEvent)(nil)),
		"TLBFindDataEvent":                  reflect.ValueOf((*vcl.TLBFindDataEvent)(nil)),
		"TLBGetDataEvent":                   reflect.ValueOf((*vcl.TLBGetDataEvent)(nil)),
		"TLBGetDataObjectEvent":             reflect.ValueOf((*vcl.TLBGetDataObjectEvent)(nil)),
		"TLVAdvancedCustomDrawEvent":        reflect.ValueOf((*vcl.TLVAdvancedCustomDrawEvent)(nil)),
		"TLVAdvancedCustomDrawItemEvent":    reflect.ValueOf((*vcl.TLVAdvancedCustomDrawItemEvent)(nil)),
		"TLVAdvancedCustomDrawSubItemEvent": reflect.ValueOf((*vcl.TLVAdvancedCustomDrawSubItemEvent)(nil)),
		"TLVChangeEvent":                    reflect.ValueOf((*vcl.TLVChangeEvent)(nil)),
		"TLVChangingEvent":                  reflect.ValueOf((*vcl.TLVChangingEvent)(nil)),
		"TLVCheckedItemEvent":               reflect.ValueOf((*vcl.TLVCheckedItemEvent)(nil)),
		"TLVColumnClickEvent":               reflect.ValueOf((*vcl.TLVColumnClickEvent)(nil)),
		"TLVColumnRClickEvent":              reflect.ValueOf((*vcl.TLVColumnRClickEvent)(nil)),
		"TLVCompareEvent":                   reflect.ValueOf((*vcl.TLVCompareEvent)(nil)),
		"TLVCustomDrawEvent":                reflect.ValueOf((*vcl.TLVCustomDrawEvent)(nil)),
		"TLVCustomDrawItemEvent":            reflect.ValueOf((*vcl.TLVCustomDrawItemEvent)(nil)),
		"TLVCustomDrawSubItemEvent":         reflect.ValueOf((*vcl.TLVCustomDrawSubItemEvent)(nil)),
		"TLVDataHintEvent":                  reflect.ValueOf((*vcl.TLVDataHintEvent)(nil)),
		"TLVDeletedEvent":                   reflect.ValueOf((*vcl.TLVDeletedEvent)(nil)),
		"TLVDrawItemEvent":                  reflect.ValueOf((*vcl.TLVDrawItemEvent)(nil)),
		"TLVEditedEvent":                    reflect.ValueOf((*vcl.TLVEditedEvent)(nil)),
		"TLVEditingEvent":                   reflect.ValueOf((*vcl.TLVEditingEvent)(nil)),
		"TLVNotifyEvent":                    reflect.ValueOf((*vcl.TLVNotifyEvent)(nil)),
		"TLVOwnerDataEvent":                 reflect.ValueOf((*vcl.TLVOwnerDataEvent)(nil)),
		"TLVOwnerDataFindEvent":             reflect.ValueOf((*vcl.TLVOwnerDataFindEvent)(nil)),
		"TLVOwnerDataHintEvent":             reflect.ValueOf((*vcl.TLVOwnerDataHintEvent)(nil)),
		"TLVSelectItemEvent":                reflect.ValueOf((*vcl.TLVSelectItemEvent)(nil)),
		"TLabel":                            reflect.ValueOf((*vcl.TLabel)(nil)),
		"TLabeledEdit":                      reflect.ValueOf((*vcl.TLabeledEdit)(nil)),
		"TLinkLabel":                        reflect.ValueOf((*vcl.TLinkLabel)(nil)),
		"TList":                             reflect.ValueOf((*vcl.TList)(nil)),
		"TListBox":                          reflect.ValueOf((*vcl.TListBox)(nil)),
		"TListColumn":                       reflect.ValueOf((*vcl.TListColumn)(nil)),
		"TListColumns":                      reflect.ValueOf((*vcl.TListColumns)(nil)),
		"TListItem":                         reflect.ValueOf((*vcl.TListItem)(nil)),
		"TListItems":                        reflect.ValueOf((*vcl.TListItems)(nil)),
		"TListView":                         reflect.ValueOf((*vcl.TListView)(nil)),
		"TMainMenu":                         reflect.ValueOf((*vcl.TMainMenu)(nil)),
		"TMargins":                          reflect.ValueOf((*vcl.TMargins)(nil)),
		"TMaskEdit":                         reflect.ValueOf((*vcl.TMaskEdit)(nil)),
		"TMeasureItemEvent":                 reflect.ValueOf((*vcl.TMeasureItemEvent)(nil)),
		"TMemo":                             reflect.ValueOf((*vcl.TMemo)(nil)),
		"TMemoryStream":                     reflect.ValueOf((*vcl.TMemoryStream)(nil)),
		"TMenuChangeEvent":                  reflect.ValueOf((*vcl.TMenuChangeEvent)(nil)),
		"TMenuDrawItemEvent":                reflect.ValueOf((*vcl.TMenuDrawItemEvent)(nil)),
		"TMenuItem":                         reflect.ValueOf((*vcl.TMenuItem)(nil)),
		"TMenuMeasureItemEvent":             reflect.ValueOf((*vcl.TMenuMeasureItemEvent)(nil)),
		"TMessageEvent":                     reflect.ValueOf((*vcl.TMessageEvent)(nil)),
		"TMiniWebview":                      reflect.ValueOf((*vcl.TMiniWebview)(nil)),
		"TMonitor":                          reflect.ValueOf((*vcl.TMonitor)(nil)),
		"TMonthCalendar":                    reflect.ValueOf((*vcl.TMonthCalendar)(nil)),
		"TMouse":                            reflect.ValueOf((*vcl.TMouse)(nil)),
		"TMouseActivateEvent":               reflect.ValueOf((*vcl.TMouseActivateEvent)(nil)),
		"TMouseEvent":                       reflect.ValueOf((*vcl.TMouseEvent)(nil)),
		"TMouseMoveEvent":                   reflect.ValueOf((*vcl.TMouseMoveEvent)(nil)),
		"TMouseWheelEvent":                  reflect.ValueOf((*vcl.TMouseWheelEvent)(nil)),
		"TMouseWheelUpDownEvent":            reflect.ValueOf((*vcl.TMouseWheelUpDownEvent)(nil)),
		"TMovedEvent":                       reflect.ValueOf((*vcl.TMovedEvent)(nil)),
		"TNotifyEvent":                      reflect.ValueOf((*vcl.TNotifyEvent)(nil)),
		"TObject":                           reflect.ValueOf((*vcl.TObject)(nil)),
		"TOpenDialog":                       reflect.ValueOf((*vcl.TOpenDialog)(nil)),
		"TOpenPictureDialog":                reflect.ValueOf((*vcl.TOpenPictureDialog)(nil)),
		"TPageControl":                      reflect.ValueOf((*vcl.TPageControl)(nil)),
		"TPageSetupDialog":                  reflect.ValueOf((*vcl.TPageSetupDialog)(nil)),
		"TPaintBox":                         reflect.ValueOf((*vcl.TPaintBox)(nil)),
		"TPanel":                            reflect.ValueOf((*vcl.TPanel)(nil)),
		"TPen":                              reflect.ValueOf((*vcl.TPen)(nil)),
		"TPicture":                          reflect.ValueOf((*vcl.TPicture)(nil)),
		"TPngImage":                         reflect.ValueOf((*vcl.TPngImage)(nil)),
		"TPopupMenu":                        reflect.ValueOf((*vcl.TPopupMenu)(nil)),
		"TPrintDialog":                      reflect.ValueOf((*vcl.TPrintDialog)(nil)),
		"TPrinter":                          reflect.ValueOf((*vcl.TPrinter)(nil)),
		"TPrinterSetupDialog":               reflect.ValueOf((*vcl.TPrinterSetupDialog)(nil)),
		"TProgressBar":                      reflect.ValueOf((*vcl.TProgressBar)(nil)),
		"TRadioButton":                      reflect.ValueOf((*vcl.TRadioButton)(nil)),
		"TRadioGroup":                       reflect.ValueOf((*vcl.TRadioGroup)(nil)),
		"TRegistry":                         reflect.ValueOf((*vcl.TRegistry)(nil)),
		"TReplaceDialog":                    reflect.ValueOf((*vcl.TReplaceDialog)(nil)),
		"TSaveDialog":                       reflect.ValueOf((*vcl.TSaveDialog)(nil)),
		"TSavePictureDialog":                reflect.ValueOf((*vcl.TSavePictureDialog)(nil)),
		"TScreen":                           reflect.ValueOf((*vcl.TScreen)(nil)),
		"TScrollBar":                        reflect.ValueOf((*vcl.TScrollBar)(nil)),
		"TScrollBox":                        reflect.ValueOf((*vcl.TScrollBox)(nil)),
		"TSectionDragEvent":                 reflect.ValueOf((*vcl.TSectionDragEvent)(nil)),
		"TSectionNotifyEvent":               reflect.ValueOf((*vcl.TSectionNotifyEvent)(nil)),
		"TSectionTrackEvent":                reflect.ValueOf((*vcl.TSectionTrackEvent)(nil)),
		"TSelectCellEvent":                  reflect.ValueOf((*vcl.TSelectCellEvent)(nil)),
		"TSetEditEvent":                     reflect.ValueOf((*vcl.TSetEditEvent)(nil)),
		"TShape":                            reflect.ValueOf((*vcl.TShape)(nil)),
		"TShortCutEvent":                    reflect.ValueOf((*vcl.TShortCutEvent)(nil)),
		"TSizeConstraints":                  reflect.ValueOf((*vcl.TSizeConstraints)(nil)),
		"TSpeedButton":                      reflect.ValueOf((*vcl.TSpeedButton)(nil)),
		"TSpinEdit":                         reflect.ValueOf((*vcl.TSpinEdit)(nil)),
		"TSplitter":                         reflect.ValueOf((*vcl.TSplitter)(nil)),
		"TStartDockEvent":                   reflect.ValueOf((*vcl.TStartDockEvent)(nil)),
		"TStartDragEvent":                   reflect.ValueOf((*vcl.TStartDragEvent)(nil)),
		"TStaticText":                       reflect.ValueOf((*vcl.TStaticText)(nil)),
		"TStatusBar":                        reflect.ValueOf((*vcl.TStatusBar)(nil)),
		"TStatusPanel":                      reflect.ValueOf((*vcl.TStatusPanel)(nil)),
		"TStatusPanels":                     reflect.ValueOf((*vcl.TStatusPanels)(nil)),
		"TStringGrid":                       reflect.ValueOf((*vcl.TStringGrid)(nil)),
		"TStringList":                       reflect.ValueOf((*vcl.TStringList)(nil)),
		"TStrings":                          reflect.ValueOf((*vcl.TStrings)(nil)),
		"TSysLinkEvent":                     reflect.ValueOf((*vcl.TSysLinkEvent)(nil)),
		"TTBAdvancedCustomDrawBtnEvent":     reflect.ValueOf((*vcl.TTBAdvancedCustomDrawBtnEvent)(nil)),
		"TTBAdvancedCustomDrawEvent":        reflect.ValueOf((*vcl.TTBAdvancedCustomDrawEvent)(nil)),
		"TTVAdvancedCustomDrawEvent":        reflect.ValueOf((*vcl.TTVAdvancedCustomDrawEvent)(nil)),
		"TTVAdvancedCustomDrawItemEvent":    reflect.ValueOf((*vcl.TTVAdvancedCustomDrawItemEvent)(nil)),
		"TTVChangedEvent":                   reflect.ValueOf((*vcl.TTVChangedEvent)(nil)),
		"TTVChangingEvent":                  reflect.ValueOf((*vcl.TTVChangingEvent)(nil)),
		"TTVCollapsingEvent":                reflect.ValueOf((*vcl.TTVCollapsingEvent)(nil)),
		"TTVCompareEvent":                   reflect.ValueOf((*vcl.TTVCompareEvent)(nil)),
		"TTVCustomDrawEvent":                reflect.ValueOf((*vcl.TTVCustomDrawEvent)(nil)),
		"TTVCustomDrawItemEvent":            reflect.ValueOf((*vcl.TTVCustomDrawItemEvent)(nil)),
		"TTVEditedEvent":                    reflect.ValueOf((*vcl.TTVEditedEvent)(nil)),
		"TTVEditingEvent":                   reflect.ValueOf((*vcl.TTVEditingEvent)(nil)),
		"TTVExpandedEvent":                  reflect.ValueOf((*vcl.TTVExpandedEvent)(nil)),
		"TTVExpandingEvent":                 reflect.ValueOf((*vcl.TTVExpandingEvent)(nil)),
		"TTVHintEvent":                      reflect.ValueOf((*vcl.TTVHintEvent)(nil)),
		"TTabChangingEvent":                 reflect.ValueOf((*vcl.TTabChangingEvent)(nil)),
		"TTabGetImageEvent":                 reflect.ValueOf((*vcl.TTabGetImageEvent)(nil)),
		"TTabSheet":                         reflect.ValueOf((*vcl.TTabSheet)(nil)),
		"TTaskDialog":                       reflect.ValueOf((*vcl.TTaskDialog)(nil)),
		"TTaskDialogBaseButtonItem":         reflect.ValueOf((*vcl.TTaskDialogBaseButtonItem)(nil)),
		"TTaskDialogButtonItem":             reflect.ValueOf((*vcl.TTaskDialogButtonItem)(nil)),
		"TTaskDialogButtons":                reflect.ValueOf((*vcl.TTaskDialogButtons)(nil)),
		"TTaskDialogRadioButtonItem":        reflect.ValueOf((*vcl.TTaskDialogRadioButtonItem)(nil)),
		"TTaskDlgClickEvent":                reflect.ValueOf((*vcl.TTaskDlgClickEvent)(nil)),
		"TTaskDlgTimerEvent":                reflect.ValueOf((*vcl.TTaskDlgTimerEvent)(nil)),
		"TThreadProc":                       reflect.ValueOf((*vcl.TThreadProc)(nil)),
		"TThumbButtonNotifyEvent":           reflect.ValueOf((*vcl.TThumbButtonNotifyEvent)(nil)),
		"TThumbPreviewItemRequestEvent":     reflect.ValueOf((*vcl.TThumbPreviewItemRequestEvent)(nil)),
		"TTimer":                            reflect.ValueOf((*vcl.TTimer)(nil)),
		"TToolBar":                          reflect.ValueOf((*vcl.TToolBar)(nil)),
		"TToolButton":                       reflect.ValueOf((*vcl.TToolButton)(nil)),
		"TTrackBar":                         reflect.ValueOf((*vcl.TTrackBar)(nil)),
		"TTrayIcon":                         reflect.ValueOf((*vcl.TTrayIcon)(nil)),
		"TTreeNode":                         reflect.ValueOf((*vcl.TTreeNode)(nil)),
		"TTreeNodes":                        reflect.ValueOf((*vcl.TTreeNodes)(nil)),
		"TTreeView":                         reflect.ValueOf((*vcl.TTreeView)(nil)),
		"TUDChangingEvent":                  reflect.ValueOf((*vcl.TUDChangingEvent)(nil)),
		"TUDClickEvent":                     reflect.ValueOf((*vcl.TUDClickEvent)(nil)),
		"TUnDockEvent":                      reflect.ValueOf((*vcl.TUnDockEvent)(nil)),
		"TUpDown":                           reflect.ValueOf((*vcl.TUpDown)(nil)),
		"TValueListEditor":                  reflect.ValueOf((*vcl.TValueListEditor)(nil)),
		"TWebJSExternalEvent":               reflect.ValueOf((*vcl.TWebJSExternalEvent)(nil)),
		"TWebTitleChangeEvent":              reflect.ValueOf((*vcl.TWebTitleChangeEvent)(nil)),
		"TWinControl":                       reflect.ValueOf((*vcl.TWinControl)(nil)),
		"TWindowPreviewItemRequestEvent":    reflect.ValueOf((*vcl.TWindowPreviewItemRequestEvent)(nil)),
		"TWndProcEvent":                     reflect.ValueOf((*vcl.TWndProcEvent)(nil)),
		"TXButton":                          reflect.ValueOf((*vcl.TXButton)(nil)),
		"Window":                            reflect.ValueOf((*vcl.Window)(nil)),

		// interface wrapper definitions
		"_IComponent":  reflect.ValueOf((*_github_com_ying32_govcl_vcl_IComponent)(nil)),
		"_IControl":    reflect.ValueOf((*_github_com_ying32_govcl_vcl_IControl)(nil)),
		"_IObject":     reflect.ValueOf((*_github_com_ying32_govcl_vcl_IObject)(nil)),
		"_IWinControl": reflect.ValueOf((*_github_com_ying32_govcl_vcl_IWinControl)(nil)),
	}
}

// _github_com_ying32_govcl_vcl_IComponent is an interface wrapper for IComponent type
type _github_com_ying32_govcl_vcl_IComponent struct {
	WAs                func() vcl.TAs
	WAssign            func(a0 vcl.IObject)
	WClassName         func() string
	WClassType         func() types.TClass
	WComponentCount    func() int32
	WComponentIndex    func() int32
	WComponents        func(a0 int32) *vcl.TComponent
	WEquals            func(a0 vcl.IObject) bool
	WFindComponent     func(a0 string) *vcl.TComponent
	WFree              func()
	WGetHashCode       func() int32
	WGetNamePath       func() string
	WHasParent         func() bool
	WInheritsFrom      func(a0 types.TClass) bool
	WInstance          func() uintptr
	WInstanceSize      func() int32
	WIs                func() vcl.TIs
	WIsValid           func() bool
	WName              func() string
	WOwner             func() *vcl.TComponent
	WSetComponentIndex func(a0 int32)
	WSetName           func(a0 string)
	WSetTag            func(a0 int)
	WTag               func() int
	WToString          func() string
	WUnsafeAddr        func() unsafe.Pointer
}

func (W _github_com_ying32_govcl_vcl_IComponent) As() vcl.TAs             { return W.WAs() }
func (W _github_com_ying32_govcl_vcl_IComponent) Assign(a0 vcl.IObject)   { W.WAssign(a0) }
func (W _github_com_ying32_govcl_vcl_IComponent) ClassName() string       { return W.WClassName() }
func (W _github_com_ying32_govcl_vcl_IComponent) ClassType() types.TClass { return W.WClassType() }
func (W _github_com_ying32_govcl_vcl_IComponent) ComponentCount() int32   { return W.WComponentCount() }
func (W _github_com_ying32_govcl_vcl_IComponent) ComponentIndex() int32   { return W.WComponentIndex() }
func (W _github_com_ying32_govcl_vcl_IComponent) Components(a0 int32) *vcl.TComponent {
	return W.WComponents(a0)
}
func (W _github_com_ying32_govcl_vcl_IComponent) Equals(a0 vcl.IObject) bool { return W.WEquals(a0) }
func (W _github_com_ying32_govcl_vcl_IComponent) FindComponent(a0 string) *vcl.TComponent {
	return W.WFindComponent(a0)
}
func (W _github_com_ying32_govcl_vcl_IComponent) Free()               { W.WFree() }
func (W _github_com_ying32_govcl_vcl_IComponent) GetHashCode() int32  { return W.WGetHashCode() }
func (W _github_com_ying32_govcl_vcl_IComponent) GetNamePath() string { return W.WGetNamePath() }
func (W _github_com_ying32_govcl_vcl_IComponent) HasParent() bool     { return W.WHasParent() }
func (W _github_com_ying32_govcl_vcl_IComponent) InheritsFrom(a0 types.TClass) bool {
	return W.WInheritsFrom(a0)
}
func (W _github_com_ying32_govcl_vcl_IComponent) Instance() uintptr      { return W.WInstance() }
func (W _github_com_ying32_govcl_vcl_IComponent) InstanceSize() int32    { return W.WInstanceSize() }
func (W _github_com_ying32_govcl_vcl_IComponent) Is() vcl.TIs            { return W.WIs() }
func (W _github_com_ying32_govcl_vcl_IComponent) IsValid() bool          { return W.WIsValid() }
func (W _github_com_ying32_govcl_vcl_IComponent) Name() string           { return W.WName() }
func (W _github_com_ying32_govcl_vcl_IComponent) Owner() *vcl.TComponent { return W.WOwner() }
func (W _github_com_ying32_govcl_vcl_IComponent) SetComponentIndex(a0 int32) {
	W.WSetComponentIndex(a0)
}
func (W _github_com_ying32_govcl_vcl_IComponent) SetName(a0 string)          { W.WSetName(a0) }
func (W _github_com_ying32_govcl_vcl_IComponent) SetTag(a0 int)              { W.WSetTag(a0) }
func (W _github_com_ying32_govcl_vcl_IComponent) Tag() int                   { return W.WTag() }
func (W _github_com_ying32_govcl_vcl_IComponent) ToString() string           { return W.WToString() }
func (W _github_com_ying32_govcl_vcl_IComponent) UnsafeAddr() unsafe.Pointer { return W.WUnsafeAddr() }

// _github_com_ying32_govcl_vcl_IControl is an interface wrapper for IControl type
type _github_com_ying32_govcl_vcl_IControl struct {
	WAction            func() *vcl.TAction
	WAlign             func() types.TAlign
	WAnchors           func() types.TSet
	WAs                func() vcl.TAs
	WAssign            func(a0 vcl.IObject)
	WBoundsRect        func() types.TRect
	WCaption           func() string
	WClassName         func() string
	WClassType         func() types.TClass
	WClientHeight      func() int32
	WClientOrigin      func() types.TPoint
	WClientRect        func() types.TRect
	WClientWidth       func() int32
	WColor             func() types.TColor
	WComponentCount    func() int32
	WComponentIndex    func() int32
	WComponents        func(a0 int32) *vcl.TComponent
	WControlState      func() types.TSet
	WControlStyle      func() types.TSet
	WCursor            func() types.TCursor
	WEnabled           func() bool
	WEquals            func(a0 vcl.IObject) bool
	WFindComponent     func(a0 string) *vcl.TComponent
	WFloating          func() bool
	WFont              func() *vcl.TFont
	WFree              func()
	WGetHashCode       func() int32
	WGetNamePath       func() string
	WHasParent         func() bool
	WHeight            func() int32
	WHide              func()
	WHint              func() string
	WInheritsFrom      func(a0 types.TClass) bool
	WInstance          func() uintptr
	WInstanceSize      func() int32
	WInvalidate        func()
	WIs                func() vcl.TIs
	WIsValid           func() bool
	WLeft              func() int32
	WMargins           func() *vcl.TMargins
	WName              func() string
	WOwner             func() *vcl.TComponent
	WParent            func() *vcl.TWinControl
	WParentToClient    func(Point types.TPoint, AParent vcl.IWinControl) types.TPoint
	WPerform           func(a0 uint32, a1 uintptr, a2 int) int
	WPopupMenu         func() *vcl.TPopupMenu
	WRealign           func()
	WRefresh           func()
	WRepaint           func()
	WScreenToClient    func(Point types.TPoint) types.TPoint
	WSendToBack        func()
	WSetAction         func(a0 vcl.IComponent)
	WSetAlign          func(a0 types.TAlign)
	WSetAnchors        func(a0 types.TSet)
	WSetBounds         func(a0 int32, a1 int32, a2 int32, a3 int32)
	WSetBoundsRect     func(a0 types.TRect)
	WSetCaption        func(a0 string)
	WSetClientHeight   func(a0 int32)
	WSetClientWidth    func(a0 int32)
	WSetColor          func(a0 types.TColor)
	WSetComponentIndex func(a0 int32)
	WSetControlState   func(a0 types.TSet)
	WSetControlStyle   func(a0 types.TSet)
	WSetCursor         func(a0 types.TCursor)
	WSetEnabled        func(a0 bool)
	WSetFloating       func(a0 bool)
	WSetFont           func(a0 *vcl.TFont)
	WSetHeight         func(a0 int32)
	WSetHint           func(a0 string)
	WSetLeft           func(a0 int32)
	WSetMargins        func(a0 *vcl.TMargins)
	WSetName           func(a0 string)
	WSetParent         func(a0 vcl.IWinControl)
	WSetPopupMenu      func(a0 vcl.IComponent)
	WSetShowHint       func(a0 bool)
	WSetTag            func(a0 int)
	WSetTextBuf        func(a0 string)
	WSetTop            func(a0 int32)
	WSetVisible        func(a0 bool)
	WSetWidth          func(a0 int32)
	WShow              func()
	WShowHint          func() bool
	WTag               func() int
	WToString          func() string
	WTop               func() int32
	WUnsafeAddr        func() unsafe.Pointer
	WVisible           func() bool
	WWidth             func() int32
}

func (W _github_com_ying32_govcl_vcl_IControl) Action() *vcl.TAction       { return W.WAction() }
func (W _github_com_ying32_govcl_vcl_IControl) Align() types.TAlign        { return W.WAlign() }
func (W _github_com_ying32_govcl_vcl_IControl) Anchors() types.TSet        { return W.WAnchors() }
func (W _github_com_ying32_govcl_vcl_IControl) As() vcl.TAs                { return W.WAs() }
func (W _github_com_ying32_govcl_vcl_IControl) Assign(a0 vcl.IObject)      { W.WAssign(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) BoundsRect() types.TRect    { return W.WBoundsRect() }
func (W _github_com_ying32_govcl_vcl_IControl) Caption() string            { return W.WCaption() }
func (W _github_com_ying32_govcl_vcl_IControl) ClassName() string          { return W.WClassName() }
func (W _github_com_ying32_govcl_vcl_IControl) ClassType() types.TClass    { return W.WClassType() }
func (W _github_com_ying32_govcl_vcl_IControl) ClientHeight() int32        { return W.WClientHeight() }
func (W _github_com_ying32_govcl_vcl_IControl) ClientOrigin() types.TPoint { return W.WClientOrigin() }
func (W _github_com_ying32_govcl_vcl_IControl) ClientRect() types.TRect    { return W.WClientRect() }
func (W _github_com_ying32_govcl_vcl_IControl) ClientWidth() int32         { return W.WClientWidth() }
func (W _github_com_ying32_govcl_vcl_IControl) Color() types.TColor        { return W.WColor() }
func (W _github_com_ying32_govcl_vcl_IControl) ComponentCount() int32      { return W.WComponentCount() }
func (W _github_com_ying32_govcl_vcl_IControl) ComponentIndex() int32      { return W.WComponentIndex() }
func (W _github_com_ying32_govcl_vcl_IControl) Components(a0 int32) *vcl.TComponent {
	return W.WComponents(a0)
}
func (W _github_com_ying32_govcl_vcl_IControl) ControlState() types.TSet   { return W.WControlState() }
func (W _github_com_ying32_govcl_vcl_IControl) ControlStyle() types.TSet   { return W.WControlStyle() }
func (W _github_com_ying32_govcl_vcl_IControl) Cursor() types.TCursor      { return W.WCursor() }
func (W _github_com_ying32_govcl_vcl_IControl) Enabled() bool              { return W.WEnabled() }
func (W _github_com_ying32_govcl_vcl_IControl) Equals(a0 vcl.IObject) bool { return W.WEquals(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) FindComponent(a0 string) *vcl.TComponent {
	return W.WFindComponent(a0)
}
func (W _github_com_ying32_govcl_vcl_IControl) Floating() bool      { return W.WFloating() }
func (W _github_com_ying32_govcl_vcl_IControl) Font() *vcl.TFont    { return W.WFont() }
func (W _github_com_ying32_govcl_vcl_IControl) Free()               { W.WFree() }
func (W _github_com_ying32_govcl_vcl_IControl) GetHashCode() int32  { return W.WGetHashCode() }
func (W _github_com_ying32_govcl_vcl_IControl) GetNamePath() string { return W.WGetNamePath() }
func (W _github_com_ying32_govcl_vcl_IControl) HasParent() bool     { return W.WHasParent() }
func (W _github_com_ying32_govcl_vcl_IControl) Height() int32       { return W.WHeight() }
func (W _github_com_ying32_govcl_vcl_IControl) Hide()               { W.WHide() }
func (W _github_com_ying32_govcl_vcl_IControl) Hint() string        { return W.WHint() }
func (W _github_com_ying32_govcl_vcl_IControl) InheritsFrom(a0 types.TClass) bool {
	return W.WInheritsFrom(a0)
}
func (W _github_com_ying32_govcl_vcl_IControl) Instance() uintptr        { return W.WInstance() }
func (W _github_com_ying32_govcl_vcl_IControl) InstanceSize() int32      { return W.WInstanceSize() }
func (W _github_com_ying32_govcl_vcl_IControl) Invalidate()              { W.WInvalidate() }
func (W _github_com_ying32_govcl_vcl_IControl) Is() vcl.TIs              { return W.WIs() }
func (W _github_com_ying32_govcl_vcl_IControl) IsValid() bool            { return W.WIsValid() }
func (W _github_com_ying32_govcl_vcl_IControl) Left() int32              { return W.WLeft() }
func (W _github_com_ying32_govcl_vcl_IControl) Margins() *vcl.TMargins   { return W.WMargins() }
func (W _github_com_ying32_govcl_vcl_IControl) Name() string             { return W.WName() }
func (W _github_com_ying32_govcl_vcl_IControl) Owner() *vcl.TComponent   { return W.WOwner() }
func (W _github_com_ying32_govcl_vcl_IControl) Parent() *vcl.TWinControl { return W.WParent() }
func (W _github_com_ying32_govcl_vcl_IControl) ParentToClient(Point types.TPoint, AParent vcl.IWinControl) types.TPoint {
	return W.WParentToClient(Point, AParent)
}
func (W _github_com_ying32_govcl_vcl_IControl) Perform(a0 uint32, a1 uintptr, a2 int) int {
	return W.WPerform(a0, a1, a2)
}
func (W _github_com_ying32_govcl_vcl_IControl) PopupMenu() *vcl.TPopupMenu { return W.WPopupMenu() }
func (W _github_com_ying32_govcl_vcl_IControl) Realign()                   { W.WRealign() }
func (W _github_com_ying32_govcl_vcl_IControl) Refresh()                   { W.WRefresh() }
func (W _github_com_ying32_govcl_vcl_IControl) Repaint()                   { W.WRepaint() }
func (W _github_com_ying32_govcl_vcl_IControl) ScreenToClient(Point types.TPoint) types.TPoint {
	return W.WScreenToClient(Point)
}
func (W _github_com_ying32_govcl_vcl_IControl) SendToBack()                 { W.WSendToBack() }
func (W _github_com_ying32_govcl_vcl_IControl) SetAction(a0 vcl.IComponent) { W.WSetAction(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetAlign(a0 types.TAlign)    { W.WSetAlign(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetAnchors(a0 types.TSet)    { W.WSetAnchors(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetBounds(a0 int32, a1 int32, a2 int32, a3 int32) {
	W.WSetBounds(a0, a1, a2, a3)
}
func (W _github_com_ying32_govcl_vcl_IControl) SetBoundsRect(a0 types.TRect)   { W.WSetBoundsRect(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetCaption(a0 string)           { W.WSetCaption(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetClientHeight(a0 int32)       { W.WSetClientHeight(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetClientWidth(a0 int32)        { W.WSetClientWidth(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetColor(a0 types.TColor)       { W.WSetColor(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetComponentIndex(a0 int32)     { W.WSetComponentIndex(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetControlState(a0 types.TSet)  { W.WSetControlState(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetControlStyle(a0 types.TSet)  { W.WSetControlStyle(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetCursor(a0 types.TCursor)     { W.WSetCursor(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetEnabled(a0 bool)             { W.WSetEnabled(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetFloating(a0 bool)            { W.WSetFloating(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetFont(a0 *vcl.TFont)          { W.WSetFont(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetHeight(a0 int32)             { W.WSetHeight(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetHint(a0 string)              { W.WSetHint(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetLeft(a0 int32)               { W.WSetLeft(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetMargins(a0 *vcl.TMargins)    { W.WSetMargins(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetName(a0 string)              { W.WSetName(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetParent(a0 vcl.IWinControl)   { W.WSetParent(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetPopupMenu(a0 vcl.IComponent) { W.WSetPopupMenu(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetShowHint(a0 bool)            { W.WSetShowHint(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetTag(a0 int)                  { W.WSetTag(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetTextBuf(a0 string)           { W.WSetTextBuf(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetTop(a0 int32)                { W.WSetTop(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetVisible(a0 bool)             { W.WSetVisible(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) SetWidth(a0 int32)              { W.WSetWidth(a0) }
func (W _github_com_ying32_govcl_vcl_IControl) Show()                          { W.WShow() }
func (W _github_com_ying32_govcl_vcl_IControl) ShowHint() bool                 { return W.WShowHint() }
func (W _github_com_ying32_govcl_vcl_IControl) Tag() int                       { return W.WTag() }
func (W _github_com_ying32_govcl_vcl_IControl) ToString() string               { return W.WToString() }
func (W _github_com_ying32_govcl_vcl_IControl) Top() int32                     { return W.WTop() }
func (W _github_com_ying32_govcl_vcl_IControl) UnsafeAddr() unsafe.Pointer     { return W.WUnsafeAddr() }
func (W _github_com_ying32_govcl_vcl_IControl) Visible() bool                  { return W.WVisible() }
func (W _github_com_ying32_govcl_vcl_IControl) Width() int32                   { return W.WWidth() }

// _github_com_ying32_govcl_vcl_IObject is an interface wrapper for IObject type
type _github_com_ying32_govcl_vcl_IObject struct {
	WAs           func() vcl.TAs
	WClassName    func() string
	WClassType    func() types.TClass
	WEquals       func(a0 vcl.IObject) bool
	WFree         func()
	WGetHashCode  func() int32
	WInheritsFrom func(a0 types.TClass) bool
	WInstance     func() uintptr
	WInstanceSize func() int32
	WIs           func() vcl.TIs
	WIsValid      func() bool
	WToString     func() string
	WUnsafeAddr   func() unsafe.Pointer
}

func (W _github_com_ying32_govcl_vcl_IObject) As() vcl.TAs                { return W.WAs() }
func (W _github_com_ying32_govcl_vcl_IObject) ClassName() string          { return W.WClassName() }
func (W _github_com_ying32_govcl_vcl_IObject) ClassType() types.TClass    { return W.WClassType() }
func (W _github_com_ying32_govcl_vcl_IObject) Equals(a0 vcl.IObject) bool { return W.WEquals(a0) }
func (W _github_com_ying32_govcl_vcl_IObject) Free()                      { W.WFree() }
func (W _github_com_ying32_govcl_vcl_IObject) GetHashCode() int32         { return W.WGetHashCode() }
func (W _github_com_ying32_govcl_vcl_IObject) InheritsFrom(a0 types.TClass) bool {
	return W.WInheritsFrom(a0)
}
func (W _github_com_ying32_govcl_vcl_IObject) Instance() uintptr          { return W.WInstance() }
func (W _github_com_ying32_govcl_vcl_IObject) InstanceSize() int32        { return W.WInstanceSize() }
func (W _github_com_ying32_govcl_vcl_IObject) Is() vcl.TIs                { return W.WIs() }
func (W _github_com_ying32_govcl_vcl_IObject) IsValid() bool              { return W.WIsValid() }
func (W _github_com_ying32_govcl_vcl_IObject) ToString() string           { return W.WToString() }
func (W _github_com_ying32_govcl_vcl_IObject) UnsafeAddr() unsafe.Pointer { return W.WUnsafeAddr() }

// _github_com_ying32_govcl_vcl_IWinControl is an interface wrapper for IWinControl type
type _github_com_ying32_govcl_vcl_IWinControl struct {
	WAction                  func() *vcl.TAction
	WAlign                   func() types.TAlign
	WAlignDisabled           func() bool
	WAnchors                 func() types.TSet
	WAs                      func() vcl.TAs
	WAssign                  func(a0 vcl.IObject)
	WBoundsRect              func() types.TRect
	WBrush                   func() *vcl.TBrush
	WCanFocus                func() bool
	WCaption                 func() string
	WClassName               func() string
	WClassType               func() types.TClass
	WClientHeight            func() int32
	WClientOrigin            func() types.TPoint
	WClientRect              func() types.TRect
	WClientToParent          func(Point types.TPoint, AParent vcl.IWinControl) types.TPoint
	WClientToScreen          func(Point types.TPoint) types.TPoint
	WClientWidth             func() int32
	WColor                   func() types.TColor
	WComponentCount          func() int32
	WComponentIndex          func() int32
	WComponents              func(a0 int32) *vcl.TComponent
	WContainsControl         func(a0 vcl.IControl) bool
	WControlCount            func() int32
	WControlState            func() types.TSet
	WControlStyle            func() types.TSet
	WControls                func(index int32) *vcl.TControl
	WCursor                  func() types.TCursor
	WDisableAlign            func()
	WDockClientCount         func() int32
	WDockClients             func(a0 int32) *vcl.TControl
	WDoubleBuffered          func() bool
	WEnableAlign             func()
	WEnabled                 func() bool
	WEquals                  func(a0 vcl.IObject) bool
	WFindChildControl        func(a0 string) *vcl.TControl
	WFindComponent           func(a0 string) *vcl.TComponent
	WFlipChildren            func(a0 bool)
	WFloating                func() bool
	WFocused                 func() bool
	WFont                    func() *vcl.TFont
	WFree                    func()
	WGetHashCode             func() int32
	WGetNamePath             func() string
	WHandle                  func() uintptr
	WHandleAllocated         func() bool
	WHasParent               func() bool
	WHeight                  func() int32
	WHide                    func()
	WHint                    func() string
	WInheritsFrom            func(a0 types.TClass) bool
	WInsertControl           func(a0 vcl.IControl)
	WInstance                func() uintptr
	WInstanceSize            func() int32
	WInvalidate              func()
	WIs                      func() vcl.TIs
	WIsValid                 func() bool
	WLeft                    func() int32
	WMargins                 func() *vcl.TMargins
	WMouseInClient           func() bool
	WName                    func() string
	WOwner                   func() *vcl.TComponent
	WPaintTo                 func(DC uintptr, X int32, Y int32)
	WParent                  func() *vcl.TWinControl
	WParentDoubleBuffered    func() bool
	WParentToClient          func(Point types.TPoint, AParent vcl.IWinControl) types.TPoint
	WParentWindow            func() uintptr
	WPerform                 func(a0 uint32, a1 uintptr, a2 int) int
	WPopupMenu               func() *vcl.TPopupMenu
	WRealign                 func()
	WRefresh                 func()
	WRemoveControl           func(a0 vcl.IControl)
	WRepaint                 func()
	WScaleBy                 func(M int32, D int32)
	WScreenToClient          func(Point types.TPoint) types.TPoint
	WScrollBy                func(DeltaX int32, DeltaY int32)
	WSendToBack              func()
	WSetAction               func(a0 vcl.IComponent)
	WSetAlign                func(a0 types.TAlign)
	WSetAnchors              func(a0 types.TSet)
	WSetBounds               func(a0 int32, a1 int32, a2 int32, a3 int32)
	WSetBoundsRect           func(a0 types.TRect)
	WSetCaption              func(a0 string)
	WSetClientHeight         func(a0 int32)
	WSetClientWidth          func(a0 int32)
	WSetColor                func(a0 types.TColor)
	WSetComponentIndex       func(a0 int32)
	WSetControlState         func(a0 types.TSet)
	WSetControlStyle         func(a0 types.TSet)
	WSetCursor               func(a0 types.TCursor)
	WSetDoubleBuffered       func(a0 bool)
	WSetEnabled              func(a0 bool)
	WSetFloating             func(a0 bool)
	WSetFocus                func()
	WSetFont                 func(a0 *vcl.TFont)
	WSetHandle               func(a0 uintptr)
	WSetHeight               func(a0 int32)
	WSetHint                 func(a0 string)
	WSetLeft                 func(a0 int32)
	WSetMargins              func(a0 *vcl.TMargins)
	WSetName                 func(a0 string)
	WSetParent               func(a0 vcl.IWinControl)
	WSetParentDoubleBuffered func(a0 bool)
	WSetParentWindow         func(a0 uintptr)
	WSetPopupMenu            func(a0 vcl.IComponent)
	WSetShowHint             func(a0 bool)
	WSetTabOrder             func(a0 int16)
	WSetTabStop              func(a0 bool)
	WSetTag                  func(a0 int)
	WSetTextBuf              func(a0 string)
	WSetTop                  func(a0 int32)
	WSetUseDockManager       func(value bool)
	WSetVisible              func(a0 bool)
	WSetWidth                func(a0 int32)
	WShow                    func()
	WShowHint                func() bool
	WShowing                 func() bool
	WTabOrder                func() int16
	WTabStop                 func() bool
	WTag                     func() int
	WToString                func() string
	WTop                     func() int32
	WUnsafeAddr              func() unsafe.Pointer
	WUpdate                  func()
	WUpdateControlState      func()
	WUseDockManager          func() bool
	WVisible                 func() bool
	WWidth                   func() int32
}

func (W _github_com_ying32_govcl_vcl_IWinControl) Action() *vcl.TAction    { return W.WAction() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Align() types.TAlign     { return W.WAlign() }
func (W _github_com_ying32_govcl_vcl_IWinControl) AlignDisabled() bool     { return W.WAlignDisabled() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Anchors() types.TSet     { return W.WAnchors() }
func (W _github_com_ying32_govcl_vcl_IWinControl) As() vcl.TAs             { return W.WAs() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Assign(a0 vcl.IObject)   { W.WAssign(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) BoundsRect() types.TRect { return W.WBoundsRect() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Brush() *vcl.TBrush      { return W.WBrush() }
func (W _github_com_ying32_govcl_vcl_IWinControl) CanFocus() bool          { return W.WCanFocus() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Caption() string         { return W.WCaption() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ClassName() string       { return W.WClassName() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ClassType() types.TClass { return W.WClassType() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ClientHeight() int32     { return W.WClientHeight() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ClientOrigin() types.TPoint {
	return W.WClientOrigin()
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ClientRect() types.TRect { return W.WClientRect() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ClientToParent(Point types.TPoint, AParent vcl.IWinControl) types.TPoint {
	return W.WClientToParent(Point, AParent)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ClientToScreen(Point types.TPoint) types.TPoint {
	return W.WClientToScreen(Point)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ClientWidth() int32    { return W.WClientWidth() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Color() types.TColor   { return W.WColor() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ComponentCount() int32 { return W.WComponentCount() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ComponentIndex() int32 { return W.WComponentIndex() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Components(a0 int32) *vcl.TComponent {
	return W.WComponents(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ContainsControl(a0 vcl.IControl) bool {
	return W.WContainsControl(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ControlCount() int32      { return W.WControlCount() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ControlState() types.TSet { return W.WControlState() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ControlStyle() types.TSet { return W.WControlStyle() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Controls(index int32) *vcl.TControl {
	return W.WControls(index)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) Cursor() types.TCursor { return W.WCursor() }
func (W _github_com_ying32_govcl_vcl_IWinControl) DisableAlign()         { W.WDisableAlign() }
func (W _github_com_ying32_govcl_vcl_IWinControl) DockClientCount() int32 {
	return W.WDockClientCount()
}
func (W _github_com_ying32_govcl_vcl_IWinControl) DockClients(a0 int32) *vcl.TControl {
	return W.WDockClients(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) DoubleBuffered() bool       { return W.WDoubleBuffered() }
func (W _github_com_ying32_govcl_vcl_IWinControl) EnableAlign()               { W.WEnableAlign() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Enabled() bool              { return W.WEnabled() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Equals(a0 vcl.IObject) bool { return W.WEquals(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) FindChildControl(a0 string) *vcl.TControl {
	return W.WFindChildControl(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) FindComponent(a0 string) *vcl.TComponent {
	return W.WFindComponent(a0)
}

func (W _github_com_ying32_govcl_vcl_IWinControl) FlipChildren(a0 bool)  { W.WFlipChildren(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) Floating() bool        { return W.WFloating() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Focused() bool         { return W.WFocused() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Font() *vcl.TFont      { return W.WFont() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Free()                 { W.WFree() }
func (W _github_com_ying32_govcl_vcl_IWinControl) GetHashCode() int32    { return W.WGetHashCode() }
func (W _github_com_ying32_govcl_vcl_IWinControl) GetNamePath() string   { return W.WGetNamePath() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Handle() uintptr       { return W.WHandle() }
func (W _github_com_ying32_govcl_vcl_IWinControl) HandleAllocated() bool { return W.WHandleAllocated() }
func (W _github_com_ying32_govcl_vcl_IWinControl) HasParent() bool       { return W.WHasParent() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Height() int32         { return W.WHeight() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Hide()                 { W.WHide() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Hint() string          { return W.WHint() }
func (W _github_com_ying32_govcl_vcl_IWinControl) InheritsFrom(a0 types.TClass) bool {
	return W.WInheritsFrom(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) InsertControl(a0 vcl.IControl) {
	W.WInsertControl(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) Instance() uintptr      { return W.WInstance() }
func (W _github_com_ying32_govcl_vcl_IWinControl) InstanceSize() int32    { return W.WInstanceSize() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Invalidate()            { W.WInvalidate() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Is() vcl.TIs            { return W.WIs() }
func (W _github_com_ying32_govcl_vcl_IWinControl) IsValid() bool          { return W.WIsValid() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Left() int32            { return W.WLeft() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Margins() *vcl.TMargins { return W.WMargins() }
func (W _github_com_ying32_govcl_vcl_IWinControl) MouseInClient() bool    { return W.WMouseInClient() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Name() string           { return W.WName() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Owner() *vcl.TComponent { return W.WOwner() }
func (W _github_com_ying32_govcl_vcl_IWinControl) PaintTo(DC uintptr, X int32, Y int32) {
	W.WPaintTo(DC, X, Y)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) Parent() *vcl.TWinControl { return W.WParent() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ParentDoubleBuffered() bool {
	return W.WParentDoubleBuffered()
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ParentToClient(Point types.TPoint, AParent vcl.IWinControl) types.TPoint {
	return W.WParentToClient(Point, AParent)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ParentWindow() uintptr { return W.WParentWindow() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Perform(a0 uint32, a1 uintptr, a2 int) int {
	return W.WPerform(a0, a1, a2)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) PopupMenu() *vcl.TPopupMenu { return W.WPopupMenu() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Realign()                   { W.WRealign() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Refresh()                   { W.WRefresh() }
func (W _github_com_ying32_govcl_vcl_IWinControl) RemoveControl(a0 vcl.IControl) {
	W.WRemoveControl(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) Repaint()                 { W.WRepaint() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ScaleBy(M int32, D int32) { W.WScaleBy(M, D) }
func (W _github_com_ying32_govcl_vcl_IWinControl) ScreenToClient(Point types.TPoint) types.TPoint {
	return W.WScreenToClient(Point)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) ScrollBy(DeltaX int32, DeltaY int32) {
	W.WScrollBy(DeltaX, DeltaY)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SendToBack()                 { W.WSendToBack() }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetAction(a0 vcl.IComponent) { W.WSetAction(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetAlign(a0 types.TAlign)    { W.WSetAlign(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetAnchors(a0 types.TSet)    { W.WSetAnchors(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetBounds(a0 int32, a1 int32, a2 int32, a3 int32) {
	W.WSetBounds(a0, a1, a2, a3)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetBoundsRect(a0 types.TRect) { W.WSetBoundsRect(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetCaption(a0 string)         { W.WSetCaption(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetClientHeight(a0 int32)     { W.WSetClientHeight(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetClientWidth(a0 int32)      { W.WSetClientWidth(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetColor(a0 types.TColor)     { W.WSetColor(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetComponentIndex(a0 int32) {
	W.WSetComponentIndex(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetControlState(a0 types.TSet) {
	W.WSetControlState(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetControlStyle(a0 types.TSet) {
	W.WSetControlStyle(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetCursor(a0 types.TCursor) { W.WSetCursor(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetDoubleBuffered(a0 bool) {
	W.WSetDoubleBuffered(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetEnabled(a0 bool)           { W.WSetEnabled(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetFloating(a0 bool)          { W.WSetFloating(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetFocus()                    { W.WSetFocus() }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetFont(a0 *vcl.TFont)        { W.WSetFont(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetHandle(a0 uintptr)         { W.WSetHandle(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetHeight(a0 int32)           { W.WSetHeight(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetHint(a0 string)            { W.WSetHint(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetLeft(a0 int32)             { W.WSetLeft(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetMargins(a0 *vcl.TMargins)  { W.WSetMargins(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetName(a0 string)            { W.WSetName(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetParent(a0 vcl.IWinControl) { W.WSetParent(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetParentDoubleBuffered(a0 bool) {
	W.WSetParentDoubleBuffered(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetParentWindow(a0 uintptr) { W.WSetParentWindow(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetPopupMenu(a0 vcl.IComponent) {
	W.WSetPopupMenu(a0)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetShowHint(a0 bool)  { W.WSetShowHint(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetTabOrder(a0 int16) { W.WSetTabOrder(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetTabStop(a0 bool)   { W.WSetTabStop(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetTag(a0 int)        { W.WSetTag(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetTextBuf(a0 string) { W.WSetTextBuf(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetTop(a0 int32)      { W.WSetTop(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetUseDockManager(value bool) {
	W.WSetUseDockManager(value)
}
func (W _github_com_ying32_govcl_vcl_IWinControl) SetVisible(a0 bool)         { W.WSetVisible(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) SetWidth(a0 int32)          { W.WSetWidth(a0) }
func (W _github_com_ying32_govcl_vcl_IWinControl) Show()                      { W.WShow() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ShowHint() bool             { return W.WShowHint() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Showing() bool              { return W.WShowing() }
func (W _github_com_ying32_govcl_vcl_IWinControl) TabOrder() int16            { return W.WTabOrder() }
func (W _github_com_ying32_govcl_vcl_IWinControl) TabStop() bool              { return W.WTabStop() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Tag() int                   { return W.WTag() }
func (W _github_com_ying32_govcl_vcl_IWinControl) ToString() string           { return W.WToString() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Top() int32                 { return W.WTop() }
func (W _github_com_ying32_govcl_vcl_IWinControl) UnsafeAddr() unsafe.Pointer { return W.WUnsafeAddr() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Update()                    { W.WUpdate() }
func (W _github_com_ying32_govcl_vcl_IWinControl) UpdateControlState()        { W.WUpdateControlState() }
func (W _github_com_ying32_govcl_vcl_IWinControl) UseDockManager() bool       { return W.WUseDockManager() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Visible() bool              { return W.WVisible() }
func (W _github_com_ying32_govcl_vcl_IWinControl) Width() int32               { return W.WWidth() }
