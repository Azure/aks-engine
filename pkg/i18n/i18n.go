// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package i18n

import (
	"fmt"
	"os"
	"strings"

	"path"

	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
)

func loadSystemLanguage() string {
	language := os.Getenv("LANG")
	if language == "" {
		return defaultLanguage
	}

	// Posix locale name usually has the ll_CC.encoding syntax.
	parts := strings.Split(language, ".")
	if len(parts) == 0 {
		return defaultLanguage
	}
	if _, ok := supportedTranslations[parts[0]]; ok {
		return parts[0]
	}
	return defaultLanguage
}

// LoadTranslations loads translation files and sets the locale to
// the system locale. It should be called by the main program.
func LoadTranslations() (*gotext.Locale, error) {
	lang := loadSystemLanguage()
	SetLanguage(lang)

	dir := path.Join(defaultLocalDir, lang, defaultMessageDir)
	translationFiles := []string{
		path.Join(dir, fmt.Sprintf("%s.mo", defaultDomain)),
		path.Join(dir, fmt.Sprintf("%s.po", defaultDomain)),
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return nil, err
		}
	}

	for _, file := range translationFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			data, err := Asset(file)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(file, data, 0600)
			if err != nil {
				return nil, err
			}
		}
	}

	locale := gotext.NewLocale(defaultLocalDir, lang)
	Initialize(locale)

	return locale, nil
}

// Initialize is the translation initialization function shared by the main program and package.
func Initialize(locale *gotext.Locale) error {
	if locale == nil {
		return errors.New("Initialize expected locale but got nil")
	}
	locale.AddDomain(defaultDomain)
	return nil
}

// SetLanguage sets the program's current locale. If the language is not
// supported, then the default locale is used.
func SetLanguage(language string) {
	if _, ok := supportedTranslations[language]; ok {
		gotext.SetLanguage(language)
		return
	}
	gotext.SetLanguage(defaultLanguage)
}

// GetLanguage queries the program's current locale.
func GetLanguage() string {
	return gotext.GetLanguage()
}

// Translator is a wrapper over gotext's Locale and provides interface to
// translate text string and produce translated error
type Translator struct {
	Locale *gotext.Locale
}

// T translates a text string, based on GNU's gettext library.
func (t *Translator) T(msgid string, vars ...interface{}) string {
	if t.Locale == nil {
		return fmt.Sprintf(msgid, vars...)
	}
	return t.Locale.GetD(defaultDomain, msgid, vars...)
}

// NT translates a text string into the appropriate plural form, based on GNU's gettext library.
func (t *Translator) NT(msgid, msgidPlural string, n int, vars ...interface{}) string {
	if t.Locale == nil {
		return fmt.Sprintf(msgidPlural, vars...)
	}
	return t.Locale.GetND(defaultDomain, msgid, msgidPlural, n, vars...)
}

// Errorf produces an error with a translated error string.
func (t *Translator) Errorf(msgid string, vars ...interface{}) error {
	return errors.New(t.T(msgid, vars...))
}

// NErrorf produces an error with a translated error string in the appropriate plural form.
func (t *Translator) NErrorf(msgid, msgidPlural string, n int, vars ...interface{}) error {
	return errors.New(t.NT(msgid, msgidPlural, n, vars...))
}
