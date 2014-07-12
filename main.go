package main

func main() {
	d := Devcards{Name: "Root"}
	d.AddChild(Devcards{Name: "Child 1"})
	d.AddChild(Devcards{Name: "Child 2"})
	d2 := Devcards{Name: "Child 3"}
	d2.AddChild(Devcards{Name: "Grandchild 1"})
	d.AddChild(d2)
	d.AddChild(Devcards{Name: "Child 4"})
	d.Serve()
}
