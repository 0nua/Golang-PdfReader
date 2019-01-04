package PdfReader

import "github.com/ledongthuc/pdf"

func Read(path string) string {
	file, reader, err := pdf.Open(path)
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
