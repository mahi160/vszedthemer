package models

type VSCodeTheme struct {
	Author string            `json:"author"`
	Name   string            `json:"name"`
	Type   string            `json:"type"`
	Colors VSCodeThemeColors `json:"colors"`
}

type VSCodeThemeColors struct {
	MenubarSelectionForeground                   string `json:"menubar.selectionForeground"`
	MenuSelectionBackground                      string `json:"menu.selectionBackground"`
	MenuSelectionForeground                      string `json:"menu.selectionForeground"`
	MenubarSelectionBackground                   string `json:"menubar.selectionBackground"`
	EditorBackground                             string `json:"editor.background"`
	EditorForeground                             string `json:"editor.foreground"`
	EditorFindMatchBackground                    string `json:"editor.findMatchBackground"`
	EditorFindMatchHighlightBackground           string `json:"editor.findMatchHighlightBackground"`
	EditorFindRangeHighlightBackground           string `json:"editor.findRangeHighlightBackground"`
	EditorHoverHighlightBackground               string `json:"editor.hoverHighlightBackground"`
	EditorInactiveSelectionBackground            string `json:"editor.inactiveSelectionBackground"`
	EditorLineHighlightBackground                string `json:"editor.lineHighlightBackground"`
	EditorLineHighlightBorder                    string `json:"editor.lineHighlightBorder"`
	EditorRangeHighlightBackground               string `json:"editor.rangeHighlightBackground"`
	EditorSelectionBackground                    string `json:"editor.selectionBackground"`
	EditorSelectionHighlightBackground           string `json:"editor.selectionHighlightBackground"`
	EditorWordHighlightStrongBackground          string `json:"editor.wordHighlightStrongBackground"`
	EditorWordHighlightBackground                string `json:"editor.wordHighlightBackground"`
	EditorBracketMatchBackground                 string `json:"editorBracketMatch.background"`
	EditorBracketMatchBorder                     string `json:"editorBracketMatch.border"`
	EditorCodeLensForeground                     string `json:"editorCodeLens.foreground"`
	EditorCursorForeground                       string `json:"editorCursor.foreground"`
	EditorErrorBorder                            string `json:"editorError.border"`
	EditorErrorForeground                        string `json:"editorError.foreground"`
	EditorHintForeground                         string `json:"editorHint.foreground"`
	EditorGutterBackground                       string `json:"editorGutter.background"`
	EditorGutterAddedBackground                  string `json:"editorGutter.addedBackground"`
	EditorGutterDeletedBackground                string `json:"editorGutter.deletedBackground"`
	EditorGutterModifiedBackground               string `json:"editorGutter.modifiedBackground"`
	EditorGroupBackground                        string `json:"editorGroup.background"`
	EditorGroupBorder                            string `json:"editorGroup.border"`
	EditorGroupDropBackground                    string `json:"editorGroup.dropBackground"`
	EditorGroupHeaderNoTabsBackground            string `json:"editorGroupHeader.noTabsBackground"`
	EditorGroupHeaderTabsBackground              string `json:"editorGroupHeader.tabsBackground"`
	EditorGroupHeaderTabsBorder                  string `json:"editorGroupHeader.tabsBorder"`
	EditorHoverWidgetBackground                  string `json:"editorHoverWidget.background"`
	EditorHoverWidgetBorder                      string `json:"editorHoverWidget.border"`
	EditorIndentGuideBackground                  string `json:"editorIndentGuide.background"`
	EditorInlayHintForeground                    string `json:"editorInlayHint.foreground"`
	EditorInlayHintBackground                    string `json:"editorInlayHint.background"`
	EditorLineNumberForeground                   string `json:"editorLineNumber.foreground"`
	EditorLinkActiveForeground                   string `json:"editorLink.activeForeground"`
	EditorMarkerNavigationBackground             string `json:"editorMarkerNavigation.background"`
	EditorMarkerNavigationErrorBackground        string `json:"editorMarkerNavigationError.background"`
	EditorMarkerNavigationWarningBackground      string `json:"editorMarkerNavigationWarning.background"`
	EditorOverviewRulerBorder                    string `json:"editorOverviewRuler.border"`
	EditorOverviewRulerCommonContentForeground   string `json:"editorOverviewRuler.commonContentForeground"`
	EditorOverviewRulerCurrentContentForeground  string `json:"editorOverviewRuler.currentContentForeground"`
	EditorOverviewRulerIncomingContentForeground string `json:"editorOverviewRuler.incomingContentForeground"`
	EditorRulerForeground                        string `json:"editorRuler.foreground"`
	EditorSuggestWidgetBackground                string `json:"editorSuggestWidget.background"`
	EditorSuggestWidgetBorder                    string `json:"editorSuggestWidget.border"`
	EditorSuggestWidgetForeground                string `json:"editorSuggestWidget.foreground"`
	EditorSuggestWidgetHighlightForeground       string `json:"editorSuggestWidget.highlightForeground"`
	EditorSuggestWidgetSelectedBackground        string `json:"editorSuggestWidget.selectedBackground"`
	EditorWarningBorder                          string `json:"editorWarning.border"`
	EditorWarningForeground                      string `json:"editorWarning.foreground"`
	EditorWhitespaceForeground                   string `json:"editorWhitespace.foreground"`
	EditorWidgetBackground                       string `json:"editorWidget.background"`
	EditorWidgetBorder                           string `json:"editorWidget.border"`
	ErrorForeground                              string `json:"errorForeground"`
	ExtensionButtonProminentBackground           string `json:"extensionButton.prominentBackground"`
	ExtensionButtonProminentForeground           string `json:"extensionButton.prominentForeground"`
	ExtensionButtonProminentHoverBackground      string `json:"extensionButton.prominentHoverBackground"`
	FocusBorder                                  string `json:"focusBorder"`
	Foreground                                   string `json:"foreground"`
	InputBackground                              string `json:"input.background"`
	InputBorder                                  string `json:"input.border"`
	InputForeground                              string `json:"input.foreground"`
	InputPlaceholderForeground                   string `json:"input.placeholderForeground"`
	InputOptionActiveBorder                      string `json:"inputOption.activeBorder"`
	InputValidationErrorBackground               string `json:"inputValidation.errorBackground"`
	InputValidationErrorBorder                   string `json:"inputValidation.errorBorder"`
	InputValidationInfoBackground                string `json:"inputValidation.infoBackground"`
	InputValidationInfoBorder                    string `json:"inputValidation.infoBorder"`
	InputValidationWarningBackground             string `json:"inputValidation.warningBackground"`
	InputValidationWarningBorder                 string `json:"inputValidation.warningBorder"`
	ListActiveSelectionBackground                string `json:"list.activeSelectionBackground"`
	ListActiveSelectionForeground                string `json:"list.activeSelectionForeground"`
	ListDropBackground                           string `json:"list.dropBackground"`
	ListFocusBackground                          string `json:"list.focusBackground"`
	ListFocusForeground                          string `json:"list.focusForeground"`
	ListHighlightForeground                      string `json:"list.highlightForeground"`
	ListHoverBackground                          string `json:"list.hoverBackground"`
	ListHoverForeground                          string `json:"list.hoverForeground"`
	ListInactiveSelectionBackground              string `json:"list.inactiveSelectionBackground"`
	ListInactiveSelectionForeground              string `json:"list.inactiveSelectionForeground"`
	MenuBackground                               string `json:"menu.background"`
	MergeBorder                                  string `json:"merge.border"`
	MergeCommonContentBackground                 string `json:"merge.commonContentBackground"`
	MergeCommonHeaderBackground                  string `json:"merge.commonHeaderBackground"`
	MergeCurrentContentBackground                string `json:"merge.currentContentBackground"`
	MergeCurrentHeaderBackground                 string `json:"merge.currentHeaderBackground"`
	MergeIncomingContentBackground               string `json:"merge.incomingContentBackground"`
	MergeIncomingHeaderBackground                string `json:"merge.incomingHeaderBackground"`
	NotificationCenterBorder                     string `json:"notificationCenter.border"`
	NotificationCenterHeaderForeground           string `json:"notificationCenterHeader.foreground"`
	NotificationCenterHeaderBackground           string `json:"notificationCenterHeader.background"`
	NotificationToastBorder                      string `json:"notificationToast.border"`
	NotificationsForeground                      string `json:"notifications.foreground"`
	NotificationsBackground                      string `json:"notifications.background"`
	NotificationsBorder                          string `json:"notifications.border"`
	NotificationLinkForeground                   string `json:"notificationLink.foreground"`
	PanelBackground                              string `json:"panel.background"`
	PanelBorder                                  string `json:"panel.border"`
	PanelTitleActiveBorder                       string `json:"panelTitle.activeBorder"`
	PanelTitleActiveForeground                   string `json:"panelTitle.activeForeground"`
	PanelTitleInactiveForeground                 string `json:"panelTitle.inactiveForeground"`
	PeekViewBorder                               string `json:"peekView.border"`
	PeekViewEditorBackground                     string `json:"peekViewEditor.background"`
	PeekViewEditorMatchHighlightBackground       string `json:"peekViewEditor.matchHighlightBackground"`
	PeekViewEditorGutterBackground               string `json:"peekViewEditorGutter.background"`
	PeekViewResultBackground                     string `json:"peekViewResult.background"`
	PeekViewResultFileForeground                 string `json:"peekViewResult.fileForeground"`
	PeekViewResultLineForeground                 string `json:"peekViewResult.lineForeground"`
	PeekViewResultMatchHighlightBackground       string `json:"peekViewResult.matchHighlightBackground"`
	PeekViewResultSelectionBackground            string `json:"peekViewResult.selectionBackground"`
	PeekViewResultSelectionForeground            string `json:"peekViewResult.selectionForeground"`
	PeekViewTitleBackground                      string `json:"peekViewTitle.background"`
	PeekViewTitleDescriptionForeground           string `json:"peekViewTitleDescription.foreground"`
	PeekViewTitleLabelForeground                 string `json:"peekViewTitleLabel.foreground"`
	PickerGroupBorder                            string `json:"pickerGroup.border"`
	PickerGroupForeground                        string `json:"pickerGroup.foreground"`
	ProgressBarBackground                        string `json:"progressBar.background"`
	ScrollbarShadow                              string `json:"scrollbar.shadow"`
	ScrollbarSliderActiveBackground              string `json:"scrollbarSlider.activeBackground"`
	ScrollbarSliderBackground                    string `json:"scrollbarSlider.background"`
	ScrollbarSliderHoverBackground               string `json:"scrollbarSlider.hoverBackground"`
	SelectionBackground                          string `json:"selection.background"`
	SettingsCheckboxBorder                       string `json:"settings.checkboxBorder"`
	SideBarBackground                            string `json:"sideBar.background"`
	SideBarBorder                                string `json:"sideBar.border"`
	SideBarForeground                            string `json:"sideBar.foreground"`
	SideBarSectionHeaderBackground               string `json:"sideBarSectionHeader.background"`
	SideBarSectionHeaderForeground               string `json:"sideBarSectionHeader.foreground"`
	SideBarTitleForeground                       string `json:"sideBarTitle.foreground"`
	StatusBarBackground                          string `json:"statusBar.background"`
	StatusBarBorder                              string `json:"statusBar.border"`
	StatusBarDebuggingBackground                 string `json:"statusBar.debuggingBackground"`
	StatusBarDebuggingBorder                     string `json:"statusBar.debuggingBorder"`
	StatusBarDebuggingForeground                 string `json:"statusBar.debuggingForeground"`
	StatusBarForeground                          string `json:"statusBar.foreground"`
	StatusBarNoFolderBackground                  string `json:"statusBar.noFolderBackground"`
	StatusBarNoFolderBorder                      string `json:"statusBar.noFolderBorder"`
	StatusBarNoFolderForeground                  string `json:"statusBar.noFolderForeground"`
	StatusBarItemActiveBackground                string `json:"statusBarItem.activeBackground"`
	StatusBarItemHoverBackground                 string `json:"statusBarItem.hoverBackground"`
	StatusBarItemProminentBackground             string `json:"statusBarItem.prominentBackground"`
	StatusBarItemProminentHoverBackground        string `json:"statusBarItem.prominentHoverBackground"`
	StatusBarItemRemoteBackground                string `json:"statusBarItem.remoteBackground"`
	StatusBarItemRemoteForeground                string `json:"statusBarItem.remoteForeground"`
	TabActiveBackground                          string `json:"tab.activeBackground"`
	TabActiveForeground                          string `json:"tab.activeForeground"`
	TabBorder                                    string `json:"tab.border"`
	TabActiveBorder                              string `json:"tab.activeBorder"`
	TabInactiveBackground                        string `json:"tab.inactiveBackground"`
	TabInactiveForeground                        string `json:"tab.inactiveForeground"`
	TabUnfocusedActiveForeground                 string `json:"tab.unfocusedActiveForeground"`
	TabUnfocusedInactiveForeground               string `json:"tab.unfocusedInactiveForeground"`
	TerminalAnsiBlack                            string `json:"terminal.ansiBlack"`
	TerminalAnsiBlue                             string `json:"terminal.ansiBlue"`
	TerminalAnsiBrightBlack                      string `json:"terminal.ansiBrightBlack"`
	TerminalAnsiBrightBlue                       string `json:"terminal.ansiBrightBlue"`
	TerminalAnsiBrightCyan                       string `json:"terminal.ansiBrightCyan"`
	TerminalAnsiBrightGreen                      string `json:"terminal.ansiBrightGreen"`
	TerminalAnsiBrightMagenta                    string `json:"terminal.ansiBrightMagenta"`
	TerminalAnsiBrightRed                        string `json:"terminal.ansiBrightRed"`
	TerminalAnsiBrightWhite                      string `json:"terminal.ansiBrightWhite"`
	TerminalAnsiBrightYellow                     string `json:"terminal.ansiBrightYellow"`
	TerminalAnsiCyan                             string `json:"terminal.ansiCyan"`
	TerminalAnsiGreen                            string `json:"terminal.ansiGreen"`
	TerminalAnsiMagenta                          string `json:"terminal.ansiMagenta"`
	TerminalAnsiRed                              string `json:"terminal.ansiRed"`
	TerminalAnsiWhite                            string `json:"terminal.ansiWhite"`
	TerminalAnsiYellow                           string `json:"terminal.ansiYellow"`
	TerminalBackground                           string `json:"terminal.background"`
	TerminalForeground                           string `json:"terminal.foreground"`
	TerminalCursorBackground                     string `json:"terminalCursor.background"`
	TerminalCursorForeground                     string `json:"terminalCursor.foreground"`
	GitDecorationModifiedResourceForeground      string `json:"gitDecoration.modifiedResourceForeground"`
	GitDecorationDeletedResourceForeground       string `json:"gitDecoration.deletedResourceForeground"`
	GitDecorationUntrackedResourceForeground     string `json:"gitDecoration.untrackedResourceForeground"`
	GitDecorationIgnoredResourceForeground       string `json:"gitDecoration.ignoredResourceForeground"`
	GitDecorationConflictingResourceForeground   string `json:"gitDecoration.conflictingResourceForeground"`
	TextBlockQuoteBackground                     string `json:"textBlockQuote.background"`
	TextBlockQuoteBorder                         string `json:"textBlockQuote.border"`
	TextCodeBlockBackground                      string `json:"textCodeBlock.background"`
	TextLinkActiveForeground                     string `json:"textLink.activeForeground"`
	TextLinkForeground                           string `json:"textLink.foreground"`
	TextPreformatForeground                      string `json:"textPreformat.foreground"`
	TextSeparatorForeground                      string `json:"textSeparator.foreground"`
	TitleBarActiveBackground                     string `json:"titleBar.activeBackground"`
	TitleBarActiveForeground                     string `json:"titleBar.activeForeground"`
	TitleBarInactiveBackground                   string `json:"titleBar.inactiveBackground"`
	TitleBarInactiveForeground                   string `json:"titleBar.inactiveForeground"`
	WalkThroughEmbeddedEditorBackground          string `json:"walkThrough.embeddedEditorBackground"`
	WelcomePageButtonBackground                  string `json:"welcomePage.buttonBackground"`
	WelcomePageButtonHoverBackground             string `json:"welcomePage.buttonHoverBackground"`
	WidgetShadow                                 string `json:"widget.shadow"`
}
