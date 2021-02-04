# contekan `awk`

Contekan `awk` untuk pemula dari materi belajar https://en.wikibooks.org/wiki/An_Awk_Primer

## Dokumentasi

```bash
$ man awk
```

coins.txt
```
gold     1    1986  USA                 American Eagle
gold     1    1908  Austria-Hungary     Franz Josef 100 Korona
silver  10    1981  USA                 ingot
gold     1    1984  Switzerland         ingot
gold     1    1979  RSA                 Krugerrand
gold     0.5  1981  RSA                 Krugerrand
gold     0.1  1986  PRC                 Panda
silver   1    1986  USA                 Liberty dollar
gold     0.25 1986  USA                 Liberty 5-dollar piece
silver   0.5  1986  USA                 Liberty 50-cent piece
silver   1    1987  USA                 Constitution dollar
gold     0.25 1987  USA                 Constitution 5-dollar piece
gold     1    1988  Canada              Maple Leaf
hello
```

## Pola

Untuk menggunakan awk ada pola/ _pattern_ yang harus kita ketahui

```bash
# simpel
$ awk 'search1 { action1 }'

# full
$ awk 'BEGIN { initialisasi } search1 { action1 } search2 { action2 } ... END { final_action }'
```

## Cara penggunaan

```bash
# Untuk menampilkan seluruh data yang ada kata `gold`
$ awk /silver/ coins.txt
silver  10    1981  USA                 ingot
silver   1    1986  USA                 Liberty dollar
silver   0.5  1986  USA                 Liberty 50-cent piece
silver   1    1987  USA                 Constitution dollar

# Untuk menampilkan tahun dan negara saja
$ awk '/silver/ { print $3, $4 }' coins.txt
1981 USA
1986 USA
1986 USA
1987 USA

# $0 = Print baris
# $1 = Print field yang pertama
# $2 = Print field yang kedua
# dan seterusnya
```

### Variable

Kita bisa juga membuat variable di awk

```
$ awk 'BEGIN { count = 0 } { count += 1} END { print count }' coins.txt
13
```

### Kondisional

#### if-else

```bash
$ awk '{ if($1 == "silver") { print $0 } }' coins.txt
silver  10    1981  USA                 ingot
silver   1    1986  USA                 Liberty dollar
silver   0.5  1986  USA                 Liberty 50-cent piece
silver   1    1987  USA                 Constitution dollar

$ awk '{ if($1 == "silver") { print $1 } else if($1 == "gold") { print $3 } else { print "Tidak di handle" } }' coins.txt
1986
1908
silver
1984
1979
1981
1986
silver
1986
silver
silver
1987
1988
```
