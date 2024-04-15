package operations

type Operation struct {
	Type string `json: "type"`
	Has_succeeded bool `json: "has_succeeded"`
}

func (o Operation) InitializeOperation(Type string, Has_succeeded bool) (Operation, error){
	o.Type = Type
	o.Has_succeeded = Has_succeeded
	return o, nil
}