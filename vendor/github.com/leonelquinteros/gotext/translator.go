/*
 * Copyright (c) 2018 DeineAgentur UG https://www.deineagentur.com. All rights reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for full license information.
 */

package gotext

import "net/textproto"

// Translator interface is used by Locale and Po objects.Translator
// It contains all methods needed to parse translation sources and obtain corresponding translations.
// Also implements gob.GobEncoder/gob.DobDecoder interfaces to allow serialization of Locale objects.
type Translator interface {
	ParseFile(f string)
	Parse(buf []byte)
	Get(str string, vars ...interface{}) string
	GetN(str, plural string, n int, vars ...interface{}) string
	GetC(str, ctx string, vars ...interface{}) string
	GetNC(str, plural string, n int, ctx string, vars ...interface{}) string

	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

// TranslatorEncoding is used as intermediary storage to encode Translator objects to Gob.
type TranslatorEncoding struct {
	// Headers storage
	Headers textproto.MIMEHeader

	// Language header
	Language string

	// Plural-Forms header
	PluralForms string

	// Parsed Plural-Forms header values
	Nplurals int
	Plural   string

	// Storage
	Translations map[string]*Translation
	Contexts     map[string]map[string]*Translation
}

// GetTranslator is used to recover a Translator object after unmarshaling the TranslatorEncoding object.
// Internally uses a Po object as it should be switcheable with Mo objects without problem.
// External Translator implementations should be able to serialize into a TranslatorEncoding object in order to unserialize into a Po-compatible object.
func (te *TranslatorEncoding) GetTranslator() Translator {
	po := new(Po)
	po.Headers = te.Headers
	po.Language = te.Language
	po.PluralForms = te.PluralForms
	po.nplurals = te.Nplurals
	po.plural = te.Plural
	po.translations = te.Translations
	po.contexts = te.Contexts

	return po
}
