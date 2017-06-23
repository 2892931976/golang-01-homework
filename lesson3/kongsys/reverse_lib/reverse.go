package reverse

func Reverse(s string) string {
   /* can not chinese
   strs := []byte(s)
   for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
        strs[i], strs[j] = strs[j], strs[i]
   }
   return string(strs)
   */
   runes := []rune(s)
   for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
       runes[i], runes[j] = runes[j], runes[i]
   }
   return string(runes)
}
