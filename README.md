# Sleeping Dogs Radio Song Replacer

This tool allows you to replace the songs on the radio stations in **Sleeping Dogs: Definitive Edition**.

---

## Requirements

1. **Your music converted to `.wav` format**
   - Convert your audio to **44100Hz, Signed 16-bit PCM, Stereo `.wav`**
   - You can use tools such as:
     - **Audacity**
     - **ffmpeg** (`ffmpeg -i input.mp3 -ar 44100 -ac 2 -sample_fmt s16 output.wav`)

2. **Wwise 2013.2**
   - Sleeping Dogs uses **Wwise 2013.2**, and replacement audio must be encoded with that exact version to work correctly.
   - You can download it here:
     https://archive.org/details/wwise-2013.2.10-windows

---

## Converting `.wav` Files to `.wem`

1. Launch **Wwise 2013.2**.
2. Create a new project.
3. Import your `.wav` files.
4. In the **Project Explorer**, navigate to: ShareSets → Conversion Settings → Vorbis.
5. Double-click **Vorbis Quality High**.
6. Click **Edit...** next to **Windows®**.
7. Set the following:

- **Seek table granularity (sample frames):** `1024`

8. Click **OK**, then click **Convert...**
9. Your `.wem` files will be output to: `<your project folder>/.cache/Windows/SFX/`
10. Rename your `.wem` files to match the **song IDs** you are replacing.

## Replacing the Songs

1. Go to your Sleeping Dogs installation directory.
2. Navigate to: `/data/Audio/SD2/`
3. Create a folder for your replacements and place all renamed `.wem` files in that folder.
4. Run the replacer tool: `replacer.exe -f SFX.pck -t replacements -o new_SFX.pck`

Where:
| Option | Meaning |
|--------|---------|
| `-f SFX.pck` | The original audio pack from the game |
| `-t replacements` | Directory containing your `.wem` files |
| `-o new_SFX.pck` | Output file with your custom songs |

5. **Back up your original** `SFX.pck`
6. Replace it: new_SFX.pck → rename to → SFX.pck

## Radio stations

### Boosey & Hawkes

| ID           | Artist       | Title                                                                     |
| ------------ | ------------ | ------------------------------------------------------------------------- |
| `1035836851` | Tchaikovsky  | The Nutcracker March                                                      |
| `233155423`  | Khachaturian | Sabre Dance                                                               |
| `517691373`  | Offenbach    | Orpheus in the Underworld, Act II: Can-Can                                |
| `414761973`  | Edvard Grieg | Morning Mood from the play Peer Gynt                                      |
| `11795260`   | Debussy      | Suite bergamasque: III. Clair de lune                                     |
| `415124579`  | Brahms       | Lullaby (Wiegenlied), Op. 49, No. 4                                       |
| `437623531`  | Beethoven    | Bagatelle in A minor ‘Für Elise’                                          |
| `989722079`  | Wagner       | Die Walkure: Ride of the Valkyries                                        |
| `820758683`  | Shostakovich | he Gadfly Suite / Five Days – Five Nights Suite – XII. Finale             |
| `70192315`   | Rachmaninoff | Vocalise, Op. 34, No. 14                                                  |
| `882187216`  | Pachelbel    | Canon                                                                     |
| `797274795`  | Mozart       | Piano Sonata No. 11 in A major, K. 331: III. Rondo alla turca: Allegretto |
| `752488945`  | Handel       | Messiah (Choruses) – Hallelujah                                           |
| `754242689`  | Bach         | Orchestral Suite No. 3 in D major, BWV 1068: II. Air, “Air on a G String” |
| `25478506`   | Verdi        | Rigoletto, Act III: La donna è mobile                                     |

### Daptone Radio

| ID          | Artist          | Title                                        |
| ----------- | --------------- | -------------------------------------------- |
| `714775338` |                 |                                              |
| `818423379` |                 |                                              |
| `815263223` |                 |                                              |
| `288288967` |                 |                                              |
| `640562369` |                 |                                              |
| `271053776` |                 |                                              |
| `841884139` |                 |                                              |
| `42868061`  |                 |                                              |
| `953563238` |                 |                                              |
| `41136723`  |                 |                                              |
| `716596295` |                 |                                              |
| `345990911` |                 |                                              |
| `8034663`   |                 |                                              |
| `283222090` | Charles Bradley | This Love Ain’t Big Enough for the Two of Us |
| `347445644` |                 |                                              |

### Softly Radio

| ID          | Artist              | Title         |
| ----------- | ------------------- | ------------- |
| `622753165` | Bei Bei & Shawn Lee | East          |
| `935324680` | Bei Bei & Shawn Lee | Hot Thursday  |
| `153030135` | Bei Bei & Shawn Lee | Into the Wind |
| `941230234` |                     |               |
| `101952233` |                     |               |
| `652411807` |                     |               |
| `207849330` |                     |               |
| `241881013` |                     |               |
| `481578451` |                     |               |
| `125962848` |                     |               |
| `140059496` |                     |               |
| `258199652` |                     |               |
| `794864812` |                     |               |
| `547919333` |                     |               |
| `922951351` |                     |               |
| `582140120` |                     |               |
| `149910693` |                     |               |
| `79150297`  |                     |               |
| `447631139` |                     |               |
| `216857275` |                     |               |
| `82423709`  |                     |               |
| `306677132` |                     |               |
| `26626127`  |                     |               |
| `987071588` |                     |               |
| `173659439` |                     |               |
| `944123865` |                     |               |
| `853284021` |                     |               |
| `975592254` |                     |               |
| `228832971` |                     |               |
| `76921858`  |                     |               |
| `26011993`  |                     |               |
| `536695358` |                     |               |
| `241516949` |                     |               |
| `624036716` |                     |               |
| `307971095` |                     |               |
| `281972048` |                     |               |
| `806900474` |                     |               |
| `229580733` |                     |               |

### Warp Radio

| ID           | Artist        | Title    |
| ------------ | ------------- | -------- |
| `71468293`   |               |          |
| `285695891`  | Africa HiTech | Lash Out |
| `756688699`  |               |          |
| `486816462`  |               |          |
| `669687173`  |               |          |
| `151555984`  |               |          |
| `10361450`   |               |          |
| `203525034`  |               |          |
| `64771429`   |               |          |
| `1013348789` |               |          |
| `519922173`  |               |          |
| `641371525`  |               |          |
| `895845900`  |               |          |
| `860265115`  |               |
| `883002255`  |               |          |
| `731977943`  |               |          |
| `17700560`   |               |          |
| `704417318`  |               |          |
| `934384691`  |               |          |

### Ninja Tuna Radio

| ID          | Artist | Title |
| ----------- | ------ | ----- |
| `28462829`  |        |       |
| `66320576`  |        |       |
| `906315897` | Bonobo | Kiara |
| `621641940` |        |       |
| `241297753` |        |       |
| `664950307` |        |       |
| `234941887` |        |       |
| `239024824` |        |       |
| `182429299` |        |       |
| `409081372` |        |       |
| `731679828` |        |       |
| `421052977` |        |       |
| `405300075` |        |       |
| `306295686` |        |       |
| `585648168` |        |       |
| `461101014` |        |       |
| `642310041` |        |       |
| `529411208` |        |       |
| `341244601` |        |       |
| `396657866` |        |       |

### H-Klub Radio

| ID          | Artist | Title |
| ----------- | ------ | ----- |
| `198543761` |        |       |
| `759776786` |        |       |
| `751151192` |        |       |
| `249635153` |        |       |
| `65354999`  |        |       |
| `206432563` |        |       |
| `255500584` |        |       |
| `622161138` |        |       |
| `803206285` |        |       |
| `785509598` |        |       |
| `421091422` |        |       |
| `63805603`  |        |       |
| `215466551` |        |       |
| `283634792` |        |       |
| `403668441` |        |       |
| `379924086` |        |       |
| `506526332` |        |       |

### Kerrang! Radio

| ID          | Artist         | Title            |
| ----------- | -------------- | ---------------- |
| `339841295` | Animal Kingdom | Get Away With It |
| `239192466` |                |                  |
| `507978427` |                |                  |
| `68149590`  |                |                  |
| `454645781` |                |                  |
| `720728071` |                |                  |
| `172332623` |                |                  |
| `544450863` |                |                  |
| `438876933` |                |                  |
| `731459862` |                |                  |
| `547178483` |                |                  |
| `958160664` |                |                  |
| `133821945` |                |                  |
| `743732539` |                |                  |
| `83592991`  |                |                  |
| `841758129` |                |                  |
| `345237391` |                |                  |
| `168643679` |                |                  |
| `309950624` |                |                  |
| `193527120` |                |                  |

### Sagittarius FM

| ID           | Artist            | Title                 |
| ------------ | ----------------- | --------------------- |
| `887359693`  | Climax Blues Band | Couldn’t Get it Right |
| `924397233`  |                   |                       |
| `516605372`  |                   |                       |
| `172131873`  |                   |                       |
| `447148372`  |                   |                       |
| `679069746`  |                   |                       |
| `299179807`  |                   |                       |
| `1044093422` |                   |                       |
| `220766727`  |                   |                       |
| `423003502`  |                   |                       |
| `555831153`  |                   |                       |
| `479677429`  |                   |                       |
| `37782963`   |                   |                       |
| `615133034`  |                   |

### Roadrunner Records

| ID           | Artist | Title |
| ------------ | ------ | ----- |
| `347608615`  |        |       |
| `1033051382` |        |       |
| `713642521`  |        |       |
| `643230639`  |        |       |
| `101167119`  |        |       |
| `425185674`  |        |       |
| `1014785108` |        |       |
| `427737063`  |        |       |
| `367770460`  |        |       |
| `101160181`  |        |       |
| `75278522`   |        |       |
| `415119689`  |        |       |
| `625245472`  |        |       |
| `621142893`  |        |       |
| `583322465`  |        |       |

### Real FM

| ID           | Artist | Title |
| ------------ | ------ | ----- |
| `1031166622` |        |       |
| `501528023`  |        |       |
| `513380233`  |        |       |
| `531212375`  |        |       |
| `133914539`  |        |       |
| `18038500`   |        |       |
| `1068598042` |        |       |
| `285526788`  |        |       |
| `217355793`  |        |       |
| `26526930`   |        |       |
| `489793766`  |        |       |
| `783175834`  |        |       |
| `588421527`  |        |       |
| `51782487`   |        |       |
| `298683518`  |        |       |
| `397484266`  |        |       |
| `107505628`  |        |       |
