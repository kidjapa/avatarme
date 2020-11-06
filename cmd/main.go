package cmd

import (
    "errors"
    "fmt"
    "github.com/spf13/cobra"
    "kidjapa-avatarme/avatarme"
    "os"
)

type InitialArguments struct {
    Str      string
    FileName string
    Dir      string
    Width    uint
    Height   uint
    Columns  uint
    Lines    uint
    Border   uint
}

var initialArgument InitialArguments

var rootCmd = &cobra.Command{
    Use:   "avatarme",
    Short: "create a png avatarme picture",
    Long:  ``,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Generating image")
        fmt.Println("Args passed: ", initialArgument)
        if initialArgument.Dir != "." {
            if !_validateDirectory(initialArgument.Dir) {
                panic(errors.New("invalid directory or directory folder not exists"))
            }
        }
        a := avatarme.New(
            initialArgument.Str,
            initialArgument.Dir,
            initialArgument.FileName,
            initialArgument.Width,
            initialArgument.Height,
            initialArgument.Columns,
            initialArgument.Lines,
            initialArgument.Border,
        )
        a.GenerateImage()
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.Flags().StringVarP(&initialArgument.Str, "str", "s", "testavatar", "Define the string for create avatar hash")
    rootCmd.Flags().StringVarP(&initialArgument.FileName, "file-name", "f", "image", "Define the image name for save avatar image file")
    rootCmd.Flags().StringVarP(&initialArgument.Dir, "dir", "d", ".", "Define the directory need to save the avatar. If not passed, the binary directory will be used")

    rootCmd.Flags().UintVarP(&initialArgument.Columns, "columns", "C", 4, "Define the block columns quantity.")
    rootCmd.Flags().UintVarP(&initialArgument.Lines, "lines", "L", 4, "Define the block lines quantity.")
    rootCmd.Flags().UintVarP(&initialArgument.Width, "width", "W", 1024, "Define the output image width size.")
    rootCmd.Flags().UintVarP(&initialArgument.Height, "height", "H", 1024, "Define the output image height size.")
    rootCmd.Flags().UintVarP(&initialArgument.Border, "border", "B", 30, "Define the border distance (pixel) in image size.")
}

func _validateDirectory(dir string) bool {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return false
    }
    return true
}
