package jpush

const (
	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"
)

type Audience struct {
	Object interface{}
}

func AllAudience() *Audience {
	return &Audience{
		Object: "all",
	}
}

func (this *Audience) SetID(Object []string) {
	this.set(ID, Object)
}

func (this *Audience) SetTag(Object []string) {
	this.set(TAG, Object)
}

func (this *Audience) SetTagAnd(Object []string) {
	this.set(TAG_AND, Object)
}

func (this *Audience) SetAlias(Object []string) {
	this.set(ALIAS, Object)
}

func (this *Audience) set(key string, Object []string) {
	if this.Object == nil {
		this.Object = map[string][]string{key: Object}
	} else {
		switch this.Object.(type) {
		case map[string][]string:
			this.Object.(map[string][]string)[key] = Object
		default:
		}
	}
}

func NewAudience() *Audience {
	return &Audience{}
}
