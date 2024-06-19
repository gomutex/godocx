package stypes

import (
	"encoding/xml"
	"errors"
)

// BorderStyle represents the possible values for border styles.
type BorderStyle string

const (
	BorderStyleNil                    BorderStyle = "nil"                    // No Border
	BorderStyleNone                   BorderStyle = "none"                   // No Border
	BorderStyleSingle                 BorderStyle = "single"                 // Single Line Border
	BorderStyleThick                  BorderStyle = "thick"                  // Single Line Border
	BorderStyleDouble                 BorderStyle = "double"                 // Double Line Border
	BorderStyleDotted                 BorderStyle = "dotted"                 // Dotted Line Border
	BorderStyleDashed                 BorderStyle = "dashed"                 // Dashed Line Border
	BorderStyleDotDash                BorderStyle = "dotDash"                // Dot Dash Line Border
	BorderStyleDotDotDash             BorderStyle = "dotDotDash"             // Dot Dot Dash Line Border
	BorderStyleTriple                 BorderStyle = "triple"                 // Triple Line Border
	BorderStyleThinThickSmallGap      BorderStyle = "thinThickSmallGap"      // Thin, Thick Line Border
	BorderStyleThickThinSmallGap      BorderStyle = "thickThinSmallGap"      // Thick, Thin Line Border
	BorderStyleThinThickThinSmallGap  BorderStyle = "thinThickThinSmallGap"  // Thin, Thick, Thin Line Border
	BorderStyleThinThickMediumGap     BorderStyle = "thinThickMediumGap"     // Thin, Thick Line Border
	BorderStyleThickThinMediumGap     BorderStyle = "thickThinMediumGap"     // Thick, Thin Line Border
	BorderStyleThinThickThinMediumGap BorderStyle = "thinThickThinMediumGap" // Thin, Thick, Thin Line Border
	BorderStyleThinThickLargeGap      BorderStyle = "thinThickLargeGap"      // Thin, Thick Line Border
	BorderStyleThickThinLargeGap      BorderStyle = "thickThinLargeGap"      // Thick, Thin Line Border
	BorderStyleThinThickThinLargeGap  BorderStyle = "thinThickThinLargeGap"  // Thin, Thick, Thin Line Border
	BorderStyleWave                   BorderStyle = "wave"                   // Wavy Line Border
	BorderStyleDoubleWave             BorderStyle = "doubleWave"             // Double Wave Line Border
	BorderStyleDashSmallGap           BorderStyle = "dashSmallGap"           // Dashed Line Border
	BorderStyleDashDotStroked         BorderStyle = "dashDotStroked"         // Dash Dot Strokes Line Border
	BorderStyleThreeDEmboss           BorderStyle = "threeDEmboss"           // 3D Embossed Line Border
	BorderStyleThreeDEngrave          BorderStyle = "threeDEngrave"          // 3D Engraved Line Border
	BorderStyleOutset                 BorderStyle = "outset"                 // Outset Line Border
	BorderStyleInset                  BorderStyle = "inset"                  // Inset Line Border
	BorderStyleApples                 BorderStyle = "apples"                 // Apples Art Border
	BorderStyleArchedScallops         BorderStyle = "archedScallops"         // Arched Scallops Art Border
	BorderStyleBabyPacifier           BorderStyle = "babyPacifier"           // Baby Pacifier Art Border
	BorderStyleBabyRattle             BorderStyle = "babyRattle"             // Baby Rattle Art Border
	BorderStyleBalloons3Colors        BorderStyle = "balloons3Colors"        // Three Color Balloons Art Border
	BorderStyleBalloonsHotAir         BorderStyle = "balloonsHotAir"         // Hot Air Balloons Art Border
	BorderStyleBasicBlackDashes       BorderStyle = "basicBlackDashes"       // Black Dash Art Border
	BorderStyleBasicBlackDots         BorderStyle = "basicBlackDots"         // Black Dot Art Border
	BorderStyleBasicBlackSquares      BorderStyle = "basicBlackSquares"      // Black Square Art Border
	BorderStyleBasicThinLines         BorderStyle = "basicThinLines"         // Thin Line Art Border
	BorderStyleBasicWhiteDashes       BorderStyle = "basicWhiteDashes"       // White Dash Art Border
	BorderStyleBasicWhiteDots         BorderStyle = "basicWhiteDots"         // White Dot Art Border
	BorderStyleBasicWhiteSquares      BorderStyle = "basicWhiteSquares"      // White Square Art Border
	BorderStyleBasicWideInline        BorderStyle = "basicWideInline"        // Wide Inline Art Border
	BorderStyleBasicWideMidline       BorderStyle = "basicWideMidline"       // Wide Midline Art Border
	BorderStyleBasicWideOutline       BorderStyle = "basicWideOutline"       // Wide Outline Art Border
	BorderStyleBats                   BorderStyle = "bats"                   // Bats Art Border
	BorderStyleBirds                  BorderStyle = "birds"                  // Birds Art Border
	BorderStyleBirdsFlight            BorderStyle = "birdsFlight"            // Birds Flying Art Border
	BorderStyleCabins                 BorderStyle = "cabins"                 // Cabin Art Border
	BorderStyleCakeSlice              BorderStyle = "cakeSlice"              // Cake Art Border
	BorderStyleCandyCorn              BorderStyle = "candyCorn"              // Candy Corn Art Border
	BorderStyleCelticKnotwork         BorderStyle = "celticKnotwork"         // Knot Work Art Border
	BorderStyleCertificateBanner      BorderStyle = "certificateBanner"      // Certificate Banner Art Border
	BorderStyleChainLink              BorderStyle = "chainLink"              // Chain Link Art Border
	BorderStyleChampagneBottle        BorderStyle = "champagneBottle"        // Champagne Bottle Art Border
	BorderStyleCheckedBarBlack        BorderStyle = "checkedBarBlack"        // Black and White Bar Art Border
	BorderStyleCheckedBarColor        BorderStyle = "checkedBarColor"        // Color Checked Bar Art Border
	BorderStyleCheckered              BorderStyle = "checkered"              // Checkerboard Art Border
	BorderStyleChristmasTree          BorderStyle = "christmasTree"          // Christmas Tree Art Border
	BorderStyleCirclesLines           BorderStyle = "circlesLines"           // Circles And Lines Art Border
	BorderStyleCirclesRectangles      BorderStyle = "circlesRectangles"      // Circles and Rectangles Art Border
	BorderStyleClassicalWave          BorderStyle = "classicalWave"          // Wave Art Border
	BorderStyleClocks                 BorderStyle = "clocks"                 // Clocks Art Border
	BorderStyleCompass                BorderStyle = "compass"                // Compass Art Border
	BorderStyleConfetti               BorderStyle = "confetti"               // Confetti Art Border
	BorderStyleConfettiGrays          BorderStyle = "confettiGrays"          // Confetti Art Border
	BorderStyleConfettiOutline        BorderStyle = "confettiOutline"        // Confetti Art Border
	BorderStyleConfettiStreamers      BorderStyle = "confettiStreamers"      // Confetti Streamers Art Border
	BorderStyleConfettiWhite          BorderStyle = "confettiWhite"          // Confetti Art Border
	BorderStyleCornerTriangles        BorderStyle = "cornerTriangles"        // Corner Triangle Art Border
	BorderStyleCouponCutoutDashes     BorderStyle = "couponCutoutDashes"     // Dashed Line Art Border
	BorderStyleCouponCutoutDots       BorderStyle = "couponCutoutDots"       // Dotted Line Art Border
	BorderStyleCrazyMaze              BorderStyle = "crazyMaze"              // Maze Art Border
	BorderStyleCreaturesButterfly     BorderStyle = "creaturesButterfly"     // Butterfly Art Border
	BorderStyleCreaturesFish          BorderStyle = "creaturesFish"          // Fish Art Border
	BorderStyleCreaturesInsects       BorderStyle = "creaturesInsects"       // Insects Art Border
	BorderStyleCreaturesLadyBug       BorderStyle = "creaturesLadyBug"       // Ladybug Art Border
	BorderStyleCrossStitch            BorderStyle = "crossStitch"            // Cross-stitch Art Border
	BorderStyleCup                    BorderStyle = "cup"                    // Cupid Art Border
	BorderStyleDecoArch               BorderStyle = "decoArch"               // Archway Art Border
	BorderStyleDecoArchColor          BorderStyle = "decoArchColor"          // Color Archway Art Border
	BorderStyleDecoBlocks             BorderStyle = "decoBlocks"             // Blocks Art Border
	BorderStyleDiamondsGray           BorderStyle = "diamondsGray"           // Gray Diamond Art Border
	BorderStyleDoubleD                BorderStyle = "doubleD"                // Double D Art Border
	BorderStyleDoubleDiamonds         BorderStyle = "doubleDiamonds"         // Diamond Art Border
	BorderStyleEarth1                 BorderStyle = "earth1"                 // Earth Art Border
	BorderStyleEarth2                 BorderStyle = "earth2"                 // Earth Art Border
	BorderStyleEclipsingSquares1      BorderStyle = "eclipsingSquares1"      // Shadowed Square Art Border
	BorderStyleEclipsingSquares2      BorderStyle = "eclipsingSquares2"      // Shadowed Square Art Border
	BorderStyleEggsBlack              BorderStyle = "eggsBlack"              // Painted Egg Art Border
	BorderStyleFans                   BorderStyle = "fans"                   // Fans Art Border
	BorderStyleFilm                   BorderStyle = "film"                   // Film Reel Art Border
	BorderStyleFirecrackers           BorderStyle = "firecrackers"           // Firecracker Art Border
	BorderStyleFlowersBlockPrint      BorderStyle = "flowersBlockPrint"      // Flowers Art Border
	BorderStyleFlowersDaisies         BorderStyle = "flowersDaisies"         // Daisy Art Border
	BorderStyleFlowersModern1         BorderStyle = "flowersModern1"         // Flowers Art Border
	BorderStyleFlowersModern2         BorderStyle = "flowersModern2"         // Flowers Art Border
	BorderStyleFlowersPansy           BorderStyle = "flowersPansy"           // Pansy Art Border
	BorderStyleFlowersRedRose         BorderStyle = "flowersRedRose"         // Red Rose Art Border
	BorderStyleFlowersRoses           BorderStyle = "flowersRoses"           // Roses Art Border
	BorderStyleFlowersTeacup          BorderStyle = "flowersTeacup"          // Flowers in a Teacup Art Border
	BorderStyleFlowersTiny            BorderStyle = "flowersTiny"            // Small Flower Art Border
	BorderStyleGems                   BorderStyle = "gems"                   // Gems Art Border
	BorderStyleGingerbreadMan         BorderStyle = "gingerbreadMan"         // Gingerbread Man Art Border
	BorderStyleGradient               BorderStyle = "gradient"               // Triangle Gradient Art Border
	BorderStyleHandmade1              BorderStyle = "handmade1"              // Handmade Art Border
	BorderStyleHandmade2              BorderStyle = "handmade2"              // Handmade Art Border
	BorderStyleHeartBalloon           BorderStyle = "heartBalloon"           // Heart-Shaped Balloon Art Border
	BorderStyleHeartGray              BorderStyle = "heartGray"              // Gray Heart Art Border
	BorderStyleHearts                 BorderStyle = "hearts"                 // Hearts Art Border
	BorderStyleHeebieJeebies          BorderStyle = "heebieJeebies"          // Pattern Art Border
	BorderStyleHolly                  BorderStyle = "holly"                  // Holly Art Border
	BorderStyleHouseFunky             BorderStyle = "houseFunky"             // House Art Border
	BorderStyleHypnotic               BorderStyle = "hypnotic"               // Circular Art Border
	BorderStyleIceCreamCones          BorderStyle = "iceCreamCones"          // Ice Cream Cone Art Border
	BorderStyleLightBulb              BorderStyle = "lightBulb"              // Light Bulb Art Border
	BorderStyleLightning1             BorderStyle = "lightning1"             // Lightning Art Border
	BorderStyleLightning2             BorderStyle = "lightning2"             // Lightning Art Border
	BorderStyleMapPins                BorderStyle = "mapPins"                // Map Pins Art Border
	BorderStyleMapleLeaf              BorderStyle = "mapleLeaf"              // Maple Leaf Art Border
	BorderStyleMapleMuffins           BorderStyle = "mapleMuffins"           // Muffin Art Border
	BorderStyleMarquee                BorderStyle = "marquee"                // Marquee Art Border
	BorderStyleMarqueeToothed         BorderStyle = "marqueeToothed"         // Marquee Art Border
	BorderStyleMoons                  BorderStyle = "moons"                  // Moon Art Border
	BorderStyleMosaic                 BorderStyle = "mosaic"                 // Mosaic Art Border
	BorderStyleMusicNotes             BorderStyle = "musicNotes"             // Musical Note Art Border
	BorderStyleNorthwest              BorderStyle = "northwest"              // Patterned Art Border
	BorderStyleOvals                  BorderStyle = "ovals"                  // Oval Art Border
	BorderStylePackages               BorderStyle = "packages"               // Package Art Border
	BorderStylePalmsBlack             BorderStyle = "palmsBlack"             // Black Palm Tree Art Border
	BorderStylePalmsColor             BorderStyle = "palmsColor"             // Color Palm Tree Art Border
	BorderStylePaperClips             BorderStyle = "paperClips"             // Paper Clip Art Border
	BorderStylePapyrus                BorderStyle = "papyrus"                // Papyrus Art Border
	BorderStylePartyFavor             BorderStyle = "partyFavor"             // Party Favor Art Border
	BorderStylePartyGlass             BorderStyle = "partyGlass"             // Party Glass Art Border
	BorderStylePencils                BorderStyle = "pencils"                // Pencils Art Border
	BorderStylePeople                 BorderStyle = "people"                 // Character Art Border
	BorderStylePeopleWaving           BorderStyle = "peopleWaving"           // Waving Character Border
	BorderStylePeopleHats             BorderStyle = "peopleHats"             // Character With Hat Art Border
	BorderStylePoinsettias            BorderStyle = "poinsettias"            // Poinsettia Art Border
	BorderStylePostageStamp           BorderStyle = "postageStamp"           // Postage Stamp Art Border
	BorderStylePumpkin1               BorderStyle = "pumpkin1"               // Pumpkin Art Border
	BorderStylePushPinNote2           BorderStyle = "pushPinNote2"           // Push Pin Art Border
	BorderStylePushPinNote1           BorderStyle = "pushPinNote1"           // Push Pin Art Border
	BorderStylePyramids               BorderStyle = "pyramids"               // Pyramid Art Border
	BorderStylePyramidsAbove          BorderStyle = "pyramidsAbove"          // Pyramid Art Border
	BorderStyleQuadrants              BorderStyle = "quadrants"              // Quadrants Art Border
	BorderStyleRings                  BorderStyle = "rings"                  // Rings Art Border
	BorderStyleSafari                 BorderStyle = "safari"                 // Safari Art Border
	BorderStyleSawtooth               BorderStyle = "sawtooth"               // Saw tooth Art Border
	BorderStyleSawtoothGray           BorderStyle = "sawtoothGray"           // Gray Saw tooth Art Border
	BorderStyleScaredCat              BorderStyle = "scaredCat"              // Scared Cat Art Border
	BorderStyleSeattle                BorderStyle = "seattle"                // Umbrella Art Border
	BorderStyleShadowedSquares        BorderStyle = "shadowedSquares"        // Shadowed Squares Art Border
	BorderStyleSharksTeeth            BorderStyle = "sharksTeeth"            // Shark Tooth Art Border
	BorderStyleShorebirdTracks        BorderStyle = "shorebirdTracks"        // Bird Tracks Art Border
	BorderStyleSkyrocket              BorderStyle = "skyrocket"              // Rocket Art Border
	BorderStyleSnowflakeFancy         BorderStyle = "snowflakeFancy"         // Snowflake Art Border
	BorderStyleSnowflakes             BorderStyle = "snowflakes"             // Snowflake Art Border
	BorderStyleSombrero               BorderStyle = "sombrero"               // Sombrero Art Border
	BorderStyleSouthwest              BorderStyle = "southwest"              // Southwest-themed Art Border
	BorderStyleStars                  BorderStyle = "stars"                  // Stars Art Border
	BorderStyleStarsTop               BorderStyle = "starsTop"               // Stars On Top Art Border
	BorderStyleStars3d                BorderStyle = "stars3d"                // 3-D Stars Art Border
	BorderStyleStarsBlack             BorderStyle = "starsBlack"             // Stars Art Border
	BorderStyleStarsShadowed          BorderStyle = "starsShadowed"          // Stars With Shadows Art Border
	BorderStyleSun                    BorderStyle = "sun"                    // Sun Art Border
	BorderStyleSwirligig              BorderStyle = "swirligig"              // Whirligig Art Border
	BorderStyleTornPaper              BorderStyle = "tornPaper"              // Torn Paper Art Border
	BorderStyleTornPaperBlack         BorderStyle = "tornPaperBlack"         // Black Torn Paper Art Border
	BorderStyleTrees                  BorderStyle = "trees"                  // Tree Art Border
	BorderStyleTriangleParty          BorderStyle = "triangleParty"          // Triangle Art Border
	BorderStyleTriangles              BorderStyle = "triangles"              // Triangles Art Border
	BorderStyleTribal1                BorderStyle = "tribal1"                // Tribal Art Border One
	BorderStyleTribal2                BorderStyle = "tribal2"                // Tribal Art Border Two
	BorderStyleTribal3                BorderStyle = "tribal3"                // Tribal Art Border Three
	BorderStyleTribal4                BorderStyle = "tribal4"                // Tribal Art Border Four
	BorderStyleTribal5                BorderStyle = "tribal5"                // Tribal Art Border Five
	BorderStyleTribal6                BorderStyle = "tribal6"                // Tribal Art Border Six
	BorderStyleTwistedLines1          BorderStyle = "twistedLines1"          // Twisted Lines Art Border
	BorderStyleTwistedLines2          BorderStyle = "twistedLines2"          // Twisted Lines Art Border
	BorderStyleVine                   BorderStyle = "vine"                   // Vine Art Border
	BorderStyleWaveline               BorderStyle = "waveline"               // Wavy Line Art Border
	BorderStyleWeavingAngles          BorderStyle = "weavingAngles"          // Weaving Angles Art Border
	BorderStyleWeavingBraid           BorderStyle = "weavingBraid"           // Weaving Braid Art Border
	BorderStyleWeavingRibbon          BorderStyle = "weavingRibbon"          // Weaving Ribbon Art Border
	BorderStyleWeavingStrips          BorderStyle = "weavingStrips"          // Weaving Strips Art Border
	BorderStyleWhiteFlowers           BorderStyle = "whiteFlowers"           // White Flowers Art Border
	BorderStyleWoodwork               BorderStyle = "woodwork"               // Woodwork Art Border
	BorderStyleXIllusions             BorderStyle = "xIllusions"             // Crisscross Art Border
	BorderStyleZanyTriangles          BorderStyle = "zanyTriangles"          // Triangle Art Border
	BorderStyleZigZag                 BorderStyle = "zigZag"                 // Zigzag Art Border
	BorderStyleZigZagStitch           BorderStyle = "zigZagStitch"           // Zigzag stitch Art Border
)

// BorderStyleFromStr converts a string to a BorderStyle.
func BorderStyleFromStr(value string) (BorderStyle, error) {
	switch value {
	case "nil":
		return BorderStyleNil, nil
	case "none":
		return BorderStyleNone, nil
	case "single":
		return BorderStyleSingle, nil
	case "thick":
		return BorderStyleThick, nil
	case "double":
		return BorderStyleDouble, nil
	case "dotted":
		return BorderStyleDotted, nil
	case "dashed":
		return BorderStyleDashed, nil
	case "dotDash":
		return BorderStyleDotDash, nil
	case "dotDotDash":
		return BorderStyleDotDotDash, nil
	case "triple":
		return BorderStyleTriple, nil
	case "thinThickSmallGap":
		return BorderStyleThinThickSmallGap, nil
	case "thickThinSmallGap":
		return BorderStyleThickThinSmallGap, nil
	case "thinThickThinSmallGap":
		return BorderStyleThinThickThinSmallGap, nil
	case "thinThickMediumGap":
		return BorderStyleThinThickMediumGap, nil
	case "thickThinMediumGap":
		return BorderStyleThickThinMediumGap, nil
	case "thinThickThinMediumGap":
		return BorderStyleThinThickThinMediumGap, nil
	case "thinThickLargeGap":
		return BorderStyleThinThickLargeGap, nil
	case "thickThinLargeGap":
		return BorderStyleThickThinLargeGap, nil
	case "thinThickThinLargeGap":
		return BorderStyleThinThickThinLargeGap, nil
	case "wave":
		return BorderStyleWave, nil
	case "doubleWave":
		return BorderStyleDoubleWave, nil
	case "dashSmallGap":
		return BorderStyleDashSmallGap, nil
	case "dashDotStroked":
		return BorderStyleDashDotStroked, nil
	case "threeDEmboss":
		return BorderStyleThreeDEmboss, nil
	case "threeDEngrave":
		return BorderStyleThreeDEngrave, nil
	case "outset":
		return BorderStyleOutset, nil
	case "inset":
		return BorderStyleInset, nil
	case "apples":
		return BorderStyleApples, nil
	case "archedScallops":
		return BorderStyleArchedScallops, nil
	case "babyPacifier":
		return BorderStyleBabyPacifier, nil
	case "babyRattle":
		return BorderStyleBabyRattle, nil
	case "balloons3Colors":
		return BorderStyleBalloons3Colors, nil
	case "balloonsHotAir":
		return BorderStyleBalloonsHotAir, nil
	case "basicBlackDashes":
		return BorderStyleBasicBlackDashes, nil
	case "basicBlackDots":
		return BorderStyleBasicBlackDots, nil
	case "basicBlackSquares":
		return BorderStyleBasicBlackSquares, nil
	case "basicThinLines":
		return BorderStyleBasicThinLines, nil
	case "basicWhiteDashes":
		return BorderStyleBasicWhiteDashes, nil
	case "basicWhiteDots":
		return BorderStyleBasicWhiteDots, nil
	case "basicWhiteSquares":
		return BorderStyleBasicWhiteSquares, nil
	case "basicWideInline":
		return BorderStyleBasicWideInline, nil
	case "basicWideMidline":
		return BorderStyleBasicWideMidline, nil
	case "basicWideOutline":
		return BorderStyleBasicWideOutline, nil
	case "bats":
		return BorderStyleBats, nil
	case "birds":
		return BorderStyleBirds, nil
	case "birdsFlight":
		return BorderStyleBirdsFlight, nil
	case "cabins":
		return BorderStyleCabins, nil
	case "cakeSlice":
		return BorderStyleCakeSlice, nil
	case "candyCorn":
		return BorderStyleCandyCorn, nil
	case "celticKnotwork":
		return BorderStyleCelticKnotwork, nil
	case "certificateBanner":
		return BorderStyleCertificateBanner, nil
	case "chainLink":
		return BorderStyleChainLink, nil
	case "champagneBottle":
		return BorderStyleChampagneBottle, nil
	case "checkedBarBlack":
		return BorderStyleCheckedBarBlack, nil
	case "checkedBarColor":
		return BorderStyleCheckedBarColor, nil
	case "checkered":
		return BorderStyleCheckered, nil
	case "christmasTree":
		return BorderStyleChristmasTree, nil
	case "circlesLines":
		return BorderStyleCirclesLines, nil
	case "circlesRectangles":
		return BorderStyleCirclesRectangles, nil
	case "classicalWave":
		return BorderStyleClassicalWave, nil
	case "clocks":
		return BorderStyleClocks, nil
	case "compass":
		return BorderStyleCompass, nil
	case "confetti":
		return BorderStyleConfetti, nil
	case "confettiGrays":
		return BorderStyleConfettiGrays, nil
	case "confettiOutline":
		return BorderStyleConfettiOutline, nil
	case "confettiStreamers":
		return BorderStyleConfettiStreamers, nil
	case "confettiWhite":
		return BorderStyleConfettiWhite, nil
	case "cornerTriangles":
		return BorderStyleCornerTriangles, nil
	case "couponCutoutDashes":
		return BorderStyleCouponCutoutDashes, nil
	case "couponCutoutDots":
		return BorderStyleCouponCutoutDots, nil
	case "crazyMaze":
		return BorderStyleCrazyMaze, nil
	case "creaturesButterfly":
		return BorderStyleCreaturesButterfly, nil
	case "creaturesFish":
		return BorderStyleCreaturesFish, nil
	case "creaturesInsects":
		return BorderStyleCreaturesInsects, nil
	case "creaturesLadyBug":
		return BorderStyleCreaturesLadyBug, nil
	case "crossStitch":
		return BorderStyleCrossStitch, nil
	case "cup":
		return BorderStyleCup, nil
	case "decoArch":
		return BorderStyleDecoArch, nil
	case "decoArchColor":
		return BorderStyleDecoArchColor, nil
	case "decoBlocks":
		return BorderStyleDecoBlocks, nil
	case "diamondsGray":
		return BorderStyleDiamondsGray, nil
	case "doubleD":
		return BorderStyleDoubleD, nil
	case "doubleDiamonds":
		return BorderStyleDoubleDiamonds, nil
	case "earth1":
		return BorderStyleEarth1, nil
	case "earth2":
		return BorderStyleEarth2, nil
	case "eclipsingSquares1":
		return BorderStyleEclipsingSquares1, nil
	case "eclipsingSquares2":
		return BorderStyleEclipsingSquares2, nil
	case "eggsBlack":
		return BorderStyleEggsBlack, nil
	case "fans":
		return BorderStyleFans, nil
	case "film":
		return BorderStyleFilm, nil
	case "firecrackers":
		return BorderStyleFirecrackers, nil
	case "flowersBlockPrint":
		return BorderStyleFlowersBlockPrint, nil
	case "flowersDaisies":
		return BorderStyleFlowersDaisies, nil
	case "flowersModern1":
		return BorderStyleFlowersModern1, nil
	case "flowersModern2":
		return BorderStyleFlowersModern2, nil
	case "flowersPansy":
		return BorderStyleFlowersPansy, nil
	case "flowersRedRose":
		return BorderStyleFlowersRedRose, nil
	case "flowersRoses":
		return BorderStyleFlowersRoses, nil
	case "flowersTeacup":
		return BorderStyleFlowersTeacup, nil
	case "flowersTiny":
		return BorderStyleFlowersTiny, nil
	case "gems":
		return BorderStyleGems, nil
	case "gingerbreadMan":
		return BorderStyleGingerbreadMan, nil
	case "gradient":
		return BorderStyleGradient, nil
	case "handmade1":
		return BorderStyleHandmade1, nil
	case "handmade2":
		return BorderStyleHandmade2, nil
	case "heartBalloon":
		return BorderStyleHeartBalloon, nil
	case "heartGray":
		return BorderStyleHeartGray, nil
	case "hearts":
		return BorderStyleHearts, nil
	case "heebieJeebies":
		return BorderStyleHeebieJeebies, nil
	case "holly":
		return BorderStyleHolly, nil
	case "houseFunky":
		return BorderStyleHouseFunky, nil
	case "hypnotic":
		return BorderStyleHypnotic, nil
	case "iceCreamCones":
		return BorderStyleIceCreamCones, nil
	case "lightBulb":
		return BorderStyleLightBulb, nil
	case "lightning1":
		return BorderStyleLightning1, nil
	case "lightning2":
		return BorderStyleLightning2, nil
	case "mapPins":
		return BorderStyleMapPins, nil
	case "mapleLeaf":
		return BorderStyleMapleLeaf, nil
	case "mapleMuffins":
		return BorderStyleMapleMuffins, nil
	case "marquee":
		return BorderStyleMarquee, nil
	case "marqueeToothed":
		return BorderStyleMarqueeToothed, nil
	case "moons":
		return BorderStyleMoons, nil
	case "mosaic":
		return BorderStyleMosaic, nil
	case "musicNotes":
		return BorderStyleMusicNotes, nil
	case "northwest":
		return BorderStyleNorthwest, nil
	case "ovals":
		return BorderStyleOvals, nil
	case "packages":
		return BorderStylePackages, nil
	case "palmsBlack":
		return BorderStylePalmsBlack, nil
	case "palmsColor":
		return BorderStylePalmsColor, nil
	case "paperClips":
		return BorderStylePaperClips, nil
	case "papyrus":
		return BorderStylePapyrus, nil
	case "partyFavor":
		return BorderStylePartyFavor, nil
	case "partyGlass":
		return BorderStylePartyGlass, nil
	case "pencils":
		return BorderStylePencils, nil
	case "people":
		return BorderStylePeople, nil
	case "peopleWaving":
		return BorderStylePeopleWaving, nil
	case "peopleHats":
		return BorderStylePeopleHats, nil
	case "poinsettias":
		return BorderStylePoinsettias, nil
	case "postageStamp":
		return BorderStylePostageStamp, nil
	case "pumpkin1":
		return BorderStylePumpkin1, nil
	case "pushPinNote2":
		return BorderStylePushPinNote2, nil
	case "pushPinNote1":
		return BorderStylePushPinNote1, nil
	case "pyramids":
		return BorderStylePyramids, nil
	case "pyramidsAbove":
		return BorderStylePyramidsAbove, nil
	case "quadrants":
		return BorderStyleQuadrants, nil
	case "rings":
		return BorderStyleRings, nil
	case "safari":
		return BorderStyleSafari, nil
	case "sawtooth":
		return BorderStyleSawtooth, nil
	case "sawtoothGray":
		return BorderStyleSawtoothGray, nil
	case "scaredCat":
		return BorderStyleScaredCat, nil
	case "seattle":
		return BorderStyleSeattle, nil
	case "shadowedSquares":
		return BorderStyleShadowedSquares, nil
	case "sharksTeeth":
		return BorderStyleSharksTeeth, nil
	case "shorebirdTracks":
		return BorderStyleShorebirdTracks, nil
	case "skyrocket":
		return BorderStyleSkyrocket, nil
	case "snowflakeFancy":
		return BorderStyleSnowflakeFancy, nil
	case "snowflakes":
		return BorderStyleSnowflakes, nil
	case "sombrero":
		return BorderStyleSombrero, nil
	case "southwest":
		return BorderStyleSouthwest, nil
	case "stars":
		return BorderStyleStars, nil
	case "starsTop":
		return BorderStyleStarsTop, nil
	case "stars3d":
		return BorderStyleStars3d, nil
	case "starsBlack":
		return BorderStyleStarsBlack, nil
	case "starsShadowed":
		return BorderStyleStarsShadowed, nil
	case "sun":
		return BorderStyleSun, nil
	case "swirligig":
		return BorderStyleSwirligig, nil
	case "tornPaper":
		return BorderStyleTornPaper, nil
	case "tornPaperBlack":
		return BorderStyleTornPaperBlack, nil
	case "trees":
		return BorderStyleTrees, nil
	case "triangleParty":
		return BorderStyleTriangleParty, nil
	case "triangles":
		return BorderStyleTriangles, nil
	case "tribal1":
		return BorderStyleTribal1, nil
	case "tribal2":
		return BorderStyleTribal2, nil
	case "tribal3":
		return BorderStyleTribal3, nil
	case "tribal4":
		return BorderStyleTribal4, nil
	case "tribal5":
		return BorderStyleTribal5, nil
	case "tribal6":
		return BorderStyleTribal6, nil
	case "twistedLines1":
		return BorderStyleTwistedLines1, nil
	case "twistedLines2":
		return BorderStyleTwistedLines2, nil
	case "vine":
		return BorderStyleVine, nil
	case "waveline":
		return BorderStyleWaveline, nil
	case "weavingAngles":
		return BorderStyleWeavingAngles, nil
	case "weavingBraid":
		return BorderStyleWeavingBraid, nil
	case "weavingRibbon":
		return BorderStyleWeavingRibbon, nil
	case "weavingStrips":
		return BorderStyleWeavingStrips, nil
	case "whiteFlowers":
		return BorderStyleWhiteFlowers, nil
	case "woodwork":
		return BorderStyleWoodwork, nil
	case "xIllusions":
		return BorderStyleXIllusions, nil
	case "zanyTriangles":
		return BorderStyleZanyTriangles, nil
	case "zigZag":
		return BorderStyleZigZag, nil
	case "zigZagStitch":
		return BorderStyleZigZagStitch, nil
	default:
		return "", errors.New("invalid BorderStyle value")
	}
}

// UnmarshalXMLAttr unmarshals an XML attribute into a BorderStyle.
func (b *BorderStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := BorderStyleFromStr(attr.Value)
	if err != nil {
		return err
	}

	*b = val

	return nil
}
