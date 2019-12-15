# go_api
GOで作成した簡単なAPI  
日付、体重、体脂肪率、筋肉量の記録と取得

get / => 一覧取得  
post / => 作成　(date, weight, bfp, mm)  
patch / => 更新  
delete /date => 削除 

フレームワークにechoを使用、データベースにmysql, mysql操作に標準パッケージのdatabase/sql使用。
