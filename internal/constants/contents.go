package constants

// Language constants
type Language string

const (
	LanguageEnglish Language = "en"
	LanguageHindi   Language = "hi"
)

// Visibility constants
type Visibility string

const (
	PrivateVisibility Visibility = "private"
	PublicVisibility  Visibility = "public"
)

// Content type constants
type ContentType string

const (
	ContentTypeNotes ContentType = "notes"
	ContentTypeDPP   ContentType = "dpp"
	ContentTypeVideo ContentType = "video"
)
