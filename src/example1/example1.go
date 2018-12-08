// example1 --input /tmp/file -p /var/log --path ./log -v -- -path/log apache.log
//
// TOTO: fix the error with this use case: ./bin/example1 v 10 --input /tmp/file -p /var/log --path ./log  -- -path/log apache.log

package main
import "os"
import "fmt"
import "strings"
import "beurive.com/cli"

func main() {

    var cloVerbose       bool
	var cloInput         string
	var cloPath          []string

    // CLI specification

    spec := cli.Spec{
        cli.Option{Short: "",  Long: "input", Holder: &cloInput},
        cli.Option{Short: "v", Long: "",      Holder: &cloVerbose},
        cli.Option{Short: "p", Long: "path",  Holder: &cloPath},
    };

    if cli, args, err := cli.Parse(os.Args[1:], spec); nil != err {
        fmt.Printf(`Invalid command line ! Got the error: %s`, err.Error())
    } else {
        fmt.Printf("Number of tokens: %d -> %s\n", len(cli), strings.Join(cli, ` `))
        fmt.Printf("Number of arguments: %d\n\n", len(args))

        fmt.Println("Options:\n");
        fmt.Printf("\t* cloVerbose = %s\n", (func() string { if (cloVerbose) { return `true` } else { return `false` } })())
        fmt.Printf("\t* cloInput = %s\n", cloInput)
        fmt.Printf("\t* cloPath = %s\n\n", strings.Join(cloPath, `:`))

        fmt.Println("Arguments:\n");
        for _, arg := range args {
            fmt.Printf("\t* %s\n", arg)
        }
    }
}