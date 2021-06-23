package formatters

import "FauzulPasswordManager/entities"

type PassFormatter struct {
	Id   int    `json:"id"`
	Web  string `json:"web"`
	Pass string `json:"pass"`
}

func PassFormat(data entities.Password) PassFormatter {
	return PassFormatter{
		Id:   data.Id,
		Web:  data.Web,
		Pass: data.Pass,
	}
}

func AllPasssFormat(datas []entities.Password) []PassFormatter {
	var formatter []PassFormatter
	for _, data := range datas {
		formatter = append(formatter, PassFormat(data))
	}

	return formatter
}
