package cleaner

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// noisePattern holds a compiled regex and a description for debugging.
type noisePattern struct {
	re   *regexp.Regexp
	desc string
}

// patterns is the ordered list of structural noise to strip.
// Order matters: broader patterns first, then specifics.
var patterns = []noisePattern{
	// Bracketed / parenthesised metadata — strip early so content inside
	// doesn't interfere with later patterns
	{regexp.MustCompile(`(?i)\[[^\]]*\]`), "brackets"},
	{regexp.MustCompile(`(?i)\([^)]*\)`), "parentheses"},

	// Disc / volume markers
	{regexp.MustCompile(`(?i)\bDisc\s*\d+(?:\s*of\s*\d+)?\b`), "disc number"},
	{regexp.MustCompile(`(?i)\bVol(?:ume)?\.?\s*\d+\b`), "volume number"},
	{regexp.MustCompile(`(?i)\bEpisodes?\s*[\d\-–]+\b`), "episode range"},
	{regexp.MustCompile(`(?i)\bPart\s*\d+\b`), "part number"},

	// Edition / release type
	{regexp.MustCompile(`(?i)\b(?:Standard|Limited|Collector'?s?|Ultimate|Premium|Anniversary|Deluxe|Special|Theatrical|Director'?s?\s*Cut|Extended|Unrated|Remastered|Restored|Criterion)\s*(?:Edition|Cut|Version|Release|Set)?\b`), "edition type"},

	// Format tags
	{regexp.MustCompile(`(?i)\b4K[\s\-]*UHD\b`), "4K UHD"},
	{regexp.MustCompile(`(?i)\bUHD\b`), "UHD"},
	{regexp.MustCompile(`(?i)\bBlu[\s\-]?[Rr]ay\b`), "blu-ray"},
	{regexp.MustCompile(`(?i)\bBD\b`), "BD"},
	{regexp.MustCompile(`(?i)\bDVD\b`), "DVD"},
	{regexp.MustCompile(`(?i)\bSteelbook\b`), "steelbook"},
	{regexp.MustCompile(`(?i)\bDigibook\b`), "digibook"},
	{regexp.MustCompile(`(?i)\bSlipcase\b`), "slipcase"},

	// Set / collection suffixes
	{regexp.MustCompile(`(?i)\bComplete\s+(?:Series|Season|Collection|Set|Box\s*Set)?\b`), "complete set"},
	{regexp.MustCompile(`(?i)\b\d+[\s\-]*Disc\s*(?:Set|Box)?\b`), "disc set"},
	{regexp.MustCompile(`(?i)\bBox\s*Set\b`), "box set"},
	{regexp.MustCompile(`(?i)\bCollection\b`), "collection"},

	// Region / format codes
	{regexp.MustCompile(`(?i)\bRegion\s*[A-C0-9]\b`), "region code"},
	{regexp.MustCompile(`(?i)\bNTSC\b`), "NTSC"},
	{regexp.MustCompile(`(?i)\bPAL\b`), "PAL"},

	// season labels that follow the main title
	{regexp.MustCompile(`(?i)\b(?:Season|Series|Chapter|Saga|Cour)\s*\d*\b`), "season label"},

	// Truncation artefacts
	{regexp.MustCompile(`\s*\.{2,}\s*`), "ellipsis"},
	{regexp.MustCompile(`\s*…\s*`), "unicode ellipsis"},

	// Trailing separators — run last after other noise is gone
	{regexp.MustCompile(`[\s\-–—|:,]+$`), "trailing separators"},
}

var multiSpace = regexp.MustCompile(`\s{2,}`)

// seasonRe matches the many ways a season number appears on packaging:
//
//	"Season 2", "Season Two", "S2", "Series 3", "Cour 2"
//	"The Final Season" → no number, SeasonNumber stays 0
//	"Part 2" without a season word is NOT captured here (ambiguous)
var seasonRe = regexp.MustCompile(
	`(?i)\b(?:Season|Series|Cour)\s*` +
		`(?:` +
		`(\d+)` + // numeric:  Season 2
		`|` +
		`(one|two|three|four|five|six|seven|eight|nine|ten)` + // word: Season Two
		`)`,
)

var wordToInt = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
	"six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10,
}

// ParsedTitle is the result of a full parse: clean title + any extracted metadata.
type ParsedTitle struct {
	// Title is the cleaned, noise-free title ready for an API search.
	Title string

	// SeasonNumber is the extracted season, or 0 if none was found.
	SeasonNumber int

	// HasSeason is true when a season label was present, even if no
	// number could be extracted (e.g. "The Final Season").
	HasSeason bool
}

// Parse extracts the clean title and season number from a raw packaging string.
// It captures season information before stripping it so nothing is lost.
//
// Examples:
//
//	Parse("My Hero Academia Season 4 Complete Blu-ray Box Set")
//	// → {Title: "My Hero Academia", SeasonNumber: 4, HasSeason: true}
//
//	Parse("Attack on Titan - The Final Season [Blu-ray]")
//	// → {Title: "Attack on Titan", SeasonNumber: 0, HasSeason: true}
//
//	Parse("Akira [Blu-ray]")
//	// → {Title: "Akira", SeasonNumber: 0, HasSeason: false}
func Parse(raw string) ParsedTitle {
	result := ParsedTitle{}

	// 1. Extract season before the noise-stripping pass removes it.
	if m := seasonRe.FindStringSubmatch(raw); m != nil {
		result.HasSeason = true
		switch {
		case m[1] != "": // numeric form
			n, err := strconv.Atoi(m[1])
			if err == nil {
				result.SeasonNumber = n
			}
		case m[2] != "": // word form
			result.SeasonNumber = wordToInt[strings.ToLower(m[2])]
		}
	} else if hasFinalSeason(raw) {
		// "The Final Season" / "Final Series" — season present but no number
		result.HasSeason = true
	}

	// 2. Full noise strip to get the clean title.
	result.Title = PreClean(raw)
	return result
}

// hasFinalSeason detects "The Final Season/Series/Cour" without a number.
var finalSeasonRe = regexp.MustCompile(`(?i)\b(?:final|last)\s+(?:season|series|cour)\b`)

func hasFinalSeason(s string) bool {
	return finalSeasonRe.MatchString(s)
}

// PreClean strips structural noise from a raw Blu-ray/DVD packaging string
// and returns the best single-string candidate for further processing.
//
// Example:
//
//	PreClean("Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Disc 1")
//	// → "Demon Slayer: Kimetsu no Yaiba Swordsmith Village"
func PreClean(raw string) string {
	s := raw
	for _, p := range patterns {
		s = p.re.ReplaceAllString(s, " ")
	}
	s = multiSpace.ReplaceAllString(s, " ")
	s = strings.Trim(s, " :-–—|")
	return s
}

// PreCleanDebug returns the cleaned string and a log of which patterns fired,
// useful for tuning and diagnosing unexpected results.
func PreCleanDebug(raw string) (cleaned string, log []string) {
	s := raw
	for _, p := range patterns {
		replaced := p.re.ReplaceAllString(s, " ")
		if replaced != s {
			log = append(log, p.desc)
			s = replaced
		}
	}
	s = multiSpace.ReplaceAllString(s, " ")
	s = strings.Trim(s, " :-–—|")
	return s, log
}

// Normalise returns a lowercase, accent-stripped, whitespace-collapsed version
// of text — used for comparison only, never shown to the user.
func Normalise(text string) string {
	// Decompose unicode, then drop non-ASCII
	t := norm.NFD.String(text)
	var b strings.Builder
	for _, r := range t {
		if r <= unicode.MaxASCII {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	s := multiSpace.ReplaceAllString(b.String(), " ")
	return strings.TrimSpace(s)
}
