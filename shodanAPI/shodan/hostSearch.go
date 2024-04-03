package shodan

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
)

type HostLocation struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float32 `json:"latitude"`
}

type Host struct {
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	ISP       string       `json:"isp"`
	ASN       string       `json:"asn"`
	Hostnames []string     `json:"hostnames"`
	Location  HostLocation `json:"location"`
	IP        int64        `json:"ip"`
	Domains   []string     `json:"domains"`
	Org       string       `json:"org"`
	Data      string       `json:"data"`
	Port      int          `json:"port"`
	IPString  string       `json:"ip_str"`
}

type HostSearch struct {
	Matches []Host `json:"matches"`
}

func (s *Client) HostSearch(q string) (*HostSearch, error) {
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, s.apiKey, q))
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	var ret HostSearch
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

// 导出到Excel
func SaveToExcel(hosts []Host, filename string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Hosts")
	if err != nil {
		return err
	}

	// 添加表头
	headers := []string{"IP地址", "端口", "域名", "主机名", "ISP", "ASN", "组织", "数据", "城市", "国家", "操作系统", "时间戳", "纬度", "经度"}
	row := sheet.AddRow()
	for _, header := range headers {
		cell := row.AddCell()
		cell.Value = header
	}

	// 添加数据
	for _, host := range hosts {
		row := sheet.AddRow()
		row.AddCell().Value = host.IPString
		row.AddCell().SetValue(host.Port)
		row.AddCell().Value = fmt.Sprintf("%v", host.Domains)
		row.AddCell().Value = fmt.Sprintf("%v", host.Hostnames)
		row.AddCell().Value = host.ISP
		row.AddCell().Value = host.ASN
		row.AddCell().Value = host.Org
		row.AddCell().Value = host.Data
		row.AddCell().Value = host.Location.City
		row.AddCell().Value = host.Location.CountryName
		row.AddCell().Value = host.OS
		row.AddCell().Value = host.Timestamp
		row.AddCell().SetValue(host.Location.Latitude)
		row.AddCell().SetValue(host.Location.Longitude)
	}

	err = file.Save(filename)
	if err != nil {
		return err
	}

	return nil
}
