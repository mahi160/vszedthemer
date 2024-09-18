package main

import (
	"encoding/json"
	"fmt"
	"os"

	"vszedthemer/models"

	"github.com/spf13/cobra"
)

func convertTheme(vsCodeThemeFile string) (models.ZedTheme, error) {
	vsCodeThemeFileBytes, err := os.ReadFile(vsCodeThemeFile)
	if err != nil {
		return models.ZedTheme{}, err
	}

	var vsCodeTheme models.VSCodeTheme
	err = json.Unmarshal(vsCodeThemeFileBytes, &vsCodeTheme)
	if err != nil {
		return models.ZedTheme{}, err
	}
	zedTheme := models.ZedTheme{
		Author: vsCodeTheme.Author,
		Name:   vsCodeTheme.Name,
		Themes: make([]models.ZedThemeSettings, 1),
	}

	zedTheme.Themes[0] = models.ZedThemeSettings{
		Name:       vsCodeTheme.Name,
		Appearance: vsCodeTheme.Type,
		Style: models.ZedThemeStyle{
			TerminalAnsiBlack:         vsCodeTheme.Colors.TerminalAnsiBlack,
			TerminalAnsiBlue:          vsCodeTheme.Colors.TerminalAnsiBlue,
			TerminalAnsiBrightBlack:   vsCodeTheme.Colors.TerminalAnsiBrightBlack,
			TerminalAnsiBrightBlue:    vsCodeTheme.Colors.TerminalAnsiBrightBlue,
			TerminalAnsiBrightCyan:    vsCodeTheme.Colors.TerminalAnsiBrightCyan,
			TerminalAnsiBrightGreen:   vsCodeTheme.Colors.TerminalAnsiBrightGreen,
			TerminalAnsiBrightMagenta: vsCodeTheme.Colors.TerminalAnsiBrightMagenta,
			TerminalAnsiBrightRed:     vsCodeTheme.Colors.TerminalAnsiBrightRed,
			TerminalAnsiBrightWhite:   vsCodeTheme.Colors.TerminalAnsiBrightWhite,
			TerminalAnsiBrightYellow:  vsCodeTheme.Colors.TerminalAnsiBrightYellow,
			TerminalAnsiCyan:          vsCodeTheme.Colors.TerminalAnsiCyan,
			TerminalAnsiGreen:         vsCodeTheme.Colors.TerminalAnsiGreen,
			TerminalAnsiMagenta:       vsCodeTheme.Colors.TerminalAnsiMagenta,
			TerminalAnsiRed:           vsCodeTheme.Colors.TerminalAnsiRed,
			TerminalAnsiWhite:         vsCodeTheme.Colors.TerminalAnsiWhite,
			TerminalAnsiYellow:        vsCodeTheme.Colors.TerminalAnsiYellow,
			TerminalBackground:        vsCodeTheme.Colors.TerminalBackground,
			TerminalForeground:        vsCodeTheme.Colors.TerminalForeground,
		},
	}

	return zedTheme, nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "vscode2zed",
		Short: "Converts VS Code themes to Zed themes",
		Long:  "Converts VS Code themes to Zed themes",
	}

	convertCmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert a VS Code theme",
		Long:  "Convert a VS Code theme to a Zed theme",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				cmd.Usage()
				os.Exit(1)
			}

			vsCodeThemeFile := args[0]
			zedThemeFile := args[1]

			zedTheme, err := convertTheme(vsCodeThemeFile)
			if err != nil {
				fmt.Println("Error converting theme:", err)
				os.Exit(1)
			}

			zedThemeJSON, err := json.MarshalIndent(zedTheme, "", "  ")
			if err != nil {
				fmt.Println("Error marshalling Zed theme:", err)
				os.Exit(1)
			}

			err = os.WriteFile(zedThemeFile, zedThemeJSON, 0644)
			if err != nil {
				fmt.Println("Error writing Zed theme file:", err)
				os.Exit(1)
			}

			fmt.Println("Zed theme converted successfully!")
		},
	}

	rootCmd.AddCommand(convertCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
