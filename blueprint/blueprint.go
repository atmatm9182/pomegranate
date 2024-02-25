package blueprint

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type FileSpec struct {
	Type    string
	Raw     string
	Src     string
	Entries []map[string]FileSpec
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
		var srcContents []byte
		srcContents, err = os.ReadFile(s.Src)
		if err != nil {
			err = fmt.Errorf("Could not read file '%s' because of: %s\n", s.Src, err)
		} else {
			_, err = f.Write(srcContents)
		}
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
	
	for _, entry := range s.Entries {
		for entryName, spec := range entry {
			fullPath := path.Join(name, entryName)
			if err = spec.scaffold(fullPath); err != nil {
				removeDirError := os.RemoveAll(name)
				return errors.Join(err, removeDirError)
			}
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
