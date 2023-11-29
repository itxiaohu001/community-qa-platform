package flags

import "flag"

type Flags struct {
	Config string
}

func LoadFlags() *Flags {
	c := flag.String("c", "config.json", "配置文件路径")
	flag.Parse()

	return &Flags{
		Config: *c,
	}
}
