package mapper


type Mapper interface {
	Map(trgt interface{}, dst interface{})
}

type ObjectMapper struct {

}

func NewDefaultMapper() *ObjectMapper {
	return &ObjectMapper{}
}

func (o *ObjectMapper) Map(trgt interface{}, dst interface{}) {

}
