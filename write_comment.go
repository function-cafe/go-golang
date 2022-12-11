package golang

import (
	"io"
	"strings"
)

// WriteComment writes a Go comment to the writer.
// The comment is indented by the number of tabs specified by indent.
func WriteComment(w io.Writer, text string, indent int) error {
	if text == "" {
		return nil
	}
	paragraphs := strings.Split(text, "\n")
	for i, paragraph := range paragraphs {
		if i > 0 {
			if _, err := w.Write([]byte(strings.Repeat("\t", indent))); err != nil {
				return err
			}
			if _, err := w.Write([]byte("//\n")); err != nil {
				return err
			}
		}
		sentences := strings.Split(paragraph, ".")
		for _, sentence := range sentences {
			t := strings.TrimSpace(sentence)
			if t == "" {
				continue
			}
			if _, err := w.Write([]byte(strings.Repeat("\t", indent))); err != nil {
				return err
			}
			if _, err := w.Write([]byte("// " + t + ".\n")); err != nil {
				return err
			}
		}
	}
	return nil
}
