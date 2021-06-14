/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    L, R := 1, n
    for L <= R {
        M := L + (R - L) / 2
        tmp := guess(M)
        if tmp == 0 {
            return M
        }
        if tmp == -1 {
            R = M - 1
        } else {
            L = M + 1
        }
        
    }
    return -1
}
