func cal(s string) int {
    cnt := 0
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            cnt++
        }
    }
    return cnt + 1
}

	
type pair struct {
    Name string
    Cnt  int
}

func largestWordCount(messages []string, senders []string) string {
    f := make(map[string]int)
    for i := 0; i < len(messages); i++ {
        f[senders[i]] += cal(messages[i])
    }
    
    a := make([]*pair, 0)
    
    for k, v := range f {
        a = append(a, &pair{k, v})
    }
    
    sort.SliceStable(a, func(i, j int) bool {
        if a[i].Cnt == a[j].Cnt {
            return a[i].Name > a[j].Name
        }
		return a[i].Cnt > a[j].Cnt
	})
    
    return a[0].Name
}