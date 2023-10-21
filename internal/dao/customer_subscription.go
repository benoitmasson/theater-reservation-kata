package dao

type CustomerSubscriptionDAO struct{}

// FetchCustomerSubscription simulates fetching data from Customer advantages
func (dao *CustomerSubscriptionDAO) FetchCustomerSubscription(customerID int64) bool {
	isSubscribed := false
	if customerID == 1 {
		isSubscribed = true
	}
	return isSubscribed
}
