package testdata

var (
	// A geometry for testing purposes, just c'n'p to geojson.io for
	// visualization and editing purposes. A resulting covering of the area
	// is available from (URL shorteners are bailing on this one :)):
	//   https://s2.sidewalklabs.com/regioncoverer/?center=39.732119%2C-86.480343&zoom=12&cells=886ca30c%2C886ca314%2C886ca334%2C886ca33c%2C886ca344%2C886ca34c%2C886ca354%2C886ca35c%2C886ca364%2C886ca36c%2C886ca374%2C886ca37c%2C886ca384%2C886ca38c%2C886ca394%2C886ca39c%2C886ca40c%2C886ca414%2C886ca424%2C886ca42c%2C886ca434%2C886ca43c%2C886ca444%2C886ca44c%2C886ca454%2C886ca45c%2C886ca464%2C886ca46c%2C886ca474%2C886ca47c%2C886ca484%2C886ca48c%2C886ca494%2C886ca49c%2C886ca4a4%2C886ca4ac%2C886ca4b4%2C886ca4bc%2C886ca4c4%2C886ca4cc%2C886ca4d4%2C886ca4dc%2C886ca4e4%2C886ca4ec%2C886ca4f4%2C886ca4fc%2C886ca504%2C886ca50c%2C886ca514%2C886ca51c%2C886ca524%2C886ca52c%2C886ca534%2C886ca53c%2C886ca544%2C886ca54c%2C886ca554%2C886ca55c%2C886ca564%2C886ca56c%2C886ca574%2C886ca57c%2C886ca584%2C886ca58c%2C886ca594%2C886ca59c%2C886ca5a4%2C886ca5ac%2C886ca5b4%2C886ca5bc%2C886ca5c4%2C886ca5cc%2C886ca5d4%2C886ca5dc%2C886ca5e4%2C886ca5ec%2C886ca5f4%2C886ca5fc%2C886ca604%2C886ca60c%2C886ca61c%2C886ca674%2C886ca894%2C886ca8a4%2C886ca8ac%2C886ca8b4%2C886ca8bc%2C886ca8c4%2C886ca8cc%2C886cae04%2C886cae14%2C886cae1c%2C886cae24%2C886cae2c%2C886cae34%2C886cae3c%2C886cae44%2C886cae4c%2C886cae54%2C886cae5c%2C886cae6c%2C886caedc%2C886caee4%2C886caef4%2C886caefc%2C886caf04%2C886caf0c%2C886caf14%2C886caf1c%2C886caf24%2C886caf2c%2C886caf34%2C886caf3c%2C886caf44%2C886caf4c%2C886caf54%2C886caf5c%2C886caf64%2C886caf6c%2C886caf74%2C886caf7c%2C886caf84%2C886caf8c%2C886caf94%2C886caf9c%2C886cafa4%2C886cafac%2C886cafb4%2C886cafbc%2C886cafc4%2C886cafcc%2C886cafd4%2C886cafdc%2C886cafe4%2C886cafec%2C886caff4%2C886caffc%2C886cb004%2C886cb00c%2C886cb014%2C886cb01c%2C886cb024%2C886cb02c%2C886cb034%2C886cb03c%2C886cb044%2C886cb04c%2C886cb054%2C886cb05c%2C886cb064%2C886cb06c%2C886cb074%2C886cb07c%2C886cb084%2C886cb08c%2C886cb094%2C886cb09c%2C886cb0a4%2C886cb0ac%2C886cb0b4%2C886cb0bc%2C886cb0c4%2C886cb0cc%2C886cb0d4%2C886cb0dc%2C886cb0e4%2C886cb0ec%2C886cb0f4%2C886cb0fc%2C886cb104%2C886cb10c%2C886cb114%2C886cb11c%2C886cb124%2C886cb12c%2C886cb134%2C886cb13c%2C886cb144%2C886cb14c%2C886cb15c%2C886cb164%2C886cb16c%2C886cb174%2C886cb17c%2C886cb184%2C886cb18c%2C886cb194%2C886cb19c%2C886cb1a4%2C886cb1ac%2C886cb1b4%2C886cb1bc%2C886cb1c4%2C886cb1cc%2C886cb1d4%2C886cb1dc%2C886cb1e4%2C886cb1ec%2C886cb1f4%2C886cb1fc%2C886cb204%2C886cb21c%2C886cb224%2C886cb22c%2C886cb234%2C886cb23c%2C886cb3b4%2C886cb3cc%2C886cb3d4%2C886cb3dc%2C886cb6cc%2C886cb6d4%2C886cb714%2C886cb724%2C886cb72c%2C886cb734%2C886cb73c%2C886cb744%2C886cb74c%2C886cb754%2C886cb75c%2C886cb764%2C886cb76c%2C886cb774%2C886cb77c%2C886cb784%2C886cb814%2C886cb824%2C886cb82c%2C886cb834%2C886cb83c%2C886cb844%2C886cb84c%2C886cb854%2C886cb85c%2C886cb864%2C886cb86c%2C886cb88c%2C886cb894%2C886cb8c4%2C886cb8cc%2C886cb8d4%2C886cb8dc%2C886cb8e4%2C886cb8ec%2C886cb8f4%2C886cb8fc%2C886cb904%2C886cb90c%2C886cb914%2C886cb91c%2C886cb924%2C886cb92c%2C886cb934%2C886cb93c%2C886cb944%2C886cb94c%2C886cb954%2C886cb95c%2C886cb964%2C886cb96c%2C886cb974%2C886cb97c%2C886cb984%2C886cb98c%2C886cb994%2C886cb99c%2C886cb9a4%2C886cb9ac%2C886cb9b4%2C886cb9bc%2C886cb9c4%2C886cb9cc%2C886cb9d4%2C886cb9dc%2C886cb9e4%2C886cb9ec%2C886cb9f4%2C886cb9fc%2C886cba04%2C886cba0c%2C886cba14%2C886cba1c%2C886cba24%2C886cba2c%2C886cba34%2C886cba3c%2C886cba44%2C886cba4c%2C886cba54%2C886cba5c%2C886cba64%2C886cba6c%2C886cba74%2C886cba7c%2C886cba84%2C886cba8c%2C886cba94%2C886cba9c%2C886cbaa4%2C886cbaac%2C886cbab4%2C886cbabc%2C886cbac4%2C886cbacc%2C886cbad4%2C886cbadc%2C886cbae4%2C886cbaec%2C886cbaf4%2C886cbafc%2C886cbb04%2C886cbb0c%2C886cbb14%2C886cbb1c%2C886cbb24%2C886cbb2c%2C886cbb34%2C886cbb3c%2C886cbb44%2C886cbb4c%2C886cbb54%2C886cbb5c%2C886cbb64%2C886cbb6c%2C886cbb74%2C886cbb7c%2C886cbb84%2C886cbb8c%2C886cbb94%2C886cbb9c%2C886cbba4%2C886cbbac%2C886cbbb4%2C886cbbbc%2C886cbbc4%2C886cbbcc%2C886cbbd4%2C886cbbdc%2C886cbbe4%2C886cbbec%2C886cbbf4%2C886cbbfc%2C886cbc04%2C886cbc0c%2C886cbc14%2C886cbc1c%2C886cbc24%2C886cbc2c%2C886cbc34%2C886cbc3c%2C886cbc44%2C886cbc4c%2C886cbc54%2C886cbc5c%2C886cbc64%2C886cbc6c%2C886cbc74%2C886cbc7c%2C886cbc84%2C886cbc8c%2C886cbc94%2C886cbc9c%2C886cbca4%2C886cbcac%2C886cbcb4%2C886cbcbc%2C886cbcc4%2C886cbccc%2C886cbcd4%2C886cbcdc%2C886cbce4%2C886cbcec%2C886cbcf4%2C886cbcfc%2C886cbdac%2C886cbdb4%2C886cbdcc%2C886cbdd4%2C886cbe44%2C886cbe4c%2C886cbe54%2C886cbe5c%2C886cbe64%2C886cbe6c%2C886cbe74%2C886cbe7c%2C886cbe84%2C886cbe8c%2C886cbe94%2C886cbe9c%2C886cbea4%2C886cbeac%2C886cbeb4%2C886cbebc%2C886cbec4%2C886cbecc%2C886cbed4%2C886cbedc%2C886cbee4%2C886cbeec%2C886cbef4%2C886cbefc%2C886cbf1c%2C886cbf24%2C886cbf2c%2C886cbf34
	geojson = `{
	"type": "FeatureCollection",
	"features": [
	  {
		"type": "Feature",
		"properties": {},
		"geometry": {
		  "type": "LineString",
		  "coordinates": [
			[
			  -86.54342651367186,
			  39.7642140375156
			],
			[
			  -86.46034240722656,
			  39.86600654754002
			],
			[
			  -86.31683349609374,
			  39.774769485295465
			],
			[
			  -86.36764526367188,
			  39.66227082075381
			],
			[
			  -86.52969360351562,
			  39.62525937371228
			],
			[
			  -86.57020568847656,
			  39.73676229957947
			]
		  ]
		}
	  }
	]
  }`
	Loop                           = `39.7642140375156,-86.54342651367186,39.86600654754002,-86.46034240722656,39.774769485295465,-86.31683349609374,39.66227082075381,-86.36764526367188,39.62525937371228,-86.52969360351562,39.73676229957947,-86.57020568847656`
	LoopWithOnlyTwoPoints          = `39.7642140375156,-86.54342651367186,39.86600654754002,-86.46034240722656`
	LoopWithOddNumberOfCoordinates = `39.7642140375156,-86.54342651367186,39.86600654754002,-86.46034240722656,-86.31683349609374`
)