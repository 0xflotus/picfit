package image

var Extensions = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/bmp":  "bmp",
	"image/gif":  "gif",
	"image/webp": "webp",
}

var HeaderKeys = []string{
	"Age",
	"Content-Type",
	"Last-Modified",
	"Date",
	"Etag",
}
