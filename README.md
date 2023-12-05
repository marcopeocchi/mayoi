# Mayoi 🐌

A RSS/Torznab indexer anime indexer.  
Focuses on speed and low memory usage.

![Screenshot 2023-12-04 alle 20 52 08](https://github.com/marcopeocchi/mayoi/assets/35533749/57b809ef-25ed-4007-923d-64e46d1fe200)

## A module oriented approach
Each indexer is contained in a contained unit called "module".  
Under `internal/nyaa` there's an example `Nyaa.si` indexer.  

If you want to add an indexer just copy the structure and submit a PR 😎.  
Soon a **wiki** page will be added for following the best rule when implementing an indexer.

## Performance oriented
This project aims to be an alternative to [prowlarr](https://github.com/Prowlarr/Prowlarr) when it comes to provide a **torznab** compatile indexer.
The main focus should be a low memory footprint for running **mayoi** on any device.

For instance, I'm running **mayoi** on a `Raspberry Pi 1 model A` with 256MB of RAM clocked at 1GHz and experiencing good response time.

## Docker
A docker image is availabe at `marcobaobao/mayoi` on Docker.io CR.

```sh
docker run -d --name mayoi -p 6969:6969 -v /path/to/config:/config marcobaobao/mayoi
```

## Config
Provide a config file named `config.yml` with the following structure:
```yaml
indexers:
  - https://indexer1
  - https://indexer2/rss
```
and feed it to **mayoi**

```sh
mayoi -c config.yml -d mayoi.db
```

## Currently implemented indexers
- Nyaa
- AnimeTime
