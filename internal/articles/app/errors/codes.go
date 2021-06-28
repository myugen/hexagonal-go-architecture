package errors

type ArticleErrorCode uint8

func (e ArticleErrorCode) String() string {
	return articleErrorCodeValueToName[e]
}

const (
	ArticleRetrievalErrorCode ArticleErrorCode = iota + 1
	ArticleCreationErrorCode
	ArticleUpdateErrorCode
	ArticleDeletionErrorCode
	ArticleRecoveryErrorCode
)

var articleErrorCodeValueToName = map[ArticleErrorCode]string{
	ArticleRetrievalErrorCode: "ArticleRetrievalError",
	ArticleCreationErrorCode:  "ArticleCreationError",
	ArticleUpdateErrorCode:    "ArticleUpdateError",
	ArticleDeletionErrorCode:  "ArticleDeletionError",
	ArticleRecoveryErrorCode:  "ArticleRecoveryError",
}
