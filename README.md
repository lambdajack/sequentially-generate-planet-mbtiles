# Sequentially Generate Planet Mbtiles

Catchy name right?

### _Sequentially generate and merge an entire planet.mbtiles vector tileset on low memory/power devices for free... slowly._

## TL;DR give me planet vector tiles!

1. Have [Docker]('https://docs.docker.com/get-docker/') installed.
2. Install the following:

```bash
sudo apt-get install build-essential libsqlite3-dev zlib1g-dev
```

3. Run with npx:

```bash
npx sequentially-generate-planet-mbtiles
```

4.  Rejoice - see acknowledgements below for people to thank.

## config.json (defaults shown)

### This can be supplied as follows:

```bash
npx sequentially-generate-planet-mbtiles -c /path/to/config.json
```

```json
// config.json
{
  "subRegions": [
    "africa/algeria",
    "africa/angola",
    "africa/benin",
    "africa/botswana",
    "africa/burkina-faso",
    "africa/burundi",
    "africa/cameroon",
    "africa/cape-verde",
    "africa/central-african-republic",
    "africa/chad",
    "africa/comoros",
    "africa/congo-brazzaville",
    "africa/congo-democratic-republic",
    "africa/djibouti",
    "africa/egypt",
    "africa/equatorial-guinea",
    "africa/eritrea",
    "africa/ethiopia",
    "africa/gabon",
    "africa/ghana",
    "africa/guinea",
    "africa/guinea-bissau",
    "africa/ivory-coast",
    "africa/kenya",
    "africa/lesotho",
    "africa/liberia",
    "africa/libya",
    "africa/madagascar",
    "africa/malawi",
    "africa/mali",
    "africa/mauritania",
    "africa/mauritius",
    "africa/morocco",
    "africa/mozambique",
    "africa/namibia",
    "africa/niger",
    "africa/nigeria",
    "africa/rwanda",
    "africa/saint-helena-ascension-and-tristan-da-cunha",
    "africa/sao-tome-and-principe",
    "africa/senegal-and-gambia",
    "africa/seychelles",
    "africa/sierra-leone",
    "africa/somalia",
    "africa/south-africa",
    "africa/south-sudan",
    "africa/sudan",
    "africa/swaziland",
    "africa/tanzania",
    "africa/togo",
    "africa/tunisia",
    "africa/uganda",
    "africa/zambia",
    "africa/zimbabwe",
    "antarctica",
    "asia/afghanistan",
    "asia/armenia",
    "asia/azerbaijan",
    "asia/bangladesh",
    "asia/bhutan",
    "asia/bhutan",
    "asia/cambodia",
    "asia/china",
    "asia/east-timor",
    "asia/gcc-states",
    "asia/india",
    "asia/indonesia",
    "asia/iran",
    "asia/iraq",
    "asia/isreal-and-palestine",
    "asia/japan",
    "asia/jordan",
    "asia/kazakhstan",
    "asia/kyrgyzstan",
    "asia/laos",
    "asia/lebanon",
    "asia/malaysia-singapore-brunei",
    "asia/maldives",
    "asia/mongolia",
    "asia/myanmar",
    "asia/nepal",
    "asia/north-korea",
    "asia/pakistan",
    "asia/philippines",
    "asia/russia",
    "asia/south-korea",
    "asia/sri-lanka",
    "asia/syria",
    "asia/taiwan",
    "asia/tajikistan",
    "asia/thailand",
    "asia/turkmenistan",
    "asia/uzbekistan",
    "asia/vietnam",
    "asia/yemen",
    "australia-oceania/american-oceania",
    "australia-oceania/australia",
    "australia-oceania/cook-islands",
    "australia-oceania/fiji",
    "australia-oceania/ile-de-clipperton",
    "australia-oceania/kiribati",
    "australia-oceania/marshall-islands",
    "australia-oceania/micronesia",
    "australia-oceania/nauru",
    "australia-oceania/new-caledonia",
    "australia-oceania/new-zealand",
    "australia-oceania/niue",
    "australia-oceania/palau",
    "australia-oceania/papua-new-guinea",
    "australia-oceania/pitcairn-islands",
    "australia-oceania/polynesie-francaise",
    "australia-oceania/samoa",
    "australia-oceania/solomon-islands",
    "australia-oceania/tokelau",
    "australia-oceania/tonga",
    "australia-oceania/tuvalu",
    "australia-oceania/vanuatu",
    "australia-oceania/wallis-et-futuna",
    "central-america/bahamas",
    "central-america/belize",
    "central-america/costa-rica",
    "central-america/costa-rica",
    "central-america/el-salvador",
    "central-america/guatemala",
    "central-america/haiti-and-domrep",
    "central-america/honduras",
    "central-america/jamacia",
    "central-america/nicaragua",
    "central-america/panama",
    "europe/albania",
    "europe/andorra",
    "europe/austria",
    "europe/azores",
    "europe/belarus",
    "europe/belgium",
    "europe/bosnia-herzegovina",
    "europe/bulgaria",
    "europe/croatia",
    "europe/cyprus",
    "europe/czech-republic",
    "europe/denmark",
    "europe/estonia",
    "europe/faroe-islands",
    "europe/finland",
    "europe/france",
    "europe/georgia",
    "europe/germany",
    "europe/great-britain",
    "europe/greece",
    "europe/guernsey-jersey",
    "europe/hungary",
    "europe/iceland",
    "europe/ireland-and-northern-ireland",
    "europe/isle-of-man",
    "europe/italy",
    "europe/kosovo",
    "europe/latvia",
    "europe/liechtenstein",
    "europe/lithuania",
    "europe/luxembourg",
    "europe/macedonia",
    "europe/malta",
    "europe/moldova",
    "europe/monaco",
    "europe/montenegro",
    "europe/netherlands",
    "europe/norway",
    "europe/poland",
    "europe/portugal",
    "europe/romania",
    "europe/russia",
    "europe/serbia",
    "europe/slovakia",
    "europe/slovenia",
    "europe/spain",
    "europe/sweden",
    "europe/switzerland",
    "europe/turkey",
    "europe/ukraine",
    "north-america/canada",
    "north-america/greenland",
    "north-america/mexico",
    "north-america/us",
    "south-america/argentina",
    "south-america/bolivia",
    "south-america/brazil",
    "south-america/chile",
    "south-america/colombia",
    "south-america/ecuador",
    "south-america/paraguay",
    "south-america/peru",
    "south-america/suriname",
    "south-america/uruguay",
    "south-america/venezuela"
  ],
  "keepDownloadedFiles": true,
  "keepSubRegionMbtiles": false,
  "tileZoomLevel": 12
}
```

**_subRegions_** - Defaults to downloading the each of the largest sub regions provided by Geofabrik in order to create vector tiles for the entire planet. Entries must be in the correct format according to the GEOFABRIK (https://download.geofabrik.de/) download api's. e.g. "australia-oceania-latest.osm.pbf" should be "australia-oceania"; "chad-latest.osm.pbf" should be "africa/chad"; "europe" will be downloaded from https://download.geofabrik.de/europe.html.

**_keepDownloadedFiles_** - If true, downloaded files will be kept in the pbf directory. If false, they will be deleted. Files will not be downloaded if they are already present. `True` will use over twice the disk space upon completion. We would recommend that this option is selected if you foresee multiple attempts/downloads in your future - be kind to Geofabrik <3.

**_keepSubRegionMbtiles_** - If true, each sub region mbtiles file (e.g. asia.mbtiles) will be kept, further drastically increasing required disk space. This may be particularly useful on old or slow hardware that has the tendancy to crash or give up!

**_tileZoomLevel_** - Default 12. This sets the amount of detail you will see. 14 is as detailed as most people would ever want but be warned, each point increases the amount of work required by a factor of 4. Please see: https://wiki.openstreetmap.org/wiki/Zoom_levels. For some perspective, to render Antarctica, on a 6 year old, 4 cores/8gb RAM, pc at the default zoom level 12, takes about 15 hours. at zoom level 14, estimates about 9 _days_; zoom level 10 takes less than an hour. While this programme is doing its best to give vector tiles to everyone with low spec hardware (and it certainly will), ask yourself how long you are prepared to wait. If you need a high zoom number with lots of map detail for a project you are working on, consider setting the programme off well in advance while your project is still in development; that way, minimal if no time will be wasted. Setting any higher will take weeks to process on mid spec hardware.

## Why?

There are some wonderful options out there for generating and serving your own map data and there are many reasons to want to do so. My reason, and the inspiration for this programme was cost. It is expensive to use a paid tile server option after less users using it than you might think. The problem is, when trying to host your own, a lot of research has shown me that almost all solutions for self generating tiles for a map server require hugely expensive hardware to even complete (it's not uncommon to see requirements for 64 cores and 128gb RAM!). Indeed the largest I've seen wanted 150gb of the stuff!. For generating the planet that is. If you want a small section of the world, then it is much easier. But I need the planet - so what to do? Generate smaller sections of the world, then combine them.

That's where this comes in. It does not appear to be a simple, convenient or well documented at least, process of getting everything setup to do this 'bit by bit' approach. It's not too challenging, but it is time consuming, and without a script anyway it requires rather frequent attention on your part.

**_This programme aims to be a simple set and forget, one liner which gives anyone - even those who are not the most technically minded, or just can't be bothered - a way to get a full-featured and bang up to date set of vector tiles for the entire planet ON SMALL HARDWARE._**

It's also designed (work in progress) to be fail safe - meaning that if your hardward (or our software) does crash mid process, you have not lost all your data, and you are able to start again from a point mid-way through.

It's a work in progress - but it works - again, slowly. I'll do what I can to make it much more robust as time goes on.

This also uses the maptiler mbtiles spec, meaning when you serve the files with something like tileserver-gl, you don't have to worry about setting up styles, as the basic one will be automatically available. Use the -s option to automatically serve the files when done on `http://localhost:8080`. (-s not yet implemented).

We make extensive use of openmaptiles, which in theory, does not require a huge amount of RAM, but I have tried it on a few high spec 'consumer' machines (circa. £2000-3000) and the process is never able to complete (and if it fails - you have to start all over again mostly - at least from the parts which took the longest anyway). I have spoken with a few people who have had a similar experience. That's why this has been made, to work on hardware as low as 4gb/4cores. If anyone can test any lower (who has the time though?) please let me know!

## Requirements

### Hardware

1. About 500gb clear disk space for the entire planet. Probably an SSD unless you like pain, suffering and dying of old age.
2. Probably about 8gb of RAM if you will be downloading whole continents - less if you adjust the config file to download smaller chunks at a time.
3. Time. As above, this has been written to massively streamline the process of getting a planetary vector tile set for the average person who might not have the strongest hardware or the desire to spend £££ on a 64 core 128gb RAM server. Unfortunately, if you cut out the cost, you increase the time. By a lot. Expect the entire planet to take DAYS on average hardware.

### Software

1. Have the following installed:
   ```bash
   sudo apt-get install build-essential libsqlite3-dev zlib1g-dev
   ```
2. Docker

### Run an 'low load' test

```bash
sequentially-generate-planet-mbtiles -t
```

Add the `-t` options to use the presupplied test-config.json. This test will generate low zoom levels of a small area in Africa and even on the lowest powered hardware should not take more and 20 minutes to run - often less than 5 minutes.

## Things to look out for

1. If starting the process halfway through (e.g. it crashed and you are resuming), the terminal may ask your permission when it comes to writing over certain files, since they were created with sudo privileges.

## How to serve?

We would recommend something like [tileserver-gl]('https://github.com/maptiler/tileserver-gl). Further reading can be found [here]('https://wiki.openstreetmap.org/wiki/MBTiles') (openstreetmap wiki).

## FAQ

1. **How long will this take?** Low spec hardware? Whole planet? Days/weeks. A few days for reasonable hardward. Small sections can be done in as little as a few minutes.
2. **Why do I have to run part of the programme with 'sudo' privileges?** You might not have to depending on your system, but most modern linux systems require sudo for commands like `make install`, which are required here. Therefore, we run those commands as sudo as a catch-all.
3. **Do I have to download the entire planet?** Not at all. Simply remove/change the `config.json` `subRegions` array to include only the areas you want. Once downloaded, they will be merged together into a single file called `planet.mbtiles`. You can then rename that file to something more appropriate.
4. **It's running, but my pbf folder is empty - should I be worried?** Check the openmaptiles/data folder. If your config has selected to delete files downloaded, then they will be moved rather than copied.
5. **Ubuntu only?** Nope! It should work on any distro as long as the dependancies are installed.
6. **Does 'low spec' mean I can run it on my toaster?** Maybe, but mostly not. But you can happily run it on you 4core8gb ram home pc without too much trouble. Just time.
7. **_Why javascript and not bash or something?_** Two reasons - 1: _"Anything that can be Written in JavaScript, will Eventually be Written in JavaScript"_; 2: The pure simplicity of typing `npx command` is unparralelled and is rather system agnostic. This is written to be simple, and tile servers are often used by web developers, so they are likely to have node and npm already available to use right away. In short, simplicity in use.

## Acknowledgements

Please take the time to thank the folks over at [openmaptiles]('https://github.com/openmaptiles/openmaptiles') and [tippecanoe]('https://github.com/mapbox/tippecanoe'). They are the reason any of this is possible in the first place.

## Prefer not to use npx?

```bash
git clone https://github.com/lambdajack/sequentially-generate-planet-mbtiles
cd sequentially-generate-planet-mbtiles
node main.js -c /path/to/your/config.json # omit -c if you want to use the defaults.
```

## Contributions

All welcome! Feature request, pull request, bug reports/fixes etc - go for it.

We'd like to make this tool quite robust moving forward - since we needed it for a current project of ours, we have released it notwithstanding the current rough-and-ready nature.

Feedback on this one is much appreciated :D.

## Development

We already depend on others enough for this, so the programme is written without any npm dependancies. This may change in the future though if we want to make the terminal pretty etc. Simply clone the repo and you're good to go. Just make sure you have the dependancies installed as above.

Use the provided `development-config.json` as it is preconfigured to keep downloaded data and only download very small regions for quick testing.

## Todo

1. TS conversion before significant improvement or features added.
2. Extra error handling for if one of the third party processes should fail.
3. The ability to select different system drives for downloading/generating files.
4. Write tests before significant future development.
5. Make the console prettier.
6. Add option to include or not ocean tiles -o.
7. Add automatically serve on completion option -s.
8. Typos and grammatical errors on readme - add section about distritubing the load for tile generation - maybe.
9. Add a success.json file showing which regions were successful, and ask the user if they would like to retry any failed regions
10. Add sleep function at points in stdout user may wish to read and hold it there for a time.
11. Remove overwrite planet.mbtiles - document that it is for the user to ensure the data is clean or only contains what they want.
12. Add realistic schedules - two weeks for zoom level 12 - add the wiki information (directly?)
