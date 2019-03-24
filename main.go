package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/awalterschulze/gographviz/ast"
	"github.com/rivo/tview"
	"github.com/theonlyjohnny/dotCli/logger"
	"github.com/theonlyjohnny/dotCli/utils"
)

var (
	log = logger.Log
)

func getFiles(paths []string) []*os.File {
	files := []*os.File{}
	for _, path := range paths {
		splitName := strings.Split(path, ".")
		if splitName[len(splitName)-1] != "dot" {
			log.Debugf("Skipping %q because doesn't end in .dot", path)
			continue
		}
		log.Debugf("Checking if %q exists", path)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			if file, err := os.Open(path); err == nil {
				files = append(files, file)
			}
		}
	}
	return files
}

func main() {
	files := getFiles(os.Args[1:])
	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}
	log.Infof("Operating on: %q", fileNames)

	for i, file := range files {
		fileName := fileNames[i]
		contents, err := ioutil.ReadAll(file)
		file.Close()
		if err != nil {
			log.Errorf("Unable to read from %q: %s", fileName, err.Error())
			continue
		}
		graph, err := gographviz.Parse(contents)
		if err != nil {
			log.Errorf("Could not parse contents of %q into Graph: %s", fileName, err.Error())
		}
		renderErr := renderGraph(graph)
		if renderErr != nil {
			log.Errorf("Could not render contents of %q: %s", fileName, renderErr.Error())
		}
	}
}

func getMap(input interface{}) (*map[string]string, error) {
	attrList, ok := input.(ast.AttrList)
	if !ok {
		return nil, errors.New("input doesn't implement ast.AttrList")
	}
	attrMap := attrList.GetMap()
	log.Debugf("from %v got map: %v", input, attrMap)
	return &attrMap, nil
}

func addGraphAttrs(view *tview.Box, attrsMap map[string]string) error {
	for k, v := range attrsMap {
		if k == "bb" {
			rect := utils.GetRectFromCommaString(v)

			log.Debugf("new Rect: %v", rect)
			log.Debugf("attr def: %v", v)

			view.SetRect(rect[0], rect[1], rect[2], rect[3])
		}
	}
	return nil
}

func renderGraph(graph *ast.Graph) error {
	view := tview.NewBox().SetBorder(true).SetTitle(graph.ID.String())
	for _, stmt := range graph.StmtList {
		switch attrs := stmt.(type) {
		case ast.GraphAttrs:
			attrsMap := ast.AttrList(attrs).GetMap()
			addGraphAttrs(view, attrsMap)
		// case ast.NodeStmt:
		// addNodeStmt(view, attrs)
		default:
			log.Warnf("Unknown statement type: %#v", stmt)
		}
	}
	runnable := tview.NewApplication().SetRoot(view, false)
	log.Debug("made runnable:", runnable)
	// return nil
	return runnable.Run()
}
