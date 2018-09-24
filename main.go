package main

func main() {
	lc := NewLolichan()
	defer lc.db.Close()

	cli := CLI{lc}
	cli.Run()
}
