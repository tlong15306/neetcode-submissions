class Solution {
    public int[] getConcatenation(int[] nums) {
        int n = nums.length; // Khai báo n
        int[] ans = new int[2 * n]; // Sửa 2n thành 2 * n
        
        for(int i = 0; i < 2 * n; i++){
            if(i < n){
                ans[i] = nums[i];
            } else {
                ans[i] = nums[i - n];
            }
        }
        
        return ans; // Trả về mảng kết quả
    }
}