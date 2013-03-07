// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package weather

// List of weather codes mapped to descriptions, as defined at
// http://www.worldweatheronline.com/feed/wwoConditionCodes.txt
//
// Code partly generated with:
//
// $ wget http://www.worldweatheronline.com/feed/wwoConditionCodes.txt
// $ cat wwoConditionCodes.txt | awk '{ print "case \""$1"\":\nreturn \""$2"\"" }' > codes.go

func codeName(code string) string {
	switch code {
	case "395":
		return "Moderate"
	case "392":
		return "Patchy"
	case "389":
		return "Moderate"
	case "386":
		return "Patchy"
	case "377":
		return "Moderate"
	case "374":
		return "Light"
	case "371":
		return "Moderate"
	case "368":
		return "Light"
	case "365":
		return "Moderate"
	case "362":
		return "Light"
	case "359":
		return "Torrential"
	case "356":
		return "Moderate"
	case "353":
		return "Light"
	case "350":
		return "Ice"
	case "338":
		return "Heavy"
	case "335":
		return "Patchy"
	case "332":
		return "Moderate"
	case "329":
		return "Patchy"
	case "326":
		return "Light"
	case "323":
		return "Patchy"
	case "320":
		return "Moderate"
	case "317":
		return "Light"
	case "314":
		return "Moderate"
	case "311":
		return "Light"
	case "308":
		return "Heavy"
	case "305":
		return "Heavy"
	case "302":
		return "Moderate"
	case "299":
		return "Moderate"
	case "296":
		return "Light"
	case "293":
		return "Patchy"
	case "284":
		return "Heavy"
	case "281":
		return "Freezing"
	case "266":
		return "Light"
	case "263":
		return "Patchy"
	case "260":
		return "Freezing"
	case "248":
		return "Fog"
	case "230":
		return "Blizzard"
	case "227":
		return "Blowing"
	case "200":
		return "Thundery"
	case "185":
		return "Patchy"
	case "182":
		return "Patchy"
	case "179":
		return "Patchy"
	case "176":
		return "Patchy"
	case "143":
		return "Mist"
	case "122":
		return "Overcast"
	case "119":
		return "Cloudy"
	case "116":
		return "Partly"
	case "113":
		return "Clear/Sunny"
	}

	return "Unknown condition (" + code + ")"
}
