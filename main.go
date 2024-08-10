package main

import (
	"fim/core"
	chatmodel "fim/fim_chat/models"
	groupmodel "fim/fim_group/models"
	usermodel "fim/fim_user/models"
	"flag"
	"log"
	"sync"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

type Options struct {
	DB bool
}

func main() {
	var opt Options

	//这行代码使用flag包的BoolVar函数来定义一个命令行标志。
	//这个标志的名称是"db"，默认值是false，并且会将解析的值存储到opt.DB中。
	//标志的描述也是"db"。

	//flag.Parse()：解析命令行参数，并将解析的值赋给之前定义的标志变量
	//（在这个例子中是opt.DB）。

	//if opt.DB ：检查opt.DB的值。如果为true（表示用户在命令行中使用了-db标志），
	//则执行大括号内的代码。
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		db := core.InitMysql()
		db.AutoMigrate(
			&usermodel.FriendModel{},
			&usermodel.UserModel{},
			&usermodel.UserConfigModel{},
			&usermodel.FriendVerifyModel{},

			&groupmodel.GroupModel{},
			&groupmodel.GroupMembersModel{},
			&groupmodel.GroupMsgModel{},
			&groupmodel.GroupVerifyModel{},

			&chatmodel.ChatModel{},
		)
	}
	rate := uint64(100) // 设置QPS值
	duration := 4 * time.Second

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://127.0.0.1:8888/api/user/search_friend",
	})

	attacker := vegeta.NewAttacker()
	pacer := vegeta.ConstantPacer{Freq: int(rate), Per: time.Second}
	var wg sync.WaitGroup

	// 创建一个切片来存储每个请求的时间戳
	var timestamps []time.Time
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for res := range attacker.Attack(targeter, pacer, duration, "Test") {
				log.Println(res)
				// 记录每个请求的时间戳
				mu.Lock()
				timestamps = append(timestamps, time.Now())
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// 计算并输出每秒的请求数
	qps := calculateQPS(timestamps)
	logQPS(qps)
	averageQPS := calculateAverageQPS(qps)
	log.Printf("平均每秒的请求数: %.2f\n", averageQPS)
}
func calculateQPS(timestamps []time.Time) map[string]int {
	qps := make(map[string]int)
	for _, t := range timestamps {
		sec := t.Format("2006-01-02 15:04:05")
		qps[sec]++
	}
	return qps
}
func logQPS(qps map[string]int) {
	log.Println("QPS:")
	for sec, count := range qps {
		log.Printf("%s: %d\n", sec, count)
	}
}
func calculateAverageQPS(qps map[string]int) float64 {
	totalRequests := 0
	for _, count := range qps {
		totalRequests += count
	}
	totalSeconds := len(qps)
	if totalSeconds == 0 {
		return 0
	}
	return float64(totalRequests) / float64(totalSeconds)
}
