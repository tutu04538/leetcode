/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

package main

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	/*
			排序法
		var output [][]string
		m := make(map[string][]string)
		for _, str := range strs {
			b := []byte(str)
			sort.Slice(b, func(i, j int) bool {
				return b[i] < b[j]
			})
			key := string(b)
			m[key] = append(m[key], str)
			// 下面的代码并不会修改 map 中的 []string
			// c := m[string(b)]
			// c = append(c, str)
		}

		for _, v := range m {
			output = append(output, v)
		}
		return output
	*/

	// 计数法
	var output [][]string
	m := make(map[[26]int32][]string)

	for _, str := range strs {
		var cnt [26]int32
		for _, c := range str {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}

	for _, v := range m {
		output = append(output, v)
	}
	return output

}

// @lc code=end
