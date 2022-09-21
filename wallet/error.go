package wallet

// InternalWalletError wraps errors encountered with individual Wallets and Accounts
type InternalWalletError struct {
	Message  string
	Affected []int
	Err      error
}

// Unwrap returns the embedded error
func (err InternalWalletError) Unwrap() error {
	return err.Err
}

func (err InternalWalletError) Error() string {
	m := err.Message
	if err.Err != nil {
		m += " : " + err.Err.Error()
	}
	return m
}

// MultiWalletError wraps errors encountered with the Multiwallet
type MultiWalletError struct {
	Message string
	Err     error
}

func (err MultiWalletError) Error() string {
	m := err.Message
	if err.Err != nil {
		m += " : " + err.Err.Error()
	}
	return m
}

// Unwrap returns the embedded error
func (err MultiWalletError) Unwrap() error {
	return err.Err
}
