package stack

import "container/list"

func longestValidParentheses(s string) int {
	return 0
}

// minOperations 1598. 文件夹操作日志搜集器
func minOperations(logs []string) int {
	stack := list.New()

	for _, item := range logs {
		if item == "./" {
			continue
		}

		if item == "../" {
			if stack.Len() > 0 {
				stack.Remove(stack.Front())
			}

			continue
		}

		stack.PushFront(item)
	}

	return stack.Len()
}
