package dnspod

type IResponseWrapper interface {
	GetCode() int
	GetMessage() string
	GetData() interface{}
}

type RecordCreateDataWrapper struct {
	Record struct {
		Id     string      `json:"id"`
		Name   string      `json:"name"`
		Status string      `json:"status"`
		Weight interface{} `json:"weight"`
	} `json:"record"`
}

type RecordCreateWrapper struct {
	Code     int                      `json:"code"`
	CodeDesc string                   `json:"codeDesc"`
	Message  string                   `json:"message"`
	Data     *RecordCreateDataWrapper `json:"data"`
}

func (r *RecordCreateWrapper) GetCode() int {
	return r.Code
}

func (r *RecordCreateWrapper) GetMessage() string {
	return r.Message
}

func (r *RecordCreateWrapper) GetData() interface{} {
	return r.Data
}

type RecordListDataWrapper struct {
	Domain struct {
		Id         string   `json:"id"`
		Name       string   `json:"name"`
		Punycode   string   `json:"punycode"`
		Grade      string   `json:"grade"`
		Owner      string   `json:"owner"`
		ExtStatus  string   `json:"ext_status"`
		Ttl        int      `json:"ttl"`
		MinTtl     int      `json:"min_ttl"`
		DnspodNs   []string `json:"dnspod_ns"`
		Status     string   `json:"status"`
		QProjectId int      `json:"q_project_id"`
	} `json:"domain"`
	Info struct {
		SubDomains  string `json:"sub_domains"`
		RecordTotal string `json:"record_total"`
	} `json:"info"`
	Records []struct {
		Id         int    `json:"id"`
		Ttl        int    `json:"ttl"`
		Value      string `json:"value"`
		Enabled    int    `json:"enabled"`
		Status     string `json:"status"`
		UpdatedOn  string `json:"updated_on"`
		QProjectId int    `json:"q_project_id"`
		Name       string `json:"name"`
		Line       string `json:"line"`
		LineId     string `json:"line_id"`
		Type       string `json:"type"`
		Remark     string ` json:"remark"`
		Mx         int    `json:"mx"`
		Hold       string `json:"hold,omitempty"`
	} `json:"records"`
}

type RecordListWrapper struct {
	Code     int                    `json:"code"`
	CodeDesc string                 `json:"codeDesc"`
	Message  string                 `json:"message"`
	Data     *RecordListDataWrapper `json:"data"`
}

func (r *RecordListWrapper) GetCode() int {
	return r.Code
}

func (r *RecordListWrapper) GetMessage() string {
	return r.Message
}

func (r *RecordListWrapper) GetData() interface{} {
	return r.Data
}

type RecordModifyDataWrapper struct {
	Record struct {
		Id     string      `json:"id"`
		Name   string      `json:"name"`
		Status string      `json:"status"`
		Weight interface{} `json:"weight"`
	} `json:"record"`
}

type RecordModifyWrapper struct {
	Code     int         `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

func (r *RecordModifyWrapper) GetCode() int {
	return r.Code
}

func (r *RecordModifyWrapper) GetMessage() string {
	return r.Message
}

func (r *RecordModifyWrapper) GetData() interface{} {
	return r.Data
}
