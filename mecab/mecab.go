package mecab

import (
	"fmt"

	mecab "github.com/shogo82148/go-mecab"
)



func parse(args map[string]string) {
	mecab, err := mecab.New(args)
	if err != nil {
			panic(err)
	}
	defer mecab.Destroy()

	node, err := mecab.ParseToNode(text)

	for ; !node.IsZero(); node = node.Next() {
			fmt.Printf("%s\t%s\n", node.Surface(), node.Feature())
	}
}