#

a tools rename rss file to same as video name  

```bash
usage: rename <old(video name)> <new(rss name)> [max]

max: counter, default 12
```

eg:  

rename rss pattern: `[NOIR][BDrip][%02d.*][X264_AAC][.*]` to `Noir - S01E%02d - .*`

```bash
$ rename_vs "Noir - S01E%02d - .*" "[NOIR][BDrip][%02d.*][X264_AAC][.*]" 26

mv "[NOIR][BDrip][01][X264_AAC][xxx].ass" "Noir - S01E01 - xxx.ass"
mv "[NOIR][BDrip][02][X264_AAC][xxx].ass" "Noir - S01E02 - xxx.ass"
    ...
    ...
mv "[NOIR][BDrip][26][X264_AAC][xxx].ass" "Noir - S01E26 - xxx.ass"
```
