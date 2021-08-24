package functions

/*
Ошибки, использующиеся в программе:
1) wrongColNum - не указано количество строк в ключе -f.
2) noCommands - комаанда некорректна или вообще не задана.
3) emptyResult - в результате работы возвращён "пустой" срез. Проблема в запросе.
4) wrongKeysSetSF - одновременно выбраны ключи -s и -f.
5) wrongKeysSetDF - одновременно выбраны ключи -d и -f.
*/

var wrongColNum = "Selected wrong number of column in 'f' key. Please insert column number in command, like 'f1'"
var noCommands = "You entered wrong command! Wasn't enter the keys. Please entered your command again, like 'cut -f1sd'"
var emptyResult = "Received the empty result. Please, check your command!"
var wrongKeysSetSF = "Keys 's' and 'f' can't work together! Please check your command and choose 's' or 'f' command ONLY!"
var wrongKeysSetDF = "Keys 'd' and 'f' can't work together! Please check your command and choose 'd' or 'f' command ONLY!"
