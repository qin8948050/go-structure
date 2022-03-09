/*
n个数需要n-1趟排序 
 */
func BubbleSort(s []string) []string{
	for i:=1;i<=len(s)-1;i++{
		for j:=0;j<=len(s)-1-i;j++{
			if s[j] > s[j+1] {
				s[j],s[j+1]=s[j+1],s[j]
			}
		}
	}
	return s
}