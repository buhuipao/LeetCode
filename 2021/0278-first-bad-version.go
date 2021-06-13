/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 二分查找
func firstBadVersion(n int) int {
    L, R := 1, n
    var ans int
    for L <= R {
        M := L + (R - L) / 2
        if isBadVersion(M) {
            ans = M
            R = M - 1
        } else {
            L = M + 1
        }
    }
    return ans
}
