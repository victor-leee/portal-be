package response_error

import "strings"

const (
	levelSeparator = " || "
)

type PortalError struct {
	InternalError error
	ResponseCode  *int64
}

func (p *PortalError) Error() string {
	errBuilder := strings.Builder{}
	e, ok := p, true
	for ok {
		errBuilder.WriteString(e.InternalError.Error())
		e, ok = p.InternalError.(*PortalError)
		if ok {
			errBuilder.WriteString(levelSeparator)
		}
	}

	return errBuilder.String()
}

func (p *PortalError) Code() int64 {
	e, ok := p, true
	for ok {
		if e.ResponseCode != nil {
			return *e.ResponseCode
		}
		e, ok = p.InternalError.(*PortalError)
	}

	return 0
}
