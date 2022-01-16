package corporateflightbookings

import rangeadditon "leetcode/370rangeAdditon"

func CorpFlightBookings(bookings [][]int, n int) []int {
	//increment里航班index要减1
	return rangeadditon.GetModifiedArray(n, bookings)
}
