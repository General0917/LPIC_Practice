---
title: dockerのコンテナの中に入る
tags: Docker docker-compose
author: General0917
slide: false
---
# 概要
webアプリケーションを勉強中で、dockerが良いということを知ったので、古いパソコンにUbuntuを入れて色々調べながら、何とか環境を構築したのですが、いざ動かそうとしたときに`rails`コマンドとかが実行できず、そこら辺を理解するのに時間がかかったので、まとめました。

#実行できなかった理由
当たり前なんですが、railsの実行環境はコンテナの中にあるので、コンテナに入る必要がありました。
そして、その方法を色々調べた結果、2つの方法があることが分かりました。

# dockerコマンド
```
docker container exec -it コンテナ名 bash
```
コンテナ名は

```
docker-compose ps
```

で見ることができます。また、

```
docker container ps
```
より得られるコンテナIDを入力してもいけました。

# docker-composeコマンド
```
docker-compose exec サービス名 bash
```
ここでのコンテナ名は`docker-compose.yml`で入力したサービス名となっています。

# 補足
`exec`コマンドは、起動中のコンテナ内で、指定したコマンドを実行するコマンドで、`bash`や`sh`を実行することで、コンテナ環境とカーネル等をつなげてくれます。
このとき、`docker container exec`の方は端末の起動オプション`t`と標準入力を開き続けるオプション`i`を指定する必要があります。
`docker-compose exec`の方はデフォルトで端末が割り与えられるようになっているみたいなので、オプションをはいらないです。
コンテナ内のコマンドを1回だけ実行したい場合は、`bash`を`rails -v`などに書き換えればOKです。

# 参考
[docker-compose exec](https://matsuand.github.io/docs.docker.jp.onthefly/compose/reference/exec/)
[【初心者向け】シェル・ターミナル・コンソールの違いとは？](https://eng-entrance.com/linux-basic-shell-terminal-console)
[docker runとdocker execの違いの解説](https://www.memotansu.jp/docker/852/)
