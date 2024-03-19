package blueprint

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/atmatm9182/pomegranate/blueprint/options"
	"github.com/atmatm9182/pomegranate/util"
)

type Scaffolder struct {
	logger *log.Logger
	prefix string
}

func NewScaffolder(opts *options.ScaffoldingOptions) Scaffolder {
	return Scaffolder{
		logger: opts.GetLogger(),
		prefix: opts.ScaffoldPrefix,
	}
}

func (s *Scaffolder) Scaffold(b *Blueprint) error {
	concatSrcWithPath(b.absolutePath, b.Project.Files)

	s.logger.Printf("Scaffolding blueprint for project '%s'\n", b.Project.Name)

	var err error
	if s.prefix != options.DefaultScaffoldPrefix {
		s.logCreating(s.prefix)
		if err = os.MkdirAll(s.prefix, 0777); err != nil {
			return err
		}

		defer func() {
			if err != nil {
				os.RemoveAll(s.prefix)
			}
		}()
	}

	for name, spec := range b.Project.Files {
		if err = s.scaffold(&spec, name); err != nil {
			return err
		}
	}

	s.logger.Println("Scaffolding success!")
	return nil
}

func concatSrcWithPath(absolutePath string, files map[string]FileSpec) {
	for fileName, fileSpec := range files {
		if len(fileSpec.Src) != 0 {
			fileSpec.Src = path.Join(absolutePath, fileSpec.Src)
			files[fileName] = fileSpec
		}

		if fileSpec.Type == "dir" {
			concatSrcWithPath(path.Join(absolutePath, fileName), fileSpec.Entries)
		}
	}
}

func (s *Scaffolder) copyDir(src string, dest string) error {
	srcDirEntries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	var destDir *os.File
	destDir, err = os.Open(dest)
	if err != nil {
		return err
	}
	defer destDir.Close()

	for _, entry := range srcDirEntries {
		var entryInfo os.FileInfo
		entryInfo, err = entry.Info()
		if err != nil {
			return err
		}

		entryName := entryInfo.Name()
		srcEntryPath := path.Join(src, entryName)
		destEntryPath := path.Join(dest, entryName)

		s.logCopying(srcEntryPath, destEntryPath)
		if entryInfo.IsDir() {
			err = s.copyDir(srcEntryPath, destEntryPath)
		} else {
			err = copyFile(srcEntryPath, destEntryPath)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	var destFile *os.File
	destFile, err = os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}

func (s *Scaffolder) logCreating(name string) {
	s.logger.Printf("Creating '%s'...\n", name)
}

func (s *Scaffolder) logCopying(src, dest string) {
	s.logger.Printf("Copying '%s' to '%s'...\n", src, dest)
}

func (s *Scaffolder) scaffoldFile(spec *FileSpec, name string) error {
	name = path.Join(s.prefix, name)

	s.logCreating(name)
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if len(spec.Raw) != 0 {
		_, err = f.Write([]byte(spec.Raw))
	} else if len(spec.Src) != 0 {
		s.logCopying(spec.Src, name)
		err = copyFile(spec.Src, name)
	} else {
		err = fmt.Errorf("Neither 'raw' or 'src' are specified for file %s", name)
	}

	if err != nil {
		removeErr := os.Remove(name)
		return errors.Join(err, removeErr)
	}

	return nil
}

func (s *Scaffolder) scaffoldDir(spec *FileSpec, name string) error {
	name = path.Join(s.prefix, name)

	// do not create target directory if it already exists because we don't want to override existing files
	if util.FileExists(name) {
		return fmt.Errorf("Could not create directory %s, because it already exists", name)
	}

	err := os.MkdirAll(name, 0777)
	if err != nil {
		return err
	}

	if len(spec.Src) != 0 {
		s.logCopying(spec.Src, name)
		err = s.copyDir(spec.Src, name)
		if err != nil {
			return errors.Join(err, os.RemoveAll(name))
		}
	}

	for entryName, spec := range spec.Entries {
		fullPath := path.Join(name, entryName)
		if err = s.scaffold(&spec, fullPath); err != nil {
			removeDirError := os.RemoveAll(name)
			return errors.Join(err, removeDirError)
		}
	}

	return nil
}

func (s *Scaffolder) scaffold(spec *FileSpec, name string) error {
	switch spec.Type {
	case "file":
		s.logger.Printf("Scaffolding file '%s'...\n", name)
		return s.scaffoldFile(spec, name)
	case "dir":
		s.logger.Printf("Scaffolding dir '%s'...\n", name)
		return s.scaffoldDir(spec, name)
	default:
		return fmt.Errorf("Entry type '%s' is not a valid entry type", spec.Type)
	}
}
