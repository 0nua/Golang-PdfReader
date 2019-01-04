package PdfReader

import "ledongthuc/pdf/pdf"

func Read(name string) string {
	file, reader, err := pdf.Open("./pdf/" + name)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	text := ""
	for i := 1; i <= reader.NumPage(); i++ {
		content := reader.Page(i).Content().Text
		for _, item := range content {
			text += item.S
		}
	}

	return text
}
