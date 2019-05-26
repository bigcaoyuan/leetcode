import "fmt"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

func largestRectangleArea(heights []int) int {
	s := &Stack{d: []int{}}
	heights = append(heights, 0)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for !s.Empty() && heights[i] < heights[s.Top()] {
			fmt.Println(i, heights[i], heights[s.Top()])
			top := s.Top()
			s.Pop()
			if s.Empty() {
				maxArea = max(maxArea, heights[top]*i)
			} else {
				maxArea = max(maxArea, heights[top]*(i-s.Top()-1))
			}
		}
		s.Push(i)
	}
	return maxArea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Stack struct {
	d []int
}

func (s *Stack) Pop() {
	if len(s.d) == 0 {
		return
	}
	s.d = s.d[:len(s.d)-1]
	return
}

func (s *Stack) Push(val int) {
	s.d = append(s.d, val)
}
func (s *Stack) Top() int {
	if len(s.d) == 0 {
		return 0
	}
	return s.d[len(s.d)-1]
}

func (s *Stack) Empty() bool {
	return len(s.d) == 0
}
