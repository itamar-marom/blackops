package utils

func CheckError(e error) {
	// cobra.CheckErr(e)
	if e != nil {
		panic(e)
	}
}
