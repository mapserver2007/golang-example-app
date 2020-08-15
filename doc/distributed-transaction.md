# 分散トランザクション

参考：
* [マイクロサービスにおける決済トランザクション管理](https://engineering.mercari.com/blog/entry/2019-06-07-155849/#%E5%88%86%E6%95%A3%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AB%E3%81%8A%E3%81%91%E3%82%8B%E3%83%87%E3%83%BC%E3%82%BF%E6%95%B4%E5%90%88%E6%80%A7%E6%8B%85%E4%BF%9D%E3%81%AE%E9%9B%A3%E3%81%97%E3%81%95)
* https://github.com/lysu/go-saga/blob/master/saga.go

## 既知の対策
* 2フェーズコミット
    * 分散システム内のすべてのノードの合意をとるためのアルゴリズム
    * coodinator(調整者)がすべてのcohorts(参加者)の調整を行う
* 2フェーズコミットの問題点
    * 実行中にすべての参加者がブロックされるため性能が悪い、リソースがロックされたままになることがある
    * 分散されているすべてのサービスで完了にならないのでコミットまでに時間がかかる
    * かんたんにスケールしない

## Sagaパターン
* Sagaパターンは複数のサービスにまたがるトランザクションを実現するパターンの1つ
* 各小サービスでは通常のトランザクションで処理、全体として大きなトランザクションを構成する
* SEC(Saga Execution Coodinator)がログの書き込みとサービス間の調整を行う
* ログは各サービスにおける小規模のトランザクションとコミットを書き込んでいく
* もしロールバックする場合はその書き込まれたログを逆順にたどってもとに戻す