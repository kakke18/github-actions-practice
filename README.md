# github-actions-practice

GitHub Actions練習用として、
Slackにメッセージを送信するプログラムを用いる。

## Usage
1. [リポジトリの暗号化されたシークレットの作成](https://docs.github.com/ja/free-pro-team@latest/actions/reference/encrypted-secrets#%E3%83%AA%E3%83%9D%E3%82%B8%E3%83%88%E3%83%AA%E3%81%AE%E6%9A%97%E5%8F%B7%E5%8C%96%E3%81%95%E3%82%8C%E3%81%9F%E3%82%B7%E3%83%BC%E3%82%AF%E3%83%AC%E3%83%83%E3%83%88%E3%81%AE%E4%BD%9C%E6%88%90)を参考に`$WEBHOOK_URL`を設定
2. 適宜`job.yaml`を編集。[参考](https://docs.github.com/ja/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#on)
3. pushする
