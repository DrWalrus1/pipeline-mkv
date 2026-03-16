package cleaner_test

import (
	"testing"

	"github.com/DrWalrus1/pipelinemkv/internal/metadata/cleaner"
)

var cleanTests = []struct {
	name    string
	input   string
	wantSub string // cleaned output should contain this substring
}{
	{
		name:    "your original example — truncated arc name",
		input:   "Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1",
		wantSub: "Demon Slayer: Kimetsu no Yaiba",
	},
	{
		name:    "brackets and region code",
		input:   "Spirited Away [Studio Ghibli] 4K UHD Blu-ray Limited Collector's Edition Steelbook (Region B)",
		wantSub: "Spirited Away",
	},
	{
		name:    "season label",
		input:   "My Hero Academia Season 4 Complete Series Blu-ray Box Set",
		wantSub: "My Hero Academia",
	},
	{
		name:    "remaster + DVD",
		input:   "Neon Genesis Evangelion - Perfect Collection Remastered Edition DVD",
		wantSub: "Neon Genesis Evangelion",
	},
	{
		name:    "minimal noise",
		input:   "Akira [Blu-ray]",
		wantSub: "Akira",
	},
	{
		name:    "extended edition with year",
		input:   "The Lord of the Rings: The Fellowship of the Ring - Extended Edition 4K UHD [Blu-ray] (2001)",
		wantSub: "The Lord of the Rings: The Fellowship of the Ring",
	},
	{
		name:    "year in parens + format",
		input:   "Dune: Part One (2021) [4K UHD + Blu-ray]",
		wantSub: "Dune",
	},
	{
		name:    "unicode ellipsis",
		input:   "Attack on Titan - The Final Season Part 2… [Blu-ray]",
		wantSub: "Attack on Titan",
	},
}

func TestPreClean(t *testing.T) {
	for _, tt := range cleanTests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleaner.PreClean(tt.input)
			if !contains(got, tt.wantSub) {
				t.Errorf("\n  input : %q\n  got   : %q\n  want  : contains %q", tt.input, got, tt.wantSub)
			}
		})
	}
}

func TestPreCleanDebug(t *testing.T) {
	input := "Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1"
	cleaned, log := cleaner.PreCleanDebug(input)
	t.Logf("cleaned : %q", cleaned)
	t.Logf("fired   : %v", log)
	if len(log) == 0 {
		t.Error("expected at least one pattern to fire")
	}
}

var parseTests = []struct {
	name          string
	input         string
	wantTitle     string
	wantSeason    int
	wantHasSeason bool
}{
	{
		name:          "numeric season",
		input:         "My Hero Academia Season 4 Complete Series Blu-ray Box Set",
		wantTitle:     "My Hero Academia",
		wantSeason:    4,
		wantHasSeason: true,
	},
	{
		name:          "word season",
		input:         "Some Show Season Two [Blu-ray]",
		wantTitle:     "Some Show",
		wantSeason:    2,
		wantHasSeason: true,
	},
	{
		name:          "final season — no number",
		input:         "Attack on Titan - The Final Season [Blu-ray]",
		wantTitle:     "Attack on Titan",
		wantSeason:    0,
		wantHasSeason: true,
	},
	{
		name:          "series label",
		input:         "Fullmetal Alchemist Brotherhood Series 2 DVD",
		wantTitle:     "Fullmetal Alchemist Brotherhood",
		wantSeason:    2,
		wantHasSeason: true,
	},
	{
		name:          "no season — film",
		input:         "Akira [Blu-ray]",
		wantTitle:     "Akira",
		wantSeason:    0,
		wantHasSeason: false,
	},
	{
		name:          "your original example",
		input:         "Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1",
		wantTitle:     "Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc",
		wantSeason:    0,
		wantHasSeason: false,
	},
}

func TestParse(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleaner.Parse(tt.input)
			if !contains(got.Title, tt.wantTitle) {
				t.Errorf("Title: got %q, want contains %q", got.Title, tt.wantTitle)
			}
			if got.SeasonNumber != tt.wantSeason {
				t.Errorf("SeasonNumber: got %d, want %d", got.SeasonNumber, tt.wantSeason)
			}
			if got.HasSeason != tt.wantHasSeason {
				t.Errorf("HasSeason: got %v, want %v", got.HasSeason, tt.wantHasSeason)
			}
		})
	}
}

func TestNormalise(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Déjà Vu", "deja vu"},
		{"Spirited Away", "spirited away"},
		{"AKIRA", "akira"},
		{"  extra   spaces  ", "extra spaces"},
	}
	for _, c := range cases {
		got := cleaner.Normalise(c.in)
		if got != c.want {
			t.Errorf("Normalise(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func contains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && (s == sub || len(s) > 0 && containsStr(s, sub)))
}

func containsStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
