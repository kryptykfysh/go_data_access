package models

import (
  "database/sql"
  "log"
  "time"
)

type Account struct {
  Id                sql.NullInt64
  RemoteId          sql.NullInt64
  ParentId          sql.NullInt64
  DataSourceId      sql.NullInt64
  Test              sql.NullBool
  CompanyName       sql.NullString
  DirectURL         sql.NullString
  RemoteCreatedAt   time.Time
  BrandId           sql.NullInt64
  ABN               sql.NullString
  ACN               sql.NullString
  Balance           sql.NullFloat64
  Login             sql.NullString
  Password          sql.NullString
}

func (db *DB) AllAccounts() ([]*Account, error) {
  rows, err := db.Query("SELECT * FROM accounts")
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  defer rows.Close()

  accs := make([]*Account, 0)
  for rows.Next() {
    acc := new(Account)
    err := rows.Scan(&acc.Id, &acc.RemoteId, &acc.ParentId, &acc.DataSourceId, &acc.Test, &acc.CompanyName, &acc.DirectURL, &acc.RemoteCreatedAt, &acc.BrandId, &acc.ABN, &acc.ACN, &acc.Balance, &acc.Login, &acc.Password)
    if err != nil {
      log.Fatal(err)
      return nil, err
    }
    accs = append(accs, acc)
  }
  if err = rows.Err(); err != nil {
    return nil, err
  }
  return accs, nil
}
