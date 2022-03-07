package utils

import "github.com/spf13/cobra"

func CheckError(e error) {
	cobra.CheckErr(e)
	// if e != nil {
	// 	panic(e)
	// }
}
