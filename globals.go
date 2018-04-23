package main

type session struct{
	runs int
	deletions int
	files map[string] string
	revealed map[string] bool
}

func b2s(a bool) string{
	if a {
		return "true"
	} else {
		return "false"
	}
}
