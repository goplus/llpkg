package emoji

import (
	"github.com/goplus/lib/py"
	_ "unsafe"
)

const LLGoPackage = "py.emoji"
//
//     Replace emoji names in a string with Unicode codes.
//         >>> import emoji
//         >>> print(emoji.emojize("Python is fun :thumbsup:", language='alias'))
//         Python is fun 👍
//         >>> print(emoji.emojize("Python is fun :thumbs_up:"))
//         Python is fun 👍
//         >>> print(emoji.emojize("Python is fun {thumbs_up}", delimiters = ("{", "}")))
//         Python is fun 👍
//         >>> print(emoji.emojize("Python is fun :red_heart:", variant="text_type"))
//         Python is fun ❤
//         >>> print(emoji.emojize("Python is fun :red_heart:", variant="emoji_type"))
//         Python is fun ❤️ # red heart, not black heart
//
//     :param string: String contains emoji names.
//     :param delimiters: (optional) Use delimiters other than _DEFAULT_DELIMITER. Each delimiter
//         should contain at least one character that is not part of a-zA-Z0-9 and ``_-&.()!?#*+,``.
//         See ``emoji.core._EMOJI_NAME_PATTERN`` for the regular expression of unsafe characters.
//     :param variant: (optional) Choose variation selector between "base"(None), VS-15 ("text_type") and VS-16 ("emoji_type")
//     :param language: Choose language of emoji name: language code 'es', 'de', etc. or 'alias'
//         to use English aliases
//     :param version: (optional) Max version. If set to an Emoji Version,
//         all emoji above this version will be ignored.
//     :param handle_version: (optional) Replace the emoji above ``version``
//         instead of ignoring it. handle_version can be either a string or a
//         callable; If it is a callable, it's passed the Unicode emoji and the
//         data dict from :data:`EMOJI_DATA` and must return a replacement string
//         to be used::
//
//             handle_version('\U0001F6EB', {
//                 'en' : ':airplane_departure:',
//                 'status' : fully_qualified,
//                 'E' : 1,
//                 'alias' : [':flight_departure:'],
//                 'de': ':abflug:',
//                 'es': ':avión_despegando:',
//                 ...
//             })
//
//     :raises ValueError: if ``variant`` is neither None, 'text_type' or 'emoji_type'
//
//
//
//go:linkname Emojize py.emojize
func Emojize(string *py.Object, delimiters *py.Object, variant *py.Object, language *py.Object, version *py.Object, handleVersion *py.Object) *py.Object
//
//     Replace Unicode emoji in a string with emoji shortcodes. Useful for storage.
//         >>> import emoji
//         >>> print(emoji.emojize("Python is fun :thumbs_up:"))
//         Python is fun 👍
//         >>> print(emoji.demojize("Python is fun 👍"))
//         Python is fun :thumbs_up:
//         >>> print(emoji.demojize("icode is tricky 😯", delimiters=("__", "__")))
//         Unicode is tricky __hushed_face__
//
//     :param string: String contains Unicode characters. MUST BE UNICODE.
//     :param delimiters: (optional) User delimiters other than ``_DEFAULT_DELIMITER``
//     :param language: Choose language of emoji name: language code 'es', 'de', etc. or 'alias'
//         to use English aliases
//     :param version: (optional) Max version. If set to an Emoji Version,
//         all emoji above this version will be removed.
//     :param handle_version: (optional) Replace the emoji above ``version``
//         instead of removing it. handle_version can be either a string or a
//         callable ``handle_version(emj: str, data: dict) -> str``; If it is
//         a callable, it's passed the Unicode emoji and the data dict from
//         :data:`EMOJI_DATA` and must return a replacement string  to be used.
//         The passed data is in the form of::
//
//             handle_version('\U0001F6EB', {
//                 'en' : ':airplane_departure:',
//                 'status' : fully_qualified,
//                 'E' : 1,
//                 'alias' : [':flight_departure:'],
//                 'de': ':abflug:',
//                 'es': ':avión_despegando:',
//                 ...
//             })
//
//
//
//go:linkname Demojize py.demojize
func Demojize(string *py.Object, delimiters *py.Object, language *py.Object, version *py.Object, handleVersion *py.Object) *py.Object
//
//     Find unicode emoji in a string. Yield each emoji as a named tuple
//     :class:`Token` ``(chars, EmojiMatch)`` or `:class:`Token` ``(chars, EmojiMatchZWJNonRGI)``.
//     If ``non_emoji`` is True, also yield all other characters as
//     :class:`Token` ``(char, char)`` .
//
//     :param string: String to analyze
//     :param non_emoji: If True also yield all non-emoji characters as Token(char, char)
//     :param join_emoji: If True, multiple EmojiMatch are merged into a single
//         EmojiMatchZWJNonRGI if they are separated only by a ZWJ.
//
//
//go:linkname Analyze py.analyze
func Analyze(string *py.Object, nonEmoji *py.Object, joinEmoji *py.Object) *py.Object
//
//     Returns the location and emoji in list of dict format.
//         >>> emoji.emoji_list("Hi, I am fine. 😁")
//         [{'match_start': 15, 'match_end': 16, 'emoji': '😁'}]
//
//
//go:linkname EmojiList py.emoji_list
func EmojiList(string *py.Object) *py.Object
// Returns distinct list of emojis from the string.
//
//go:linkname DistinctEmojiList py.distinct_emoji_list
func DistinctEmojiList(string *py.Object) *py.Object
//
//     Returns the count of emojis in a string.
//
//     :param unique: (optional) True if count only unique emojis
//
//
//go:linkname EmojiCount py.emoji_count
func EmojiCount(string *py.Object, unique *py.Object) *py.Object
//
//     Replace Unicode emoji in a customizable string.
//
//     :param string: String contains Unicode characters. MUST BE UNICODE.
//     :param replace: (optional) replace can be either a string or a callable;
//         If it is a callable, it's passed the Unicode emoji and the data dict from
//         :data:`EMOJI_DATA` and must return a replacement string to be used.
//         replace(str, dict) -> str
//     :param version: (optional) Max version. If set to an Emoji Version,
//         only emoji above this version will be replaced.
//
//
//go:linkname ReplaceEmoji py.replace_emoji
func ReplaceEmoji(string *py.Object, replace *py.Object, version *py.Object) *py.Object
//
//     Returns True if the string is a single emoji, and it is "recommended for
//     general interchange" by Unicode.org.
//
//
//go:linkname IsEmoji py.is_emoji
func IsEmoji(string *py.Object) *py.Object
//
//     Returns True if the string contains only emojis.
//     This might not imply that `is_emoji` for all the characters, for example,
//     if the string contains variation selectors.
//
//
//go:linkname PurelyEmoji py.purely_emoji
func PurelyEmoji(string *py.Object) *py.Object
//
//     Returns the Emoji Version of the emoji.
//
//     See https://www.unicode.org/reports/tr51/#Versioning for more information.
//         >>> emoji.version("😁")
//         0.6
//         >>> emoji.version(":butterfly:")
//         3
//
//     :param string: An emoji or a text containing an emoji
//     :raises ValueError: if ``string`` does not contain an emoji
//
//
//go:linkname Version py.version
func Version(string *py.Object) *py.Object
// Generate dict containing all fully-qualified and component emoji name for a language
//     The dict is only generated once per language and then cached in _EMOJI_UNICODE[lang]
//
//go:linkname GetEmojiUnicodeDict py.get_emoji_unicode_dict
func GetEmojiUnicodeDict(lang *py.Object) *py.Object
// Generate dict containing all fully-qualified and component aliases
//     The dict is only generated once and then cached in _ALIASES_UNICODE
//
//go:linkname GetAliasesUnicodeDict py.get_aliases_unicode_dict
func GetAliasesUnicodeDict() *py.Object
