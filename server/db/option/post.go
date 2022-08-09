package option

type PostOption struct {
	SizeRange            *RangeOption
	SearchNameAndContent string
	SearchName           string
	SearchContent        string
	SearchWriter         string
	SearchComment        string
}

type CommentOption struct {
	SizeRange *RangeOption
	PostId    int
}
