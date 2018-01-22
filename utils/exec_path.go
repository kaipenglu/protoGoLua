package utils

import (
    "path/filepath"
    "os"
)

var exePath string = ""
func ExecFilePath() string {
    if(exePath == "") {
        exePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
    }
    return exePath
}
