# Doc

## TL;DR give me planet vector tiles!

1. Have [Docker]('https://docs.docker.com/get-docker/') installed.
2. Install the following:

```bash
sudo apt-get install build-essential libsqlite3-dev zlib1g-dev
```

3. `npx sequentially-generate-planet-mbtiles`

4. [Rejoice]('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmedia.giphy.com%2Fmedia%2FWIg8P0VNpgH8Q%2Fgiphy.gif&f=1&nofb=1') - see acknowledgements below for people to thank.

## config.js (defaults shown)

### This can be supplied as follows:

```bash
npx sequentially-generate-planet-mbtiles -c config.js
```

```javascript
// config.js
{
  subRegions: [
    "africa",
    "antarctica",
    "asia",
    "australia-oceania",
    "central-america",
    "europe",
    "north-america",
    "south-america",
  ];
  keepDownloadedFiles: false;
  keepSubRegionMbtiles: false;
}
```

**_subRegions_** - Defaults to downloading the each of the largest sub regions provided by Geofabrik in order to create vector tiles for the entire planet. Entries must be in the correct format according to the GEOFABRIK (https://download.geofabrik.de/) download api's. e.g. "australia-oceania-latest.osm.pbf" should be "australia-oceania"; "chad-latest.osm.pbf" should be "africa/chad"; "europe" will be downloaded from https://download.geofabrik.de/europe.html.

**_keepDownloadedFiles_** - If true, downloaded files will be kept in the pbf directory. If false, they will be deleted. Files will not be downloaded if they are already present. `True` will use over twice the disk space upon completion. We would recommend that this option is selected if you foresee multiple attempts/downloads in your future - be kind to Geofabrik <3.

**_keepSubRegionMbtiles_** - If true, each sub region mbtiles file (e.g. asia.mbtiles) will be kept, further drastically increasing required disk space. This may be particularly useful on old or slow hardware that has the tendancy to crash or give up!.

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
2. **Do I have to download the entire planet?** Not at all. Simply remove/change the `config.js` `subRegions` array to include only the areas you want. Once downloaded, they will be merged together into a single file called `planet.mbtiles`. You can then rename that file to something more appropriate.

## Todo

1. Extra error handling for if one of the third party processes should fail.
2. Config file for area inputs.
