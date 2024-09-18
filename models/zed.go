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
	EditorBackground       string `json:"editor.background"`
	EditorForeground       string `json:"editor.foreground"`
	EditorGutterBackground string `json:"editor.gutter.background"`
	EditorGutterForeground string `json:"editor.gutter.foreground"`

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
}
