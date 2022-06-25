package blogLogic

import (
	"fmt"
	"io/ioutil"
	dataStruct "simpleCollabBlog/simple-blog-backend/dataStructs"
	"strings"

	"github.com/russross/blackfriday"
)

func GetHeadline(str string) (result string) {
	startToken := "<h1>"
	endToken := "</h1>"
	s := strings.Index(str, startToken)
	if s == -1 {
		return
	}
	s += len(startToken)

	e := strings.Index(str[s:], endToken)
	if e == -1 {
		return
	}
	e += s + e + 1
	return str[s:e]
}

func CreateArticle(article string) *dataStruct.TestDataStruct {
	newArticle := dataStruct.TestDataStruct{
		Title:   GetHeadline(article),
		Content: article,
	}

	return &newArticle
}

func MdToArticleStruct(pathToMd string) (outputArticle *dataStruct.TestDataStruct) {
	content, err := ioutil.ReadFile(pathToMd)
	if err != nil {
		fmt.Print("%s", err)
	}

	contentByte := blackfriday.MarkdownCommon(content)
	outputArticle = CreateArticle(string(contentByte))

	return
}
