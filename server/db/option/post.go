package option

type QueryType int
type ClassificationType int

const (
	QueryUndefined = iota
	QueryTitleOrContent
	QueryTitle
	QueryContent
	QueryWriter
	QueryComment
)

const (
	ClassificationALL = iota
	ClassificationPopular
	ClassificationNotice
)

var QueryTypeMap = map[string]QueryType{
	"Undefined":      QueryUndefined,
	"TitleOrContent": QueryTitleOrContent,
	"Title":          QueryTitle,
	"Content":        QueryContent,
	"Writer":         QueryWriter,
	"Comment":        QueryComment,
}

var ClassificationTypeMap = map[string]ClassificationType{
	"All":     ClassificationALL,
	"Popular": ClassificationPopular,
	"Notice":  ClassificationNotice,
}

type PostOption struct {
	SizeRange          *RangeOption
	QueryType          QueryType
	ClassificationType ClassificationType
	Query              string
	Tags               []string
}

type CommentOption struct {
	SizeRange *RangeOption
	PostId    int
}
