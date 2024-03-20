package main

import (
	"fmt"
	"log"

	pigimongo "github.com/ikasoba/pigimongo-db/core"
)

type DiaryEntry struct {
	Id      string `pigimongo:"Id_"`
	Content string
}

func main() {
	db, err := pigimongo.NewDatabase(":memory:")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("\n1. メモを追加する")
		fmt.Println("2. メモを検索する")
		fmt.Println("3. 特定の既存メモを更新する")
		fmt.Println("4. 特定の既存メモを消去する")
		fmt.Println("5. どろん")
		fmt.Print("操作したいものを番号で指示してね: ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			log.Println("何か読み込めんかった...:", err)
			continue
		}

		switch choice {
		case 1:
			addDiaryEntry(db)
		case 2:
			viewDiaryEntry(db)
		case 3:
			updateDiaryEntry(db)
		case 4:
			removeDiaryEntry(db)
		case 5:
			fmt.Println("夢に出てきてあげるよ...")
			return
		default:
			fmt.Println("何か面白い指示をしたようだね。")
		}
	}
}

func addDiaryEntry(db *pigimongo.Database) {
	var content string
	fmt.Print("メモを入力してね: ")
	if _, err := fmt.Scanln(&content); err != nil {
		log.Println("一体何書いたんや...:", err)
		return
	}

	entry := DiaryEntry{Content: content}
	err := db.Add(entry)
	if err != nil {
		log.Println("DBに追加できなかった...:", err)
		return
	}

	fmt.Println("メモが正常に追加されました！!")
}

func viewDiaryEntry(db *pigimongo.Database) {
	var content string
	fmt.Print("調べたい単語を入力してね: ")
	if _, err := fmt.Scanln(&content); err != nil {
		log.Println("読み込めへん...:", err)
		return
	}

	entry := &DiaryEntry{}
	err := db.Find(entry, `Content == ?`, content)
	if err != nil {
		log.Println("メモ無かったわ。:", err)
		return
	}
	//IDの実装まだやってないけどそのうち
	fmt.Printf("メモ見つかったで。:\nID: %s\nContent: %s\n", entry.Id, entry.Content)
}

func updateDiaryEntry(db *pigimongo.Database) {
	var oldContent, newContent string
	fmt.Print("更新したいメモ入力してね: ")
	if _, err := fmt.Scanln(&oldContent); err != nil {
		log.Println("読み込めへん...:", err)
		return
	}

	fmt.Print("どんな内容にする？: ")
	if _, err := fmt.Scanln(&newContent); err != nil {
		log.Println("読み込めへん...:", err)
		return
	}

	err := db.Update(DiaryEntry{Content: newContent}, `Content == ?`, oldContent)
	if err != nil {
		log.Println("DBの更新に失敗したわ。:", err)
		return
	}

	fmt.Println("メモの更新に成功したで!")
}

func removeDiaryEntry(db *pigimongo.Database) {
	var content string
	fmt.Print("削除したいメモ教えて: ")
	if _, err := fmt.Scanln(&content); err != nil {
		log.Println("読み込めへん...:", err)
		return
	}

	err := db.Remove(`Content == ?`, content)
	if err != nil {
		log.Println("DBから取り除かれへんかったわ...:", err)
		return
	}

	fmt.Println("DBから除去できたで!")
}
