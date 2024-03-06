package cmd

import (
	"fmt"
	"os"
	_"path"
	_"path/filepath"
	"testing"
)

func TestOsExecutable(t *testing.T) {
	str, _ := os.Executable()
	// str = path.Join(filepath.Dir(str), "conf")
	fmt.Println(str)
}