package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	minYear = 2015
	maxYear = 2021
)

var (
	fYear    = flag.Uint("year", 0, "specifies year")
	fDay     = flag.Uint("day", 0, "specifies day")
	fLang    = flag.String("lang", "", "programming language of solutions")
	fBasedir = flag.String("basedir", "", `base directory of solutions (default "../<value of -lang>")`)
)

//go:embed templates/**
var tmplFS embed.FS

type templateData struct {
	Year      uint
	FullYear  string
	Day       uint
	PaddedDay string
	FullDay   string
}

func imain() int {
	flag.Parse()

	lang := *fLang
	if lang == "" {
		fmt.Println("a programming language must be specified with the -lang flag")
		return 1
	}

	year := *fYear
	if year == 0 {
		fmt.Println("a year must be specified with the -year flag")
		return 1
	}
	if year < minYear || year > maxYear {
		fmt.Printf("invalid year %d: the year must be a year on which an AoC event was held\n", year)
		return 1
	}

	day := *fDay
	if day == 0 {
		fmt.Println("a day must be specified with the -day flag")
		return 1
	}
	if day > 25 {
		fmt.Printf("invalid day %d: the day must be in the range 1-25 (both inclusive)\n", day)
		return 1
	}

	basedir := *fBasedir
	if basedir == "" {
		basedir = fmt.Sprintf("../%s", lang)
	}

	fullYear := fmt.Sprintf("year%d", year)
	paddedDay := fmt.Sprintf("%02d", day)
	fullDay := fmt.Sprintf("day%s", paddedDay)
	data := templateData{
		Year:      year,
		FullYear:  fullYear,
		Day:       day,
		PaddedDay: paddedDay,
		FullDay:   fullDay,
	}

	subFS, err := fs.Sub(tmplFS, "templates")
	if err != nil {
		panic(`no "templates" directory in embedded template FS`)
	}
	if _, err := subFS.Open(lang); err != nil {
		fmt.Printf("couldn't find templates: %v", err)
		return 1
	}

	walkFn := func(path string, e fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		outPath := path
		outPath = strings.Replace(outPath, "YYYY", fmt.Sprint(year), -1)
		outPath = strings.Replace(outPath, "DD", paddedDay, -1)
		outPath = strings.TrimSuffix(outPath, ".tmpl")
		outPath = filepath.Join(basedir, outPath)
		if e.IsDir() {
			if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
				return fmt.Errorf("error creating directory %s: %w", outPath, err)
			}
			return nil
		}
		tmpl, err := template.ParseFS(subFS, path)
		if err != nil {
			return fmt.Errorf("error parsing template: %w", err)
		}
		outFile, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("error creating output file: %w", err)
		}
		defer func() {
			if err := outFile.Close(); err != nil {
				fmt.Printf("error closing output file: %+v\n", err)
			}
		}()
		if err := tmpl.Execute(outFile, data); err != nil {
			return fmt.Errorf("error executing template: %w", err)
		}
		return nil
	}
	if err := fs.WalkDir(subFS, lang, walkFn); err != nil {
		fmt.Printf("writing generated files: %v", err)
		return 2
	}

	return 0
}

func main() {
	os.Exit(imain())
}
