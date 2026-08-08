public class Solution {
    public int Search(int[] nums, int target) {
        int index = -1;
        int right = nums.Length - 1;
        int left = 0;
        while (left <= right)
        {
            int mid = (left+right)/2;// it's not a good idea tho
            if (nums[mid] > target)
                right = mid - 1;
            else if (nums[mid] < target)
                left = mid + 1;
            else {
                return mid;
            }
        }
        
        return index;
    }
}
