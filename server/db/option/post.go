package option

type QueryType int

const (
	Undefined = iota
	TitleAndContent
	Title
	Content
	Writer
	Comment
)

var QueryTypeMap = map[string]QueryType{
	"Undefined":       Undefined,
	"TitleAndContent": TitleAndContent,
	"Title":           Title,
	"Content":         Content,
	"Writer":          Writer,
	"Comment":         Comment,
}

type PostOption struct {
	SizeRange *RangeOption
	QueryType QueryType
	Query     string
	Tags      []string
}

type CommentOption struct {
	SizeRange *RangeOption
	PostId    int
}
