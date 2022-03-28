[300. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
```
func lengthOfLIS(nums []int) int {
    dp:=make([]int,len(nums))
    for i:=0;i<len(nums);i++{
        dp[i]=1
    }
    max:=1
    for i:=1;i<len(nums);i++{
        for j:=0;j<i;j++{
            if nums[j]<nums[i]{
                dp[i]=com(dp[i],dp[j]+1)
            }
            if dp[i]>max{
                max=dp[i]
            }
        }
    }
    return max
}

func com(a,b int)int{
    if a>b{
        return a
    }
    return b
}
```

[674. 最长连续递增序列](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/)
```
func findLengthOfLCIS(nums []int) int {
    dp:=make([]int,len(nums))
    for i:=0;i<len(nums);i++{
        dp[i]=1
    }
    max:=1
    for i:=1;i<len(nums);i++{
        if nums[i]>nums[i-1]{
            dp[i]=dp[i-1]+1
        }
        if dp[i]>max{
            max=dp[i]
        }
    }
    return max
}
```

[718. 最长重复子数组](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/)
```
func findLength(nums1 []int, nums2 []int) int {
    max:=0
    dp:=make([][]int,len(nums1))
    for i:=0;i<len(nums1);i++{
        dp[i]=make([]int,len(nums2))
    }
    for i:=0;i<len(nums1);i++{
        if nums1[i]==nums2[0]{
            dp[i][0]=1
            max=1
        }
    }
    for j:=0;j<len(nums2);j++{
        if nums2[j]==nums1[0]{
            dp[0][j]=1
            max=1
        }
    }
    
    for i:=1;i<len(nums1);i++{
        for j:=1;j<len(nums2);j++{
            if nums1[i]==nums2[j]{
                dp[i][j]=dp[i-1][j-1]+1
            }
            if dp[i][j]>max{
                max=dp[i][j]
            }
        }
    }
    return max
}
```


```
func longestCommonSubsequence(text1 string, text2 string) int {
    max:=0
    dp:=make([][]int,len(text1)+1)
    for i:=0;i<=len(text1);i++{
        dp[i]=make([]int,len(text2)+1)
    }
    for i:=0;i<len(text1);i++{
        for j:=0;j<len(text2);j++{
            if text1[i]==text2[j]{
                dp[i+1][j+1]=dp[i][j]+1
            }else{
                dp[i+1][j+1]=com(dp[i][j+1],dp[i+1][j])
            }
            if dp[i+1][j+1]>max{
                max=dp[i+1][j+1]
            }
        }
    }
    return max

}

func com(a,b int)int{
    if a>b{
        return a
    }
    return b
}
```

[53. 最大子数组和](https://leetcode-cn.com/problems/maximum-subarray/)
```
func maxSubArray(nums []int) int {
    dp:=make([]int,len(nums))
    copy(dp,nums)
    max:=nums[0]
    for i:=1;i<len(nums);i++{
        dp[i]=com(dp[i-1]+nums[i],nums[i])
        if dp[i]>max{
            max=dp[i]
        }
    }
    return max
}

func com(a,b int)int{
    if a>b{
        return a
    }
    return b
}
```


[115. 不同的子序列](https://leetcode-cn.com/problems/distinct-subsequences/)
```
func numDistinct(s string, t string) int {
    dp:=make([][]int,len(s)+1)
    for k,_:=range dp{
        dp[k]=make([]int,len(t)+1)
    }
    for k,_:=range dp{
        dp[k][0]=1
    }
    
    for i:=0;i<len(s);i++{
        for j:=0;j<len(t);j++{
            if s[i]==t[j]{
                dp[i+1][j+1]=dp[i][j]+dp[i][j+1]
            }else{
                dp[i+1][j+1]=dp[i][j+1]
            }
            
        }
    }
   
    return dp[len(s)][len(t)]
}
```

[392. 判断子序列](https://leetcode-cn.com/problems/is-subsequence/)
```
func isSubsequence(s string, t string) bool {
    dp:=make([][]int,len(s)+1)
    for i:=0;i<=len(s);i++{
        dp[i]=make([]int,len(t)+1)
    }
    max:=0
    for i:=0;i<len(s);i++{
        for j:=0;j<len(t);j++{
            if s[i]==t[j]{
                dp[i+1][j+1]=dp[i][j]+1
            }else{
                dp[i+1][j+1]=dp[i+1][j]
            }
            if max<dp[i+1][j+1]{
                max=dp[i+1][j+1]
            }
        }
    }
    if max==len(s){
        return true
    }
    return false
}
```


[583. 两个字符串的删除操作](https://leetcode-cn.com/problems/delete-operation-for-two-strings/)
```
func minDistance(word1 string, word2 string) int {
    dp:=make([][]int,len(word1)+1)
    for k,_:=range dp{
        dp[k]=make([]int,len(word2)+1)
    } 
    for k,_:=range dp{
        dp[k][0]=k
    }
    for k,_:=range dp[0]{
        dp[0][k]=k
    }
    for k1,v1:=range word1{
        for k2,v2:=range word2{
            if v1==v2{
                dp[k1+1][k2+1]=dp[k1][k2]
            }else{
                dp[k1+1][k2+1]=min(min(dp[k1][k2]+2,dp[k1][k2+1]+1),dp[k1+1][k2]+1)
            }
        }
    }
    return dp[len(word1)][len(word2)]
}

func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}
```