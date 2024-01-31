import "strings"
func lengthOfLastWord(s string) int {
    // split string by space delimiter
    wordArray := strings.Fields(s)

    // get last word in string
    var lastWord string = wordArray[len(wordArray)-1]

    // return length of last word in word array
    return len(lastWord)
    
}