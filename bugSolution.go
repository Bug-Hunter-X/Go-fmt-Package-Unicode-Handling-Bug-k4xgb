func main() {
  unicodeString := "你好，世界!\u2605"
  fmt.Println("Original String:", unicodeString)

  //Corrected Version
  fmt.Printf("Corrected String: %%s\n", unicodeString)
} 