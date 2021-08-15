|                | tealeg | excelize | My API |
|----------------|:------:|:--------:|:----:|
| RandomGet      |   1!   |     3    |   2  |
| RandomSet      |   1!   |     3    |   2  |
| RandomSetStyle |   1!   |     3    |   2  |
| ReadBigFile    |    2   |     3    |   1  |
| UpdateBigFile  |    2!! |     3    |   1  |
| ReadHugeFile   |    -   |     -    |   1  |
| UpdateHugeFile |    -   |     -    |   1  |

* ! - does not mutate information directly, so faster get/set, but slower read/write files - sometimes it can take forever to open file.
* !! - corrupted file after saving, lost styles/formatting