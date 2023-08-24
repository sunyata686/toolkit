package toolkit

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type CdFormItem struct {
	Raw      string     `json:"raw"`
	Img      string     `json:"img"`
	Tablize  [][]string `json:"tablize"`
	Excelize [][]string `json:"excelize"`
	File     string     `json:"file"`
}
type CdMainJson struct {
	Title      string       `json:"title"`
	GroupName  string       `json:"group_name"`
	MailDstTag string       `json:"mail_dst_tag"`
	MailDsts   string       `json:"mail_dsts"`
	FormItems  []CdFormItem `json:"form_items"`
}

func PostComplexDeliverV2(postUrl string, mainJsonS CdMainJson, filePaths []string) (httpStatusCode int, resBody string, err error) {
	valMap := make(map[string]string)
	bs, err := json.Marshal(mainJsonS)
	if err != nil {
		return -1, "", errors.Wrap(err, "error while json.Marshal()")
	}
	valMap["main_json"] = string(bs)

	fileMap := make(map[string]string)
	for _, v := range filePaths {
		fileMap[v] = v
	}
	return PostForm(postUrl, valMap, fileMap)
}
