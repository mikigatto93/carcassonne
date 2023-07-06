package tile

type Terrain int

const (
	Grass Terrain = +1 // 0 is used as absence of a value
	City
	Street
	Water
)

type TileLayout [5][5]Terrain

var Base TileLayout = TileLayout{
	{Grass, Grass, Grass, Grass, Grass},
	{Grass, Grass, Grass, Grass, Grass},
	{Grass, Grass, Grass, Grass, Grass},
	{Grass, Grass, Grass, Grass, Grass},
	{Grass, Grass, Grass, Grass, Grass},
}

var RiverStart TileLayout = TileLayout{
	{Grass, Grass, Grass, Grass, Grass},
	{Grass, Water, Water, Water, Grass},
	{Water, Water, Water, Water, Grass},
	{Grass, Water, Water, Water, Grass},
	{Grass, Grass, Grass, Grass, Grass},
}
