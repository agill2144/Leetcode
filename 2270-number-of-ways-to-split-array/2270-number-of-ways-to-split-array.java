class Solution {
    public int waysToSplitArray(int[] nums) {
                long  [] left = new long  [nums.length];
        long  lSum = 0;
        for (int i = 0; i<nums.length; i++) {
            lSum += nums[i];
            left[i] = lSum;
        }
        System.out.println(Arrays.toString(left));
        long  [] right = new long[nums.length];
        long  rSum = 0;
        for (int j = nums.length-1; j>=0; j--) {
            rSum += nums[j];
            right[j] = rSum;
        }
        System.out.println(Arrays.toString(right));
        int i = 0, j = 1;
        int ans = 0;
        while (i<left.length && j<right.length) {
            if (left[i]>=right[j]) ans++;
            i++;
            j++;
        }
        return ans;
    }
}