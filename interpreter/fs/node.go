package fs

type NodeType string

const (
	Dir  NodeType = "dir"
	File NodeType = "file"
)

func (t NodeType) IsDir() bool { return t == Dir }

type Node struct {
	Type     NodeType
	Children []string
	Content  string
	Meta     Meta
	Hidden   bool
}

type Meta struct {
	Size      int64
	Modified  string
	ImagePath string // URL path for full-size image
	ThumbPath string // URL path for thumbnail
	ImgWidth  int    // natural pixel width of image
	ImgHeight int    // natural pixel height of image
	MimeType  string // MIME type for file type detection
}
