package utils

import (
	"database_lesson/dto"
	"database_lesson/global"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

func SendEmailValidate(em []string) (string, error) {
	e := email.NewEmail()
	//e.From = fmt.Sprintf("发件人笔名 <发件人邮箱>")
	//e.To = []string{"1530122381@qq.com"}
	e.From = "1357880399@qq.com"
	e.To = em
	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")
	//设置文件发送的内容
	content := fmt.Sprintf(`
	<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
	`, em[0], t, vCode)
	e.Text = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1357880399@qq.com", "oyagqeommuakidaf", "smtp.qq.com"))
	fmt.Println(111)
	return vCode, err
}

// TODO 把超时时间抽象到yaml文件
func InsertCodeIntoRedis(model *dto.CertifyDto) (err error) {
	global.RedisClient.Do("set", model.ID, model.Code)
	global.RedisClient.Do("expire", model.ID, 300)
	return nil
}
