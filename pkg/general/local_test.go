package general

import (
	"fmt"
	"github.com/qordobacode/cli-v2/pkg/types"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var (
	localService = Local{}
)

func TestLocal_Read(t *testing.T) {
	bytes, err := localService.Read("local.go")
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
}

func TestLocal_ReadNotFound(t *testing.T) {
	bytes, err := localService.Read("not_present_file")
	assert.NotNil(t, err)
	assert.Nil(t, bytes)
}

func TestLocal_BuildFileNameSimple(t *testing.T) {
	res := []struct {
		FileName         string
		FilePath         string
		Version          string
		Suffix           string
		ExpectedFileName string
		FilePathParam    []string
	}{
		{
			FileName:         "wip-core.tmp.json",
			ExpectedFileName: "wip-core.tmp.json",
		},
		{
			FileName:         "wip-core.tmp.json",
			Version:          "v-0.1.0",
			ExpectedFileName: "wip-core.tmp-v-0.1.0.json",
		},
		{
			FileName:         "wip-core.tmp.json",
			Suffix:           "original",
			ExpectedFileName: "wip-core.tmp-original.json",
		},
		{
			FileName:         "wip-core.tmp.json",
			Version:          "version2",
			Suffix:           "original",
			ExpectedFileName: "wip-core.tmp-version2-original.json",
		},
		{
			FileName:         "wip-core.tmp.json",
			FilePath:         "v510-en/three/wip-core.tmp.json",
			Version:          "version2",
			Suffix:           "original",
			FilePathParam:    []string{"en-en"},
			ExpectedFileName: "v510-en" + string(os.PathSeparator) + "three" + string(os.PathSeparator) + "wip-core.tmp-version2-original.json",
		},
	}
	for _, asset := range res {
		file := types.File{
			Filename: asset.FileName,
			Filepath: asset.FilePath,
			Version:  asset.Version,
		}
		j := types.File2Download{
			File: &file,
			Person: types.Person{
				Code: "en-kr",
				ID:   0,
			},
			ReplaceIn: "kr",
		}
		fileName := localService.BuildDirectoryFilePath(&j, asset.FilePathParam, asset.Suffix, true)
		assert.Equal(t, asset.ExpectedFileName, fileName, "asset %+v is incorrect", asset)
	}
}

func Test_FileExists(t *testing.T) {
	assert.True(t, localService.FileExists("local_test.go"))
	assert.False(t, localService.FileExists("notfound.xml"))
}

func Test_QordobaHome(t *testing.T) {
	qordobaHome, err := localService.QordobaHome()
	assert.Nil(t, err)
	assert.NotNil(t, qordobaHome)
}

func Test_FilesInFolder(t *testing.T) {
	files := localService.FilesInFolder(".", false)
	fmt.Printf("files = %v\n", files)
}

func Test_RenderTable2Stdin(t *testing.T) {
	header := []string{"column 1", "column 2"}
	data := [][]string{
		{"row1_1", "row1_2"},
		{"row2_1", "row2_2"},
	}
	localService.RenderTable2Stdin(header, data)
}

func TestLocal_LoadCached(t *testing.T) {
	bytes, err := localService.LoadCached("file.txt")
	assert.Nil(t, bytes)
	assert.NotNil(t, err)
}

func Test_Glob(t *testing.T) {
	matches, _ := filepath.Glob("../rest")
	fmt.Printf("%v\n", matches)
}
