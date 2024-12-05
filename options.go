package modbus

type options struct {
	unitId *uint8
}

type OptionFunc func(*options)

func newOptions(opts ...OptionFunc) *options {
	o := &options{}
	for idx := range opts {
		opts[idx](o)
	}
	return o
}

func UnitId(unitId uint8) OptionFunc {
	return func(o *options) {
		o.unitId = &unitId
	}
}
