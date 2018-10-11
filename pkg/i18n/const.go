package i18n

const (
	defaultLanguage   = "en_US"
	defaultDomain     = "aksengine"
	defaultLocalDir   = "translations"
	defaultMessageDir = "LC_MESSAGES"
)

var supportedTranslations = map[string]bool{
	defaultLanguage: true,
	"cs_CZ":         true,
	"de_DE":         true,
	"es_ES":         true,
	"fr_FR":         true,
	"hu_HU":         true,
	"it_IT":         true,
	"ja_JP":         true,
	"ko_KR":         true,
	"nl_NL":         true,
	"pl_PL":         true,
	"pt_BR":         true,
	"pt_PT":         true,
	"ru_RU":         true,
	"sv_SE":         true,
	"tr_TR":         true,
	"zh_CN":         true,
	"zh_TW":         true,
}
