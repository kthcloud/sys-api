package db

const (
	ColCapacities = "capacities"
	ColGpuInfo    = "gpuInfo"
	ColStats      = "stats"
	ColStatus     = "status"
	ColHosts      = "hosts"
	ColZones      = "zones"
)

type CollectionDefinition struct {
	Name          string
	Indexes       [][]string
	UniqueIndexes [][]string
}

func getCollectionDefinitions() map[string]CollectionDefinition {
	collections := []CollectionDefinition{
		{Name: ColCapacities},
		{Name: ColGpuInfo},
		{Name: ColStats},
		{Name: ColStatus},
		{Name: ColHosts, UniqueIndexes: [][]string{{"name"}}},
		{Name: ColZones, UniqueIndexes: [][]string{{"name"}}},
	}

	collectionMap := make(map[string]CollectionDefinition)
	for _, collection := range collections {
		collectionMap[collection.Name] = collection
	}

	return collectionMap
}
