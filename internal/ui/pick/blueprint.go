package pick

import (
	"github.com/ashleymorris2/booty/internal/fs"
	"github.com/ashleymorris2/booty/internal/ui/components/menu"
)

func BlueprintFrom(files []string) (string, error) {

	var results = fs.ReadMetadataFromFiles(files)

	items := make([]menu.Item, len(files))
	for res := range results {
		if res.Err != nil {
			continue
		}
		items[res.Index] = menu.NewItem(res.Item.Title, res.Item.Description, res.Item.FilePath)
	}

	m, err := menu.Show(items)
	if err != nil {
		return "", err
	}

	return m.Result, nil
}
