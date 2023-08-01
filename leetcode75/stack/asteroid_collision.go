package stack

func asteroidCollision(asteroids []int) []int {
    var stack []int
    for _, asteroid := range asteroids {
        if asteroid > 0 {
            stack = append(stack, asteroid)
        } else {
            for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -asteroid {
                stack = stack[:len(stack)-1]
            }
            if len(stack) > 0 && stack[len(stack)-1] == -asteroid {
                stack = stack[:len(stack)-1]
            } else if len(stack) == 0 || stack[len(stack)-1] < 0 {
                stack = append(stack, asteroid)
            }
        }
    }
    return stack
}
