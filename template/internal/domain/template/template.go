package template

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"
)

type TemplateType int

const (
	SMS   TemplateType = 0
	Email TemplateType = 1
)

var (
	contentRegex = regexp.MustCompile(`<body>.*{{content}}.*</body>`)
	scriptRegex  = regexp.MustCompile(`(?i)<\s*script`)
)

type Template struct {
	uuid      string
	content   string
	tmplType  TemplateType
	createdAt time.Time
}

func NewTemplate(uuid string, content string, tmplType TemplateType, createdAt time.Time) (*Template, error) {
	if tmplType != SMS && tmplType != Email {
		return nil, errors.New("invalid template type")
	}

	if content == "" {
		return nil, errors.New("empty template content")
	}

	if createdAt.IsZero() {
		return nil, errors.New("invalid template creation date")
	}

	return &Template{
		uuid:      uuid,
		content:   content,
		tmplType:  tmplType,
		createdAt: createdAt,
	}, nil
}

func UnmarshalFromDB(
	uuid string,
	content string,
	tmplType TemplateType,
	createdAt time.Time,
) (*Template, error) {
	return NewTemplate(uuid, content, tmplType, createdAt)
}

func (t Template) Validate() error {
	decoder := xml.NewDecoder(strings.NewReader(t.content))

	tags := map[string]bool{
		"html": false,
		"head": false,
		"body": false,
	}

	var syntaxError *xml.SyntaxError

	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}

			if errors.As(err, &syntaxError) {
				return fmt.Errorf("template content contains syntax error: %s", syntaxError.Msg)
			}

			return err
		}

		switch token := t.(type) {
		case xml.StartElement:
			if _, ok := tags[token.Name.Local]; ok {
				tags[token.Name.Local] = true
			}
		}
	}

	for tag, found := range tags {
		if !found {
			return fmt.Errorf("template content must contain <%s> tag", tag)
		}
	}

	if !contentRegex.MatchString(t.content) {
		return errors.New("template content must contain {{content}} tag inside <body> tag")
	}

	if scriptRegex.MatchString(t.content) {
		return errors.New("template content must not contain <script> tag")
	}

	return nil
}

func (t Template) Enrich(message string) string {
	return strings.Replace(t.content, "{{content}}", message, 1)
}

func (t Template) UUID() string {
	return t.uuid
}

func (t Template) Content() string {
	return t.content
}

func (t Template) TmplType() TemplateType {
	return t.tmplType
}

func (t Template) CreatedAt() time.Time {
	return t.createdAt
}
