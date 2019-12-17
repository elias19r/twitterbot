package text

import (
	"regexp"
	"strings"
)

// List of regular expression patterns.
var regexpPatterns = map[string]string{}

// List of *Regexp for each pattern.
var regexps = map[string]*regexp.Regexp{}

func init() {
	// Populate regexpPatterns.
	regexpPatterns["hashtag"] = `(?i:#\p{L}+)`
	regexpPatterns["punctuation"] = `(?:[!~&*:?\]\[|}{]+)`
	regexpPatterns["whitespace"] = `(?:[[:space:]]+)`
	regexpPatterns["emoji"] = `(?:[\x{1F600}-\x{1F64F}]|[\x{1F680}-\x{1F6FF}])`

	// Create *Regexp objects.
	for pattern, s := range regexpPatterns {
		regexps[pattern] = regexp.MustCompile(s)
	}
}

// Clear removes emojis, hashtags, punctuation and normalize whitespaces.
func Clear(s string) string {
	s = StripEmoji(s)
	s = StripHashtag(s)
	s = StripPunctuation(s)
	s = Normalize(s)
	return s
}

// Strip strips a specified pattern from s.
func Strip(s string, pattern string) string {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return s
	}
	return r.ReplaceAllString(s, " ")
}

// StripEmoji strips all emojis from s.
func StripEmoji(s string) string {
	return regexps["emoji"].ReplaceAllString(s, " ")
}

// StripHashtag strips all hashtags from s.
func StripHashtag(s string) string {
	return regexps["hashtag"].ReplaceAllString(s, " ")
}

// StripPunctuation strips all punctuation from s.
func StripPunctuation(s string) string {
	return regexps["punctuation"].ReplaceAllString(s, " ")
}

// Normalize trims and replaces duplicate whitespaces with a single space.
func Normalize(s string) string {
	return regexps["whitespace"].ReplaceAllString(strings.TrimSpace(s), " ")
}

// Truncate truncates s to n-1 runes, if s has more than n runes, appending "…"
func Truncate(s string, n int) string {
	r := []rune(s)

	if len(r) > n {
		r = r[0:n]
		r[len(r)-1] = '…'
	}

	return string(r)
}

// GetHashtags returns a map with the hashtags in s.
func GetHashtags(s string) (values []string, ok bool) {
	if ok = regexps["hashtag"].MatchString(s); ok {
		values = append(values, regexps["hashtag"].FindAllString(s, -1)...)
	}
	return
}
