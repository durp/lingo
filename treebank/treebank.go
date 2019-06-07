package treebank

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io"

	"github.com/chewxy/lingo"

	"bufio"
	"os"
	"strconv"
	"strings"
)

// Loader is anything that loads into a slice of SentenceTags. For future uses, to load tree banks
type Loader func(string) []SentenceTag

// LoadUniversal loads a treebank file formatted in a CONLLU format
func LoadUniversal(fileName string) []SentenceTag {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()

	return ReadConllu(f)
}

const conllu_unspecified = "_"

const (
	conllu_ID int = iota
	conllu_Form
	conllu_Lemma
	conllu_UPOS
	conllu_XPOS
	conllu_Features
	conllu_Head
	conllu_DepRel
	conllu_Deps
	conllu_Misc
)

// ReadConllu reads a file formatted in a CONLLU format
// It will skip # lines and empty nodes (Head = "_"
func ReadConllu(reader io.Reader) []SentenceTag {
	s, st, sh, sdt := reset()
	sentences := make([]SentenceTag, 0)
	sentenceCount := 0

	var usedTags lingo.TagSet
	var usedDepTypes lingo.DependencyTypeSet
	var unknownTags = make(map[string]struct{})
	var unknownDepType = make(map[string]struct{})

	colCount := 0
	for bs := bufio.NewScanner(reader); bs.Scan(); colCount++ {
		l := bs.Text()

		if len(l) == 0 {
			// then this is a new sentence
			sentences = finish(s, st, sh, sdt, sentences)
			s, st, sh, sdt = reset()

			sentenceCount++
			continue
		}

		cols := strings.Split(l, "\t")
		if strings.HasPrefix(cols[conllu_ID], "#") {
			// skip comments
			continue
		}

		head := cols[conllu_Head]
		if head == conllu_unspecified {
			// skip empty node (e.g., GAP)
			continue
		}
		h, err := strconv.Atoi(head)
		if err != nil {
			panic(errors.Wrapf(err, "cols: %v", cols)) // panic is the right option, because there is no default
		}

		word := lingo.UnescapeSpecials(cols[conllu_Form])

		var tag string
		switch lingo.BUILD_TAGSET {
		case "stanfordtags":
			tag = cols[conllu_XPOS]
		case "universaltags", "universaltagsv2":
			tag = cols[conllu_UPOS]
		default:
			panic("unknown tagset")
		}

		t, ok := StringToPOSTag(tag)
		if ok {
			usedTags[t] = struct{}{}
		} else {
			unknownTags[tag] = struct{}{}
		}

		depType := cols[conllu_DepRel]
		var dt lingo.DependencyType
		if dt, ok = StringToDependencyType(depType); ok {
			usedDepTypes[dt] = struct{}{}
		} else {
			unknownDepType[depType] = struct{}{}
		}

		lexType := StringToLexType(tag)
		lexeme := lingo.Lexeme{Value: word, LexemeType: lexType, Line: sentenceCount, Col: colCount}
		s = append(s, lexeme)
		st = append(st, t)
		sh = append(sh, h)
		sdt = append(sdt, dt)
	}

	return sentences
}

// LoadEWT loads a zipped English Web Treebank (as donated by Google)
func LoadEWT(filename string) []SentenceTag {

	r, err := zip.OpenReader(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = r.Close()
	}()

	sentences := make([]SentenceTag, 0)

	for _, f := range r.File {
		contents, err := f.Open()
		if err != nil {
			panic(err)
		}
		sentences = append(sentences, ReadConllu(contents)...)
		_ = contents.Close()
	}

	return sentences
}
