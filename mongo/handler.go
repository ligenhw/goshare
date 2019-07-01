package mongo

// GetLinks get all friend links
var (
	db          = getDb()
	GetLinks    = CreateGetCollection(db.Collection("link"))
	GetBooks    = CreateGetCollection(db.Collection("book"))
	GetProjects = CreateGetCollection(db.Collection("project"))
)
