package parse

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/awalterschulze/gographviz/ast"
)

func getFiles(paths []string) []*os.File {
	files := []*os.File{}
	for _, path := range paths {
		splitName := strings.Split(path, ".")
		if splitName[len(splitName)-1] != "dot" {
			log.Debugf("Skipping %q because doesn't end in .dot", path)
			continue
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			if file, err := os.Open(path); err == nil {
				files = append(files, file)
			}
		}
	}
	return files
}

func getGraphVizGraphs(paths []string) []*ast.Graph {

	files := getFiles(paths)
	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}
	// log.Infof("Operating on: %q", fileNames)

	var result []*ast.Graph

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
		result = append(result, graph)
	}
	return result
}
