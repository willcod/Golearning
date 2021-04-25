/*
Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/
/*
A: 对于是否要抛给上层，取决于上下层之间的约定。如果底层不需要暴露具体的实现给上层，上层也不需要知道底层是怎么实现的。
	那么底层只要给出成功或失败给上层就可以了。但是往往，我们很希望知道error发生的第一现场以及call stack等信息用来定位问题，
	那么error就应该wrap了之后往上抛。通过使用pkg/errors包，我们可以往错误里添加额外的定位信息，最后抛给上层处理。
*/

package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func daoLevel() (string, error) {
	return "", errors.Wrap(sql.ErrNoRows, "DAO level error")
}

func accessDBandErrorHappens() (string, error) {
	content, err := daoLevel()
	if err != nil {
		return "", err
	}
	return content, nil
}

func main() {
	result, err := accessDBandErrorHappens()
	fmt.Println("result:", result)
	fmt.Println("error", err)
}
