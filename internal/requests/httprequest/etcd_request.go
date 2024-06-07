package httprequest

type EtcdPutRequest struct {
	Key   string `json:"key" validate:"required,max=255"`
	Value string `json:"value" validate:"required,max=255"`
}

type EtcdRangeRequest struct {
	Key      string `json:"key" validate:"required,max=255"`
	RangeEnd string `json:"range_end" validate:"max=255"`
}
