# Sequentially Generate Planet Mbtiles

Catchy name right?

### _Sequentially generate and merge an entire planet.mbtiles vector tileset on low memory/power devices... slowly._

## TL;DR give me planet vector tiles!

1. Have [Docker]('https://docs.docker.com/get-docker/') installed.
2. Install the following:

```bash
sudo apt-get install build-essential libsqlite3-dev zlib1g-dev
```

3. `npx sequentially-generate-planet-mbtiles`

4. [Rejoice]('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmedia.giphy.com%2Fmedia%2FWIg8P0VNpgH8Q%2Fgiphy.gif&f=1&nofb=1') - see acknowledgements below for people to thank.

## config.json (defaults shown)

### This can be supplied as follows:

```bash
npx sequentially-generate-planet-mbtiles -c /path/to/config.json
```

```json
// config.json
{
  "subRegions": [
    "africa",
    "antarctica",
    "asia",
    "australia-oceania",
    "central-america",
    "europe",
    "north-america",
    "south-america"
  ],
  "keepDownloadedFiles": false,
  "keepSubRegionMbtiles": false
}
```

**_subRegions_** - Defaults to downloading the each of the largest sub regions provided by Geofabrik in order to create vector tiles for the entire planet. Entries must be in the correct format according to the GEOFABRIK (https://download.geofabrik.de/) download api's. e.g. "australia-oceania-latest.osm.pbf" should be "australia-oceania"; "chad-latest.osm.pbf" should be "africa/chad"; "europe" will be downloaded from https://download.geofabrik.de/europe.html.

**_keepDownloadedFiles_** - If true, downloaded files will be kept in the pbf directory. If false, they will be deleted. Files will not be downloaded if they are already present. `True` will use over twice the disk space upon completion. We would recommend that this option is selected if you foresee multiple attempts/downloads in your future - be kind to Geofabrik <3.

**_keepSubRegionMbtiles_** - If true, each sub region mbtiles file (e.g. asia.mbtiles) will be kept, further drastically increasing required disk space. This may be particularly useful on old or slow hardware that has the tendancy to crash or give up!.

## Why?

There are some wonderful options out there for generating and serving your own map data and there are many reasons to want to do so. My reason, and the inspiration for this programme was cost. It is expensive to use a paid tile server option after less users using it than you might think. The problem is, when trying to host your own, a lot of research has shown me that almost all solutions for self generating tiles for a map server require hugely expensive hardware to even complete (it's not uncommon to see requirements for 64 cores and 128gb RAM!). Indeed the largest I've seen wanted 150gb of the stuff!. For generating the planet that is. If you want a small section of the world, then it is much easier. But I need the planet - so what to do? Generate smaller sections of the world, then combine them.

That's where this comes in. It does not appear to be a simple, convenient or well documented at least, process of getting everything setup to do this 'bit by bit' approach. It's not too challenging, but it is time consuming, and without a script anyway it requires rather frequent attention on your part.

**_This programme aims to be a simple set and forget, one liner which gives anyone - even those who are not the most technically minded, or just can't be bothered - a way to get a full-featured and bang up to date set of vector tiles for the entire planet ON SMALL HARDWARE._**

It's also designed (work in progress) to be fail safe - meaning that if your hardward (or our software) does crash mid process, you have not lost all your data, and you are able to start again from a point mid-way through.

It's a work in progress - but it works - again, slowly. I'll do what I can to make it much more robust as time goes on.

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

## How to serve?

We would recommend something like [tileserver-gl]('https://github.com/maptiler/tileserver-gl). Further reading can be found [here]('https://wiki.openstreetmap.org/wiki/MBTiles') (openstreetmap wiki).

## FAQ

1. **Why do I have to run part of the programme with 'sudo' privileges?** You might not have to depending on your system, but most modern linux systems require sudo for commands like `make install`, which are required here. Therefore, we run those commands as sudo as a catch-all.
2. **Do I have to download the entire planet?** Not at all. Simply remove/change the `config.json` `subRegions` array to include only the areas you want. Once downloaded, they will be merged together into a single file called `planet.mbtiles`. You can then rename that file to something more appropriate.

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

## Todo

1. TS conversion before significant improvement or features added.
2. Extra error handling for if one of the third party processes should fail.
3. The ability to select different system drives for downloading/generating files.
