package data

import (
	"time"
)

/* Thread構造体  */
type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreateAt time.Time
}

/* indexハンドラ関数向けのデータベースからスレッドを抽出する関数。  */
func Threads() (threads []Thread, err error){

	// クエリ送信。行を受け取る。
	rows, err := Db.Query("SELECT id, uuid, topic, user, user_id, created_at FROM threads ORDER BY created_at DESC")

	if err != nil {
		return
	}

	// 構造体:thに受け取った行を取り込む。
	for rows.Next(){
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreateAt); err != nil {
			return
		}

		threads = append(threads, th)
	}
	rows.Close()
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", thread.Id)

	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}

	rows.Close()
	return
}
