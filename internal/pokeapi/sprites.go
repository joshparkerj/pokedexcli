package pokeapi

type IconsSprites struct {
	FrontDefault *string `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type EmeraldSprites struct {
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type HomeSprites struct {
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type RubySapphireSprites struct {
	BackDefault  *string `json:"back_default"`
	BackShiny    *string `json:"back_shiny"`
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type GoldSilverSprites struct {
	BackDefault      *string `json:"back_default"`
	BackShiny        *string `json:"back_shiny"`
	FrontDefault     *string `json:"front_default"`
	FrontShiny       *string `json:"front_shiny"`
	FrontTransparent *string `json:"front_transparent"`
}

type RedBlueYellowSprites struct {
	BackDefault      *string `json:"back_default"`
	BackGray         *string `json:"back_gray"`
	BackTransparent  *string `json:"back_transparent"`
	FrontDefault     *string `json:"front_default"`
	FrontGray        *string `json:"front_gray"`
	FrontTransparent *string `json:"front_transparent"`
}

type CrystalSprites struct {
	BackDefault           *string `json:"back_default"`
	BackShiny             *string `json:"back_shiny"`
	BackShinyTransparent  *string `json:"back_shiny_transparent"`
	BackTransparent       *string `json:"back_transparent"`
	FrontDefault          *string `json:"front_default"`
	FrontShiny            *string `json:"front_shiny"`
	FrontShinyTransparent *string `json:"front_shiny_transparent"`
	FrontTransparent      *string `json:"front_transparent"`
}

type ShowdownSprites struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type BlackWhiteSprites struct {
	Animated         ShowdownSprites `json:"animated"`
	BackDefault      *string         `json:"back_default"`
	BackFemale       *string         `json:"back_female"`
	BackShiny        *string         `json:"back_shiny"`
	BackShinyFemale  *string         `json:"back_shiny_female"`
	FrontDefault     *string         `json:"front_default"`
	FrontFemale      *string         `json:"front_female"`
	FrontShiny       *string         `json:"front_shiny"`
	FrontShinyFemale *string         `json:"front_shiny_female"`
}

type OtherSprites struct {
	DreamWorld      IconsSprites    `json:"dream_world"`
	Home            HomeSprites     `json:"home"`
	OfficialArtwork EmeraldSprites  `json:"official-artwork"`
	Showdown        ShowdownSprites `json:"showdown"`
}

type GenerationISprites struct {
	RedBlue RedBlueYellowSprites `json:"red-blue"`
	Yellow  RedBlueYellowSprites `json:"yellow"`
}

type GenerationIISprites struct {
	Crystal CrystalSprites    `json:"crystal"`
	Gold    GoldSilverSprites `json:"gold"`
	Silver  GoldSilverSprites `json:"silver"`
}

type GenerationIIISprites struct {
	Emerald          EmeraldSprites      `json:"emerald"`
	FireredLeafgreen RubySapphireSprites `json:"firered-leafgreen"`
	RubySapphire     RubySapphireSprites `json:"ruby-sapphire"`
}

type GenerationIVSprites struct {
	DiamondPearl        ShowdownSprites `json:"diamond-pearl"`
	HeartgoldSoulsilver ShowdownSprites `json:"heartgold-soulsilver"`
	Platinum            ShowdownSprites `json:"platinum"`
}

type GenerationVSprites struct {
	BlackWhite BlackWhiteSprites `json:"black-white"`
}

type GenerationVISprites struct {
	OmegarubyAlphasapphire HomeSprites `json:"omegaruby-alphasapphire"`
	XY                     HomeSprites `json:"x-y"`
}
type GenerationVIISprites struct {
	Icons             IconsSprites `json:"icons"`
	UltraSunUltraMoon HomeSprites  `json:"ultra-sun-ultra-moon"`
}
type GenerationVIIISprites struct {
	Icons IconsSprites `json:"icons"`
}

type VersionsSprites struct {
	GenerationI    GenerationISprites    `json:"generation-i"`
	GenerationII   GenerationIISprites   `json:"generation-ii"`
	GenerationIII  GenerationIIISprites  `json:"generation-iii"`
	GenerationIV   GenerationIVSprites   `json:"generation-iv"`
	GenerationV    GenerationVSprites    `json:"generation-v"`
	GenerationVI   GenerationVISprites   `json:"generation-vi"`
	GenerationVII  GenerationVIISprites  `json:"generation-vii"`
	GenerationVIII GenerationVIIISprites `json:"generation-viii"`
}

type Sprites struct {
	BackDefault      *string         `json:"back_default"`
	BackFemale       *string         `json:"back_female"`
	BackShiny        *string         `json:"back_shiny"`
	BackShinyFemale  *string         `json:"back_shiny_female"`
	FrontDefault     *string         `json:"front_default"`
	FrontFemale      *string         `json:"front_female"`
	FrontShiny       *string         `json:"front_shiny"`
	FrontShinyFemale *string         `json:"front_shiny_female"`
	Other            OtherSprites    `json:"other"`
	Versions         VersionsSprites `json:"versions"`
}
