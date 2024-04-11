# Itinerary_Prettifier Tool

The Itinerary Prettification Tool is a command-line utility designed to enhance the readability of flight itineraries.

This tool accepts a text-based itinerary from a specified file as input ```input.txt```, processes the text to improve its presentation for end-users, and saves the refined version to a new file ```output.txt```. It utilizes a CSV lookup file ```airport-lookup.csv```, which is provided alongside the program, to convert IATA and ICAO codes into their corresponding airport names.

An IATA code is denoted by a single # followed by three letters. For instance, #CDG represents the IATA code for "Charles de Gaulle International Airport". Conversely, an ICAO code is represented by double # followed by four letters. For example, ##EHAM corresponds to the ICAO code for Amsterdam Airport Schiphol in Netherlands.

Additionally, the program beautifies dates and times that adhere to the ISO 8601 standard:

Dates in the format D(2007-04-05T12:30−02:00) are displayed in the output as DD-Mmm-YYYY. E.g. "05 Apr 2007".
12-hour times specified as T12(2007-04-05T12:30−02:00) are presented as "12:30PM (-02:00)".
24-hour times indicated as T24(2007-04-05T12:30−02:00) are formatted as "12:30 (-02:00)".


## Usage
To execute the tool, use the command line with the following arguments:

Path to the input file.
Path to the output file.
Path to the airport lookup CSV file.
```$ go run . ./input.txt ./output.txt ./airport-lookup.csv```


## The -h flag can be used to display usage instructions:
```$ go run . -h```


## itinerary usage:
```go run . ./input.txt ./output.txt ./airport-lookup.csv```





