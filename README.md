go言語で簡単なSocket通信（TCP/IP）を行う（Qiitaからの引用）

Client
プログラム起動でServerと接続
標準入力からの文字列をServerに送信。Serverからの返信を出力して、無限loop。
Ctrl-Cで終了する。

Server
Clientからの接続を待ち、受信した文字列を出力して、Clientに同じものを返す。
