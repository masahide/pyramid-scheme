package lib



import (
	"fmt"
	"os"
	"flag"
	"github.com/masahide/pyramid-scheme/version"
)



var (

	//flag.BoolVar(&ps.version, "v", false, "show version")
	Version   = flag.Bool("v", false, "show version")
)


func ShowVersion() string {
	return fmt.Sprintf("pyramid-scheme version: %v-%v", version.VERSION, version.GITCOMMIT)
}

func Usage() {
	cmd := os.Args[0]
	s :=  "Usage: %s [options ...]\n"
	fmt.Printf("%s\n", ShowVersion())
	fmt.Fprintf(os.Stderr, s, cmd)
	flag.PrintDefaults()
}

