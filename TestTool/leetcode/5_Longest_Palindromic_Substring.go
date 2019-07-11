package main

// dynamic
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	boolMap := make([][]bool, len(s))
	start, maxLen := 0, 1

	for i, _ := range boolMap {
		boolMap[i] = make([]bool, len(s))
	}

	// 初始化1, 2的长度
	for i := 0; i < len(s); i++ {
		boolMap[i][i] = true

		if i < len(s)-1 && s[i] == s[i+1] {
			boolMap[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	for strLen := 3; strLen <= len(s); strLen++ {
		for i := 0; i <= len(s)-strLen; i++ {
			j := i + strLen - 1

			if boolMap[i+1][j-1] && s[i] == s[j] {
				boolMap[i][j] = true
				maxLen, start = strLen, i
			}
		}
	}
	return s[start : start+maxLen]
}

// 从中心向两侧扩展 O(n^2)
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0

	for i, _ := range s {
		nowLen := maxPand(s, i)
		if nowLen > end - start {
			start, end = i - (nowLen-1)/2, i + nowLen/2
		}
	}

	return s[start:end+1]
}

// 获得可见范围内的最大长度
func canExpand(s string, a, b int) int {
	for a >= 0 && b < len(s) && s[a] == s[b] {
		a--
		b++
	}

	return b-a-1
}

func maxPand(s string, a int) int {
	len1 := canExpand(s, a, a)
	len2 := canExpand(s, a, a+1)

	if len2 > len1 {
		return len2
	} else {
		return len1
	}
}

// manacher
func dealString(s string) string {
	s1 := make([]byte, 2*len(s)+1)

	for i := 0; i < len(s); i++ {
		s1[2*i] = '#'
		s1[2*i+1] = s[i]
	}

	return string(s1)
}

/*
int max_r=0,pos=0;
    for (int i=0;i<len;++i){
        if (max_r>i) r[i]=min(max_r-i,r[2*pos-i]);
        else r[i]=0;
        while (i+r[i]+1<len&&i-r[i]-1>=0&&s[i+r[i]+1]==s[i-r[i]-1]) r[i]++;
        if (r[i]+i>max_r){
            max_r=r[i]+i;
            pos=i;
        }
    }
 */

func longestPalindrome(s string) string {
	subLen := make([]int, len(s))
	s = dealString(s)

	max_r, pos := 0, 0
	for i := 0; i < len(s); i++ {
		if max_r > i {
			subLen[i] = minInt(max_r - i, subLen[2*pos-i])
		} else {
			subLen[i] = 0
		}

		for i + subLen[i] + 1 < len(s) && i - subLen[i] - 1 >= 0 &&
			s[i + subLen[i] + 1] == s[i - subLen[i] - 1] {
				subLen[i]++
		}

		if subLen[i]+i > max_r {
			max_r = subLen[i]+i
			pos = i
		}
	}

	return s[pos-max_r:pos+max_r]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}