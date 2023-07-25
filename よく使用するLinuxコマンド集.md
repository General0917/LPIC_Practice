---
title: よく使うLinuxコマンド
tags: Linux Linuxコマンド 初心者向け
author: arene-calix
slide: false
---
Linuxでよく使うコマンドやショートカットキーをまとめました。

対象読者は、ターミナルに苦手意識のあるエンジニアです。
初心者向けに書きましたが、そこそこテクい書き方もでてきます。
普段使いしている方にも案外発見があるかも? しれません。

## この記事のスタンス

* わかりやすさ >> 正確さ
* よく使うものを紹介 >> 網羅的に紹介
* デフォルトで使える >> 別途インストールが必要
* こんなときに使うよ、みたいな主観的なコメントを大事にしたい
* ディレクトリって何? 標準出力って何? とかまでは説明しない

## 一覧

|大項目|紹介するもの|
|---|---|
|はじめに〜コマンド入力をラクにする|Tab補完, bashショートカット|
|ディレクトリの確認|pwd, ls, tree|
|階層移動とファイル操作|cd, mkdir, touch, mv, cp, rm, tar, scp|
|テキスト処理(フィルタコマンド)|cat, wc, head, tail, sort, uniq, grep, sed, awk, xargs, less, >, >>(リダイレクト)|
|インストールまわり|apt, yum, sudo, su, echo, which, wheris, source, ., chmod, chown, systemctl|
|OSまわり|date, df, du, free, top, ps, kill, pkill, pgrep, netstat|
|その他|find, history, diff, jobs, fg, bg, &(バックグラウンド実行), &&, \$()(コマンド置換), <()(プロセス置換), \$?, for|
|書かないことにしたやつ|vi, make, curl, rsync, ssh-keygen, npm, git|
|おまけ|nice, sl|

## はじめに〜コマンド入力をラクにする

コマンド全てを手入力するのは、しんどいです。
Tab補完やbashのショートカットなどを使ってラクをしましょう。
-> <https://www.atmarkit.co.jp/ait/articles/1603/09/news019.html>

最低限、コマンド名やファイル名はTabキーを打って補完したほうがよいと思います。
bashのショートカットはいろいろありますが、個人的には

* `Ctrl+A, Ctrl+E`(行頭、行末へカーソル移動)
* `Ctrl+U, Ctrl+K`(行頭まで削除、行末まで削除)  # パスワードを打ち間違えたときに便利
* `Ctrl+W`(1単語分コマンドを削除)
* `Ctrl+L`(ターミナルの表示をクリア)
* `↑, ↓, Ctrl+R`(コマンド履歴の参照)
* `Ctrl+C`(実行中のコマンドを中止)

あたりの使用頻度が高いです。
これらはreadlineというライブラリが提供するショートカットです。  # Ctrl+Cは別かも
bashに限らず多くのコマンドラインツールで使えるため、覚えておくと何かと便利です。
(例えばpythonやmysqlでも使えます)

<br>

## ディレクトリの確認

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|pwd|今いるディレクトリの絶対パスを表示|print working directory|
|ls|ファイルやディレクトリを表示|list|
|tree|ディレクトリ構造を表示|directory tree|

### ・ **pwd**

```bash
# pwd: 今いるフォルダの絶対パスを表示
arene@~/qiita $ pwd
/home/arene/qiita
```

<br>

### ・ **ls**

```bash
# ls: 現在のフォルダにあるファイルやディレクトリを表示
$ ls
dir1  file1.txt  file2.txt  file3.txt

# ls -al:
#  * -a: 隠しファイルも表示(由来: all)
#  * -l: 詳細な情報を表示(由来: list?)
#  * とにかく全部みたいときに使う
#  * 大きいファイルがある場合はls -alhにすると見やすい
#    -h: M(メガ)、G(ギガ)などを付けてサイズを見やすくする(由来:human readable)
#  * パーミッション(左端に出るrとかwとか)や所有者(arene areneのところ)はインストールでエラーが起きた時によく見る
#    -> 変えたい場合はそれぞれchmod, chownを使う
#  * llでls -alやls -lと同じ動作になる環境もよく見かけます (キーワードはalias(エイリアス)。下記リンク参照)
$ ls -al
total 0
drwxr-xr-x 1 arene arene 4096 Nov 10 18:07 .
drwxrwxrwx 1 arene arene 4096 Nov 10 18:04 ..
-rw-r--r-- 1 arene arene    0 Nov 10 18:07 .hidden_file1.txt
-rw-r--r-- 1 arene arene    0 Nov 10 18:14 dir1
-rw-r--r-- 1 arene arene    4 Nov 10 18:04 file1.txt
-rw-r--r-- 1 arene arene    0 Nov 10 18:02 file2.txt
-rw-r--r-- 1 arene arene    0 Nov 10 18:02 file3.txt

# ls -ltr: 新しいファイルが一番下に来るように表示
#  * -t: タイムスタンプ順で表示(由来: time)
#  * -r: 逆順で表示(由来: reverse)
#  * 最新のログファイルを探す目的でよく使う
#  * 最新が一番下に来るので、ファイル数が多くても見切れない
#  * 逆に、一番古いファイルを見たい場合はls -lt
$ ls -ltr
total 0
-rw-r--r-- 1 arene arene  123 Oct 10 02:30 20191010.log
-rw-r--r-- 1 arene arene  123 Oct 11 02:30 20191011.log
-rw-r--r-- 1 arene arene  123 Oct 12 02:30 20191012.log
-rw-r--r-- 1 arene arene  123 Oct 13 02:30 20191013.log
-rw-r--r-- 1 arene arene  123 Oct 14 02:30 20191014.log
-rw-r--r-- 1 arene arene  123 Oct 15 02:30 20191015.log
-rw-r--r-- 1 arene arene  123 Oct 16 02:30 20191016.log
-rw-r--r-- 1 arene arene  123 Oct 17 02:30 20191017.log
-rw-r--r-- 1 arene arene  123 Oct 18 02:30 20191018.log
-rw-r--r-- 1 arene arene  123 Oct 19 02:30 20191019.log  ←最新ファイルが一番下に来る
$
```

余談: [世の中のエンジニアのalias設定](https://qiita.com/reireias/items/d906ab086c3bc4c22147)

<br>

### ・ **tree**  ※別途インストールが必要

```bash
# tree: ディレクトリ構造を表示
#  * ls -Rでも同様の情報が分かるけど、見づらい。
#  * sudo apt install treeやyum install treeでインストールできます
#  * 今後の説明でよく使うため紹介
$ tree
.
|-- dir1
|   |-- dir11
|   |   |-- file111.txt
|   |   `-- file112.txt
|   |-- file11.txt
|   |-- file12.txt
|   `-- file13.txt
|-- file1.txt
|-- file2.txt
`-- file3.txt
```

<br>

## 階層移動とファイル操作

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|cd|階層移動(カレントディレクトリの変更)|change directory|
|mkdir|ディレクトリの作成|make directory|
|touch|ファイルの作成、タイムスタンプ更新|??|
|mv|ファイルやディレクトリの移動|move|
|cp|ファイルやディレクトリの移動|copy|
|rm|ファイルの削除|remove|
|tar|ファイルの圧縮、展開(tar形式)|tape archives(←はじめて知った)|
|scp|ローカルマシン（コマンドを実行する側）とリモートマシン間でのコピー、リモートマシン同士でのコピー|??|

<br>

### ・ **cd**

```bash
# cd path: pathへ移動
arene@~/qiita $ ls
dir1  file1.txt  file2.txt  file3.txt
arene@~/qiita $ cd dir1/
arene@~/qiita/dir1 $ pwd
/home/arene/qiita/dir1

# cd: ログインユーザのホームディレクトリに移動
arene@~/qiita/dir1 $ cd
arene@~ $ pwd
/home/arene

# cd -: 直前にいたディレクトリへ移動
#  * 階層の離れた2つのディレクトリを行き来するときに便利
#  * pushd, popdコマンドでも似たようなことができるけど、こっちの方が好き
arene@~/qiita/dir1 $ pwd
/home/arene/qiita/dir1
arene@~/qiita/dir1 $ cd
arene@~ $ pwd
/home/arene

arene@~ $ cd -
arene@~/qiita/dir1 $ pwd
/home/arene/qiita/dir1

arene@~/qiita/dir1 $ cd -
arene@~ $ pwd
/home/arene

# cd ~/path: ログインユーザのホームディレクトリ以下のpathに移動
#  * ~はログインユーザのホームディレクトリに読み替えられる(=チルダ展開)
#  * ~xxxだとxxxユーザのホームディレクトリに読み替えられる
arene@~/qiita/dir1 $ cd ~/bin/
arene@~/bin $ pwd
/home/arene/bin
```

<br>

### ・ **mkdir**

```bash
# mkdir directory_name: ディレクトリを作成(1階層のみ)
# mkdir -p path/to/directory: 深い階層のディレクトリを一気に作成
$ ls  # 最初は何もない状態

$ mkdir dir1  # ディレクトリを作成
$ ls
dir1

$ mkdir aaa/bbb/ccc # 深い階層を一気に作るには-pオプションが必要
mkdir: cannot create directory ‘aaa/bbb/ccc’: No such file or directory
$ ls
dir1
$ mkdir -p aaa/bbb/ccc
$ tree
.
|-- aaa
|   `-- bbb
|       `-- ccc
`-- dir1
```

<br>

### ・ **touch**

```bash
# touch file_name: ファイルを新規作成 or タイムスタイプを現在時刻に更新
#  * 本来はファイルのタイムスタンプを更新するコマンドなんだけど、
#    指定したファイルが存在しない場合は新規作成するため、専らファイル作成コマンドと化している印象
$ touch file1  # 新規作成
$ ls -l
-rw-r--r-- 1 arene arene 0 Nov 10 10:10 file1
$ touch file1  # 5分後に再実行 -> タイムスタンプが更新される
$ ls -l
-rw-r--r-- 1 arene arene 0 Nov 10 10:15 file1

# touch --date="YYYYMMDD hh:mm:ss" file_name: ファイルのタイムスタンプを任意の時刻に更新
#  * 時刻がらみの動作確認をする目的でまれによく使う
#  * 関連コマンド
#    date -s "YYYYMMDD hh:mm:ss": OS時刻変更 (-sはsetの意味)
#    (--dateオプションの由来と思われる)
$ touch --date "20101010 10:10:10" file1
$ ls -l
total 0
-rw-r--r-- 1 arene arene 0 Oct 10  2010 file1

# 応用:
# ブレース展開と組み合わせると大量の試験ファイルを簡単に作れます
# ブレース展開(連番ver): {num1..num2}がnum1～num2までの連番に展開されるbashの機能
#                       (列挙verは次項(mv)を参照)
$ touch file{1..3}.txt # -> touch file1.txt file2.txt file3.txt と展開される
$ ls
file1  file2.txt  file3.txt
```

<br>

### ・ **mv**

```bash
# mv source/path destination/path: ファイルやディレクトリの移動
# mv filename_before filename_after: リネーム
#  * OS管理上リネームとファイル移動は大体同じ
#  * 確認がうるさいときは-fオプションをつける(f: force / 誤操作注意!)
$ tree
.
|-- dir1
|-- dir2
`-- file1

$ mv file1 dir1  # 移動
$ tree
.
|-- dir1
|   `-- file1
`-- dir2

$ mv dir1/file1 dir1/file1.txt  # リネーム
$ tree
.
|-- dir1
|   `-- file1.txt
`-- dir2

# 応用:
# ディレクトリ構成によっては、ブレース展開と組み合わせると簡潔に記述できます
# ブレース展開(列挙ver): cmd {aaa,bbb,ccc}がcmd aaa bbb cccに展開されるbashの機能
#                        ※スペースを入れて{aaa, bbb, ccc}としないよう注意
$ tree
.
|-- dir1
|   `-- file1.txt
`-- dir2

$ mv dir{1,2}/file1.txt # mv dir1/file1.txt dir2/file1.txtに展開される
$ tree
.
|-- dir1
`-- dir2
    `-- file1.txt

# 応用2:
# mvした後にcd !$とするとファイルの移動先へスムーズに移動できます
#  * bashでは!$を使うと、直前に実行したコマンドの末尾の引数に展開される
#  * !$は!-1$のエイリアス。!-2$だと2個前に実行したコマンドの末尾の引数に展開される
#  * !はコマンド履歴の展開を表す。$が末尾引数なのは正規表現と同じイメージ
#  * !系のコマンドは黒魔術めいててあまり好きになれないけど、
#    mv -> cd !$だけはめっちゃ使うし便利なので覚えました。(cpも同様)
#  * mv -> ls -> cd !-2$ もよくやります
arene@~/qiita $ ls
dir1  file1
arene@~/qiita $ mv file1 dir1/a/b/c/d/e/f/g/
arene@~/qiita $ cd !$   # !$は直前のコマンドの末尾の引数=a/b/c/d/e/f/g/
cd dir1/a/b/c/d/e/f/g/  # !系のコマンドを使うと展開結果が標準出力に出てくる
arene@~/qiita/dir1/a/b/c/d/e/f/g $ ls
file1
```

履歴展開の詳しい情報: <https://mseeeen.msen.jp/bash-history-expansion/>

<br>

### ・ **cp**

```bash
# cp -r source/path destination/path: ファイルやディレクトリのコピー
#  * -r: ディレクトリ以下を再帰的にコピー(由来: recursive)
#  * -f: 確認無しで強制コピー(由来: force)  <- 誤操作注意!
#  * -p: コピー前後でパーミッションを保持(由来: permission)
#  * 私は常にrを付ける派(別にあっても悪さしないし、ディレクトリかファイルかでオプション使い分けるのがめんどくさい)
#  * fは誤操作を水際で止めてくれるかもなので、基本付けない派
#  * パーミッションが大事な場面ではpの有無に気を配る(普段は手癖で付けちゃう派)
#  * 類似コマンドscpもよく使います。
#    ネットワーク越しにファイルやディレクトリをコピーできます(書式は代替同じ)
$ tree
.
|-- dir1
|   `-- file1
`-- dir2

$ cp dir1/file1 dir2  # ファイルを別フォルダにコピー
$ tree
.
|-- dir1
|   `-- file1
`-- dir2
    `-- file1

$ cp dir1/file1 dir2/file2  # リネームしつつファイルを別フォルダにコピー
$ tree
.
|-- dir1
|   `-- file1
`-- dir2
    |-- file1
    `-- file2

# 応用:
# ブレース展開と組み合わせるとバックアップファイル作成などを簡潔に記述できます
# ブレース展開(列挙ver): cmd {aaa,bbb,ccc}がcmd aaa bbb cccに展開されるbashの機能
$ ls
important_file
$ cp important_file{,.bak}  # cp important_file important_file.bakに展開される
$ ls
important_file  important_file.bak
```

<br>

### ・ **rm**

```bash
# rm -f file_name: ファイルを削除
# rm -rf directory_name: ディレクトリを削除
#  * -f: 確認無しで強制コピー(由来: force)
#  * -r: ディレクトリ以下を再帰的に削除(由来: recursive)
#  * Windowsと違って消したら戻せないので、よくよく気を付けよう
#  * うっかり「rm -rf /」を実行するとOSを含めシステムが全削除されます
#    慌ててCtrl+Cで止めたが、基本的なコマンドの実行ファイルが削除されていてまともに作業できない!
#    ...みたいなことになります。(昔、隣の席の人がやらかしてた)
#  * シェルスクリプトで rm -rf /${DIR_PATH} と書いたが${DIR_PATH}が空文字で
#    「rm -rf /」となるのはありがちかも。
#    (こういう事故が起きないよう、シェルスクリプトではset -uをつけるのがベター(下記リンク参照))
$ ls  # 初期状態
dir1  dir2  dir3  file1.txt  file2.txt  file3.txt

$ rm -f file1.txt  # ファイル名を指定して削除
$ ls
dir1  dir2  dir3  file2.txt  file3.txt

$ rm -f *.txt  # ワイルドカード(*)を使ってtxtファイルを全削除
$ ls
dir1  dir2  dir3

$ rm -f dir1  # rオプションがないとディレクトリは削除できない
rm: cannot remove 'dir1': Is a directory

$ rm -rf dir1  # ディレクトリ名を指定して削除
$ ls
dir2  dir3

$ rm -rf dir*  # ワイルドカード(*)を使って一括削除
$ ls
$ 
```

余談: [シェルスクリプトを書くときはset -euしておく](https://qiita.com/youcune/items/fcfb4ad3d7c1edf9dc96)

<br>

### ・ **tar**

```bash
# tar -czvf xxx.tgz file1 file2 dir1 : 圧縮(file1 file2 dir1をアーカイブした圧縮ファイルxxx.tgzを作成)
# tar -tzvf xxx.tgz: 圧縮ファイルに含まれるファイル名を表示(=展開のテスト)
# tar -xzvf xxx.tgz: 展開
#  * tarのオプションはハイパーややこしい
#    ...んだけど、普通に作業する分には上の3つで事足りると思う
#  * c(create), t(test), x(extract) + zvfと覚える
#  * なおアーカイブと圧縮は別の事象。
#    複数ファイルを一つにまとめるのがアーカイブ。ファイル容量を削減するのが圧縮。
$ ls  # 初期状態
dir1  dir2  file1.txt  file2.txt

$ tar czvf something.tgz dir* file*  # 圧縮
dir1/
dir2/
file1.txt
file2.txt
$ ls
dir1  dir2  file1.txt  file2.txt  something.tgz

$ rm -rf dir* file*  # 一旦元ファイルを削除
$ ls
something.tgz

$ tar tzvf something.tgz  # 中身だけ見る
drwxr-xr-x arene/arene       0 2019-11-12 00:31 dir1/
drwxr-xr-x arene/arene       0 2019-11-12 00:30 dir2/
-rw-r--r-- arene/arene       0 2019-11-12 01:00 file1.txt
-rw-r--r-- arene/arene       0 2019-11-12 01:00 file2.txt
$ ls
something.tgz

$ tar xzvf something.tgz  # 展開
dir1/
dir1/file1.txt
dir2/
file1.txt
file2.txt
$ ls
dir1  dir2  file1.txt  file2.txt  something.tgz

# 余談:
# tarは圧縮、展開できているにもかかわらず0以外の終了ステータスを返すことが結構あります。
# (ex. ファイルのタイムスタンプが未来時刻のとき)
# シェルスクリプトとかでtarの終了ステータスをとったり、set -eするときは要注意
```

<br>

### ・ **scp**

```bash
# 下記のように、リモートマシンに接続する際、機器のパスワードを求められます。
# この場合、パスワードを入力し、Enterキーをクリックすることで先に進めます。

# [root@localhost ~]# scpコマンド
# root@192.168.56.105's password:

# 下記のように、リモートマシンに接続する際、フィンガープリントの登録を求められます。
# この場合、「yes」と入力し、Enterキーをクリックすることで先に進めます。

# [root@localhost ~]# scpコマンド
# The authenticity of host '192.168.56.105 (192.168.56.105)' can't be established.
# ECDSA key fingerprint is SHA256:a6natbKJ32I7Cxyz/MRoitKc5Z/nyIwLUwfaEOZeDtE.
# Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
# Warning: Permanently added '192.168.56.105' (ECDSA) to the list of known hosts.

# scpコマンドの基本的な書式は以下になる。
scp [オプション] コピー元 コピー先

|  オプション  |  説明  |
| ---- | ---- |
|  -c  |  データの暗号化方式を指定する  |
|  -C  |  データを圧縮してコピーする  |
|  -F SSH設定ファイル  |  SSHの設定ファイルを指定する  |
|  -l 帯域  |  通信帯域を制限する（Kbit/s単位）  |
|  -o  |  SSHのオプションを使用する  |
|  -p  |  コピー元のタイムスタンプやパーミッションを保持する  |
|  -P ポート番号  |  接続に使用するポートを指定する  |
|  -q  |  実行状況などの情報を表示しない  |
|  -r  |  ディレクトリを再帰的にコピーする  |
|  -i 秘密鍵ファイル  |  ssh接続に使用する鍵ファイルを指定する  |
|  -v  |  詳細な実行状況を表示する  |

# ローカルからリモートへコピー
scp コピー元 [接続先ユーザ]@ホスト名(IPアドレス):コピー先

# [root@localhost workspace]# scp test_file_01.txt root@192.168.56.105:/home/workspace/
# root@192.168.56.105's password:
# test_file_01.txt                              100%    0     0.0KB/s   00:00

# リモートからローカルへコピー
scp [接続先ユーザ]@ホスト名(IPアドレス):コピー元 コピー先

# [root@localhost workspace]# scp root@192.168.56.105:/home/workspace/test_file_01.txt /home/workspace/
# root@192.168.56.105's password:
# test_file_01.txt                              100%    0     0.0KB/s   00:00

# リモート同士でコピー
scp [接続先ユーザ]@ホスト名(IPアドレス):コピー元 [接続先ユーザ]@ホスト名(IPアドレス):コピー先

# ここでは、例としてリモートマシン上で「test_file_01.txt」を「test_file_02.txt」にリネームしつつコピーを行っています。
# また、リモートマシン同士では、コピー元と先の２回パスワードを求められるようです。
# [root@localhost workspace]# scp root@192.168.56.105:/home/workspace/test_file_01.txt root@192.168.56.105:/home/workspace/test_file_02.txt
# root@192.168.56.105's password:
# root@192.168.56.105's password:
# test_file_01.txt                              100%    0     0.0KB/s   00:00
# Connection to 192.168.56.105 closed.

# -rオプションでのディレクトリコピー
scp -r コピー元のパス コピー先のパス

# 下記は、ローカルマシンの「test_dir_01」ディレクトリをディレクトリ配下のファイルも含めて、リモートマシンにコピーする例です。
# [root@localhost workspace]# scp -r test_dir_01/ root@192.168.56.105:/home/workspace/
# root@192.168.56.105's password:
# test_file_01.txt                              100%    0     0.0KB/s   00:00

# -q オプションで実行状況を表示せずコピー
scp -q コピー元のパス コピー先のパス

# 下記のように実行状況が表示されずに処理が行われます。
# [root@localhost workspace]# scp -q test_file_01.txt root@192.168.56.105:/home/wo
# rkspace/
# root@192.168.56.105's password:

# 参考URL
https://syutaku.blog/linux-cmd-scp/

```

<br>

## テキスト処理(フィルタコマンド)

Linuxの醍醐味、テキスト処理。
コマンドの説明に移る前に、テキスト処理の概要を説明します。

* bashにはパイプ(|)と呼ばれる構文がある
* `コマンドA | コマンドB | コマンドC`と書いたら、次の動きとなる
  * コマンドAは処理結果(文字列A)を標準出力に出す
  * コマンドBは標準入力から文字列Aを受け取り、処理結果(文字列B)を標準出力に出す
  * コマンドCは標準入力から文字列Bを受け取り、処理結果(文字列C)を標準出力に出す
* コマンドA、B、Cのような、標準入力から文字列をinputして
  標準出力へ文字列をoutputするコマンドを「フィルタコマンド」という
* フィルタコマンドを使って文字列を加工することを「テキスト処理」という
* また、パイプ等を駆使してコマンド1行でいろんな処理をすることを俗に「ワンライナー」という

・・・というわけで、Let's テキスト処理!

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|cat|ファイル内容を結合して出力|concatenate(結合)|
|wc|単語数、行数を数える|word count|
|head|先頭からn行を出力|head(先頭)|
|tail|末尾からn行を出力|tail(末尾)|
|sort|行単位でソート|sort|
|uniq|重複を排除|unique|
|grep|テキスト検索|global regular expression print|
|sed|文字列置換|stream editor|
|awk|unix伝統のプログラム言語|開発者3名の頭文字|
|xargs|標準入力をコマンドライン引数に変換|execute arguments?|
|less|標準入力をエディタっぽく表示|less is more(moreコマンドの上位互換)|
|>, >>(リダイレクト)|標準入力をファイルに書き出す||

※less, >, >>はフィルタコマンドじゃない気がしますが、テキスト処理の終端でよく使うためここで紹介します

### ・ **cat**

```bash
# cat file1: file1の内容を標準出力に出力
# cat file1 file2: file1の内容を標準出力に出力 -> file2の内容を標準出力に出力
#  * 用途1: 数行程度のファイルを、read onlyで見る
#  * 用途2: 複数のログファイルをひとまとめにして確認
#  * フィルタ処理の起点に使うことが多い
$ cat .gitignore  # 単一ファイルをdump
.DS_Store
node_modules
/dist

$ ls
access.log  error1.log  error2.log
$ cat error*.log  # error1.logとerror2.logをまとめて確認
2019/09/14 22:40:33 [emerg] 9723#9723: invalid number of arguments in "root" directive in /etc/nginx/sites-enabled/default:45
2019/09/14 22:42:24 [notice] 9777#9777: signal process started
2019/09/14 22:49:23 [notice] 9975#9975: signal process started
2019/09/14 22:49:23 [error] 9975#9975: open() "/run/nginx.pid" failed (2: No such file or directory)
2019/09/14 22:56:00 [notice] 10309#10309: signal process started
2019/09/14 22:56:10 [notice] 10312#10312: signal process started
2019/09/14 22:56:10 [error] 10312#10312: open() "/run/nginx.pid" failed (2: No such file or directory)
2019/09/14 22:56:22 [notice] 10318#10318: signal process started
2019/09/14 22:56:22 [error] 10318#10318: open() "/run/nginx.pid" failed (2: No such file or directory)
2019/12/07 21:49:50 [notice] 1499#1499: signal process started
2019/12/07 21:49:50 [error] 1499#1499: open() "/run/nginx.pid" failed (2: No such file or directory)
2019/12/07 21:51:19 [emerg] 1777#1777: invalid number of arguments in "root" directive in /etc/nginx/sites-enabled/default:45
```

<br>

### ・ **wc**

```bash
# wc -l file1: file1の行数を数える(行数 + ファイル名が出る)
# cat file1 | wc -l: file1の行数を数える(行数だけ出る)
#  * -l: 行数を数える(由来: line)
#  * -w: 単語数を数える(由来: word)
#  * -c: バイト数を数える(由来: char? / 1文字1バイトだった時代にできたオプションゆえcなのだと思われる)
#  * 正直-lしか使ったことない(バイト数とかlsで事足りるし)
#  * 簡易的に正直-lしか使ったことない(バイト数ならlsで事足りるし)
#  * プログラムから使う場合ファイル名は邪魔なのでcat | wc -lを使うケースが多いと思う
$ ls
access.log  error1.log  error2.log
$ wc -l error1.log  # 行数カウント(1)
7 error1.log
$ wc -l error2.log  # 行数カウント(2)
5 error2.log
$ wc -l error*.log  # ワイルドカード指定で複数ファイルの行数をカウント
   7 error1.log
   5 error2.log
  12 total
$ cat error*.log | wc -l  # error1.logとerror2.logを結合したうえで行数をカウント
12

# 応用: 簡易的にステップ数を数える
$ ls
dist    src    node_modules    package.json    public    tests
$ find src/ -type f | xargs cat | wc -l  # src以下の全ファイルの行数の合計をカウント
271
# 解説
#  * find src/ -type f: src以下のファイルの一覧を出力
#  * cmd1 | cmd2: cmd1が標準出力に出した内容を、標準入力として受け取ってcmd2を実行
#  * cmd1 | xargs cmd2: cmd1が標準出力に出した内容を、「コマンドライン引数」として受け取ってcmd2を実行
#  * find | xargs cat: findで検索したsrc以下のファイル全部を結合して標準出力に出す
#  * find | xargs cat | wc -l: findで検索したsrc以下のファイル全部を結合したものの行数をカウント
#  * findやxargsの項もご参照ください
```

<br>

### ・ **head**

```bash
# head -n 3 file1: file1の先頭3行を出力
#  * 言うほど使わないけど、次に紹介するtailと対になるコマンドなので紹介
#  * 用途1: 重たいファイルの頭の方だけ確認
#  * 用途2: ファイルのヘッダ行だけ取得してプログラムから使う
$ cat file1.txt
1 aaa AAA
2 bbb BBB
3 ccc CCC
4 ddd DDD
5 eee EEE
6 fff FFF
7 ggg GGG
8 hhh HHH
9 iii III
10 jjj JJJ
11 kkk KKK
12 lll LLL
13 mmm MMM

$ head -n 3 file1.txt
1 aaa AAA
2 bbb BBB
3 ccc CCC
```

<br>

### ・ **tail**

```bash
# tail -n 3 file1: file1の末尾3行を出力
#  * 用途1: 重たいログファイルの最後の方だけ見る
#  * tailといば、次のtail -fの使い方がメイン
$ cat file1.txt
1 aaa AAA
2 bbb BBB
3 ccc CCC
4 ddd DDD
5 eee EEE
6 fff FFF
7 ggg GGG
8 hhh HHH
9 iii III
10 jjj JJJ
11 kkk KKK
12 lll LLL
13 mmm MMM

$ tail -n 3 file1.txt
11 kkk KKK
12 lll LLL
13 mmm MMM


# tail -f error.log: error.logを監視して、更新された内容を出力(由来: feed?)
#  * ログ確認で大活躍 (末尾出力なので最新のログが出てくる)
#  * tail -fしてファイルの更新を監視することを俗に"tailさせる"と言う(別々の現場で聞いたことがある)
#  * ログファイルをtailさせて網を張る (適宜grepで絞り込みもしておく)
#    -> 不具合を再現させる
#    -> 不具合発生時時に出たログをピンポイントで抽出する
#    という使い方。
$ tail -f error.log  # ログ監視
-> Ctrl+Cで止めるまで、error.logの更新を待ち続ける
-> error.logが更新されたら、更新内容をそのまま出力する

$ tail -f error.log | grep 500  # 500を含むログだけ監視
-> Ctrl+Cで止めるまで、error.logの更新を待ち続ける
-> error.logに500を含むログが更新されたら、出力する
```

<br>

### ・ **sort, uniq**

```bash
# sort file1: file1を行単位でソート
# uniq file1: file1の重複業を削除
# cat file1 | sort | uniq: file1をソートして、重複業を排除
#  * sortとuniqはワンセット的なところがあるのでまとめて紹介
#  * sort -> uniqの順番が大切(下記例を参照)
#  * 元のファイルは変わらない(フィルタコマンド共通の性質)
#
#  * sortは-rで逆順ソート、-Rでランダムソート、みたいに結構オプションが多彩
#  * ls -lの実行結果をファイルサイズ順でsortする、みたいに
#    特定の列にだけ着目してソートすることも可能(この場合だとls -lSでいける)
#  * けど、正直覚えてない
#  * sortに限らず、ややこしいことをするなら普通のプログラミング言語を使った方がいい
#    (手元のcsvファイルからちょっと重複行を排除したい、みたいなイージーケースで使うくらいでいいと思う)
$ cat not_sorted_and_not_unique.txt
1 aaa AAA
3 ccc CCC
2 bbb BBB
3 ccc CCC
2 bbb BBB
1 aaa AAA
3 ccc CCC

$ cat not_sorted_and_not_unique.txt | sort
1 aaa AAA
1 aaa AAA
2 bbb BBB
2 bbb BBB
3 ccc CCC
3 ccc CCC
3 ccc CCC

$ cat not_sorted_and_not_unique.txt | sort | uniq
1 aaa AAA
2 bbb BBB
3 ccc CCC

$ cat not_sorted_and_not_unique.txt | uniq | sort  # sort -> uniqを逆にすると、期待結果が得られない
1 aaa AAA
1 aaa AAA
2 bbb BBB
2 bbb BBB
3 ccc CCC
3 ccc CCC
3 ccc CCC

# 小ネタ: 乱数生成
$ echo {1..65535} | sed 's/ /\n/g' | sort -R | head -n 1
11828
# 解説
#  * echo {1..65535}: "1 2 3 4 5 (中略) 65535"を生成(キーワード: ブレース展開(上述))
#  * sed 's/ /\n/g': スペースを改行に置換
#  * sort -R：行単位でランダムソート
#  * head -n 1：先頭1行だけ表示
#  * 1秒ほどかかるし、実用性は皆無
#  * sort -Rを初めて知って、ちょっと嬉しくなって書いてみました
#  * ただ、こういう風にコマンドをつなげて、工夫次第でいろんなことができるのがフィルタコマンドの楽しさだと思います
#    (その魅力に取りつかれたのが、いわゆる"シェル芸人")
```

もっと知りたい方へ: [sortコマンド、基本と応用とワナ](https://qiita.com/richmikan@github/items/cc4494359b1ac2f72311)

<br>

### ・ **grep**

```bash
# grep ERROR *.log: 拡張子がlogのファイルから、ERRORを含む行だけ抽出
# cat error.log | grep ERROR: error.logからERRORを含む行だけ抽出
# cat error.log | grep -2 ERROR: error.logからERRORを含む行とその前後2行を出力
# cat error.log | grep -e ERROR -e WARN: error.logからERRORまたはWARNを含む行を抽出
# cat error.log | grep ERROR | grep -v 400: error.logからERRORを含む行を抽出して、400を含む行を排除した結果を表示
#  * -e: 複数キーワードをAND条件で指定(由来: ?? たぶん違うけど、個人的にはフランス語のet(=and)だと解釈してる)
#  * -v: キーワードを含む行を排除(由来: verbose??)
#  * テキスト絞り込み、という高需要な用途ゆえワンライナーの中終盤で大活躍
#  * 正規表現も使えます
#  * 個人的にはcat | grep形式しか使わない(フィルタコマンドは全部脳死でcat | cmdする派)
$ cat file1.txt
1 aaa AAA
2 bbb BBB
3 ccc CCC
4 ddd DDD
5 eee EEE
6 fff FFF
7 ggg GGG
8 hhh HHH
9 iii III
10 jjj JJJ
11 kkk KKK
12 lll LLL
13 mmm MMM

$ cat file1.txt | grep -e CCC -e JJJ
3 ccc CCC
10 jjj JJJ

$ cat file1.txt | grep -2 -e CCC -e JJJ
1 aaa AAA
2 bbb BBB
3 ccc CCC
4 ddd DDD
5 eee EEE
--
8 hhh HHH
9 iii III
10 jjj JJJ
11 kkk KKK
12 lll LLL

$ cat file1.txt | grep -2 -e CCC -e JJJ | grep -v -e AAA -e BBB -e KKK -e LLL
3 ccc CCC
4 ddd DDD
5 eee EEE
--
8 hhh HHH
9 iii III
10 jjj JJJ
```

<br>

### ・ **sed**

```bash
# cat file1 | sed 's/BEFORE/AFTER/g': file1中のBEFOREをAFTERに一括置換
#  * s/BEFORE/AFTER/g: BEFOREをAFTERに置換(由来: substituteとglobal?)
#  * s/BEFORE/AFTER/: 1番目に出現したBEFOREをAFTERに置換
#  * viでも:%s/BEFORE/AFTER/gで一括置換できるので覚えとくと便利
#    (使用例: git rebase -i HEAD~5 -> viが開く -> :%s/pick/s/g で直近5コミットをまとめる)
#  * 一括置換だけじゃなく削除や部分置換もできるし、正規表現も使えます
#  * 元のファイルは変わりません(-iオプションを付けて上書きすることも可能)
#  * 私がさっと使えるのは一括置換ぐらいですが、シェル芸人はsedとawkを駆使してるイメージ
$ cat typo.txt  # スペルミスのあるファイル
Hello Wolrd!
Wolrd Wide Web

$ cat typo.txt | sed 's/Wolrd/World/g'  # スペルミスを直す
Hello World!
World Wide Web

$ cat typo.txt | sed 's/Wolrd/World/g' > typo_fixed.txt  # 直した結果を別ファイルに保存
$ cat typo_fixed.txt
Hello World!
World Wide Web
```

もっと知りたい方へ: [sedでこういうときはどう書く?](https://qiita.com/hirohiro77/items/7fe2f68781c41777e507)

<br>

### ・ **awk**

```bash
# cmd1 | awk '{print $5}': cmd1実行結果から、スペース区切りで5列目だけ表示
# cmd1 | awk -F ',' '{print $5}': cmd1実行結果から、カンマ区切りで5列目だけ表示
#  * ワンライナーの王様(私見)
#  * ifもforも変数も使えるし、区分としてはコマンドというよりプログラミング言語
#  * 〇〇区切りのn列目を抽出したい、みたいな要件だとドンピシャな言語
#  * なお、awkの文脈だと"n列目"じゃなくて"nフィールド目"と呼ぶ
#  * 筆者は、書けないけど、雰囲気読みだけできる人
#    (過去遭遇したレガシー案件で、独自書式のTSVをawkでパースしてXMLに変換してるのを見たことがあります)
#    (↑現職じゃないよ!)
#  * まれによく使われるので、教養として読めるようになってもいいかもしれないです
$ ls -l
total 0
drwxr-xr-x 1 arene arene 4096 Feb  4 22:40 abc
drwxr-xr-x 1 arene arene 4096 Feb  4 22:40 def
-rw-r--r-- 1 arene arene  134 Feb  4 22:50 file1.txt

arene@~/qiita/src $ ls -l | awk '{print $5}'  # 5列目だけ表示

4096
4096
134
```

<br>

### ・ **xargs**

```bash
# cmd1 | xargs cmd2: cmd1の実行結果をコマンドライン引数として受け取って、cmd2を実行
#  * cmd1 | cmd2は、cmd1の実行結果を「標準入力」として受け取ってcmd2を実行するのに対し
#    cmd1 | xargs cmd2は、cmd1の実行結果を「コマンドライン引数」として受け取ってcmd2を実行する
#  * 個人的にすごく好きなコマンド (うまく使えると賢くなった気持ちになれる)
#  * 慣れがいるけど、こいつがないとワンライナーできない要件が結構ある印象
#  * findの後に使うことが多い(複数ファイルに対する一括操作)
#  * 応用編のコマンドだし、分かりやすくて実用的な例は浮かばなかったです
$ ls -1
src
test
$ ls -1 | echo  # echoは標準入力を受け付けないので何も出ない

$ ls -1 | xargs echo  # xargsで引数として渡してやると表示できる
src test
$ ls -1 | xargs -n 1 echo  # -n 1をつけると1行ずつ渡すので、1行ずつechoされる
src
test


# 応用: 複数ファイルの一括リネーム
# find -type f dir_name | xargs -I{} mv {} {}.bak: dir_name以下のファイル全部に.bakを付ける
#  * この例だとrenameコマンドを使った方がラクなはず (なんかrenameに馴染めなくて個人的には使わない)
#  * mvやcpなど、受け取った文字列を2回以上使いたい場合は-I{}オプションを使う
$ tree  # 最初の状態
.
|-- src
|   |-- main.js
|   `-- sub.js
`-- test
    |-- main.js
    `-- sub.js

$ find test/ -type f  # findでtest以下のファイルの相対パスを表示
test/main.js
test/sub.js

$ find test/ -type f | xargs -I{} mv {} {}.test
  # 以下の内容に展開される (-I{}で、以降の{}が入力内容に置換されるようになる)
  # mv test/main.js test/main.js.bak
  # mv test/sub.js test/sub.js.bak
$ tree
.
|-- src
|   |-- main.js
|   `-- sub.js
`-- test
    |-- main.js.test
    `-- sub.js.test

$ find test/ -type f | sed 's/js.test/js/g' | xargs -I{} mv {}.test {} # 元に戻す
  # test/main.js.test をsedで test/main.js に置換したうえでmvに渡す
$ tree
.
|-- src
|   |-- main.js
|   `-- sub.js
`-- test
    |-- main.js
    `-- sub.js
```

<br>

### ・ **less**

```bash
# less file1: file1を見る(read only)
# cat file1 | cmd1 | cmd2 | less: file1をいろいろ加工した結果を見る
#  * ターミナルに出力せず、何かを見たいときにとりあえず使うコマンド
#  * 類似コマンドmoreの上位互換 (less is more!)
#  * read onlyなので安全安心 (変える気がないのにviで見るのはやめよう)
#  * viの一部キーバインドが使える
#    gg: 先頭行へ移動
#    G: 最終行へ移動
#    /pattern: patternでファイル内検索
#    q: 閉じる
#  * Fでtail -fと同様のこともできるし、名前に反してけっこう高機能
(例は下記リンク参照)
```

もっと知りたい方へ:
・[あなたはだんだん、ファイルを読むのにlessコマンドを使いたくなる](https://qiita.com/marrontan619/items/95e954972706f32be255)
・[エンジニアなら知っておきたい lessコマンドtips 11選](https://qiita.com/ine1127/items/64b5b6cf52471c3fe59c)

<br>

### ・ **>, >>(リダイレクト)**

```bash
# cmd1 >> file1: cmd1の実行結果をfile1に書き出す(追記)
# cmd1 > file1: cmd1の実行結果をfile1に書き出す(上書き)
$ cat file1.txt  # リダイレクト前
1 aaa AAA
2 bbb BBB
3 ccc CCC

$ echo "4 ddd DDD" >> file1.txt  # リダイレクト(追記)
$ cat file1.txt
1 aaa AAA
2 bbb BBB
3 ccc CCC
4 ddd DDD

$ echo "4 ddd DDD" > file1.txt  # リダイレクト(上書き)
$ cat file1.txt
4 ddd DDD



# echo "echo login!" >> ~/.bashrc: bashrc末尾に設定を追加  ※実際に叩かないでください!
# * .bashrcはbashの設定ファイルです
# * >> でなく > を使うと、上書きされて設定が消えるので要注意
# * 手順書や環境構築自動化スクリプトで、設定ファイルを編集するときにリダイレクトを使います
#   (手動で環境構築するときは>>と>のうっかりミスが怖いため、普通にファイルを開いた方がいいと思う)



# something.sh > log.txt: something.shの実行結果(標準出力)をログ出力
# something.sh > log.txt 2>&1: something.shの実行結果(標準出力+標準エラー出力)をログ出力
# something.sh >/dev/null 2>&1: something.shの実行結果をどこにも出力しないようにする
$ cat something.sh  # 1行目で標準出力、2行目で標準エラー出力にメッセージの出るシェルスクリプト
#!/bin/bash
echo standard output
syntax-error!!!! # standard error

$ ./something.sh > log.txt  # 単にリダイレクトした場合、標準出力しかリダイレクトされない
./something.sh: line 3: syntax-error!!!!: command not found
$ cat log.txt
standard output

$ ./something.sh > log.txt 2>&1  # 2>&1を足すと、両方リダイレクトされる
$ cat log.txt
standard output
./something.sh: line 3: syntax-error!!!!: command not found

$ ./something.sh >/dev/null 2>&1  # 何も出ないようにする(いわゆる「デブヌルに投げる」というやつ)
$
$ ./something.sh 2>&1 >/dev/null  # なお逆にするとうまくいかない
./something.sh: line 3: syntax-error!!!!: command not found
#  解説
#  * 1: 標準出力   2: 標準エラー出力
#  * /dev/null: ゴミ箱的な、奈落的な、OSが用意する特別な空ファイル
#  * > log.txtは1>log.txtと同じで、標準出力をログ出力している
#  * > log.txt 2>&1は、2(標準エラー出力)を1(標準出力)の向き先(=ログファイル)へ向けている
#  * 2>&1 > log.txtでうまくいかないのは↓2つを逐次実行しているから
#    (1) 2>&1: 標準エラー出力を、デフォルトの標準出力(stdout)に切り替える
#    (2) > log.txt2: 標準出力をログファイルへ向ける
#    => 結果、標準エラー出力はデフォルトの標準出力へ、標準出力はログファイルへと出力される
```

もっと知りたい方へ: [いい加減覚えよう。 `command > /dev/null 2>&1`の意味](https://qiita.com/ritukiii/items/b3d91e97b71ecd41d4ea)

<br>

## インストールまわり

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|apt, yum|コマンドのインストール|Advanced Package Tool, [Yellowdog Updater Modified](https://ja.wikipedia.org/wiki/Yellowdog_Updater_Modified)|
|sudo|ルート権限でコマンドを実行|superuser do(substitute user do)|
|su|ユーザ切り替え|substitute user|
|echo|文字列の表示|echo|
|env|環境変数の表示|environment|
|which, whereis|コマンドの場所を探す|which, where is|
|source, .|設定の反映(ファイル内容を現在のシェルで実行)|source|
|chmod|ファイル、ディレクトリのパーミッションを変更|change mode|
|chown|ファイル、ディレクトリの所有者を変更|change owner|
|systemctl|サービスの起動、停止など|system control|

### ・ **apt, yum**

```bash
# apt install git: gitをインストール(Ubuntuなど、Debian系OS)
# yum install git: gitをインストール(CentOSなど、RedHat系OS)
#  * しばしばsudo apt～形式で実行する
#  * 実行して「Permission Denied」「権限が足りません」系のメッセージがでたらとりあえずsudoする(雑)
#    (あるいはchown, chmodで権限を適切に設定する)
$ sudo apt install git
(大量のメッセージが出るが割愛)
```

<br>

### ・ **sudo**

```bash
# sudo cmd1: cmd1をrootユーザとして実行
#  * Ubuntuでは高頻度で使用 (Ubuntuはrootで作業しないでねという思想のため、一応従う)
#  * CentOSでは次のsuでrootユーザに切り替えるため、あまり使わない
#  * ユーザによっては実行できない場合がある
#    (セキュリティのしっかりした環境だとsudoできるユーザを制限している)
$ sudo vi /etc/hosts  # rootしか変更できない設定ファイルを編集
[sudo] password for arene:
$
```

<br>

### ・ **su**

```bash
# su user1: user1に切り替える(環境変数は現在のものを引き継ぐ)
# su - user1: user1に切り替える(現在の環境変数を捨てて、user1デフォルトの環境変数を利用)
# su -: rootユーザーに切り替える(現在の環境変数を捨てて、rootユーザーデフォルトの環境変数を利用)
#  * 私は常にハイフンを付けます
#    (環境変数を引き継ぐことによる思わぬミスを防ぐ意図)
#  * CentOSでsu - oracleとか、su - postgresとかよく使ってた
$ su -
Password:
#
```

<br>

### ・ **echo**

```bash
# echo abc: 文字列abcを出力
# echo $PATH: 環境変数PATHを出力
#  * 用途1: シェルスクリプトで使用方法やエラーメッセージを出力
#  * 用途2: 環境変数の確認
#  * コマンドを実行して「command not found」と出た場合は、大抵PATHが通ってないので、まずはPATHを通そう
$ echo abc
abc
$ echo $LANG
UTF-8
```

補足情報: [PATHを通すとは、環境変数とは](https://qiita.com/fuwamaki/items/3d8af42cf7abee760a81)

<br>

### ・ **env**

```bash
# env | less: 環境変数を確認
#  * envだけでも見れるが、環境変数が多い場合見切れてしまうためlessで確認
#    -> lessを開いた状態で/PATHとすると、"PATH"で検索できます
```

<br>

### ・ **which, whereis**

```bash
# which cmd: cmdの実体が置かれている場所を表示
# whereis cmd: whichのちょっと詳しい版
#  * 正直whichしか使わない
#  * 複数バージョンのnodeをインストールしたんだけど、今動いてるやつの実体はどこにあるの?
#    いらないコマンドを削除したいんだけど、こいつどこにあるの?
#    とか、そういうケースで使う
$ which ls
/bin/ls

$ ls
access.log  error1.log  error2.log  src

$ /bin/ls
access.log  error1.log  error2.log  src
```

<br>

### ・ **source, .**

```bash
# source ~/.bashrc: .bashrcを再読み込み
# . ~/.bashrc: ↑と同じ(.はsourceのエイリアス)
#  * シェルの設定ファイルを変更した後の再読み込みで使うケースが100%(自分調べ)
#  * 一応、シェルスクリプトを実行することもできる
#
#  * sourceは、引数で指定したファイルを「現在のシェル」で実行するコマンド
#  * 普通にコマンドやシェルスクリプトを実行した場合は「新たに生成した別のシェル」で処理を実行している
#    (現在のシェルの変数が汚れないようになっている)
#  * これに対し、sourceでは現在のシェルで処理を行うため、
#    処理中に変更した環境変数やalias設定が、実行終了後も引き継がれる(=設定が反映される)
$ env | grep MY_ENV  # before

$ echo "export MY_ENV=abc" >> ~/.bashrc  # 適当な環境変数を足す
$ env | grep MY_ENV  # まだ反映されていない

$ . ~/.bashrc  # sourceでbashrcを再読み込み
$ env | grep MY_ENV  # ↑で設定した環境変数が反映されている
MY_ENV=abc
```

<br>

### ・ **chmod**

```bash
# chmod 755 *.sh: shファイルに実行権限を付与
# chmod 644 *.js: jsファイルを普通に読み書きできる設定にする
#  * 謎の数字にもちゃんと意味があるんだけど、正直644と755しか使わない
#  * wとかrとか文字でも設定できるけど、私は数字派
#  * プログラムを実行して「Permission denied」と出たときは大抵実行権限がないだけなので、755へ変更すればOK
#
# 一応説明すると
#  * 755のように数字を3つ並べるのは次の3つを指定している
#    [所有者に対する権限][所有グループに対する権限][その他に対する権限]
#  * 数字の意味は次の通り
#     0: 権限なし
#     1: 実行権限
#     2: 書き込み権限
#     4: 読み込み権限
#    (7=1+2+4で全部OK, 6=2+4で読み書きのみ、といった具合)
#  * つまり755は、"所有者はなんでもできる、他の人は読み書きだけできる"という設定
#         644は、"所有者は読み書きできる、他の人は読むことだけできる"という設定
$ ls -l  # before
total 0
-rw-r--r-- 1 arene arene 0 Feb  8 23:26 abc

$ chmod 755 abc  # 実行権限を付与
$ ls -l
total 0
-rwxr-xr-x 1 arene arene 0 Feb  8 23:26 abc

$ chmod 644 abc  # 実行権限をなくす
$ ls -l
total 0
-rw-r--r-- 1 arene arene 0 Feb  8 23:26 abc



# 応用: 一括変更
# find dir1 -type f | grep sh$ | xargs chmod 755: dir1以下のshすべてに実行権限を付与
#  * findでファイルの相対パスを探す -> grepで末尾がshで終わるファイルを探す -> 探したファイルをchmodする
#  * find dir1 -type f -name "*.sh" | xargs chmod 755 でも同じ
```

もっと知りたい方へ: [Linuxの権限確認と変更(chmod)（超初心者向け）](https://qiita.com/shisama/items/5f4c4fa768642aad9e06)

<br>

### ・ **chown**

```bash
# chown user1:group1 file1: file1の所有者を変更(ユーザーをuser1, グループをgroup1にする)
# find dir1 | xargs chown user1:group1: dir1以下の全ファイルの所有者を一括変更
#  * ユーザ一覧の確認はcat /etc/passwd
#    (前半は各種ミドルウェアが足したユーザーなので、大抵は末尾数行を見れば事足りる)
#  * グループ一覧の確認はcat /etc/group (同上)
@@TODO: 具体例(いい感じに複数ユーザいる環境がなかった)
```

<br>

### ・ **systemctl**

```bash
#  * サービス、というのはfirewallやwebサーバみたいな、バックグラウンド実行されるプログラムのことを指す
#    (デーモンともいう)
#  * systemctlはサービスの起動停止などを行うコマンド
#  * これ系のコマンドは環境によって結構違うが、2020年現在、新しめの環境はsystemctl
#  * 古いlinuxだとserviceコマンド、chkconfigコマンドに分かれている
#  * Macにはないっぽい(最近デビューしたところなので知見ゼロ)
#  * 自作プログラムをサービスに登録することも勿論できます

# 起動、停止、現状確認
#  * 単に起動するだけだと、OS再起動したら止まってしまう点に注意
#  * サービス名はTabキーで補完できます
systemctl status service1      # service1の状態を確認 (生きてるか死んでるかを確認)
systemctl start service1       # service1を起動
systemctl stop service1        # service1を停止
systemctl restart service1     # service1を再起動 (停止->起動)

# 自動起動の設定
#  * enabled: OS起動時に自動起動する
#  * disabled: OS起動時に自動起動しない
systemctl list-unit-files      # サービスの一覧+自動起動するかどうかを表示
systemctl enable service1      # service1を自動起動するようにする
systemctl disable service1     # service1を自動起動しないようにする
systemctl is-enabled service1  # service1が自動起動する設定かどうか確認
```

<br>

## OSまわり

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|date|時刻の確認、設定|date|
|df|ディスク空き容量の確認|disk free|
|du|ディレクトリのサイズを確認|disk usuage|
|free|メモリの空き状況を確認|free|
|top|CPUやメモリの使用状況を確認|??|
|ps|プロセス情報の確認|process status|
|kill|PIDを指定してプロセスを停止させる(シグナルを送る)|kill|
|pkill|指定したプロセス名を持つプロセスを一括で停止させる|process kill?|
|pgrep|指定したプロセス名を持つプロセスのPIDを表示|pid grep|
|netstat|ネットワークの状況を見る|network status|

### ・ **date**

```bash
# date: 現在時を表示
# date '+%Y%m%d %H:%M:%S': YYYYMMDD hh:mm:ss形式で現在時を表示
# date -s "YYYYMMDD hh:mm:ss": OS時刻を変更(由来: set)
#  * たまに使い、使うたびに書式どんなんだっけ? となるやつ
#  * date -s "YYYYMMDD hh:mm:ss"(OS時刻変更)とtouch -d "YYYYMMDD hh:mm:ss"(ファイルのタイムスタンプ変更)
#    が同じなので、これだけ覚えるようにして、あとは都度検索するスタイルにしています
$ date
Sun Feb  9 11:00:41 JST 2020

$ date '+%Y%m%d %H:%M:%S'
20200209 11:01:13

$ date -s "20200209 11:02:00"
Sun Feb  9 11:02:00 JST 2020
```

もっと知りたい方へ: [date コマンドの日付指定頻出パターン](https://qiita.com/suzuki-kei/items/cb0a78a655fef37cb59d)

<br>

### ・ **df**

```bash
# df -h: ディスクの使用量/空き容量を単位付きで表示(由来: human readable)
# df: ディスクの使用量/空き容量を表示
#  * 基本は-hで見る
#  * hは値を丸めるため、正確な値を知りたい場合はオプション無しで叩く
#  * ファイルシステムが何か、それぞれどこにマウントされてるかも分かる
#  * ↓はWSLのUbuntuで見てるので、Cドライブがあったり、ちょっと変
@@TODO: 素のUbuntuでの実行結果に差し替える
$ df -h  # Use%が何%使ってるか。Used、Availがどれだけ使っているか、空いているか
Filesystem      Size  Used Avail Use% Mounted on
rootfs          230G  199G   31G  87% /
none            230G  199G   31G  87% /dev
none            230G  199G   31G  87% /run
none            230G  199G   31G  87% /run/lock
none            230G  199G   31G  87% /run/shm
none            230G  199G   31G  87% /run/user
cgroup          230G  199G   31G  87% /sys/fs/cgroup
C:\             230G  199G   31G  87% /mnt/c
E:\             223G  141G   83G  63% /mnt/e

$ df
Filesystem     1K-blocks      Used Available Use% Mounted on
rootfs         240312316 207873316  32439000  87% /
none           240312316 207873316  32439000  87% /dev
none           240312316 207873316  32439000  87% /run
none           240312316 207873316  32439000  87% /run/lock
none           240312316 207873316  32439000  87% /run/shm
none           240312316 207873316  32439000  87% /run/user
cgroup         240312316 207873316  32439000  87% /sys/fs/cgroup
C:\            240312316 207873316  32439000  87% /mnt/c
E:\            233322492 146962124  86360368  63% /mnt/e
```

<br>

### ・ **du**

```bash
# du -h: 各ディレクトリの容量を単位付きで表示(由来: human readable)
# du: 各ディレクトリの容量を表示
#  * lsではディレクトリのサイズを見ることができない
#  * 実サイズを見たいときはduを使う
#  * サブディレクトリが多くて見辛いときは適宜grepしたりlessしたりする
$ ls -lh  # lsだとディレクトリは一律4.0Kで表示され、実サイズは分からない
total 0
drwxr-xr-x 1 arene arene 4.0K Oct 14 08:53 dist
-rw-r--r-- 1 arene arene    0 Jan  1 10:10 file1.txt
drwxr-xr-x 1 arene arene 4.0K Oct 14 09:11 src

$ du -h  # dfで見ると、実サイズが分かる
0       ./dist/css
8.0K    ./dist/img
888K    ./dist/js
908K    ./dist
8.0K    ./src/assets
4.0K    ./src/components
4.0K    ./src/pages
16K     ./src
924K    .
```

<br>

### ・ **free**

```bash
# free -h: メモリ使用状況を単位付きで表示(由来: human readable)
# free: メモリ使用状況を表示
#  * OSによって若干表示内容が異なるらしい(新しい奴はavailableが出る)
#  * 紹介したものの、こう見ればOKという見方に自信がないです。。
#    (freeやavailableがある程度大きければ問題なし、くらいの認識)
#  * メモリが不足してないか判断するための正しい見方をご存じなら、ぜひコメントをください
$ free -h
              total        used        free      shared  buff/cache   available
Mem:           7.9G        6.8G        886M         17M        223M        980M
Swap:           24G        1.1G         22G
$ free
              total        used        free      shared  buff/cache   available
Mem:        8263508     7099428      934728       17720      229352     1030348
Swap:      25165824     1149132    24016692
```

<br>

### ・ **top**

```bash
# top: CPUやメモリの使用状況を確認
#  * デフォルトだとCPU使用率の多いプロセスが上に来る
#  * %CPUがCPU使用率。どのプロセスが高負荷かを確認できる。
#  * 右上のload averageもちょくちょく見る
#    値がCPUのコア数を超えると高負荷状態(デュアルコアなら、2以上が高負荷)
#    cpuのコア数は cat /proc/cpuinfo で確認可能
#  * ただ、4コアでload averageが3だったとしても
#    core1にタスクが集中していて、core1だけ高負荷状態、ということがありえる。
#    ↑のload averageがコア数未満ならOK、というのはあくまで一応の目安
$ top
top - 12:06:17 up 87 days, 11:55,  0 users,  load average: 0.52, 0.58, 0.59
Tasks:  13 total,   1 running,  12 sleeping,   0 stopped,   0 zombie
%Cpu(s): 10.2 us,  8.0 sy,  0.0 ni, 81.7 id,  0.0 wa,  0.1 hi,  0.0 si,  0.0 st
KiB Mem :  8263508 total,  1821072 free,  6213084 used,   229352 buff/cache
KiB Swap: 25165824 total, 23985072 free,  1180752 used.  1916692 avail Mem

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND
 5310 arene     20   0   17620   2052   1516 R   1.0  0.0   0:00.18 top
    1 root      20   0    8896    172    136 S   0.0  0.0   0:00.21 init
   74 root      20   0   19464    504    448 S   0.0  0.0   0:00.01 sshd
 1862 root      20   0   57560    344    312 S   0.0  0.0   0:00.01 nginx
 1863 www-data  20   0   58204   1036    904 S   0.0  0.0   0:00.19 nginx
 1865 www-data  20   0   58204   1036    920 S   0.0  0.0   0:00.07 nginx
 1868 www-data  20   0   58204   1036    904 S   0.0  0.0   0:00.01 nginx
 1869 www-data  20   0   58204    948    856 S   0.0  0.0   0:00.00 nginx
 1920 root      20   0    8904    224    176 S   0.0  0.0   0:00.01 init
 1921 arene     20   0   17332   4032   3896 S   0.0  0.0   0:00.32 bash
 1996 root      20   0   20220   4204   4056 S   0.0  0.1   0:00.17 sshd
 2069 arene     20   0   20488   2092   1956 S   0.0  0.0   0:05.02 sshd
 2070 arene     20   0   18828   5628   5520 S   0.0  0.1   0:11.96 bash
```

もっと知りたい方へ:
・[マルチコア時代のロードアベレージの見方](https://naoya-2.hatenadiary.org/entry/20070518/1179492085)
・[topコマンドの使い方](https://qiita.com/k0kubun/items/7368c323d90f24a00c2f)

<br>

### ・ **ps**

```bash
# ps -ef: 全てのプロセスの詳細な情報を見る(由来: every, full)
#  * 用途1: あるプロセスが生きてるかどうかチェック  (webサーバ起動してる?)
#  * 用途2: あるプロセスのPID(プロセスID)をチェック -> kill ${PID}
#  * 他にもいろいろ見られるんだろうけど、あまり知らないです
#  * 歴史的経緯でオプションが2系統に分かれてて超ややこしいけど、私は-efしか使わない
$ ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0  2019 ?        00:00:00 /init ro
root        74     1  0  2019 ?        00:00:00 /usr/sbin/sshd
root      1862     1  0  2019 ?        00:00:00 nginx:
www-data  1863  1862  0  2019 ?        00:00:00 nginx:
www-data  1865  1862  0  2019 ?        00:00:00 nginx:
www-data  1868  1862  0  2019 ?        00:00:00 nginx:
www-data  1869  1862  0  2019 ?        00:00:00 nginx:
root      1920     1  0 Feb04 tty1     00:00:00 /init ro
arene     1921  1920  0 Feb04 tty1     00:00:00 -bash
root      1996    74  0 Feb04 ?        00:00:00 sshd: arene [priv]
arene     2069  1996  0 Feb04 ?        00:00:04 sshd: arene@pts/0
arene     2070  2069  0 Feb04 pts/0    00:00:11 -bash
arene     5090  2070  0 11:13 pts/0    00:00:00 ps -ef
```

<br>

### ・ **kill**

```bash
# kill 123: プロセスIDが123のプロセスを停止させる(SIGTERMを送る)
# kill -9 123: プロセスIDが123のプロセスを問答無用で殺す(9はSIGKILLのシグナル番号)
# kill -KILL 123: -9と同じ
#  * 名前の通り、プロセスを殺す目的で使うケースが99%
#  * ただ、正確には特定のプロセスに任意のシグナルを送るコマンドで、
#    デフォルトではSIGTERM(由来: terminate)を送っている
#  * シグナルとは、OSが各プロセスへ割り込み処理を命令するために送る信号
#    (例えばCtrl+Cでコマンドを終了するときは、SIGINT(由来: interrupt)が送られている。
#     割り込み処理なので、無限ループするプログラムであっても終了できる)
#  * SIGKILL(9)は、一番強力なシグナルで、こいつを送ると問答無用でプロセスを殺せる
#  * SIGKILLは結構乱暴なので最終手段。ますは無印killを試そう。
#    (SIGKILLだと終了処理さえ許さずぶち切られるため、次回起動時に不都合が起きる場合がまれにある)
$ ps -ef | grep eternal_loop | grep -v grep  # 適当に書いた無限ループするプログラムのPIDを調べる
arene     5500  2070  0 13:00 pts/0    00:00:00 ./eternal_loop

$ kill 5500  # pidを指定してkillする
[1]+  Terminated              ./eternal_loop
$ ps -ef | grep eternal_loop | grep -v grep  # killされたことを確認
$
```

より詳しい情報:
[Linux シグナルの基礎](https://harasou.jp/2017/01/23/linux-signal/)
[SIGNAL Manページ](https://linuxjm.osdn.jp/html/LDP_man-pages/man7/signal.7.html)

<br>

### ・ **pkill**

```bash
# pkill process_name_prefix: process_name_prefixで始まるプロセスすべてを終了させる
# pkill -9 process_name_prefix: process_name_prefixで始まるプロセスすべてを問答無用で終了させる
#  * シグナルについては前項(kill)を参照
#  * ヒットした奴は全部殺されるので、実行前にps -ef | grep process_name_prefix で
#    対象プロセスを確認するのがベター
$ ps -ef | grep eternal_loop | grep -v grep  # 無限ループするプロセスがたくさんある
arene     5558  2070  0 13:13 pts/0    00:00:00 ./eternal_loop
arene     5562  2070  0 13:13 pts/0    00:00:00 ./eternal_loop2
arene     5566  2070  0 13:13 pts/0    00:00:00 ./eternal_loop3
arene     5570  2070  0 13:13 pts/0    00:00:00 ./_bak_eternal_loop

$ pkill eternal_loop  # pkillでまとめて殺す
[1]   Terminated              ./eternal_loop
[2]   Terminated              ./eternal_loop2
[3]-  Terminated              ./eternal_loop3

$ ps -ef | grep eternal_loop | grep -v grep  # 前方一致で一致した3つは殺された
arene     5570  2070  0 13:13 pts/0    00:00:00 ./_bak_eternal_loop
```

<br>

### ・ **pgrep**

```bash
# pgrep process_name_prefix: process_name_prefixで始まるプロセスすべてのPIDを出力
#  * 主にシェルスクリプトやワンライナーでPIDを動的に抽出して使いたいときに使用
#  * $()(コマンド置換)やxargsと組み合わせることが多い
$ ps -ef | grep eternal_loop | grep -v grep  # 無限ループするプロセスがたくさん
arene     5570  2070  0 13:13 pts/0    00:00:00 ./_bak_eternal_loop
arene     5590  2070  0 13:18 pts/0    00:00:00 ./eternal_loop
arene     5594  2070  0 13:18 pts/0    00:00:00 ./eternal_loop2
arene     5598  2070  0 13:18 pts/0    00:00:00 ./eternal_loop3

$ pgrep eternal_loop  # プロセスIDを抽出
5590
5594
5598
$ pgrep eternal_loop | xargs kill  # 抽出したプロセスを殺す
[5]   Terminated              ./eternal_loop
[6]-  Terminated              ./eternal_loop2
[7]+  Terminated              ./eternal_loop3

$ ps -ef | grep eternal_loop | grep -v grep  # 死亡確認
arene     5570  2070  0 13:13 pts/0    00:00:00 ./_bak_eternal_loop
```

<br>

### ・ **netstat**

```bash
# netstat -anp| less: ネットワークの状態を確認
#  * -a: すべての接続を表示(由来: all)
#  * -n: 名前解決しないで素のIPアドレス、ポート番号を表示(由来: number?)
#  * -p: プロセスIDを表示(由来: process)
#  * いろいろ見られるんだろうけど、私は-anpしか使わない
#  * LISTENING, ESTABLISHED, TIME_WAITなど各ポートの状況を確認できて素敵
@@TODO: ubuntuでの実行結果を貼る(手元のWSL環境だと見れなかった)
```

より詳しい情報: [TCP/IP通信の状態を調べる「netstat」コマンドを使いこなす](https://www.atmarkit.co.jp/ait/articles/0207/20/news003.html)

<br>

## その他

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|find|ファイルやディレクトリを探す(パスの出力)|find|
|history|コマンド履歴の参照|history|
|diff|差分確認|difference|
|jobs|実行中のジョブを確認|jobs|
|bg|指定したジョブをバックグラウンドに移す|background|
|fg|指定したジョブをフォアグラウンドに移す|foreground|
|&|バックグラウンド実行||
|&&, |||複数コマンドをつなげる|論理演算子|
|$(), <()|コマンド置換、プロセス置換||
|$?|直前のコマンドの終了ステータスを確認||
|for|ループ処理||

### ・ **find**

```bash
# find dir1 -type f: dir1以下のファイル一覧を表示
# find dir1 -type f -name "*.js": dir1以下のjsファイルの一覧を表示
# find dir1 -type d: dir1以下のディレクトリ一覧を表示
#  * オプションが豊富で、n階層下まで探す、特定の日時より古いファイルだけ探す、
#    特定のパーミッションのファイルだけ探す、とかいろいろできる
#  * だけど忘れちゃうので、さっと書けるのはこのくらい
#  * lsと違ってファイルパスが出力されるため、find xxx | xargs rm -rf みたいに一括操作に向いている
$ find src/ -type f
src/App.vue
src/assets/logo.png
src/components/HelloWorld.vue

$ find src/ -type f -name "*.png"
src/assets/logo.png

$ find src/ -type d
src/
src/assets
src/components
```

もっと知りたい方へ: [findコマンドで覚えておきたい使い方12個](https://orebibou.com/2015/03/find%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%81%A7%E8%A6%9A%E3%81%88%E3%81%A6%E3%81%8A%E3%81%8D%E3%81%9F%E3%81%84%E4%BD%BF%E3%81%84%E6%96%B912%E5%80%8B/)

<br>

### ・ **history**

```bash
# history | less: コマンド履歴を確認
#  * 用途1: がちゃがちゃ環境構築してて、結果うまく行ったんだけど結局何をどうしたんだっけ? を調べる
#  * 用途2: 初めて入る謎のサーバの用途を調べる
#  * 用途3: よく使うけどめっちゃ長くて覚えられないあのコマンドを探して再利用する
#  * 環境変数HISTSIZEでhistoryを保持する数を指定できる
#    デフォルト値は大概小さいので、大きくしておくと、何かあった時に幸せになれるかも
#  * sshやDB接続のパスワードをコマンド直打ちすると、historyを見て素抜かれるので気を付けよう
#  * 単にコマンドを再利用したいだけならCtrl+Rの方がおすすめ。
#    fzfを入れるとかなり使いやすくなります。
$ history | tail
 3325  find src/ -type f
 3326  find src/ -type d
 3327  find src/ -type f -name "*.png"
 3328  find src/ -type d | xargs ls
 3329  find src/ -type d | xargs ls-l
 3330  find src/ -type d | xargs ls -l
 3331  find src/ -type d | xargs -n 1 ls -l
 3332  find src/ -type d -ls
 3333  find src/ -type f -ls
 3334  history | tail

$ echo $HISTSIZE
10000
```

<br>

### ・ **diff**

```bash
# diff file1 file2: file1とfile2の差分を表示
# diff -r dir1 dir2: dir1とdir2の差分を表示(サブディレクトリもチェック)
#  * 環境構築とかで差分のあるなしをチェックしたい時によく使う
#  * 差分内容をしっかり見比べたい場合は、WinMergeなりMeldなり差分比較用のソフトを使った方がいい
$ ls
dist src
$ cp -pr src/ src2  # コピーして差分を見る => 差分なし(当たり前)
$ diff -r src src2

$ echo "abc" >> src2/App.vue  # わざと差分を作って差分を見る
$ diff -r src src2
diff -r src/App.vue src2/App.vue
17a18
> abc



# 応用編: ソートした結果同士を比較
$ cat unsort1.txt  # 1～5をランダムに並べたファイル
1
5
2
4
3

$ cat unsort2.txt  # 1～5をランダムに並べたファイル その2
1
2
3
5
4

$ diff <(cat unsort1.txt | sort) <(cat unsort2.txt | sort)  # sortした結果同士を比較すると差分なし
$
$ diff $(cat unsort1.txt | sort) $(cat unsort2.txt | sort)  # 似てるけどコマンド置換ではエラー
diff: extra operand '3'
diff: Try 'diff --help' for more information.
# 解説
#  * <(cmd): cmd1の実行結果を別コマンドの入力として扱う (プロセス置換)
#  * $(cmd): cmd1の実行結果を文字列として展開する (コマンド置換)
#  * プロセス置換を利用すると、ワンライナーで2つのファイルをソートした結果を比較できる
#  * プロセス置換を使わないと、一旦別ファイルに吐き出して...とやるので割とめんどくさい
#  * csvの比較などで活躍します
#  * <()はファイル的なものとして扱われる一方で、$()はコマンド中の文字列として展開される。
#    diffのようなファイルを引数に取るコマンドにはプロセス置換が適している。
```

<br>

### ・ **jobs, fg, bg**

```bash
# jobs: バックグラウンド実行中のジョブ一覧を表示
# fg 1: ジョブ1をフォアグラウンド実行に切り替える
# bg 1: ジョブ1をバックグラウンド実行に切り替える
#  * うっかりバックグラウンド実行してしまったプログラムをフォアグラウンドに戻すときに使う
#  * 代表ケースは、viでCtrl+Zを押したとき
#    (vi編集中にCtrl+Zを押すとジョブが停止され、どこにいったか分からなくなるのは初心者あるある。
#     そんなときは、落ち着いてjobs -> fgとすればOK)
#  * バックグラウンド実行すべきものをうっかり普通に実行した場合は、
#    Ctlr+Zで停止させて、jobs -> bgでバックグラウンド実行に切り替えられる
#    (...のだけど、いつもCtrl+Cで止めて、&を付けなおして再実行しちゃうので使ったことがない)
$ ./eternal_loop1 &  # 無限ループするプログラムをバックグラウンド実行
[1] 5906
$ ./eternal_loop2 &
[2] 5910
$ ps -ef | grep eternal_loop | grep -v grep
arene     5906  2070  0 18:29 pts/0    00:00:00 ./eternal_loop1
arene     5910  2070  0 18:29 pts/0    00:00:00 ./eternal_loop2

$ jobs  # jobsで見ると、2つバックグラウンド実行されていることが分かる
[1]-  Running                 ./eternal_loop1 &
[2]+  Running                 ./eternal_loop2 &

$ fg 2  # ジョブ番号2をフォアグラウンドに切り替える
./eternal_loop2
^C  # 無限ループして終わらないため、Ctrl+Cで終わらせる

$ jobs  # ジョブ番号2が終了したことを確認
[1]+  Running                 ./eternal_loop1 &
$ ps -ef | grep eternal_loop | grep -v grep
arene     5906  2070  0 18:29 pts/0    00:00:00 ./eternal_loop1
```

<br>

### ・ **& (バックグラウンド実行)**

```bash
# cmd1: cmd1をフォアグラウンドで実行
# cmd1 &: cmd1をバックグラウンドで実行
#  * 重たいバッチ処理や、一時的にwebサーバを動かしたいときは、
#    コマンドをバックグラウンド実行すると便利 (勿論、ターミナルをもう一つ立ち上げてもOK)
#  * 次の&&や、リダイレクトの2>&1と混同しやすいが別物
#  * この辺の記号系は慣れるしかない
$ ./eternal_loop1 &  # 無限ループするプログラムをバックグラウンド実行
[1] 6104

$ echo 123  # バックグラウンドで実行したため、他のコマンドを使える
123
```

<br>

### ・ **&&, ||**

```bash
# cmd1 && cmd2: cmd1が成功したら、cmd2を実行(cmd1が失敗したらそこで終わり)
# cmd1 || cmd2: cmd1が失敗したら、cmd2を実行(cmd1が成功したらそこで終わり)
#  * 用途1: ワンライナーでちょっとした逐次処理を書く
#  * 用途2: cmd1 || echo "error message"
#  * 実用的な例がパッとでてこないけど、ちょいちょい使うし、見かける
## 両方成功するケース
$ echo aaa && echo bbb
aaa
bbb
$ echo aaa || echo bbb
aaa

## 両方失敗するケース
$ echoooo aaa && echoooo bbb
echoooo: command not found
$ echoooo aaa || echoooo bbb
echoooo: command not found
echoooo: command not found
```

<br>
### ・ **$(), <() (コマンド置換、プロセス置換)**

```bash
# echo ${var1}: 変数var1の中身を出力 (変数展開)
# echo $(cmd1): cmd1の実行結果を出力 (コマンド置換)
# echo `cmd1`: ↑とだいたい同じ (コマンド置換(旧記法))
# diff <(cmd1) <(cmd2): cmd1とcmd2の実行結果を出力 (プロセス置換)
#  * ${}と$()は混同しやすい。jsのtemplateリテラルと同じ奴が変数置換
#  * $()は``の新記法。$(cmd1 $(cmd2))のようにネストさせやすいのが特徴
#    dateとかpgrepとか動的に変わる内容と組み合わせることが多い。
#  * <()はシェルスクリプトやワンライナーで、一時ファイルを使いたくなったときが使い時。
#    catやdiff, wihle read lineなど、ファイル内容を使う系のコマンドと組み合わせる。
$ cat lsByOption.sh  # ワンライナーでのいい例が浮かばなかったのでしょぼいシェルスクリプトを用意しました
#!/bin/bash
OPTION=$1
ls $(echo ${OPTION})  # 第1引数が-lなら、ls -lになる

$ ls  # 普通にlsを実行
lsByOption.sh  unsort1.txt  unsort2.txt

$ ./lsByOption.sh -l  # ls $(echo ${OPTION})がls -lになる
total 0
-rwxr-xr-x 1 arene arene 45 Feb  9 19:44 lsByOption.sh
-rw-r--r-- 1 arene arene 10 Feb  9 19:29 unsort1.txt
-rw-r--r-- 1 arene arene 10 Feb  9 19:30 unsort2.txt

$ ./lsByOption.sh -al  # ls $(echo ${OPTION})がls -alになる
total 0
drwxr-xr-x 1 arene arene 4096 Feb  9 19:44 .
drwxr-xr-x 1 arene arene 4096 Feb  9 19:28 ..
-rwxr-xr-x 1 arene arene   45 Feb  9 19:44 lsByOption.sh
-rw-r--r-- 1 arene arene   10 Feb  9 19:29 unsort1.txt
-rw-r--r-- 1 arene arene   10 Feb  9 19:30 unsort2.tx
```

もっと知りたい方へ:
・[コマンドとコマンドをつなぐ糊](https://qiita.com/greymd/items/32d4dcb6fff4832f1fc5)
・[bashのプロセス置換機能を活用して、シェル作業やスクリプト書きを効率化する](https://sechiro.hatenablog.com/entry/2013/08/15/bash%E3%81%AE%E3%83%97%E3%83%AD%E3%82%BB%E3%82%B9%E7%BD%AE%E6%8F%9B%E6%A9%9F%E8%83%BD%E3%82%92%E6%B4%BB%E7%94%A8%E3%81%97%E3%81%A6%E3%80%81%E3%82%B7%E3%82%A7%E3%83%AB%E4%BD%9C%E6%A5%AD%E3%82%84%E3%82%B9)

<br>

### ・ **$?**

```bash
# echo $?: 直前のコマンドの終了ステータスを表示
#  * シェルスクリプトで異常系の処理を書くときに使う?
#  * いきおいで挙げたけどあんまり使わないかも
$ echo 123  # OKケース
123
$ echo $?
0

$ hdskds  # NGケース
hdskds: command not found
$ echo $?
127
```

<br>

### ・ **for**

```bash
# for i in {1..10} ; do cmd1; done: cmd1を10回繰り返す
#  * よく使いたくなるけど、使うたびにどう書いたっけ? となるやる
#  * ワンライナーにするために無理やり1行にしてるけど、ちゃんと改行すると↓になる
#    for i in {1..10} ;
#    do
#      cmd1;
#    done
#  * {1..10}: ブレース展開(連番ver): 1 2 3 4 5 6 7 8 9 10に展開される
#  *          代わりに$(seq 10)でもOK
$ for i in {1..10} ; do echo $i; done
1
2
3
4
5
6
7
8
9
10
```

<br>

## 書かないことにしたやつ

よく使うけど流石にスコープ外だろう、まれに使うけど実はあんまり知らない
・・・といった理由で割愛したくなったコマンド達。
項目だけ挙げます。

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|vi|ファイルの編集|visual editor(visual interface|
|make|プログラムのコンパイル|make|
|curl|HTTPリクエストを出す|command url??|
|rsync|ネットワーク越しにディレクトリ内容を同期|remote synchronizer|
|ssh-keygen|sshの秘密鍵、公開鍵を作る|ssh key generator|
|npm|nodeのパッケージをインストールするなど|node package manager|
|git|gitを使う(雑)|イギリス英語のスラングでバカ <- はじめて知った|

<br>

## おまけ

全く使わないけど紹介してみたかったので、書きました。たのしかった。

|コマンド名|何ができる?|コマンド名は何に由来している?|
|---|---|---|
|nice|ユーザーの優先度を調整||
|sl|汽車が出る|lsの反対|

### ・ **nice**

```bash
# nice -n 20 cmd1: cmd1を優先度20で実行
#  * linuxの各ユーザはnice値という優先度を持っている
#    -20(優先度最高)～20(優先度最低)
#  * CPU負荷の高いバッチ処理を、他プロセスの隙間時間に実行したいときに使えるかも(経験ないです)
#  * 高野豊さんの「rootから/へのメッセージ」というエッセイ本で知りました。
#    日本にunixが来た当初、cpuリソースは貴重だったため
#    重たい処理をバンバン走らせる素行の悪いユーザに対して、root管理者はnice値を上げて対抗していたそうです。
#  * このエピソードが何故か分からないけど私はすごく好き
$ nice -n 20 ls
src dist
```

参考: [rootから/へのメッセージ(Amazon)](https://www.amazon.co.jp/root-%E3%83%AB%E3%83%BC%E3%83%88-%E3%81%8B%E3%82%89-%E3%81%B8%E3%81%AE%E3%83%A1%E3%83%83%E3%82%BB%E3%83%BC%E3%82%B8%E2%80%95%E3%82%B9%E3%83%BC%E3%83%91%E3%83%BC%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E3%81%8C%E8%A6%8B%E3%81%9F%E3%81%B2%E3%81%A8%E3%81%A8%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF/dp/4756107869)

<br>

### ・ **sl  ※インストールが必要**

実行するとターミナルに汽車が走ります。
typoしたときは汽車でもみて落ち着こう。
(ご丁寧にCtrl+Cを無効化してるあたり、ユーモアにあふれてて素敵)

![image.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/162371/11d6044a-84e3-d2c0-e3ee-f5c8193e8153.png)

もっと知りたい方へ: [仕事で役に立たない！Linuxネタコマンド集](https://qiita.com/ryuichi1208/items/598cb0571a2576ecd0e1)

<br>

## 最後に

コマンドライン操作は、楽しいです。
また、Linux(Unix)のシステムまわりの話も面白いです。
この記事を通して、ちょっとでも発見や興味の広がりがあれば幸いに思います。
