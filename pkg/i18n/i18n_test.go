// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package i18n

import (
	"os"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLoadTranslations(t *testing.T) {
	RegisterTestingT(t)

	origLang := os.Getenv("LANG")
	if origLang != "" && !strings.HasPrefix(origLang, "en_US") {
		// The unit test has only en_US translation files
		return
	}

	_, err := LoadTranslations()
	Expect(err).Should(BeNil())
}

func TestTranslationLanguage(t *testing.T) {
	RegisterTestingT(t)

	origLang := os.Getenv("LANG")
	os.Setenv("LANG", "en_US.UTF-8")
	_, err := LoadTranslations()
	Expect(err).Should(BeNil())

	lang := GetLanguage()
	Expect(lang).Should(Equal("en_US"))

	os.Setenv("LANG", origLang)
}

func TestTranslationLanguageDefault(t *testing.T) {
	RegisterTestingT(t)

	origLang := os.Getenv("LANG")
	os.Setenv("LANG", "ll_CC.UTF-8")
	_, err := LoadTranslations()
	Expect(err).Should(BeNil())

	lang := GetLanguage()
	Expect(lang).Should(Equal(defaultLanguage))

	os.Setenv("LANG", origLang)
}

func TestTranslations(t *testing.T) {
	RegisterTestingT(t)

	origLang := os.Getenv("LANG")
	if origLang != "" && !strings.HasPrefix(origLang, "en_US") {
		// The unit test has only en_US translation files
		return
	}

	l, err := LoadTranslations()
	Expect(err).Should(BeNil())

	translator := &Translator{
		Locale: l,
	}

	msg := translator.T("Aloha")
	Expect(msg).Should(Equal("Aloha"))

	msg = translator.T("Hello %s", "World")
	Expect(msg).Should(Equal("Hello World"))
}

func TestTranslationsPlural(t *testing.T) {
	RegisterTestingT(t)

	origLang := os.Getenv("LANG")
	if origLang != "" && !strings.HasPrefix(origLang, "en_US") {
		// The unit test has only en_US translation files
		return
	}

	l, err := LoadTranslations()
	Expect(err).Should(BeNil())

	translator := &Translator{
		Locale: l,
	}

	msg := translator.NT("There is %d parameter in resource %s", "There are %d parameters in resource %s", 1, 1, "Foo")
	Expect(msg).Should(Equal("There is 1 parameter in resource Foo"))

	msg = translator.NT("There is %d parameter in resource %s", "There are %d parameters in resource %s", 9, 9, "Foo")
	Expect(msg).Should(Equal("There are 9 parameters in resource Foo"))
}

func TestTranslationsError(t *testing.T) {
	RegisterTestingT(t)

	origLang := os.Getenv("LANG")
	if origLang != "" && !strings.HasPrefix(origLang, "en_US") {
		// The unit test has only en_US translation files
		return
	}

	l, err := LoadTranslations()
	Expect(err).Should(BeNil())

	translator := &Translator{
		Locale: l,
	}

	e := translator.Errorf("File not exists")
	Expect(e.Error()).Should(Equal("File not exists"))

	e = translator.NErrorf("There is %d error in the api model", "There are %d errors in the api model", 3, 3)
	Expect(e.Error()).Should(Equal("There are 3 errors in the api model"))
}
