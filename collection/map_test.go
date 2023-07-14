package colllection

import (
	"fmt"
	"testing"
)

/*
map(映射)是Go语言中的一种内置类型,它将键值对互相关联,我们可以通过键key获取对应的值value
类似其他语言的例如 Java 的 HashMap
*/
func TestMap(t *testing.T) {
	// 声明一个map, map是Go语言的关键字, 要避免命名冲突
	// key不为为 slice, map, function
	// 必须通过make()函数进行初始化
	var hashMap = make(map[int]string, 10)

	// 进行数据初始化
	hashMap[202180910] = "洛必达"
	hashMap[202180911] = "牛顿"
	hashMap[202180912] = "拉布尼茨"
	hashMap[202180913] = "阿基米德"
	hashMap[202180914] = "苏格拉底"

	// 输出集合
	fmt.Println(hashMap)

	// 方式2
	var hashMap1 = make(map[int]string)
	hashMap1[202180915] = "亚里士多德"
	hashMap1[202180916] = "欧几里得"
	fmt.Println(hashMap1)

	// 方式3
	var hashMap2 = map[int]string{
		202180917: "柏拉图",
		202180918: "阿波罗尼斯",
	}
	fmt.Println(hashMap2)

	// 删除集合
	delete(hashMap, 202180910)
	fmt.Println(hashMap)

	// 查找操作
	value, flag := hashMap[202180911]
	fmt.Println(value, flag)

	// 获取长度
	fmt.Println(len(hashMap))

	// 遍历map
	for k, v := range hashMap {
		fmt.Printf("%v: %v \n", k, v)
	}
}

/*
测试嵌套多重map
*/
func TestMultipleMap(t *testing.T) {
	// 参考Map<String, Map<Integer, String>>
	hashMaps := make(map[string]map[int]string)

	// 赋值
	hashMaps["班级1"] = make(map[int]string, 3)
	hashMaps["班级1"][2000158] = "张三"
	hashMaps["班级1"][2000157] = "李四"
	hashMaps["班级1"][2000159] = "王五"

	hashMaps["班级2"] = make(map[int]string, 3)
	hashMaps["班级2"][2000110] = "李小龙"
	hashMaps["班级2"][2000111] = "吕布"
	hashMaps["班级2"][2000112] = "刘备"

	for k1, hashMap := range hashMaps {
		fmt.Println(k1)
		for k2, value := range hashMap {
			fmt.Printf("%v: %v \n", k2, value)
		}
		fmt.Println()
	}

}
