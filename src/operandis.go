package computor

type IOperandis interface {
	Parse() (err error)
}
type Operandis struct {
	Left  interface{}
	Right interface{}
}

type Op struct {
	Type string
	Var  interface{}
}

func (o *Operandis) Parse() (err error) {
	if o.Left, err = parseLeft(o.Left.([]byte)); err != nil {
		return
	}
	if o.Right, err = parseRight(o.Right.([]byte)); err != nil {
		return
	}
	return
}

func parseLeft(l []byte) (i interface{}, err error) {
	return
}

func parseRight(r []byte) (i interface{}, err error) {
	return
}
