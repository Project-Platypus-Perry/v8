package constants

import "database/sql/driver"

// Language constants
type Language string

const (
	LanguageEnglish Language = "en"
	LanguageHindi   Language = "hi"
)

func (p *Language) Scan(value interface{}) error {
	*p = Language(value.(string))
	return nil
}

func (p Language) Value() (driver.Value, error) {
	return string(p), nil
}

// Visibility constants
type Visibility string

const (
	PrivateVisibility Visibility = "private"
	PublicVisibility  Visibility = "public"
)

func (p *Visibility) Scan(value interface{}) error {
	*p = Visibility(value.(string))
	return nil
}

func (p Visibility) Value() (driver.Value, error) {
	return string(p), nil
}

// Content type constants
type ContentType string

const (
	ContentTypeNotes ContentType = "notes"
	ContentTypeDPP   ContentType = "dpp"
	ContentTypeVideo ContentType = "video"
)

func (p *ContentType) Scan(value interface{}) error {
	*p = ContentType(value.(string))
	return nil
}

func (p ContentType) Value() (driver.Value, error) {
	return string(p), nil
}
