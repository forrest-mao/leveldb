package cli

import (
	"fmt"
)

var cmdDesc = map[string]string{
	"create": "create an empty leveldb",
	"import": "import data into leveldb",
	"export": "export data from leveldb",
}

var cmdUsage = map[string]string{
	"create": "leveldb -create = <LevelDB Path>",
	"import": "leveldb -import = <LevelDB Path> -data = <Data File Path>",
	"export": "leveldb -export = <LevelDB Path>",
}

var cmdExample = map[string]string{
	"create": "leveldb -create = /Users/jemy/Data/test.ldb",
	"import": "leveldb -import = /Users/jemy/Data/test.ldb -data = /Users/jemy/Data/test.txt",
	"export": "leveldb -export = /Users/jemy/Data/test.ldb >> test.dat",
}

func Help(cmds ...string) {
	if len(cmds) == 0 {
		fmt.Println("leveldb v1.0:\r\n")
		for cmd, desc := range cmdDesc {
			fmt.Println(" ", cmd)
			fmt.Println("\tDesc:", desc)
			fmt.Println("\tUsage:", cmdUsage[cmd])
			fmt.Println("\tExample:", cmdExample[cmd])
			fmt.Println()
		}
	}
}
