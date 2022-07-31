go言語で簡単なSocket通信（TCP/IP）を行う（Qiitaからの引用。若干修正。）

Client<br>
プログラム起動でServerと接続<br>
標準入力からの文字列をServerに送信。Serverからの返信を出力して、無限loop。<br>
Ctrl-Cで終了する。<br>

Server<br>
Clientからの接続を待ち、受信した文字列を出力して、Clientに同じものを返す。<br>
