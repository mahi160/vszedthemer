package models

type ZedTheme struct {
	Author string             `json:"author"`
	Name   string             `json:"name"`
	Themes []ZedThemeSettings `json:"themes"`
}
type ZedThemeSettings struct {
	Name       string        `json:"name"`
	Appearance string        `json:"appearance"`
	Style      ZedThemeStyle `json:"style"`
}
type ZedThemeStyle struct {
	Background        string `json:"background"`
	Border            string `json:"border"`
	BorderDisabled    string `json:"border.disabled"`
	BorderFocused     string `json:"border.focused"`
	BorderSelected    string `json:"border.selected"`
	BorderTransparent string `json:"border.transparent"`
	BorderVariant     string `json:"border.variant"`

	Conflict           string `json:"conflict"`
	ConflictBackground string `json:"conflict.background"`
	ConflictBorder     string `json:"conflict.border"`

	Created           string `json:"created"`
	CreatedBackground string `json:"created.background"`
	CreatedBorder     string `json:"created.border"`

	Deleted           string `json:"deleted"`
	DeletedBackground string `json:"deleted.background"`
	DeletedBorder     string `json:"deleted.border"`

	DropTargetBackground string `json:"dropTarget.background"`

	EditorActiveLineBackground             string `json:"editorActiveLineBackground"`
	EditorActiveLineNumber                 string `json:"editorActiveLineNumber"`
	EditorActiveWrapGuide                  string `json:"editor.active_wrap_guide"`
	EditorBackground                       string `json:"editor.background"`
	EditorDocumentHighlightReadBackground  string `json:"editor.document_highlight.read_background"`
	EditorDocumentHighlightWriteBackground string `json:"editor.document_highlight.write_background"`
	EditorForeground                       string `json:"editor.foreground"`
	EditorGutterBackground                 string `json:"editor.gutter.background"`
	EditorHighlightedLineBackground        string `json:"editor.highlighted_line.background"`
	EditorInvisible                        string `json:"editor.invisible"`
	EditorLineNumber                       string `json:"editor.line_number"`
	EditorSubHeaderBackground              string `json:"editor.subHeader.background"`
	EditorWrapGuide                        string `json:"editor.wrap_guide"`
	EditorActive                           string `json:"editorActive"`

	ElementBackground string `json:"element.background"`
	ElementDisabled   string `json:"element.disabled"`
	ElementHover      string `json:"element.hover"`
	ElementSelected   string `json:"element.selected"`

	ElevatedSurfaceBackground string `json:"elevated_surface.background"`

	Error           string `json:"error"`
	ErrorBackground string `json:"error.background"`
	ErrorBorder     string `json:"error.border"`

	GhostElementActive     string `json:"ghost_element.active"`
	GhostElementBackground string `json:"ghost_element.background"`
	GhostElementDisabled   string `json:"ghost_element.disabled"`
	GhostElementHover      string `json:"ghost_element.hover"`
	GhostElementSelected   string `json:"ghost_element.selected"`

	Hidden           string `json:"hidden"`
	HiddenBackground string `json:"hidden.background"`
	HiddenBorder     string `json:"hidden.border"`

	Hint           string `json:"hint"`
	HintBackground string `json:"hint.background"`
	HintBorder     string `json:"hint.border"`

	Icon            string `json:"icon"`
	IconAccent      string `json:"icon.accent"`
	IconDisabled    string `json:"icon.disabled"`
	IconMuted       string `json:"icon.muted"`
	IconPlaceholder string `json:"icon.placeholder"`

	Ignored           string `json:"ignored"`
	IgnoredBackground string `json:"ignored.background"`
	IgnoredBorder     string `json:"ignored.border"`

	Info           string `json:"info"`
	InfoBackground string `json:"info.background"`
	InfoBorder     string `json:"info.border"`

	LinkTextHover string `json:"link_text.hover"`

	Modified           string `json:"modified"`
	ModifiedBackground string `json:"modified.background"`
	ModifiedBorder     string `json:"modified.border"`

	PaneFocusedBorder string `json:"pane.focused_border"`

	PanelBackground    string `json:"panel.background"`
	PanelFocusedBorder string `json:"panel.focused_border"`

	Predictive           string `json:"predictive"`
	PredictiveBackground string `json:"predictive.background"`
	PredictiveBorder     string `json:"predictive.border"`

	Renamed           string `json:"renamed"`
	RenamedBackground string `json:"renamed.background"`
	RenamedBorder     string `json:"renamed.border"`

	ScrollbarThumbBorder          string `json:"scrollbar.thumb.border"`
	ScrollbarThumbHoverBackground string `json:"scrollbar.thumb.hover_background"`
	ScrollbarTrackBackground      string `json:"scrollbar.track.background"`
	ScrollbarTrackBorder          string `json:"scrollbar.track.border"`
	ScrollbarThumbBackground      string `json:"scrollbar_thumb.background"`

	SearchMatchBackground string `json:"search.match_background"`
	StatusBarBackground   string `json:"status_bar.background"`

	Success           string `json:"success"`
	SuccessBackground string `json:"success.background"`
	SuccessBorder     string `json:"success.border"`

	SurfaceBackground string `json:"surface.background"`

	Syntax string `json:"syntax"`

	TabActiveBackground   string `json:"tab.active_background"`
	TabInactiveBackground string `json:"tab.inactive_background"`
	TabBarBackground      string `json:"tab_bar.background"`

	TerminalAnsiBlack         string `json:"terminal.ansi.black"`
	TerminalAnsiBlue          string `json:"terminal.ansi.blue"`
	TerminalAnsiCyan          string `json:"terminal.ansi.cyan"`
	TerminalAnsiGreen         string `json:"terminal.ansi.green"`
	TerminalAnsiMagenta       string `json:"terminal.ansi.magenta"`
	TerminalAnsiRed           string `json:"terminal.ansi.red"`
	TerminalAnsiWhite         string `json:"terminal.ansi.white"`
	TerminalAnsiYellow        string `json:"terminal.ansi.yellow"`
	TerminalAnsiBrightBlack   string `json:"terminal.ansi.bright_black"`
	TerminalAnsiBrightBlue    string `json:"terminal.ansi.bright_blue"`
	TerminalAnsiBrightCyan    string `json:"terminal.ansi.bright_cyan"`
	TerminalAnsiBrightGreen   string `json:"terminal.ansi.bright_green"`
	TerminalAnsiBrightMagenta string `json:"terminal.ansi.bright_magenta"`
	TerminalAnsiBrightRed     string `json:"terminal.ansi.bright_red"`
	TerminalAnsiBrightWhite   string `json:"terminal.ansi.bright_white"`
	TerminalAnsiBrightYellow  string `json:"terminal.ansi.bright_yellow"`
	TerminalBackground        string `json:"terminal.background"`
	TerminalForeground        string `json:"terminal.foreground"`

	Text            string `json:"text"`
	TextAccent      string `json:"text.accent"`
	TextDisabled    string `json:"text.disabled"`
	TextMuted       string `json:"text.muted"`
	TextPlaceholder string `json:"text.placeholder"`

	TitleBarBackground string `json:"title_bar.background"`
	ToolbarBackground  string `json:"toolbar.background"`

	UnReachable           string `json:"unreachable"`
	UnReachableBackground string `json:"unreachable.background"`
	UnReachableBorder     string `json:"unreachable.border"`

	Warning           string `json:"warning"`
	WarningBackground string `json:"warning.background"`
	WarningBorder     string `json:"warning.border"`
}
