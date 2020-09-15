package dnspod

import (
	"testing"
	"fmt"
)

var client *Client

func init() {
	client = NewClient("", "", "")
}

func TestRecordCreate(t *testing.T) {
	r := &RecordCreate{
		Domain:     "",
		SubDomain:  "",
		RecordType: RecordTypeA,
		RecordLine: "默认",
		Value:      "",
	}

	data, resp, err := client.RecordsService.RecordCreate(r)

	fmt.Println(data)
	fmt.Println(resp)
	fmt.Println(err)
}

func TestRecordList(t *testing.T) {
	r := &RecordList{
		Domain:     "",
		SubDomain:  "",
		RecordType: RecordTypeA,
	}

	data, resp, err := client.RecordsService.RecordList(r)

	for _, r := range data.Records {
		fmt.Println(r.Id)
	}
	fmt.Println(data.Info.RecordTotal)
	fmt.Println(resp)
	fmt.Println(err)
}

func TestRecordModify(t *testing.T) {
	r := &RecordModify{
		Domain:     "",
		recordId:   0,
		SubDomain:  "",
		RecordType: RecordTypeA,
		RecordLine: "默认",
		Value:      "",
	}
	data, resp, err := client.RecordsService.RecordModify(r)

	fmt.Println(data)
	fmt.Println(resp)
	fmt.Println(err)
}
