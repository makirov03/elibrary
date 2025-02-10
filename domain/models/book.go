package models

type Book struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	FilePath   string `json:"file_path"`
	UploadedBy string `json:"uploaded_by"`
}
