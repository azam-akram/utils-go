package url_utils

import (
	"fmt"
	"log"
	"net/url"
)

func DemoUrlEncode() {

	pathWithoutEncodedCharacter := "foo bar"
	unescapedPath, err := url.PathUnescape(pathWithoutEncodedCharacter)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("unescaped path: %s\n", unescapedPath)

	pathWithEncodedCharacter := "foo%20bar"
	unescapedPathWithEncodeCharacter, err := url.PathUnescape(pathWithEncodedCharacter)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("unescaped path: %s\n", unescapedPathWithEncodeCharacter)

	// decode query by url.QueryUnescape
	/*query := "query=ab%2Bc&query2=de%24f"
	unescapedQuery, err := url.QueryUnescape(query)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("unescaped query: %s\n", unescapedQuery)

	// decode query and parse by url.ParseQuery
	parsedQuery, err := url.ParseQuery(query)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("parsed query args:")
	for key, values := range parsedQuery {
		fmt.Printf("  %s = %s\n", key, values[0])
	}*/

}
