package main

import (
	"fmt"
	"os/exec"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
)

func main()  {
	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		panic(err)
	}else{
		println(string(out))
	}
	process, err := sysinfo.Self()
	if err != nil {
		panic(err)
	}

	if handleCounter, ok := process.(types.OpenHandleCounter); ok {
		count, err := handleCounter.OpenHandleCount()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d open handles\n", count)
	}

	info, _ := sysinfo.Host()
	fmt.Println(info.FQDN())
	fmt.Println(info.Info().OS)
}