// Copyright (c) 2020 qianjunakasumi <i@qianjunakasumi.ren>
//                    qianjunakasumi <qianjunakasumi@outlook.com>
//                    [qianjunakasumi] 千橘雫霞(https://github.com/qianjunakasumi)
//
//      This Source Code Form is subject to the terms of the Mozilla Public
//      License, v. 2.0. If a copy of the MPL was not distributed with this
//      file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import "fmt"

func main() {

	var (
		Num        int // 被除数
		numberList = []int{2, 3, 4, 5, 6}
		textList   = []string{"一", "二", "三", "四", "五"}
		result     string // 结果
	)

	fmt.Println("请输入一个数字")
	if _, err := fmt.Scanln(&Num); err != nil {
		panic(err)
	}

	for i := 0; i < len(numberList); i++ {

		var (
			chuShu = numberList[i] // 除数
			yuShu  = Num % chuShu  // 余数
		)

		if yuShu != 0 {

			fmt.Printf(
				"\n第%v步，用%v除%v，得到余数%v. 因为余数不为0，所以%v不能整除%v. ",
				textList[i], chuShu, Num, yuShu, chuShu, Num,
			)
			result = "是"
			continue

		}

		fmt.Printf(
			"\n第%v步，用%v除%v，得到余数%v. 因为余数为0，所以%v能整除%v. ",
			textList[i], chuShu, Num, yuShu, chuShu, Num,
		)
		result = "不是"
		break

	}

	fmt.Printf("因此，%v%v质数\n", Num, result)
	fmt.Println("\n程序执行完毕，按任意键退出")
	_, _ = fmt.Scanln(&Num)

}
