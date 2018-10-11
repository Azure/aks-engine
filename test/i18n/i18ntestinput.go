package fake

import (
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
)

// This file is used to generate the test translation file
// 1. go-xgettext -o i18ntestinput.pot --keyword=translator.T --keyword-plural=translator.NT --msgid-bugs-address="" --sort-output test/i18n/i18ntestinput.go
// 2. go-xgettext -o i18ntestinput.err.pot --keyword=translator.Errorf --keyword-plural=translator.NErrorf --msgid-bugs-address="" --sort-output test/i18n/i18ntestinput.go
// 3. sed '1,18d' i18ntestinput.err.pot >> i18ntestinput.pot
// 4. msginit -l en_US -o i18ntestinput.po -i i18ntestinput.pot
// 5. Modify i18ntestinput.po using poedit as necessary
// Or msgfmt -c -v -o i18ntestinput.mo i18ntestinput.po
// 6. for d in "en_US"; do cp i18ntestinput.mo translations/test/$d/LC_MESSAGES/aksengine.mo; cp i18ntestinput.po translations/test/$d/LC_MESSAGES/aksengine.po; done
// 7. rm i18ntestinput.*

var (
	locale     = gotext.NewLocale("d", "l")
	translator = &i18n.Translator{
		Locale: locale,
	}
	world    = "World"
	resource = "Foo"
)

func aloha() {
	translator.T("Aloha")
}

func foo() {
	translator.T("Hello %s", world)
}

func bar() {
	translator.NT("There is %d parameter in resource %s", "There are %d parameters in resource %s", 9, 9, resource)
}

func file() error {
	return translator.Errorf("File not exists")
}

func api() error {
	return translator.NErrorf("There is %d error in the api model", "There are %d errors in the api model", 3, 3)
}
