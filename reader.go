package PdfReader

import "github.com/ledongthuc/pdf"

func Read(path string) string {
	file, reader, err := pdf.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	content := ""
	for i := 1; i <= reader.NumPage(); i++ {
		text := reader.Page(i).Content().Text
		for _, item := range text {
			content += item.S
		}
	}
	return content
}
