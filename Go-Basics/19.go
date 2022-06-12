package main

type user struct {
	name       string
	email      string
	age        int
	previleged bool
}

type admin struct {
	person user
	level  string
}

func main() {
	lisa := user{
		"Lisa",
		"Lisa@gmail.com",
		"21",
		true,
	}

	fired := admin{
		person: user{
			"Lisa",
			"Lisa@gmail.com",
			"21",
			true,
		},
		level: "super",
	}

}
