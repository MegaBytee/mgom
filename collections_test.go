package mgom

//declarative test file for all collections that will use in the mgom_test :

//collection : logs
type Log struct {
	Key   string
	Value string
}

var log_index = []IndexFactory{
	{
		T: IdxUNIQUE,
		Values: map[string]string{
			"key": "true",
		},
	},
}

//collection : posts
type Post struct {
	PostId string `bson:"post_id" json:"post_id"`
	Author string
	Title  string
	Body   string
	Views  int
}

var post_index = []IndexFactory{
	{
		T: IdxUNIQUE,
		Values: map[string]string{
			"post_id": "true",
			"author":  "false",
		},
	},
	{
		T: IdxTEXT,
		Values: map[string]string{
			"title": "9",
		},
	},
}

//collection : tags
type Tag struct {
	Value string
	Count int
}

var tag_index = []IndexFactory{
	{
		T: IdxUNIQUE,
		Values: map[string]string{
			"value": "true",
		},
	},
	{
		T: IdxTEXT,
		Values: map[string]string{
			"value": "9",
		},
	},
}
