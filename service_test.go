package wallet

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ilyakaznacheev/tiny-wallet/internal/model"
	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

var testDatabaseErr = errors.New("test error")

type testDatabaseData struct {
	dat interface{}
	err error
}
type TestDatabase struct {
	GetAllAccountsData testDatabaseData
	GetAllPaymentsData testDatabaseData
	GetAccountData     map[string]testDatabaseData
	CreatePaymentData  testDatabaseData
	CreateAccountData  testDatabaseData
}

func (db *TestDatabase) GetAllAccounts() ([]model.Account, error) {
	return db.GetAllAccountsData.dat.([]model.Account), db.GetAllAccountsData.err
}

func (db *TestDatabase) GetAllPayments() ([]model.Payment, error) {
	return db.GetAllPaymentsData.dat.([]model.Payment), db.GetAllPaymentsData.err
}

func (db *TestDatabase) GetAccount(accountID string) (*model.Account, error) {
	testData := db.GetAccountData[accountID]
	return testData.dat.(*model.Account), testData.err
}

func (db *TestDatabase) CreatePayment(p model.Payment, lastChangedFrom, lastChangedTo *time.Time) (*model.Payment, error) {
	return db.CreatePaymentData.dat.(*model.Payment), db.CreatePaymentData.err
}

func (db *TestDatabase) CreateAccount(a model.Account) (*model.Account, error) {
	return db.CreateAccountData.dat.(*model.Account), db.CreateAccountData.err
}

func TestServiceGetAllPayments(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		db      Database
		want    []model.Payment
		wantErr bool
	}{
		{
			name: "simple",
			db: &TestDatabase{
				GetAllPaymentsData: testDatabaseData{
					dat: []model.Payment{
						model.Payment{
							ID:        1,
							AccFromID: "1",
							AccToID:   "2",
							DateTime:  now,
							Amount:    12345,
							Currency:  currency.USD,
						},
						model.Payment{
							ID:        2,
							AccFromID: "2",
							AccToID:   "3",
							DateTime:  now,
							Amount:    456,
							Currency:  currency.USD,
						},
					},
					err: nil,
				},
			},
			want: []model.Payment{
				model.Payment{
					ID:        1,
					AccFromID: "1",
					AccToID:   "2",
					DateTime:  now,
					Amount:    12345,
					Currency:  currency.USD,
				},
				model.Payment{
					ID:        2,
					AccFromID: "2",
					AccToID:   "3",
					DateTime:  now,
					Amount:    456,
					Currency:  currency.USD,
				},
			},
			wantErr: false,
		},

		{
			name: "error",
			db: &TestDatabase{
				GetAllPaymentsData: testDatabaseData{
					dat: []model.Payment{},
					err: testDatabaseErr,
				},
			},
			want:    []model.Payment{},
			wantErr: true,
		},

		{
			name: "not found",
			db: &TestDatabase{
				GetAllPaymentsData: testDatabaseData{
					dat: []model.Payment{},
					err: sql.ErrNoRows,
				},
			},
			want:    []model.Payment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &walletService{
				db: tt.db,
			}
			got, err := s.GetAllPayments(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("wrong error state %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrong response value %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceGetAllAccounts(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name    string
		db      Database
		want    []model.Account
		wantErr bool
	}{
		{
			name: "simple",
			db: &TestDatabase{
				GetAllAccountsData: testDatabaseData{
					dat: []model.Account{
						model.Account{
							ID:         "1",
							LastUpdate: &now,
							Balance:    12345,
							Currency:   currency.USD,
						},
						model.Account{
							ID:         "2",
							LastUpdate: &now,
							Balance:    67890,
							Currency:   currency.USD,
						},
					},
					err: nil,
				},
			},
			want: []model.Account{
				model.Account{
					ID:         "1",
					LastUpdate: &now,
					Balance:    12345,
					Currency:   currency.USD,
				},
				model.Account{
					ID:         "2",
					LastUpdate: &now,
					Balance:    67890,
					Currency:   currency.USD,
				},
			},
			wantErr: false,
		},

		{
			name: "error",
			db: &TestDatabase{
				GetAllAccountsData: testDatabaseData{
					dat: []model.Account{},
					err: testDatabaseErr,
				},
			},
			want:    []model.Account{},
			wantErr: true,
		},

		{
			name: "not found",
			db: &TestDatabase{
				GetAllAccountsData: testDatabaseData{
					dat: []model.Account{},
					err: sql.ErrNoRows,
				},
			},
			want:    []model.Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &walletService{
				db: tt.db,
			}
			got, err := s.GetAllAccounts(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("wrong error state %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrong response value %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServicePostPayment(t *testing.T) {
	now := time.Now()
	type args struct {
		fromID string
		toID   string
		amount float64
	}
	tests := []struct {
		name    string
		args    args
		db      Database
		want    *model.Payment
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 123,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{
							ID:         "1",
							LastUpdate: &now,
							Balance:    12345,
							Currency:   currency.USD,
						},
						err: nil,
					},
					"2": testDatabaseData{
						dat: &model.Account{
							ID:         "2",
							LastUpdate: &now,
							Balance:    67890,
							Currency:   currency.USD,
						},
						err: nil,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{
						ID:        1,
						AccFromID: "1",
						AccToID:   "2",
						DateTime:  now,
						Amount:    12345,
						Currency:  currency.USD,
					},
					err: nil,
				},
			},
			want: &model.Payment{
				ID:        1,
				AccFromID: "1",
				AccToID:   "2",
				DateTime:  now,
				Amount:    12345,
				Currency:  currency.USD,
			},
			wantErr: false,
		},

		{
			name: "error one",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 123,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{},
						err: testDatabaseErr,
					},
					"2": testDatabaseData{
						dat: &model.Account{
							ID:         "2",
							LastUpdate: &now,
							Balance:    67890,
							Currency:   currency.USD,
						},
						err: nil,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{
						ID:        1,
						AccFromID: "1",
						AccToID:   "2",
						DateTime:  now,
						Amount:    12345,
						Currency:  currency.USD,
					},
					err: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "error two",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 123,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{
							ID:         "1",
							LastUpdate: &now,
							Balance:    12345,
							Currency:   currency.USD,
						},
						err: nil,
					},
					"2": testDatabaseData{
						dat: &model.Account{},
						err: testDatabaseErr,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{
						ID:        1,
						AccFromID: "1",
						AccToID:   "2",
						DateTime:  now,
						Amount:    12345,
						Currency:  currency.USD,
					},
					err: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "not found one",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 123,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{},
						err: sql.ErrNoRows,
					},
					"2": testDatabaseData{
						dat: &model.Account{
							ID:         "2",
							LastUpdate: &now,
							Balance:    67890,
							Currency:   currency.USD,
						},
						err: nil,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{
						ID:        1,
						AccFromID: "1",
						AccToID:   "2",
						DateTime:  now,
						Amount:    12345,
						Currency:  currency.USD,
					},
					err: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "not found two",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 123,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{
							ID:         "1",
							LastUpdate: &now,
							Balance:    12345,
							Currency:   currency.USD,
						},
						err: nil,
					},
					"2": testDatabaseData{
						dat: &model.Account{},
						err: sql.ErrNoRows,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{
						ID:        1,
						AccFromID: "1",
						AccToID:   "2",
						DateTime:  now,
						Amount:    12345,
						Currency:  currency.USD,
					},
					err: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "currency error",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 123,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{
							ID:         "1",
							LastUpdate: &now,
							Balance:    12345,
							Currency:   currency.USD,
						},
						err: nil,
					},
					"2": testDatabaseData{
						dat: &model.Account{
							ID:         "2",
							LastUpdate: &now,
							Balance:    67890,
							Currency:   currency.CAD,
						},
						err: nil,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{},
					err: nil,
				},
			},
			want:    &model.Payment{},
			wantErr: true,
		},

		{
			name: "amount error",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 456,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{
							ID:         "1",
							LastUpdate: &now,
							Balance:    123,
							Currency:   currency.USD,
						},
						err: nil,
					},
					"2": testDatabaseData{
						dat: &model.Account{
							ID:         "2",
							LastUpdate: &now,
							Balance:    567,
							Currency:   currency.USD,
						},
						err: nil,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{},
					err: nil,
				},
			},
			want:    &model.Payment{},
			wantErr: true,
		},

		{
			name: "creation error",
			args: args{
				fromID: "1",
				toID:   "2",
				amount: 123,
			},
			db: &TestDatabase{
				GetAccountData: map[string]testDatabaseData{
					"1": testDatabaseData{
						dat: &model.Account{
							ID:         "1",
							LastUpdate: &now,
							Balance:    12345,
							Currency:   currency.USD,
						},
						err: nil,
					},
					"2": testDatabaseData{
						dat: &model.Account{
							ID:         "2",
							LastUpdate: &now,
							Balance:    67890,
							Currency:   currency.USD,
						},
						err: nil,
					},
				},
				CreatePaymentData: testDatabaseData{
					dat: &model.Payment{},
					err: testDatabaseErr,
				},
			},
			want:    &model.Payment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &walletService{
				db: tt.db,
			}
			got, err := s.PostPayment(context.Background(), tt.args.fromID, tt.args.toID, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("wrong error state %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrong response value %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walletService_PostAccount(t *testing.T) {
	now := time.Now()
	type args struct {
		id      string
		balance float64
		curr    string
	}
	tests := []struct {
		name    string
		args    args
		db      Database
		want    *model.Account
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				id:      "1",
				balance: 123.45,
				curr:    "USD",
			},
			db: &TestDatabase{
				CreateAccountData: testDatabaseData{
					dat: &model.Account{
						ID:         "1",
						LastUpdate: &now,
						Balance:    12345,
						Currency:   currency.USD,
					},
					err: nil,
				},
			},
			want: &model.Account{
				ID:         "1",
				LastUpdate: &now,
				Balance:    12345,
				Currency:   currency.USD,
			},
			wantErr: false,
		},

		{
			name: "error currency",
			args: args{
				id:      "1",
				balance: 123.45,
				curr:    "AAA",
			},
			db:      &TestDatabase{},
			want:    &model.Account{},
			wantErr: true,
		},

		{
			name: "error amount",
			args: args{
				id:      "1",
				balance: -123.45,
				curr:    "USD",
			},
			db:      &TestDatabase{},
			want:    &model.Account{},
			wantErr: true,
		},

		{
			name: "error creation",
			args: args{
				id:      "1",
				balance: 123.45,
				curr:    "USD",
			},
			db: &TestDatabase{
				CreateAccountData: testDatabaseData{
					dat: &model.Account{},
					err: testDatabaseErr,
				},
			},
			want:    &model.Account{},
			wantErr: true,
		},

		{
			name: "error already exists",
			args: args{
				id:      "1",
				balance: 123.45,
				curr:    "USD",
			},
			db: &TestDatabase{
				CreateAccountData: testDatabaseData{
					dat: &model.Account{},
					err: model.ErrRowExists,
				},
			},
			want:    &model.Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &walletService{
				db: tt.db,
			}
			got, err := s.PostAccount(context.Background(), tt.args.id, tt.args.balance, tt.args.curr)
			if (err != nil) != tt.wantErr {
				t.Errorf("walletService.PostAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walletService.PostAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
