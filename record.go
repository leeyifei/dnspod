package dnspod

import (
	"strconv"
	"net/http"
)

const (
	RecordTypeA     = "A"
	RecordTypeCNAME = "CNAME"
)

type RecordCreate struct {
	Domain     string
	SubDomain  string
	RecordType string
	RecordLine string
	Value      string
	Ttl        int
	Mx         int
}

type RecordList struct {
	Domain     string
	Offset     int
	Length     int
	SubDomain  string
	RecordType string
	QProjectId int
}

type RecordModify struct {
	Domain     string
	RecordId   int
	SubDomain  string
	RecordType string
	RecordLine string
	Value      string
	Ttl        int
	Mx         int
}

type RecordsService struct {
	ctx *Client
}

func (d *RecordsService) RecordCreate(r *RecordCreate) (*RecordCreateDataWrapper, *http.Response, error) {
	data := map[string]string{
		"Action":     "RecordCreate",
		"domain":     r.Domain,
		"subDomain":  r.SubDomain,
		"recordType": r.RecordType,
		"recordLine": r.RecordLine,
		"value":      r.Value,
	}
	if r.Ttl > 0 {
		data["ttl"] = strconv.Itoa(r.Ttl)
	}
	if r.Mx > 0 {
		data["mx"] = strconv.Itoa(r.Mx)
	}

	var v IResponseWrapper
	v = new(RecordCreateWrapper)
	resp, err := d.ctx.Get(data, v)
	if err != nil {
		return nil, resp, err
	}

	result := v.GetData().(*RecordCreateDataWrapper)

	return result, resp, err
}

func (d *RecordsService) RecordList(r *RecordList) (*RecordListDataWrapper, *http.Response, error) {
	data := map[string]string{
		"Action": "RecordList",
		"domain": r.Domain,
	}

	if r.Offset > 0 {
		data["offset"] = strconv.Itoa(r.Offset)
	}

	if r.Length > 0 {
		data["length"] = strconv.Itoa(r.Length)
	}

	if r.SubDomain != "" {
		data["subDomain"] = r.SubDomain
	}

	if r.RecordType != "" {
		data["recordType"] = r.RecordType
	}

	if r.QProjectId > 0 {
		data["qProjectId"] = strconv.Itoa(r.QProjectId)
	}

	var v IResponseWrapper
	v = new(RecordListWrapper)
	resp, err := d.ctx.Get(data, v)
	if err != nil {
		return nil, resp, err
	}

	result := v.GetData().(*RecordListDataWrapper)
	return result, resp, nil
}

func (d *RecordsService) RecordModify(r *RecordModify) (interface{}, *http.Response, error) {
	data := map[string]string{
		"Action":     "RecordModify",
		"domain":     r.Domain,
		"recordId":   strconv.Itoa(r.RecordId),
		"subDomain":  r.SubDomain,
		"recordType": r.RecordType,
		"recordLine": r.RecordLine,
		"value":      r.Value,
	}

	if r.Ttl > 0 {
		data["ttl"] = strconv.Itoa(r.Ttl)
	}

	if r.Mx > 0 {
		data["mx"] = strconv.Itoa(r.Mx)
	}

	var v IResponseWrapper
	v = new(RecordModifyWrapper)

	resp, err := d.ctx.Get(data, v)
	if err != nil {
		return nil, resp, err
	}

	result := v.GetData()
	return result, resp, nil
}
