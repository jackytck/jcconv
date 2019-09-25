package cmd

func must(e error) {
	if e != nil {
		panic(e)
	}
}
