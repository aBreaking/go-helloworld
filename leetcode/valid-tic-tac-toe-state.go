package main

import (
	"fmt"
	"strings"
)

/**
794. 有效的井字游戏
给你一个字符串数组 board 表示井字游戏的棋盘。当且仅当在井字游戏过程中，棋盘有可能达到 board 所显示的状态时，才返回 true 。

井字游戏的棋盘是一个 3 x 3 数组，由字符 ' '，'X' 和 'O' 组成。字符 ' ' 代表一个空位。

以下是井字游戏的规则：

玩家轮流将字符放入空位（' '）中。
玩家 1 总是放字符 'X' ，而玩家 2 总是放字符 'O' 。
'X' 和 'O' 只允许放置在空位中，不允许对已放有字符的位置进行填充。
当有 3 个相同（且非空）的字符填充任何行、列或对角线时，游戏结束。
当所有位置非空时，也算为游戏结束。
如果游戏结束，玩家不允许再放置字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-tic-tac-toe-state
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

	var board = []string{"XOX", "OXO", "XOX"}

	fmt.Println(validTicTacToe(board))

}

/**
假设玩家一X数量为xn，玩家二O数量为on，分一下三种情况：
1. 没有玩家赢，必然有 0<=xn-on<=1
2. X赢，那么必然满足 xn-on=1
3. O赢，那么必然满足 xn=on

然后就是赢的算法了，直线、对角线 分别计算
1. 三横三纵
2. 两条对角线
*/
func validTicTacToe(board []string) bool {

	var xn, on int
	for _, s := range board {
		xn += strings.Count(s, "X")
		on += strings.Count(s, "O")
	}
	if xn-on > 1 || xn-on < 0 {
		return false
	}

	if win(board, "X") && xn-on != 1 {
		return false
	}

	if win(board, "O") && xn != on {
		return false
	}

	return true
}

func win(board []string, c string) bool {
	for i := 0; i < 3; i++ {
		if c == string(board[i][0]) && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return true
		}
		if c == string(board[0][i]) && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return true
		}
	}

	if c == string(board[0][0]) && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return true
	}

	if c == string(board[0][2]) && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return true
	}

	return false
}
