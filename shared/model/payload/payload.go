package payload

import (
	"demo4/shared/driver"
)

type Payload struct {
	Data      interface{}            `json:"data"`
	Publisher driver.ApplicationData `json:"publisher"`
	TraceID   string                 `json:"traceId"`
}
