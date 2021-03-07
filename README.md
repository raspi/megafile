# megafile
MegaFile (.MEG) reader.

File format information from [OpenRA](https://github.com/OpenRA/OpenRA) and [official Electronic Arts Command & Conquer Remastered Collection source code](https://github.com/electronicarts/CnC_Remastered_Collection)


1. Buy Command & Conquer™ Remastered Collection
1. Extract music files from `MUSIC.MEG`
1. 🎶🤘😎🤘🎶

# Extract files:

```shell
% ./megafile -x ./steamapps/common/CnCRemastered/Data/MUSIC.MEG 
file #001(idx:193)/203: data/audio/music/tdr_mus_reaching_out.wav 5760586 bytes at offset 13016
file #002(idx:13)/203: data/audio/music/rab_mus_the_plan_am.wav 9472126 bytes at offset 5773604
file #003(idx:32)/203: data/audio/music/rac_mus_hell_march_retaliation_remix.wav 2555916 bytes at offset 15245732

...

file #201(idx:135)/203: data/audio/music/tdc_mus_great_shot_extended.wav 678166 bytes at offset 1442382672
file #202(idx:41)/203: data/audio/music/rac_mus_reload_fire.wav 1970506 bytes at offset 1443060840
file #203(idx:37)/203: data/audio/music/rac_mus_mud_retaliation_remix.wav 3463536 bytes at offset 1445031348
```

Convert all to FLAC:

    % find . -iname "*.wav" -exec ffmpeg -hide_banner -i "{}" "{}.flac" \;

# Build executable:

    make

# Release:

    make release
