package msg

import(
    "time"
    "encoding/xml"
)

type Response struct {
    XMLName xml.Name `xml:"xml"`
    msgBase
    ArticleCount int `xml:",omitempty"`
    Articles []*item `xml:"Articles>item,omitempty"`
    FuncFlag int
}

type item struct {
    XMLName xml.Name `xml:"item"`
    Title string
    Description string
    PicUrl string
    Url string
}

func Marshal(resp Response) (data []byte ,err error) {
    resp.CreateTime /= time.Second
    data, err = xml.Marshal(resp)
    return
}
