class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        hash = {}
        for index, num in enumerate(nums):
            minus = target - num
            try:
                hash[minus]
                return [hash[minus], index]
            except KeyError:
                hash[num] = index
