func productExceptSelf(nums []int) []int {
    ans := make([]int, len(nums))
    pre, post := 1, 1
    for i := range nums {
        ans[i] = pre
        pre = pre * nums[i]
    }

    for i := len(nums)-1; i >= 0; i-- {
        ans[i] = ans[i] * post
        post = post * nums[i]
    }

    return ans
}