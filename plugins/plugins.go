package plugins

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/serainville/gologger"
)

var consoleLog = gologger.GetLogger(gologger.BASIC, gologger.ColoredLog)

type Plugin struct {
	name    string
	group   string
	value   string
	file    string
	builtin bool
}

type Plugins struct {
	Plugins []*Plugin
}

func (p *Plugins) Discover() {
	path := "./plugins"

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		// Directory exists. Find all executable files!
		files, _ := ioutil.ReadDir(path)

		// Add executable files to plugins list.
		for _, f := range files {
			fileInfo, _ := os.Stat(path + "/" + f.Name())
			if isExecutable(fileInfo) != false {
				plugin := new(Plugin)

				plugin.file = path + "/" + f.Name()
				plugin.builtin = false
				plugin.name = getPluginName(f.Name())

				p.Plugins = append(p.Plugins, plugin)

			}
		}
	}
}

func (p *Plugins) List() {
	var plugins string
	if len(p.Plugins) > 0 {
		for _, pl := range p.Plugins {
			plugins = plugins + pl.name + ", "
		}
		plugins = plugins[0 : len(plugins)-2]
	}
	consoleLog.Info("Found Plugins: [" + plugins + "]")

}

func getPluginName(file string) string {
	var extension = filepath.Ext(file)
	return file[0 : len(file)-len(extension)]
}

func isExecutable(file os.FileInfo) bool {
	mode := file.Mode().String()

	bits := strings.Split(mode, "")
	if bits[3] == "x" || bits[6] == "x" || bits[9] == "x" {
		return true
	}
	return false

}
