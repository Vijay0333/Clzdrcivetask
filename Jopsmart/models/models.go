package models

type Car struct {
  ID          int64  `gorm:"primary_key;auto_increment"`
  LicensePlate string `gorm:"unique"`
  Model       string `gorm:"nullable"`
  EntryTime   time.Time
  Status      string
  ExitTime    time.Time `gorm:"null"`
}
