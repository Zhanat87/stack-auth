package models
import (
	"time"
	"gopkg.in/mgo.v2/bson"
)
/*
@link https://golang.org/pkg/encoding/json/
omitempty - если значение поля отсутствует, то это поле не показывается просто

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

The "string" option signals that a field is stored as JSON inside a JSON-encoded string.
It applies only to fields of string, floating point, integer, or boolean types.
This extra level of encoding is sometimes used when communicating with JavaScript programs:
Int64String int64 `json:",string"
 */
type (
	User struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName string `json:"firstname"`
		LastName string `json:"lastname"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty "`
		ChatRoom string `json:"chatroom"`
	}
	Task struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy string `json:"createdby"`
		Name string `json:"name"`
		Description string `json:"description"`
		CreatedOn time.Time `json:"createdon,omitempty"`
		Due time.Time `json:"due,omitempty"`
		Status string `json:"status,omitempty"`
		Tags []string `json:"tags,omitempty"`
	}
	TaskNote struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskId bson.ObjectId `json:"taskid"`
		Description string `json:"description"`
		CreatedOn time.Time `json:"createdon,omitempty"`
	}
)