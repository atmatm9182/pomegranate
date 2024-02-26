package blueprint

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

type FileSpec struct {
	Type    string
	Raw     string
	Src     string
	Entries map[string]FileSpec
}

func (s *FileSpec) scaffoldFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if len(s.Raw) != 0 {
		_, err = f.Write([]byte(s.Raw))
	} else if len(s.Src) != 0 {
		err = copyFile(s.Src, name)
	} else {
		err = fmt.Errorf("Neither 'raw' or 'src' are specified for file %s", name)
	}
	
	if err != nil {
		removeErr := os.Remove(name)
		return errors.Join(err, removeErr)
	}
	
	return nil
}

func (s *FileSpec) scaffoldDir(name string) error {
	err := os.Mkdir(name, 0777)
	if err != nil {
		return err
	}

	if len(s.Src) != 0 {
		err = copyDir(s.Src, name)
		if err != nil {
			return errors.Join(err, os.RemoveAll(name))
		}
	}
	
	for entryName, spec := range s.Entries {
		fullPath := path.Join(name, entryName)
		if err = spec.scaffold(fullPath); err != nil {
			removeDirError := os.RemoveAll(name)
			return errors.Join(err, removeDirError)
		}
	}

	return nil
}

func (s *FileSpec) scaffold(name string) error {
	switch s.Type {
	case "file":
		return s.scaffoldFile(name)
	case "dir":
		return s.scaffoldDir(name)
	default:
		return fmt.Errorf("Entry type '%s' is not a valid entry type", s.Type)
	}
}

type Blueprint struct {
	Project struct {
		Name  string
		Files map[string]FileSpec
	}
}

func (b *Blueprint) Scaffold() error {
	for name, spec := range b.Project.Files {
		if err := spec.scaffold(name); err != nil {
			return err
		}
	}

	return nil
}

func copyDir(src string, dest string) error {
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
		fullEntryPath := path.Join(src, entryName)
		destEntryPath := path.Join(dest, entryName)

		if entryInfo.IsDir() {
			err = copyDir(fullEntryPath, destEntryPath)
		} else {
			err = copyFile(fullEntryPath, destEntryPath)
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
