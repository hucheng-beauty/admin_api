package password

import (
	"crypto/sha512"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

// Generate 生成密文
func Generate(pwd string) string {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	// 通过盐值解决彩虹表问题
	salt, encodedPwd := password.Encode(pwd, options)
	// $算法$盐值%密文
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

// Verify 校验密码
func Verify(encryptedPwd, pwd string) bool {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	passwordInfo := strings.Split(encryptedPwd, "$")
	return password.Verify(pwd, passwordInfo[2], passwordInfo[3], options)
}
