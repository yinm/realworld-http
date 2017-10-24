resp, err := http.Get("http://...")
if err != nil {
	// エラー発生
	panic(err)
}

// スコープが抜けたところで必ずクローズ
defer resp.Body.Close()

// ioutil.ReadAllでサーバーレスポンスを最後まで一括で読み切る
body, err := ioutil.ReadAll(resp.Body)
