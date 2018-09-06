// worldjob
package main

import (
	"fmt"
	_ "io"
	"io/ioutil"
	"net/http"
	_ "os"
	"regexp"
	"strings"
	"xlsx" //https://github.com/tealeg/xlsx
)

type worldJobTag struct {
	rctntcSj              string
	rctntcSprtQualfCn     string
	dsptcNationScd        string
	dsptcKsco             string
	joDemandCareerStleScd string
	joDemandAcdmcrScd     string
	rctntcEndDay          string
	linkUrl               string
	directApply           string
}

var worldJob = make(map[int]worldJobTag)

var (
	PageIndex int
)

func StripTags(html string) string {
	html = strings.Replace(html, "&#13;", "", -1)
	html = strings.Replace(html, "&lt;", "", -1) //html = strings.Replace(html, "&lt;", "<", -1)
	html = strings.Replace(html, "&gt;", "", -1) //html = strings.Replace(html, "&gt;", ">", -1)
	return html
}

func main() {

	//https://www.data.go.kr/dataset/3038249/openapi.do

	spage := "http://www.worldjob.or.kr/openapi/openapi.do?"
	sdobType := "1"    //1:해외취업,2:해외연수,3:해외인턴,4:해외봉사,5:해외창업
	sdsptcKsco := "01" //직종별코드(해외취업,연수만 사용)01:전산,컴퓨터,02:전기/전자,06:기계/금속,07:건설/토목,08:사무/서비스,09:의료,10:기타
	scontinent := "1"  //대륙별코드 1:아시아,2:북아메리카, 3:남아메리카,4:유럽,5:오세아니아,6:아프리카
	sepmt61 := "Y"     //일자리Best20(해외취업만 사용)Y,N
	//spageIndex := "10" //페이징숫자
	sshowItemListCount := "1000" //한번에보여질리스트갯수출력결과
	sUrl := fmt.Sprintf("%sdobType=%s&dsptcKsco=%s&continent=%s&showItemListCount=%s&sepmt61=%s", spage, sdobType, sdsptcKsco, scontinent, sshowItemListCount, sepmt61)
	//fmt.Println(sUrl)

	res, err := http.Get(sUrl)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return
	}

	var crawl = string(body)
	//fmt.Println(string(body))

	var pattern = regexp.MustCompile(`<ITEM>([\w\W]+?)</ITEM>`)
	data := pattern.FindAllString(crawl, -1)

	var rctntcSj = regexp.MustCompile(`<rctntcSj>([\w\W]+?)</rctntcSj>`)
	var rctntcSprtQualfCn = regexp.MustCompile(`<rctntcSprtQualfCn>([\w\W]+?)</rctntcSprtQualfCn>`)
	var dsptcNationScd = regexp.MustCompile(`<dsptcNationScd>([\w\W]+?)</dsptcNationScd>`)
	var dsptcKsco = regexp.MustCompile(`<dsptcKsco>([\w\W]+?)</dsptcKsco>`)
	var joDemandCareerStleScd = regexp.MustCompile(`<joDemandCareerStleScd>([\w\W]+?)</joDemandCareerStleScd>`)
	var joDemandAcdmcrScd = regexp.MustCompile(`<joDemandAcdmcrScd>([\w\W]+?)</joDemandAcdmcrScd>`)
	var rctntcEndDay = regexp.MustCompile(`<rctntcEndDay>([\w\W]+?)</rctntcEndDay>`)
	var linkUrl = regexp.MustCompile(`<linkUrl>([\w\W]+?)</linkUrl>`)
	var directApply = regexp.MustCompile(`<directApply>([\w\W]+?)</directApply>`)

	if data != nil {

		for _, val := range data {

			td1 := rctntcSj.FindAllString(val, -1)
			td2 := rctntcSprtQualfCn.FindAllString(val, -1)
			td3 := dsptcNationScd.FindAllString(val, -1)
			td4 := dsptcKsco.FindAllString(val, -1)
			td5 := joDemandCareerStleScd.FindAllString(val, -1)
			td6 := joDemandAcdmcrScd.FindAllString(val, -1)
			td7 := rctntcEndDay.FindAllString(val, -1)
			td8 := linkUrl.FindAllString(val, -1)
			td9 := directApply.FindAllString(val, -1)

			if len(td1) == 0 || len(td2) == 0 || len(td3) == 0 || len(td4) == 0 || len(td5) == 0 || len(td6) == 0 || len(td7) == 0 || len(td8) == 0 || len(td9) == 0 {
				continue
			}

			td1_S := strings.Index(td1[0], ">")
			td1_E := strings.LastIndex(td1[0], "<")

			td2_S := strings.Index(td2[0], ">")
			td2_E := strings.LastIndex(td2[0], "<")

			td3_S := strings.Index(td3[0], ">")
			td3_E := strings.LastIndex(td3[0], "<")

			td4_S := strings.Index(td4[0], ">")
			td4_E := strings.LastIndex(td4[0], "<")

			td5_S := strings.Index(td5[0], ">")
			td5_E := strings.LastIndex(td5[0], "<")

			td6_S := strings.Index(td6[0], ">")
			td6_E := strings.LastIndex(td6[0], "<")

			td7_S := strings.Index(td7[0], ">")
			td7_E := strings.LastIndex(td7[0], "<")

			td8_S := strings.Index(td8[0], ">")
			td8_E := strings.LastIndex(td8[0], "<")

			td9_S := strings.Index(td9[0], ">")
			td9_E := strings.LastIndex(td9[0], "<")

			PageIndex++

			worldJob[PageIndex] = worldJobTag{string(td1[0][td1_S+1 : td1_E]),
				StripTags(string(td2[0][td2_S+1 : td2_E])),
				string(td3[0][td3_S+1 : td3_E]),
				string(td4[0][td4_S+1 : td4_E]),
				string(td5[0][td5_S+1 : td5_E]),
				string(td6[0][td6_S+1 : td6_E]),
				string(td7[0][td7_S+1 : td7_E]),
				string(td8[0][td8_S+1 : td8_E]),
				string(td9[0][td9_S+1 : td9_E])}

			//link := fmt.Sprintf("[%d]%s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n", PageIndex, string(td1[0][td1_S+1:td1_E]), StripTags(string(td2[0][td2_S+1:td2_E])), string(td3[0][td3_S+1:td3_E]), string(td4[0][td4_S+1:td4_E]), string(td5[0][td5_S+1:td5_E]), string(td6[0][td6_S+1:td6_E]), string(td7[0][td7_S+1:td7_E]), string(td8[0][td8_S+1:td8_E]), string(td9[0][td9_S+1:td9_E]))
			//fmt.Println(link)
		}

		//엑셀로 저장하기
		var file *xlsx.File
		var sheet *xlsx.Sheet
		var row *xlsx.Row
		var cell1, cell2, cell3, cell4, cell5, cell6, cell7, cell8, cell9 *xlsx.Cell
		var err error

		file = xlsx.NewFile()
		sheet, err = file.AddSheet("Sheet1")
		if err != nil {
			fmt.Printf(err.Error())
		}

		for _, val := range worldJob {
			//link := fmt.Sprintf("%s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n %s\r\n", val.rctntcSj, val.rctntcSprtQualfCn, val.dsptcNationScd, val.dsptcKsco, val.joDemandCareerStleScd, val.joDemandAcdmcrScd, val.rctntcEndDay, val.linkUrl, val.directApply)
			//fmt.Println(link)

			row = sheet.AddRow()

			cell1 = row.AddCell()
			cell1.Value = val.rctntcSj

			cell2 = row.AddCell()
			cell2.Value = val.rctntcSprtQualfCn

			cell3 = row.AddCell()
			cell3.Value = val.dsptcNationScd

			cell4 = row.AddCell()
			cell4.Value = val.dsptcKsco

			cell5 = row.AddCell()
			cell5.Value = val.joDemandCareerStleScd

			cell6 = row.AddCell()
			cell6.Value = val.joDemandAcdmcrScd

			cell7 = row.AddCell()
			cell7.Value = val.rctntcEndDay

			cell8 = row.AddCell()
			cell8.Value = val.linkUrl

			cell9 = row.AddCell()
			cell9.Value = val.directApply
		}

		err = file.Save("daeseong.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}

		/*
			for key, val := range worldJob {
				fmt.Println(key, val)
			}
		*/
	}

}
